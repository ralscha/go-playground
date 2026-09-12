package main

import (
	"bytes"
	"encoding/hex"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func TestDeriveKeyCompatibility(t *testing.T) {
	// Fixed PBKDF2-HMAC-SHA256 vector, independently calculated with .NET.
	// Preserve the parameters used by existing encrypted files.
	salt := []byte("12345678")
	key, gotSalt, err := deriveKey("password", salt)
	if err != nil {
		t.Fatal(err)
	}
	const want = "3473dcdd934a9897f0bcbcb7fa8c17ca5151950d0ca0e5267b1a32ba23a8f8c0"
	if hex.EncodeToString(key) != want || !bytes.Equal(gotSalt, salt) {
		t.Fatalf("derived key or salt changed: key=%x salt=%x", key, gotSalt)
	}
}

func TestS3CopyRoundTrip(t *testing.T) {
	for _, password := range []string{"", "example-password"} {
		name := "plain"
		if password != "" {
			name = "encrypted"
		}
		t.Run(name, func(t *testing.T) {
			var mu sync.Mutex
			var object []byte
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/test-bucket/nested/object.txt" {
					t.Errorf("unexpected object path: %s", r.URL.Path)
					http.NotFound(w, r)
					return
				}
				mu.Lock()
				defer mu.Unlock()
				switch r.Method {
				case http.MethodPut:
					var err error
					object, err = io.ReadAll(r.Body)
					if err != nil {
						t.Error(err)
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					w.Header().Set("ETag", `"test-etag"`)
				case http.MethodGet:
					w.Header().Set("Content-Length", strconv.Itoa(len(object)))
					_, _ = w.Write(object)
				default:
					t.Errorf("unexpected S3 method: %s", r.Method)
					w.WriteHeader(http.StatusMethodNotAllowed)
				}
			}))
			defer server.Close()

			client := s3.New(s3.Options{
				Region:                     "eu-central-2",
				BaseEndpoint:               &server.URL,
				UsePathStyle:               true,
				Credentials:                aws.AnonymousCredentials{},
				HTTPClient:                 server.Client(),
				RequestChecksumCalculation: aws.RequestChecksumCalculationWhenRequired,
			})
			dir := t.TempDir()
			source := filepath.Join(dir, "source.txt")
			target := filepath.Join(dir, "download.txt")
			payload := bytes.Repeat([]byte("S3 transfer example\n"), 137)
			if err := os.WriteFile(source, payload, 0o600); err != nil {
				t.Fatal(err)
			}
			const objectURL = "s3://test-bucket/nested/object.txt"
			s3Upload(t.Context(), client, source, objectURL, password)
			mu.Lock()
			stored := bytes.Clone(object)
			mu.Unlock()
			if password == "" && !bytes.Equal(stored, payload) {
				t.Fatal("uploaded object differs from source")
			}
			if password != "" && (bytes.Equal(stored, payload) || len(stored) != len(payload)+16+8) {
				t.Fatal("encrypted object must contain ciphertext followed by its IV and salt")
			}
			s3Download(t.Context(), client, objectURL, target, password)
			got, err := os.ReadFile(target)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(got, payload) {
				t.Fatal("downloaded content differs from source")
			}
			for _, path := range []string{source + ".enc", target + ".enc"} {
				if _, err := os.Stat(path); !os.IsNotExist(err) {
					t.Errorf("temporary encrypted file remains: %s (%v)", path, err)
				}
			}
		})
	}
}
