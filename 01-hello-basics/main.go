package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Go Basics Tutorial ===")
	fmt.Println()

	// Variables and basic types
	demonstrateVariables()

	// Functions
	demonstrateFunctions()

	// Control flow
	demonstrateControlFlow()

	// Collections
	demonstrateCollections()
}

// demonstrateVariables shows different ways to declare and use variables
func demonstrateVariables() {
	fmt.Println("--- Variables ---")

	// Explicit type declaration
	var name string = "Alice"
	var age int = 30
	var height float64 = 5.6

	// Type inference
	var city = "San Francisco"

	// Short declaration (most common in Go)
	country := "USA"
	isActive := true

	// Multiple declarations
	var (
		firstName = "John"
		lastName  = "Doe"
	)

	// Zero values - Go initializes all variables with zero values
	var count int        // 0
	var empty string     // ""
	var enabled bool     // false
	var pointer *int     // nil

	fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, age, height)
	fmt.Printf("Location: %s, %s\n", city, country)
	fmt.Printf("Active: %v\n", isActive)
	fmt.Printf("Full name: %s %s\n", firstName, lastName)
	fmt.Printf("Zero values - count: %d, empty: '%s', enabled: %v, pointer: %v\n\n",
		count, empty, enabled, pointer)
}

// demonstrateFunctions shows various function patterns
func demonstrateFunctions() {
	fmt.Println("--- Functions ---")

	// Simple function call
	greeting := greet("Bob")
	fmt.Println(greeting)

	// Multiple return values
	sumResult, product := calculate(5, 3)
	fmt.Printf("Sum: %d, Product: %d\n", sumResult, product)

	// Named return values
	result := divide(10, 2)
	fmt.Printf("Division result: %.2f\n", result)

	// Variadic functions
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Total sum: %d\n\n", total)
}

// greet returns a greeting message
func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// calculate demonstrates multiple return values
func calculate(a, b int) (int, int) {
	return a + b, a * b
}

// divide demonstrates named return values
func divide(a, b float64) (result float64) {
	result = a / b
	return // naked return uses named result variable
}

// sum demonstrates variadic functions (variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// demonstrateControlFlow shows if/else, switch, and loops
func demonstrateControlFlow() {
	fmt.Println("--- Control Flow ---")

	// If/else
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C or below")
	}

	// If with initialization statement
	if val := score * 2; val > 150 {
		fmt.Printf("Double score %d is high\n", val)
	}

	// Switch statement
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend")
	default:
		fmt.Println("Unknown day")
	}

	// Switch without condition (like if/else chain)
	number := 42
	switch {
	case number < 0:
		fmt.Println("Negative")
	case number == 0:
		fmt.Println("Zero")
	case number > 0:
		fmt.Println("Positive")
	}

	// For loop (the only loop in Go)
	fmt.Print("Counting: ")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// While-style loop
	counter := 0
	for counter < 3 {
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}

	// Infinite loop with break
	accumulator := 0
	for {
		accumulator++
		if accumulator > 5 {
			break
		}
	}
	fmt.Printf("Accumulator reached: %d\n\n", accumulator)
}

// demonstrateCollections shows arrays, slices, and maps
func demonstrateCollections() {
	fmt.Println("--- Collections ---")

	// Arrays - fixed size
	var numbers [3]int = [3]int{1, 2, 3}
	fmt.Printf("Array: %v\n", numbers)

	// Slices - dynamic size (most common)
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Printf("Slice: %v\n", fruits)

	// Slice operations
	fruits = append(fruits, "date")           // Add element
	fmt.Printf("After append: %v\n", fruits)

	subset := fruits[1:3]                     // Slice from index 1 to 3 (exclusive)
	fmt.Printf("Subset [1:3]: %v\n", subset)

	// Make slice with capacity
	colors := make([]string, 0, 5)            // length 0, capacity 5
	colors = append(colors, "red", "green", "blue")
	fmt.Printf("Colors: %v (len: %d, cap: %d)\n", colors, len(colors), cap(colors))

	// Maps - key-value pairs
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Carol": 28,
	}
	fmt.Printf("Ages map: %v\n", ages)

	// Access map value
	aliceAge := ages["Alice"]
	fmt.Printf("Alice's age: %d\n", aliceAge)

	// Check if key exists
	if age, exists := ages["David"]; exists {
		fmt.Printf("David's age: %d\n", age)
	} else {
		fmt.Println("David not found in map")
	}

	// Iterate over map
	fmt.Println("All ages:")
	for name, age := range ages {
		fmt.Printf("  %s: %d\n", name, age)
	}

	// Delete from map
	delete(ages, "Bob")
	fmt.Printf("After deleting Bob: %v\n", ages)

	// String manipulation
	text := "Hello, Go!"
	fmt.Printf("\nOriginal: %s\n", text)
	fmt.Printf("Upper: %s\n", strings.ToUpper(text))
	fmt.Printf("Contains 'Go': %v\n", strings.Contains(text, "Go"))
	fmt.Printf("Split: %v\n", strings.Split(text, ", "))
}
