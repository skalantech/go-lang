package main

import (
	"fmt"
	"sync"
)

// Constants
const N_THREADS = 6

// The order in which threads should execute
var Ordre = [N_THREADS]int{2, 3, 1, 4, 5, 0}

// Channels for synchronization
var channels [N_THREADS]chan bool

// Function that each goroutine will run
func FonctionThread(threadNum int, wg *sync.WaitGroup) {
	// Wait for the channel signal to start
	<-channels[threadNum]

	// Critical section: thread-specific work
	fmt.Printf("Je suis le thread %d.\n", threadNum)

	// Find the index of the current thread in the order array
	for i := 0; i < N_THREADS; i++ {
		if Ordre[i] == threadNum && i+1 < N_THREADS {
			// Signal the next thread in the order
			channels[Ordre[i+1]] <- true
			break
		}
	}

	// Mark the goroutine as done
	wg.Done()
}

func main() {
	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Initialize channels
	for i := 0; i < N_THREADS; i++ {
		channels[i] = make(chan bool)
	}

	// Create goroutines
	for i := 0; i < N_THREADS; i++ {
		wg.Add(1)
		go FonctionThread(i, &wg)
	}

	// Signal the first thread to start
	channels[Ordre[0]] <- true

	// Wait for all goroutines to finish
	wg.Wait()
}
