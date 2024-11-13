package main

import (
	"fmt"
	"sync"
)

const MAX = 1000000 // Number of items to produce

func producer(buffer chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= MAX; i++ {
		buffer <- i // Send item to the buffer (blocks if full)
	}
	close(buffer) // Close the channel when done
}

func consumer(buffer chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range buffer { // Receives items from the buffer
		item++
		// Consumer logic (e.g., processing item)
	}
}

func main() {
	buffer := make(chan int, 1) // Buffer with capacity of 1 (similar to the original C buffer)
	var wg sync.WaitGroup

	// Start producer
	wg.Add(1)
	go producer(buffer, &wg)

	// Start consumer
	wg.Add(1)
	go consumer(buffer, &wg)

	// Wait for both producer and consumer to finish
	wg.Wait()

	fmt.Println("Processing completed.")
}
