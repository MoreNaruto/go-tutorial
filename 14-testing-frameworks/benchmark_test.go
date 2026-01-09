package main

import (
	"testing"
)

// Basic benchmark - measures Add function performance
func BenchmarkAdd(b *testing.B) {
	calc := NewCalculator()

	// b.N is automatically adjusted by the testing framework
	// to get reliable timing results
	for i := 0; i < b.N; i++ {
		calc.Add(2, 3)
	}
}

// Benchmark with setup - use b.ResetTimer() to exclude setup time
func BenchmarkAddWithSetup(b *testing.B) {
	// Setup code (not measured)
	calc := NewCalculator()
	numbers := []int{1, 2, 3, 4, 5}

	// Reset timer to exclude setup time
	b.ResetTimer()

	// Only this loop is measured
	for i := 0; i < b.N; i++ {
		calc.Add(numbers[0], numbers[1])
	}
}

// Benchmark with memory allocation reporting
func BenchmarkSum(b *testing.B) {
	calc := NewCalculator()
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b.ResetTimer()
	b.ReportAllocs() // Report memory allocations

	for i := 0; i < b.N; i++ {
		calc.Sum(numbers)
	}
}

// Benchmark with different input sizes (sub-benchmarks)
func BenchmarkSumVariousSizes(b *testing.B) {
	calc := NewCalculator()

	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run(benchName(size), func(b *testing.B) {
			numbers := make([]int, size)
			for i := range numbers {
				numbers[i] = i
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				calc.Sum(numbers)
			}
		})
	}
}

// Helper function for benchmark names
func benchName(size int) string {
	if size < 1000 {
		return string(rune('0' + size/100))
	}
	return "large"
}

// Benchmark comparing two approaches
func BenchmarkMultiply(b *testing.B) {
	calc := NewCalculator()

	b.Run("small numbers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			calc.Multiply(2, 3)
		}
	})

	b.Run("large numbers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			calc.Multiply(1000000, 2000000)
		}
	})
}

// Benchmark with timer control
func BenchmarkDivideWithTimerControl(b *testing.B) {
	calc := NewCalculator()

	for i := 0; i < b.N; i++ {
		// Expensive setup (stop timer)
		b.StopTimer()
		a, b_val := 100, 2
		b.StartTimer()

		// Only measure the actual operation
		calc.Divide(a, b_val)
	}
}

// Benchmark table-driven tests
func BenchmarkArithmeticOperations(b *testing.B) {
	calc := NewCalculator()

	benchmarks := []struct {
		name string
		fn   func() int
	}{
		{"Add", func() int { return calc.Add(10, 20) }},
		{"Subtract", func() int { return calc.Subtract(20, 10) }},
		{"Multiply", func() int { return calc.Multiply(10, 20) }},
		{"Abs", func() int { return calc.Abs(-50) }},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.fn()
			}
		})
	}
}

// Benchmark for slice operations
func BenchmarkAverage(b *testing.B) {
	calc := NewCalculator()

	b.Run("small slice", func(b *testing.B) {
		numbers := []int{1, 2, 3, 4, 5}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			calc.Average(numbers)
		}
	})

	b.Run("medium slice", func(b *testing.B) {
		numbers := make([]int, 100)
		for i := range numbers {
			numbers[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			calc.Average(numbers)
		}
	})

	b.Run("large slice", func(b *testing.B) {
		numbers := make([]int, 10000)
		for i := range numbers {
			numbers[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			calc.Average(numbers)
		}
	})
}

// Benchmark with memory allocation tracking
func BenchmarkMaxAllocations(b *testing.B) {
	calc := NewCalculator()

	// This benchmark will show memory allocations
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		numbers := []int{1, 5, 3, 9, 2, 8, 4, 7, 6}
		calc.Max(numbers)
	}
}

// Benchmark with pre-allocated slice (better performance)
func BenchmarkMaxNoAllocations(b *testing.B) {
	calc := NewCalculator()
	numbers := []int{1, 5, 3, 9, 2, 8, 4, 7, 6}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		calc.Max(numbers)
	}
}

// Benchmark comparing implementations
func BenchmarkIsEven(b *testing.B) {
	calc := NewCalculator()

	b.Run("even number", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			calc.IsEven(42)
		}
	})

	b.Run("odd number", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			calc.IsEven(41)
		}
	})

	b.Run("zero", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			calc.IsEven(0)
		}
	})

	b.Run("negative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			calc.IsEven(-42)
		}
	})
}

// Parallel benchmark - runs test function in parallel
func BenchmarkAddParallel(b *testing.B) {
	calc := NewCalculator()

	// RunParallel runs the benchmark in parallel
	// Useful for testing concurrent access
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			calc.Add(2, 3)
		}
	})
}

// Parallel benchmark with sub-benchmarks
func BenchmarkCalculatorParallel(b *testing.B) {
	calc := NewCalculator()

	b.Run("Add", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				calc.Add(10, 20)
			}
		})
	})

	b.Run("Multiply", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				calc.Multiply(10, 20)
			}
		})
	})
}

// Benchmark with bytes per operation
func BenchmarkSumBytes(b *testing.B) {
	calc := NewCalculator()
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Calculate bytes per operation
	bytesPerOp := int64(len(numbers) * 8) // 8 bytes per int

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		calc.Sum(numbers)
		b.SetBytes(bytesPerOp)
	}
}

// Benchmark demonstrating performance optimization
// Compare two implementations
func BenchmarkPowerNaive(b *testing.B) {
	calc := NewCalculator()

	for i := 0; i < b.N; i++ {
		calc.Power(2, 10)
	}
}

// Example of what NOT to do in benchmarks
func BenchmarkBadExample(b *testing.B) {
	// DON'T: Create new instance in the loop
	// This measures allocation overhead, not the operation
	for i := 0; i < b.N; i++ {
		calc := NewCalculator()
		calc.Add(2, 3)
	}
}

// Correct way
func BenchmarkGoodExample(b *testing.B) {
	calc := NewCalculator() // Create once
	b.ResetTimer()          // Reset to exclude setup

	for i := 0; i < b.N; i++ {
		calc.Add(2, 3) // Only measure the operation
	}
}

// Benchmark with setup and teardown
func BenchmarkWithSetupTeardown(b *testing.B) {
	// Setup (not measured)
	calc := NewCalculator()
	testData := make([]int, 1000)
	for i := range testData {
		testData[i] = i
	}

	b.ResetTimer() // Start measuring here

	for i := 0; i < b.N; i++ {
		calc.Sum(testData)
	}

	b.StopTimer() // Stop measuring

	// Teardown (not measured)
	testData = nil
}

// Benchmark comparing error handling overhead
func BenchmarkErrorHandling(b *testing.B) {
	calc := NewCalculator()

	b.Run("no error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := calc.Divide(10, 2)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("with error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := calc.Divide(10, 0)
			if err == nil {
				b.Fatal("expected error")
			}
		}
	})
}

/*
To run benchmarks:

# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkAdd

# Run with memory stats
go test -bench=. -benchmem

# Run for longer (more accurate)
go test -bench=. -benchtime=10s

# Run multiple times
go test -bench=. -count=5

# Compare benchmarks
go test -bench=. > old.txt
# ... make changes ...
go test -bench=. > new.txt
benchstat old.txt new.txt  # Requires golang.org/x/perf/cmd/benchstat

# Run only benchmarks matching pattern
go test -bench=Add

# Run with CPU profiling
go test -bench=. -cpuprofile=cpu.prof

# Run with memory profiling
go test -bench=. -memprofile=mem.prof

# Parallel execution
go test -bench=. -cpu=1,2,4,8
*/
