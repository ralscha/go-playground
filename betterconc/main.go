package main

import (
	"context"
	"fmt"
	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/iter"
	"github.com/sourcegraph/conc/pool"
	"io"
	"log"
	"net/http"
	"sync"
)

func main() {
	// waitGroupExample()
	// concWaitGroupExample()

	jokes, err := fetchJokesPool(context.Background(), []string{"Programming", "Miscellaneous", "Dark", "Pun", "Spooky", "Christmas"})
	if err != nil {
		log.Fatal(err)
	}
	for _, joke := range jokes {
		fmt.Println(joke)
	}
	fmt.Println("=====================================")
	jokes, err = fetchJokes2(context.Background(), []string{"Programming", "Miscellaneous", "Dark", "Pun", "Spooky", "Christmas"})
	if err != nil {
		log.Fatal(err)
	}
	for _, joke := range jokes {
		fmt.Println(joke)
	}

}

func fetchJoke(ctx context.Context, category string) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		fmt.Sprintf("https://v2.jokeapi.dev/joke/%s", category),
		nil,
	)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	b, err := io.ReadAll(resp.Body)
	return string(b), err
}

func fetchJokesPool(ctx context.Context, categories []string) ([]string, error) {
	p := pool.NewWithResults[string]().WithMaxGoroutines(5).WithContext(ctx)
	for _, category := range categories {
		category := category
		p.Go(func(ctx context.Context) (string, error) {
			return fetchJoke(ctx, category)
		})
	}
	return p.Wait()
}

func fetchJokes2(ctx context.Context, categories []string) ([]string, error) {
	return iter.MapErr(categories, func(category *string) (string, error) {
		return fetchJoke(ctx, *category)
	})
}

func waitGroupExample() {
	var wg sync.WaitGroup
	wg.Add(1)
	go doSomethingThatMightPanic2(&wg)
	wg.Wait()
}

func concWaitGroupExample() {
	var wg conc.WaitGroup
	wg.Go(doSomethingThatMightPanic)
	// panics with a nice stack trace
	wg.Wait()
}

func doSomethingThatMightPanic() {
	panic("oh no")
}

func doSomethingThatMightPanic2(wg *sync.WaitGroup) {
	defer wg.Done()
	panic("oh no")
}
