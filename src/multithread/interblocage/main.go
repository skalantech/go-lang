package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Constants similar to the C code
const (
	N_MUTEX    = 10
	N_THREADS  = 2
	N_PIGE     = 1
	SLEEP_TIME = 3 * time.Millisecond
)

var mutexes [N_MUTEX]sync.Mutex

// Generates unique random numbers
func numerosDifferentsHasard(nNumero, max int) []int {
	selected := make(map[int]bool)
	var result []int
	for len(result) < nNumero {
		candidate := rand.Intn(max)
		if !selected[candidate] {
			selected[candidate] = true
			result = append(result, candidate)
		}
	}
	return result
}

// Goroutine function simulating thread behavior
func threadFunction(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	iteration := 0

	for {
		fmt.Printf("------ Iteration %d for thread %d ------\n", iteration, id)
		iteration++

		// Select random mutexes
		tableau := numerosDifferentsHasard(N_PIGE, N_MUTEX)

		// Lock the selected mutexes
		for _, m := range tableau {
			fmt.Printf("Thread %d acquiring mutex %d\n", id, m)
			mutexes[m].Lock()
		}

		fmt.Printf("All mutexes acquired by thread %d\n", id)
		time.Sleep(SLEEP_TIME)

		// Unlock the selected mutexes
		for _, m := range tableau {
			fmt.Printf("Thread %d releasing mutex %d\n", id, m)
			mutexes[m].Unlock()
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	// Create threads (goroutines in Go)
	for i := 0; i < N_THREADS; i++ {
		fmt.Printf("Main starting thread %d\n", i)
		wg.Add(1)
		go threadFunction(i, &wg)
	}

	// Wait indefinitely (equivalent to pthread_join in the infinite loop case)
	wg.Wait()
}
