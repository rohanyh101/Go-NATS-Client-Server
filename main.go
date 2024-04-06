package main

import (
	"log"
	"sync"
)

func main() {
	log.Println("Starting...")

	js, err := JetStreamInit()
	if err != nil {
		log.Fatalf("Failed to initialize JetStream: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		publishReviews(js)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		consumeReviews(js)
	}()

	wg.Wait()

	log.Println("Exit...")
}
