package main

import (
	"sync"
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	resultChan := make(chan int)

	go sum(numbers, resultChan)

	result := <-resultChan
	expected := 15

	if result != expected {
		t.Errorf("sum(%v) = %d, want %d", numbers, result, expected)
	}
}

func TestConcurrentSum(t *testing.T) {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5, 6}

	chan1 := make(chan int)
	chan2 := make(chan int)

	go sum(numbers1, chan1)
	go sum(numbers2, chan2)

	result1 := <-chan1
	result2 := <-chan2
	total := result1 + result2

	expected := 21
	if total != expected {
		t.Errorf("concurrent sum = %d, want %d", total, expected)
	}
}

func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 3)

	// Should not block - buffer has capacity
	ch <- 1
	ch <- 2
	ch <- 3

	// Verify values
	if v := <-ch; v != 1 {
		t.Errorf("Expected 1, got %d", v)
	}
	if v := <-ch; v != 2 {
		t.Errorf("Expected 2, got %d", v)
	}
	if v := <-ch; v != 3 {
		t.Errorf("Expected 3, got %d", v)
	}
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	count := 0
	sum := 0
	for val := range ch {
		count++
		sum += val
	}

	if count != 3 {
		t.Errorf("Expected 3 values, got %d", count)
	}
	if sum != 6 {
		t.Errorf("Expected sum of 6, got %d", sum)
	}

	// Reading from closed channel returns zero value
	val, ok := <-ch
	if ok {
		t.Error("Expected channel to be closed")
	}
	if val != 0 {
		t.Errorf("Expected zero value, got %d", val)
	}
}

func TestSelectWithTimeout(t *testing.T) {
	ch := make(chan string)

	// Don't send anything - should timeout
	select {
	case <-ch:
		t.Error("Should not receive from channel")
	case <-time.After(50 * time.Millisecond):
		// Expected timeout
	}
}

func TestSelectDefault(t *testing.T) {
	ch := make(chan int)

	// Non-blocking receive with default
	select {
	case val := <-ch:
		t.Errorf("Should not receive value, got %d", val)
	default:
		// Expected: no value available
	}
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	counter := 0
	var mu sync.Mutex // Protect counter from race condition

	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	if counter != numWorkers {
		t.Errorf("Expected counter to be %d, got %d", numWorkers, counter)
	}
}

func TestConcurrentWorkers(t *testing.T) {
	var wg sync.WaitGroup
	results := make(chan int, 10)

	// Launch workers
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			results <- id * 2
		}(i)
	}

	// Close results channel after all workers complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	sum := 0
	for result := range results {
		sum += result
	}

	expected := (2 + 4 + 6 + 8 + 10)
	if sum != expected {
		t.Errorf("Sum of results = %d, want %d", sum, expected)
	}
}

func TestProducer(t *testing.T) {
	ch := make(chan int, 5)
	count := 10

	go producer(ch, count)

	received := 0
	for range ch {
		received++
	}

	if received != count {
		t.Errorf("Expected to receive %d values, got %d", count, received)
	}
}

// Benchmark goroutine creation
func BenchmarkGoroutineCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
		wg.Wait()
	}
}

// Benchmark channel send/receive
func BenchmarkChannelSendReceive(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			<-ch
		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- i
	}
}

// Benchmark buffered channel
func BenchmarkBufferedChannel(b *testing.B) {
	ch := make(chan int, 100)
	go func() {
		for i := 0; i < b.N; i++ {
			<-ch
		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- i
	}
}
