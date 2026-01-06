package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple name", "Alice", "Hello, Alice!"},
		{"empty name", "", "Hello, !"},
		{"name with spaces", "John Doe", "Hello, John Doe!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := greet(tt.input)
			if result != tt.expected {
				t.Errorf("greet(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name            string
		a, b            int
		expectedSum     int
		expectedProduct int
	}{
		{"positive numbers", 5, 3, 8, 15},
		{"with zero", 5, 0, 5, 0},
		{"negative numbers", -2, 3, 1, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum, product := calculate(tt.a, tt.b)
			if sum != tt.expectedSum {
				t.Errorf("calculate(%d, %d) sum = %d, want %d", tt.a, tt.b, sum, tt.expectedSum)
			}
			if product != tt.expectedProduct {
				t.Errorf("calculate(%d, %d) product = %d, want %d", tt.a, tt.b, product, tt.expectedProduct)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"simple division", 10, 2, 5.0},
		{"decimal result", 7, 2, 3.5},
		{"divide by one", 42, 1, 42.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := divide(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("divide(%.2f, %.2f) = %.2f, want %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"multiple numbers", []int{1, 2, 3, 4, 5}, 15},
		{"single number", []int{42}, 42},
		{"empty slice", []int{}, 0},
		{"with negatives", []int{-5, 10, -3}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sum(tt.numbers...)
			if result != tt.expected {
				t.Errorf("sum(%v) = %d, want %d", tt.numbers, result, tt.expected)
			}
		})
	}
}

// Benchmark example
func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		sum(numbers...)
	}
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		greet("Benchmark")
	}
}
