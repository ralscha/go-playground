package main

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yml api.yml

import (
	"context"
	"fmt"
	"log"
)

func main() {
	client, err := NewClientWithResponses("https://api.chucknorris.io")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	response, err := client.GetRandomFactWithResponse(ctx)
	if err != nil {
		log.Fatalf("Failed to get random fact: %v", err)
	}

	if response.StatusCode() != 200 {
		log.Fatalf("API returned status %d: %s", response.StatusCode(), response.Status())
	}

	if response.JSON200 != nil {
		fact := response.JSON200
		fmt.Println("Chuck Norris Fact:")
		fmt.Println("==================================================")

		if fact.Value != nil {
			fmt.Printf("Fact: %s\n", *fact.Value)
		}

		if fact.Id != nil {
			fmt.Printf("ID: %s\n", *fact.Id)
		}

		if fact.IconUrl != nil && *fact.IconUrl != "" {
			fmt.Printf("Icon: %s\n", *fact.IconUrl)
		}

		if fact.Url != nil && *fact.Url != "" {
			fmt.Printf("URL: %s\n", *fact.Url)
		}
	} else {
		fmt.Println("No fact data received")
	}
}
