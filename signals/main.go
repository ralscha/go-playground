package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	write()
}

func write() {
	// This program creates a numbers.txt file and populates it with a lot of
	// numbers, one per line. Errors omitted for brevity.

	// Create a file and wrap it in a buffered writer.
	f, _ := os.Create("numbers.txt")
	bw := bufio.NewWriter(f)

	// Close this channel to stop the number generator.
	stopch := make(chan struct{})

	// Use a WaitGroup to wait for the number generator to stop.
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Run the number generator in a separate Go routine.
	go func() {
		defer wg.Done()
	L:
		for i := 0; i < 1000000000; i++ {
			_, err := fmt.Fprintf(bw, "%d\n", i)
			if err != nil {
				log.Fatal(err)
			}
			select {
			case <-stopch:
				// The stopch channel has been closed. Break the loop.
				break L
			default:
			}
		}
	}()

	// Make a signal channel. Register SIGINT.
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)

	// In a separate Go routine, wait for the signal. On signal, close the
	// stopch channel so that the number generator stops.
	go func() {
		<-sigch
		fmt.Println("Interrupted.")
		close(stopch)
	}()

	// Wait for the number generator to stop.
	wg.Wait()

	fmt.Println("Flushing.")

	// Flush buffered writer. Close file.
	err := bw.Flush()
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Exiting.")
}

func web() {
	server := http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Simulate a slow HTTP response.
			time.Sleep(10 * time.Second)
			_, err := io.WriteString(w, "Hello")
			if err != nil {
				log.Fatal(err)
			}
		}),
	}

	// Start the HTTP server in a separate Go routine.
	go func() {
		fmt.Println("Listening for HTTP connections.")
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Make a signal channel. Register SIGINT.
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)

	// Wait for the signal.
	<-sigch

	fmt.Println("Interrupted. Exiting.")

	// Trigger a shutdown and allow 13 seconds to drain connections. Ignoring
	// CancelFunc for brevity.
	ctx, cancel := context.WithTimeout(context.Background(), 13*time.Second)
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
	cancel()
}

func notify() {
	// Make a signal-based context. The stop function, when called, unregisters
	// the signals and restores the default signal behaviour.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	// Wait for the signal.
	<-ctx.Done()
	stop() // After calling stop, another SIGINT will terminate the program.

	fmt.Println("Interrupted. Exiting.")

	// Long clean-up code goes here.
	time.Sleep(5 * time.Second)
}

func simple() {
	fmt.Println("Waiting for signal.")

	// Make a buffered channel.
	sigch := make(chan os.Signal, 1)

	// Register the signals that you want to handle.
	signal.Notify(sigch, os.Interrupt)

	// Wait for the signal.
	<-sigch

	fmt.Println("Received interrupt. Exiting.")
}
