package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	// Wait for the signal
	<-ch
	// Stop receiving signals
	signal.Stop(ch)

	fmt.Println("Shutting down...")
	// Simulate a slow shutdown
	time.Sleep(3 * time.Second)
	// Shutdown has completed
	fmt.Println("Shutdown complete")
}
