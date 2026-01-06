package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Advanced Concurrency Patterns ===")
	fmt.Println()

	// Worker Pool Pattern
	demonstrateWorkerPool()

	// Fan-Out/Fan-In Pattern
	demonstrateFanOutFanIn()

	// Pipeline Pattern
	demonstratePipeline()
}

// Worker Pool Pattern
func demonstrateWorkerPool() {
	fmt.Println("--- Worker Pool Pattern ---")

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start workers
	numWorkers := 3
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}

	fmt.Println()
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(50 * time.Millisecond)
		results <- job * 2
	}
}

// Fan-Out/Fan-In Pattern
func demonstrateFanOutFanIn() {
	fmt.Println("--- Fan-Out/Fan-In Pattern ---")

	// Input channel
	input := make(chan int, 5)

	// Fan-out: Multiple workers processing input
	numWorkers := 3
	workers := make([]<-chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		workers[i] = fanOutWorker(input)
	}

	// Fan-in: Merge results from all workers
	results := fanIn(workers...)

	// Send input
	go func() {
		for i := 1; i <= 6; i++ {
			input <- i
		}
		close(input)
	}()

	// Receive results
	for result := range results {
		fmt.Printf("Fan-in result: %d\n", result)
	}

	fmt.Println()
}

func fanOutWorker(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			time.Sleep(30 * time.Millisecond)
			output <- num * num
		}
	}()
	return output
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Pipeline Pattern
func demonstratePipeline() {
	fmt.Println("--- Pipeline Pattern ---")

	// Stage 1: Generate numbers
	numbers := generate(1, 2, 3, 4, 5)

	// Stage 2: Square numbers
	squares := square(numbers)

	// Stage 3: Filter even numbers
	evens := filterEven(squares)

	// Consume results
	for num := range evens {
		fmt.Printf("Pipeline result: %d\n", num)
	}

	fmt.Println()
}

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()
	return out
}
