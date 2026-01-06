# 01 - Hello Basics

**Level:** Beginner

## What This Project Teaches

This project introduces fundamental Go concepts:
- Variable declaration and types
- Functions and multiple return values
- Control flow (if/else, switch, for loops)
- Collections (arrays, slices, maps)
- Basic string manipulation

## Why This Structure?

This is a **flat, simple structure** appropriate for learning basics:
- `main.go` - Main program demonstrating all concepts
- `basics_test.go` - Unit tests for testable functions

No complex folder structure is needed. The focus is on learning Go syntax and core features.

## Key Concepts Explained

### Variables and Types

Go is statically typed with type inference:

```go
var name string = "Alice"  // Explicit type
var age = 30               // Type inference
city := "NYC"              // Short declaration (most common)
```

**Zero values:** All variables have default values:
- `int`, `float64`: `0`
- `string`: `""`
- `bool`: `false`
- Pointers, slices, maps: `nil`

### Functions

Go functions are first-class values and can return multiple values:

```go
func calculate(a, b int) (int, int) {
    return a + b, a * b
}
```

**Variadic functions** accept variable number of arguments:

```go
func sum(numbers ...int) int {
    // numbers is a slice inside the function
}
```

### Control Flow

**For loops** are the only loop construct in Go:

```go
// Standard loop
for i := 0; i < 10; i++ { }

// While-style loop
for condition { }

// Infinite loop
for { }
```

**Switch statements** don't fall through by default (no `break` needed):

```go
switch day {
case "Monday", "Tuesday":
    // Handle weekday
default:
    // Handle other
}
```

### Collections

**Slices** are dynamic arrays (most commonly used):

```go
fruits := []string{"apple", "banana"}
fruits = append(fruits, "cherry")  // Add elements
subset := fruits[1:3]              // Slice operation
```

**Maps** are key-value data structures:

```go
ages := map[string]int{"Alice": 30, "Bob": 25}
age, exists := ages["Alice"]  // Check if key exists
delete(ages, "Bob")           // Remove key
```

## Running the Code

Run the program to see all demonstrations:

```bash
go run main.go
```

Expected output includes demonstrations of:
- Variable declarations and zero values
- Function calls with multiple return values
- Control flow examples
- Collection operations

## Running Tests

Execute all tests:

```bash
go test ./...
```

Run with verbose output:

```bash
go test -v
```

Run with coverage:

```bash
go test -cover
```

Run benchmarks:

```bash
go test -bench=.
```

## What You'll Learn

After working through this project, you'll understand:

1. How to declare and use variables in Go
2. Different ways to write functions and return values
3. Go's approach to control flow
4. Working with slices and maps
5. How to write table-driven tests
6. Basic benchmarking

## Next Steps

After mastering these basics, move on to:
- **02-structs-interfaces** - Learn about Go's type system
- **03-error-handling** - Understand Go's error patterns
- **04-basic-concurrency** - Explore goroutines and channels

## Common Patterns Demonstrated

- **Table-driven tests:** The idiomatic Go way to write comprehensive tests
- **Short variable declarations:** Using `:=` for concise code
- **Range loops:** Iterating over slices and maps
- **Zero values:** Understanding Go's default initialization
- **Multiple return values:** Particularly useful for error handling (covered in next projects)
