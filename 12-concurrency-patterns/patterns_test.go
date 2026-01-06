package main

import (
	"sync"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start worker
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(1, jobs, results, &wg)

	// Send jobs
	jobs <- 5
	jobs <- 10
	close(jobs)

	// Wait and close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Verify results
	count := 0
	for range results {
		count++
	}

	if count != 2 {
		t.Errorf("Expected 2 results, got %d", count)
	}
}

func TestGenerate(t *testing.T) {
	ch := generate(1, 2, 3)

	nums := []int{}
	for num := range ch {
		nums = append(nums, num)
	}

	if len(nums) != 3 {
		t.Errorf("Expected 3 numbers, got %d", len(nums))
	}
}

func TestSquare(t *testing.T) {
	input := generate(2, 3, 4)
	output := square(input)

	results := []int{}
	for num := range output {
		results = append(results, num)
	}

	expected := []int{4, 9, 16}
	for i, v := range results {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
	}
}

func TestFilterEven(t *testing.T) {
	input := generate(1, 2, 3, 4)
	output := filterEven(input)

	results := []int{}
	for num := range output {
		results = append(results, num)
	}

	if len(results) != 2 {
		t.Errorf("Expected 2 even numbers, got %d", len(results))
	}
}

func TestPipeline(t *testing.T) {
	// Build pipeline
	numbers := generate(1, 2, 3, 4, 5)
	squares := square(numbers)
	evens := filterEven(squares)

	// Collect results
	results := []int{}
	for num := range evens {
		results = append(results, num)
	}

	// Expected: 4 (2*2) and 16 (4*4)
	if len(results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(results))
	}
}

func TestFanIn(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()

	merged := fanIn(ch1, ch2)

	results := []int{}
	for num := range merged {
		results = append(results, num)
	}

	if len(results) != 4 {
		t.Errorf("Expected 4 results, got %d", len(results))
	}
}

func BenchmarkWorkerPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jobs := make(chan int, 100)
		results := make(chan int, 100)

		var wg sync.WaitGroup
		for w := 1; w <= 5; w++ {
			wg.Add(1)
			go worker(w, jobs, results, &wg)
		}

		for j := 0; j < 100; j++ {
			jobs <- j
		}
		close(jobs)

		go func() {
			wg.Wait()
			close(results)
		}()

		for range results {
		}
	}
}

func BenchmarkPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		squares := square(numbers)
		evens := filterEven(squares)

		for range evens {
		}
	}
}
