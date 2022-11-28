package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/seqsense/s3sync"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		println("Usage: s3sync source target number_of_parallel_jobs")
		return
	}

	source := os.Args[1]
	target := os.Args[2]
	numberOfParallelJobs, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-2"),
	})

	syncManager := s3sync.New(sess, s3sync.WithParallel(numberOfParallelJobs))

	err = syncManager.Sync(source, target)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

}
