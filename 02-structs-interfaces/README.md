# 02 - Structs and Interfaces

**Level:** Beginner

## What This Project Teaches

This project covers Go's type system fundamentals:
- Defining and using structs
- Methods with value and pointer receivers
- Interfaces and polymorphism
- Struct embedding (composition over inheritance)
- Anonymous structs

## Why This Structure?

Like the basics project, this uses a **flat structure** appropriate for focused learning:
- `main.go` - Demonstrates structs, methods, interfaces, and composition
- `structs_test.go` - Comprehensive tests for all types and methods

The code is organized by concept, making it easy to follow the progression from simple structs to interfaces and composition.

## Key Concepts Explained

### Structs

Structs are Go's way of creating custom types with named fields:

```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

// Create struct instances
person1 := Person{FirstName: "Alice", LastName: "Smith", Age: 30}
person2 := &Person{FirstName: "Bob", LastName: "Jones", Age: 25}
```

**Zero values:** Uninitialized struct fields get their type's zero value.

### Methods

Methods are functions with a receiver argument:

```go
// Value receiver - operates on a copy
func (p Person) FullName() string {
    return p.FirstName + " " + p.LastName
}

// Pointer receiver - can modify the original struct
func (p *Person) HaveBirthday() {
    p.Age++
}
```

**When to use pointer receivers:**
- When you need to modify the struct
- For large structs (avoid copying)
- For consistency (if any method needs a pointer, use pointers for all)

### Interfaces

Interfaces define behavior without implementation details:

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

**Implicit implementation:** Types implement interfaces automatically by having the required methods. No explicit declaration needed.

```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Circle now implements Shape interface automatically
```

**Polymorphism:**

```go
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

// Works with any type that implements Shape
printShapeInfo(Circle{Radius: 5})
printShapeInfo(Rectangle{Width: 4, Height: 6})
```

### Composition

Go uses **composition over inheritance** via struct embedding:

```go
type Employee struct {
    Person           // Embedded struct
    EmployeeID int
    Department string
}

employee := Employee{
    Person: Person{FirstName: "Alice", LastName: "Smith", Age: 28},
    EmployeeID: 12345,
    Department: "Engineering",
}

// Embedded fields and methods are promoted
fmt.Println(employee.FirstName)      // Access Person field directly
fmt.Println(employee.FullName())     // Call Person method directly
```

**Benefits of composition:**
- Clear relationships
- No inheritance complexity
- Flexible and explicit
- Multiple embeddings possible

## Running the Code

Run the program to see demonstrations:

```bash
go run main.go
```

You'll see examples of:
- Struct creation and field access
- Method calls with different receivers
- Interface polymorphism
- Struct embedding and field promotion

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

1. How to define and use structs
2. Difference between value and pointer receivers
3. How interfaces enable polymorphism in Go
4. Why Go favors composition over inheritance
5. How to write tests for types and methods
6. Anonymous structs for ad-hoc data structures

## Design Patterns Demonstrated

- **Interface segregation:** Small, focused interfaces (Shape has only 2 methods)
- **Composition:** Employee embeds Person rather than inheriting from it
- **Polymorphism:** printShapeInfo works with any Shape implementation
- **Encapsulation:** Methods operate on specific types

## Common Pitfalls Avoided

1. **Forgetting pointer receivers:** Methods that need to modify the struct must use pointer receivers
2. **Over-using interfaces:** Interfaces should be discovered from usage, not predefined
3. **Inheritance thinking:** Go doesn't have inheritance; embrace composition
4. **Unnecessary interface declarations:** Types implement interfaces implicitly

## Next Steps

After mastering structs and interfaces, move on to:
- **03-error-handling** - Learn Go's error patterns and how interfaces enable error handling
- **04-basic-concurrency** - Use interfaces for concurrent patterns
- **08-testing-strategies** - Advanced testing with interfaces and mocks

## Real-World Usage

These patterns appear everywhere in Go:
- **Standard library:** `io.Reader`, `io.Writer`, `http.Handler` are all interfaces
- **Database drivers:** All implement `database/sql/driver` interfaces
- **HTTP servers:** Handlers implement `http.Handler` interface
- **Testing:** Mock objects implement the same interfaces as real implementations
