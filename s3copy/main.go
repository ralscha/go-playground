package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	if len(os.Args) < 3 {
		println("Usage: s3copy source target [password]")
		return
	}

	source := os.Args[1]
	target := os.Args[2]

	var password string
	if len(os.Args) > 3 {
		password = os.Args[3]
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	if strings.HasPrefix(target, "s3://") {
		s3Upload(ctx, s3Client, source, target, password)
	} else {
		s3Download(ctx, s3Client, source, target, password)
	}

}

func s3Upload(ctx context.Context, s3client *s3.Client, source, target, password string) {
	println("Uploading", source, "to", target)
	bucket, s3key := parseS3Url(target)

	inFile, err := os.Open(source)
	checkError(err)

	if password != "" {
		key, salt, err := deriveKey(password, nil)
		checkError(err)
		block, err := aes.NewCipher(key)
		checkError(err)

		iv := make([]byte, block.BlockSize())
		_, err = rand.Read(iv)
		checkError(err)

		outFile, err := os.Create(source + ".enc")
		checkError(err)

		buf := make([]byte, 1024)
		stream := cipher.NewCTR(block, iv)
		for {
			n, err := inFile.Read(buf)
			if n > 0 {
				stream.XORKeyStream(buf, buf[:n])
				_, err := outFile.Write(buf[:n])
				checkError(err)
			}

			if err == io.EOF {
				break
			}

			checkError(err)
		}
		_, err = outFile.Write(iv)
		checkError(err)

		_, err = outFile.Write(salt)
		checkError(err)

		err = outFile.Close()
		checkError(err)

		err = inFile.Close()
		checkError(err)

		inFile, err = os.Open(source + ".enc")
		checkError(err)
	}

	transfers := transfermanager.New(s3client)
	_, err = transfers.UploadObject(ctx, &transfermanager.UploadObjectInput{
		Bucket: &bucket,
		Key:    &s3key,
		Body:   inFile,
	})
	checkError(err)

	err = inFile.Close()
	checkError(err)

	if password != "" {
		err = os.Remove(source + ".enc")
		checkError(err)
	}

}

func s3Download(ctx context.Context, s3client *s3.Client, source, target, password string) {
	println("Downloading", source, "to", target)
	bucket, key := parseS3Url(source)

	transfers := transfermanager.New(s3client)
	var fileName string
	if password != "" {
		fileName = target + ".enc"
	} else {
		fileName = target
	}
	outFile, err := os.Create(fileName)
	checkError(err)

	_, err = transfers.DownloadObject(ctx, &transfermanager.DownloadObjectInput{
		Bucket:   &bucket,
		Key:      &key,
		WriterAt: outFile,
	})
	checkError(err)

	err = outFile.Close()
	checkError(err)

	if password != "" {
		inFile, err := os.Open(fileName)
		checkError(err)

		fi, err := inFile.Stat()
		checkError(err)

		salt := make([]byte, 8)
		_, err = inFile.ReadAt(salt, fi.Size()-int64(len(salt)))
		checkError(err)

		key, _, err := deriveKey(password, salt)
		checkError(err)
		block, err := aes.NewCipher(key)
		checkError(err)

		iv := make([]byte, block.BlockSize())
		msgLen := fi.Size() - int64(len(iv)) - int64(len(salt))
		_, err = inFile.ReadAt(iv, msgLen)
		checkError(err)

		outFile, err := os.Create(target)
		checkError(err)

		buf := make([]byte, 1024)
		stream := cipher.NewCTR(block, iv)
		for {
			n, err := inFile.Read(buf)
			if n > 0 {
				if n > int(msgLen) {
					n = int(msgLen)
				}
				msgLen -= int64(n)
				stream.XORKeyStream(buf, buf[:n])
				_, err := outFile.Write(buf[:n])
				checkError(err)
			}

			if err == io.EOF {
				break
			}

			checkError(err)
		}
		err = outFile.Close()
		checkError(err)

		err = inFile.Close()
		checkError(err)

		err = os.Remove(fileName)
		checkError(err)
	}

}

func deriveKey(passphrase string, salt []byte) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, 8)
		_, err := rand.Read(salt)
		if err != nil {
			return nil, nil, err
		}
	}
	key, err := pbkdf2.Key(sha256.New, passphrase, salt, 10_000, 32)
	return key, salt, err
}

func parseS3Url(url string) (bucket, key string) {
	url = strings.TrimPrefix(url, "s3://")
	parts := strings.SplitN(url, "/", 2)
	return parts[0], parts[1]
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}
