package main

// Test case: 3 producers, 3 consumers, 50 items
func test3Prod3Cons50Items() {
	queue := NewQueue(QueueCapacity, 50)
	createProducersConsumers(queue, 3, 3)
}

// Test case: 5 producers, 10 consumers, 100 items
func test5Prod10Cons100Items() {
	queue := NewQueue(QueueCapacity, 100)
	createProducersConsumers(queue, 5, 10)
}

// Test case: 10 producers, 5 consumers, 100 items
func test10Prod5Cons100Items() {
	queue := NewQueue(QueueCapacity, 100)
	createProducersConsumers(queue, 10, 5)
}

// Test case: 20 producers, 20 consumers, 500 items
func test20Prod20Cons500Items() {
	queue := NewQueue(QueueCapacity, 500)
	createProducersConsumers(queue, 20, 20)
}
