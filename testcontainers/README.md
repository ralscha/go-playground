# Redis with Testcontainers

Run `go run .` with Docker available to start a temporary Redis 8.2 container,
write and read a key with `go-redis/v9`, and demonstrate a missing key.

The example uses `testcontainers.Run` with functional options and resolves the
mapped Redis port. It closes the Redis client and terminates the container when
the example completes, including on errors.
