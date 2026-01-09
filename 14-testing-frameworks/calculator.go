package main

import (
	"errors"
	"math"
)

// Calculator provides basic arithmetic operations
type Calculator struct{}

// NewCalculator creates a new Calculator instance
func NewCalculator() *Calculator {
	return &Calculator{}
}

// Add returns the sum of two numbers
func (c *Calculator) Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of two numbers
func (c *Calculator) Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two numbers
func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two numbers
// Returns error if dividing by zero
func (c *Calculator) Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return float64(a) / float64(b), nil
}

// Power returns a raised to the power of b
func (c *Calculator) Power(a, b int) float64 {
	return math.Pow(float64(a), float64(b))
}

// IsEven checks if a number is even
func (c *Calculator) IsEven(n int) bool {
	return n%2 == 0
}

// Abs returns the absolute value of a number
func (c *Calculator) Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Sum returns the sum of a slice of numbers
func (c *Calculator) Sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// Average returns the average of a slice of numbers
// Returns error if slice is empty
func (c *Calculator) Average(numbers []int) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("cannot calculate average of empty slice")
	}
	sum := c.Sum(numbers)
	return float64(sum) / float64(len(numbers)), nil
}

// Max returns the maximum value in a slice
// Returns error if slice is empty
func (c *Calculator) Max(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("cannot find max of empty slice")
	}
	max := numbers[0]
	for _, n := range numbers[1:] {
		if n > max {
			max = n
		}
	}
	return max, nil
}

// Min returns the minimum value in a slice
// Returns error if slice is empty
func (c *Calculator) Min(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("cannot find min of empty slice")
	}
	min := numbers[0]
	for _, n := range numbers[1:] {
		if n < min {
			min = n
		}
	}
	return min, nil
}
