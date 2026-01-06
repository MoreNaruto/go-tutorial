package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Structs and Interfaces Tutorial ===")
	fmt.Println()

	// Demonstrate basic structs
	demonstrateStructs()

	// Demonstrate methods
	demonstrateMethods()

	// Demonstrate interfaces
	demonstrateInterfaces()

	// Demonstrate composition
	demonstrateComposition()
}

// Person represents a person with basic information
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// demonstrateStructs shows struct creation and usage
func demonstrateStructs() {
	fmt.Println("--- Structs ---")

	// Create struct with field names (preferred)
	person1 := Person{
		FirstName: "Alice",
		LastName:  "Smith",
		Age:       30,
	}
	fmt.Printf("Person 1: %+v\n", person1)

	// Create struct with positional arguments (not recommended)
	person2 := Person{"Bob", "Jones", 25}
	fmt.Printf("Person 2: %s %s, Age: %d\n", person2.FirstName, person2.LastName, person2.Age)

	// Zero value struct
	var person3 Person
	fmt.Printf("Person 3 (zero value): %+v\n", person3)

	// Anonymous struct (useful for one-off data structures)
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("Config: %+v\n", config)

	// Pointer to struct
	person4 := &Person{FirstName: "Carol", LastName: "White", Age: 35}
	fmt.Printf("Person 4 (pointer): %+v\n", person4)

	// Modifying struct fields
	person4.Age = 36
	fmt.Printf("Person 4 after birthday: %+v\n\n", person4)
}

// FullName returns the person's full name
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// HaveBirthday increments the person's age (pointer receiver to modify the struct)
func (p *Person) HaveBirthday() {
	p.Age++
}

// IsAdult checks if the person is an adult
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// demonstrateMethods shows methods on structs
func demonstrateMethods() {
	fmt.Println("--- Methods ---")

	person := Person{
		FirstName: "David",
		LastName:  "Brown",
		Age:       17,
	}

	// Value receiver method
	fmt.Printf("Full name: %s\n", person.FullName())
	fmt.Printf("Is adult: %v\n", person.IsAdult())

	// Pointer receiver method (modifies the struct)
	fmt.Printf("Age before birthday: %d\n", person.Age)
	person.HaveBirthday()
	fmt.Printf("Age after birthday: %d\n", person.Age)
	fmt.Printf("Is adult now: %v\n\n", person.IsAdult())
}

// Shape is an interface that defines behavior for geometric shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle represents a circle shape
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the perimeter of a circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle represents a rectangle shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// printShapeInfo accepts any type that implements the Shape interface
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// demonstrateInterfaces shows interface usage and polymorphism
func demonstrateInterfaces() {
	fmt.Println("--- Interfaces ---")

	// Create shapes
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Both types implement the Shape interface
	fmt.Println("Circle:")
	printShapeInfo(circle)

	fmt.Println("Rectangle:")
	printShapeInfo(rectangle)

	// Interface slice - polymorphic collection
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 5, Height: 10},
		Circle{Radius: 7},
	}

	fmt.Println("\nAll shapes:")
	totalArea := 0.0
	for i, shape := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		printShapeInfo(shape)
		totalArea += shape.Area()
	}
	fmt.Printf("Total area of all shapes: %.2f\n\n", totalArea)
}

// Address represents a physical address
type Address struct {
	Street  string
	City    string
	ZipCode string
}

// Employee embeds Person (composition, not inheritance)
type Employee struct {
	Person           // Embedded struct - promotes fields and methods
	EmployeeID int
	Department string
	Address    Address // Nested struct
}

// demonstrateComposition shows struct embedding and composition
func demonstrateComposition() {
	fmt.Println("--- Composition ---")

	// Create employee with embedded Person
	employee := Employee{
		Person: Person{
			FirstName: "Emma",
			LastName:  "Wilson",
			Age:       28,
		},
		EmployeeID: 12345,
		Department: "Engineering",
		Address: Address{
			Street:  "123 Main St",
			City:    "San Francisco",
			ZipCode: "94102",
		},
	}

	// Access embedded struct fields directly (promotion)
	fmt.Printf("Employee: %s %s\n", employee.FirstName, employee.LastName)
	fmt.Printf("Employee ID: %d\n", employee.EmployeeID)
	fmt.Printf("Department: %s\n", employee.Department)

	// Access embedded struct methods directly
	fmt.Printf("Full name (from Person method): %s\n", employee.FullName())
	fmt.Printf("Is adult: %v\n", employee.IsAdult())

	// Access nested struct fields
	fmt.Printf("Address: %s, %s %s\n", employee.Address.Street, employee.Address.City, employee.Address.ZipCode)

	// Pointer receiver method still works on embedded struct
	employee.HaveBirthday()
	fmt.Printf("Age after birthday: %d\n", employee.Age)
}
