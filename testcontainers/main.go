package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	redisC, err := testcontainers.Run(ctx, "redis:8.2-alpine",
		testcontainers.WithExposedPorts("6379/tcp"),
		testcontainers.WithWaitStrategy(wait.ForLog("Ready to accept connections")),
	)
	defer func() {
		if err := testcontainers.TerminateContainer(redisC); err != nil {
			log.Printf("terminate Redis container: %v", err)
		}
	}()
	if err != nil {
		return err
	}

	endpoint, err := redisC.PortEndpoint(ctx, "6379/tcp", "")
	if err != nil {
		return err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     endpoint,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		return err
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		return err
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		return err
	} else {
		fmt.Println("key2", val2)
	}
	return nil
}
