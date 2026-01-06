package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Basic Concurrency Tutorial ===")
	fmt.Println()

	// Demonstrate basic goroutines
	demonstrateGoroutines()

	// Demonstrate channels
	demonstrateChannels()

	// Demonstrate buffered channels
	demonstrateBufferedChannels()

	// Demonstrate select statement
	demonstrateSelect()

	// Demonstrate WaitGroups
	demonstrateWaitGroup()
}

// worker simulates work by sleeping
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

// demonstrateGoroutines shows basic goroutine usage
func demonstrateGoroutines() {
	fmt.Println("--- Goroutines ---")

	// Sequential execution
	fmt.Println("Sequential:")
	worker(1)
	worker(2)

	fmt.Println("\nConcurrent:")
	// Concurrent execution with goroutines
	go worker(3)
	go worker(4)
	go worker(5)

	// Wait for goroutines to finish
	// In real code, use WaitGroups or channels (shown later)
	time.Sleep(200 * time.Millisecond)

	// Anonymous goroutine
	go func() {
		fmt.Println("Anonymous goroutine executed")
	}()

	time.Sleep(50 * time.Millisecond)
	fmt.Println()
}

// sum calculates sum and sends result on channel
func sum(numbers []int, result chan int) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	result <- total // Send result on channel
}

// demonstrateChannels shows channel creation and usage
func demonstrateChannels() {
	fmt.Println("--- Channels ---")

	// Create a channel
	numbers := []int{1, 2, 3, 4, 5}
	resultChan := make(chan int)

	// Start goroutine that sends on channel
	go sum(numbers, resultChan)

	// Receive from channel (blocks until value is ready)
	result := <-resultChan
	fmt.Printf("Sum of %v = %d\n", numbers, result)

	// Split work across multiple goroutines
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5, 6}

	chan1 := make(chan int)
	chan2 := make(chan int)

	go sum(numbers1, chan1)
	go sum(numbers2, chan2)

	// Receive from both channels
	result1 := <-chan1
	result2 := <-chan2
	fmt.Printf("Partial sums: %d + %d = %d\n", result1, result2, result1+result2)

	// Channel direction (send-only, receive-only)
	messages := make(chan string)
	go sendMessage(messages) // send-only in function
	msg := <-messages
	fmt.Printf("Received: %s\n", msg)

	// Closing channels
	jobs := make(chan int, 3)
	jobs <- 1
	jobs <- 2
	jobs <- 3
	close(jobs) // Signal no more values will be sent

	// Range over closed channel
	fmt.Print("Jobs: ")
	for job := range jobs {
		fmt.Printf("%d ", job)
	}
	fmt.Println()
	fmt.Println()
}

// sendMessage demonstrates send-only channel parameter
func sendMessage(ch chan<- string) {
	ch <- "Hello from goroutine"
}

// producer sends values on channel
func producer(ch chan int, count int) {
	for i := 1; i <= count; i++ {
		ch <- i
		time.Sleep(50 * time.Millisecond)
	}
	close(ch)
}

// demonstrateBufferedChannels shows buffered channel usage
func demonstrateBufferedChannels() {
	fmt.Println("--- Buffered Channels ---")

	// Unbuffered channel (capacity 0)
	// Sending blocks until receiver is ready
	unbuffered := make(chan int)
	go func() {
		unbuffered <- 42
	}()
	fmt.Printf("Received from unbuffered: %d\n", <-unbuffered)

	// Buffered channel (capacity > 0)
	// Can send without receiver up to capacity
	buffered := make(chan int, 3)
	buffered <- 1
	buffered <- 2
	buffered <- 3
	// buffered <- 4  // Would block because buffer is full

	fmt.Printf("Buffered channel contents: %d, %d, %d\n",
		<-buffered, <-buffered, <-buffered)

	// Producer-consumer with buffered channel
	productChan := make(chan int, 5)
	go producer(productChan, 10)

	fmt.Print("Consuming: ")
	for product := range productChan {
		fmt.Printf("%d ", product)
		time.Sleep(75 * time.Millisecond) // Slow consumer
	}
	fmt.Println()
	fmt.Println()
}

// demonstrateSelect shows select statement for multiple channels
func demonstrateSelect() {
	fmt.Println("--- Select Statement ---")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "message from channel 1"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "message from channel 2"
	}()

	// Select waits on multiple channel operations
	// Proceeds with whichever channel is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}

	// Select with timeout
	timeoutChan := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		timeoutChan <- "result"
	}()

	select {
	case res := <-timeoutChan:
		fmt.Println("Received:", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout: operation took too long")
	}

	// Select with default (non-blocking)
	messages := make(chan string, 1)
	messages <- "buffered"

	select {
	case msg := <-messages:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received")
	}

	// Try again with empty channel
	select {
	case msg := <-messages:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message available (non-blocking)")
	}

	fmt.Println()
}

// workerWithWG performs work and signals completion via WaitGroup
func workerWithWG(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion when function returns

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(id*50) * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

// demonstrateWaitGroup shows WaitGroup for goroutine synchronization
func demonstrateWaitGroup() {
	fmt.Println("--- WaitGroup ---")

	var wg sync.WaitGroup

	// Launch multiple goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment counter
		go workerWithWG(i, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("All workers completed")

	// WaitGroup with anonymous functions
	fmt.Println("\nConcurrent tasks:")
	var wg2 sync.WaitGroup
	tasks := []string{"task-A", "task-B", "task-C"}

	for _, task := range tasks {
		wg2.Add(1)
		go func(taskName string) {
			defer wg2.Done()
			fmt.Printf("Processing %s\n", taskName)
			time.Sleep(50 * time.Millisecond)
		}(task) // Pass task as parameter to avoid closure issues
	}

	wg2.Wait()
	fmt.Println("All tasks completed")
}
