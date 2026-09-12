package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/seqsense/s3sync/v2"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if err := run(ctx, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("usage: s3sync source target number_of_parallel_jobs")
	}
	numberOfParallelJobs, err := strconv.Atoi(args[2])
	if err != nil || numberOfParallelJobs <= 0 {
		return fmt.Errorf("number_of_parallel_jobs must be a positive integer")
	}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithDefaultRegion("eu-central-2"))
	if err != nil {
		return fmt.Errorf("load AWS configuration: %w", err)
	}

	syncManager := s3sync.New(cfg, s3sync.WithParallel(numberOfParallelJobs))
	return syncManager.Sync(ctx, args[0], args[1])
}
