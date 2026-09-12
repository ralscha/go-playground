# S3 copy

Uploads or downloads a file using AWS SDK for Go v2's
[`transfermanager`](https://github.com/aws/aws-sdk-go-v2/discussions/3306).
Configure AWS credentials and region through the standard environment variables
or shared configuration files before running:

```sh
go run . ./file.txt s3://your-bucket/file.txt
go run . s3://your-bucket/file.txt ./download.txt
```

An optional third argument supplies a password for the example's encryption.
Key derivation uses Go's `crypto/pbkdf2`, retaining the existing SHA-256,
10,000-iteration, 32-byte key and ciphertext/IV/8-byte-salt file format.
The AES-CTR format does not authenticate the ciphertext or detect wrong passwords.

`go test ./...` checks plain and encrypted transfers against a local HTTP server
and verifies key compatibility without AWS credentials.
