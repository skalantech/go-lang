package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const QueueCapacity = 10

// Queue structure using channels
type Queue struct {
	items          chan int
	itemsProduced  int
	itemsConsumed  int
	totalItemsTest int
	mutex          sync.Mutex
	wg             sync.WaitGroup
}

// Initialize a new queue
func NewQueue(capacity, totalItemsTest int) *Queue {
	return &Queue{
		items:          make(chan int, capacity),
		totalItemsTest: totalItemsTest,
	}
}

// Producer function
func producer(id int, queue *Queue) {
	defer queue.wg.Done()
	for {
		queue.mutex.Lock()
		if queue.itemsProduced >= queue.totalItemsTest {
			queue.mutex.Unlock()
			break
		}
		queue.itemsProduced++
		item := rand.Intn(1000)
		queue.mutex.Unlock()

		// Produce item and send it to the queue
		queue.items <- item
		fmt.Printf("Producer %d produced item %d\n", id, item)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

// Consumer function
func consumer(id int, queue *Queue) {
	defer queue.wg.Done()
	for {
		queue.mutex.Lock()
		if queue.itemsConsumed >= queue.totalItemsTest {
			queue.mutex.Unlock()
			break
		}
		queue.itemsConsumed++
		queue.mutex.Unlock()

		// Consume item from the queue
		item := <-queue.items
		fmt.Printf("Consumer %d consumed item %d\n", id, item)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func createProducersConsumers(queue *Queue, numProducers, numConsumers int) {
	// Start producers
	for i := 0; i < numProducers; i++ {
		queue.wg.Add(1)
		go producer(i+1, queue)
	}

	// Start consumers
	for i := 0; i < numConsumers; i++ {
		queue.wg.Add(1)
		go consumer(i+1, queue)
	}

	queue.wg.Wait()
	close(queue.items)

	fmt.Printf("Total items produced: %d\n", queue.itemsProduced)
	fmt.Printf("Total items consumed: %d\n", queue.itemsConsumed)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Choose which test to run
	// test3Prod3Cons50Items()
	// test5Prod10Cons100Items()
	// test10Prod5Cons100Items()
	test20Prod20Cons500Items()
}
