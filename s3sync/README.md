# S3 sync

Synchronizes local directories and S3 prefixes using
[`s3sync/v2`](https://github.com/seqsense/s3sync/blob/master/MIGRATION.md)
and AWS SDK for Go v2.

```sh
go run . ./local-directory s3://your-bucket/prefix 4
go run . s3://your-bucket/prefix ./local-directory 4
```

The third argument is the positive number of parallel jobs. Configure credentials
and region through the standard AWS environment variables or shared configuration
files (`AWS_PROFILE` is supported). The region defaults to `eu-central-2` when no
region is configured. Ctrl+C cancels the sync through its context.

Upstream `s3sync/v2` v2.0.0 still depends internally on the deprecated AWS v2
`feature/s3/manager` package. Removing that transitive dependency requires an
upstream migration; this example no longer uses AWS SDK v1.
