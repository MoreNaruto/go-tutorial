package main

import (
	"math"
	"testing"
)

func TestPersonFullName(t *testing.T) {
	tests := []struct {
		name     string
		person   Person
		expected string
	}{
		{
			name:     "normal name",
			person:   Person{FirstName: "John", LastName: "Doe", Age: 30},
			expected: "John Doe",
		},
		{
			name:     "empty first name",
			person:   Person{FirstName: "", LastName: "Smith", Age: 25},
			expected: " Smith",
		},
		{
			name:     "single name",
			person:   Person{FirstName: "Madonna", LastName: "", Age: 40},
			expected: "Madonna ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.person.FullName()
			if result != tt.expected {
				t.Errorf("FullName() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestPersonIsAdult(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		expected bool
	}{
		{"adult", 18, true},
		{"older adult", 30, true},
		{"minor", 17, false},
		{"young minor", 5, false},
		{"edge case zero", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			person := Person{Age: tt.age}
			result := person.IsAdult()
			if result != tt.expected {
				t.Errorf("IsAdult() for age %d = %v, want %v", tt.age, result, tt.expected)
			}
		})
	}
}

func TestPersonHaveBirthday(t *testing.T) {
	person := Person{FirstName: "Test", LastName: "User", Age: 20}

	person.HaveBirthday()
	if person.Age != 21 {
		t.Errorf("After one birthday, age = %d, want 21", person.Age)
	}

	person.HaveBirthday()
	person.HaveBirthday()
	if person.Age != 23 {
		t.Errorf("After three birthdays, age = %d, want 23", person.Age)
	}
}

func TestCircleArea(t *testing.T) {
	tests := []struct {
		name     string
		radius   float64
		expected float64
	}{
		{"radius 1", 1, math.Pi},
		{"radius 2", 2, 4 * math.Pi},
		{"radius 5", 5, 25 * math.Pi},
		{"radius 0", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			circle := Circle{Radius: tt.radius}
			result := circle.Area()
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("Circle.Area() = %f, want %f", result, tt.expected)
			}
		})
	}
}

func TestCirclePerimeter(t *testing.T) {
	circle := Circle{Radius: 5}
	expected := 2 * math.Pi * 5
	result := circle.Perimeter()

	if math.Abs(result-expected) > 0.0001 {
		t.Errorf("Circle.Perimeter() = %f, want %f", result, expected)
	}
}

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		name     string
		width    float64
		height   float64
		expected float64
	}{
		{"square", 5, 5, 25},
		{"rectangle", 4, 6, 24},
		{"zero width", 0, 10, 0},
		{"zero height", 10, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rect := Rectangle{Width: tt.width, Height: tt.height}
			result := rect.Area()
			if result != tt.expected {
				t.Errorf("Rectangle.Area() = %f, want %f", result, tt.expected)
			}
		})
	}
}

func TestRectanglePerimeter(t *testing.T) {
	rect := Rectangle{Width: 4, Height: 6}
	expected := 2 * (4 + 6)
	result := rect.Perimeter()

	if result != float64(expected) {
		t.Errorf("Rectangle.Perimeter() = %f, want %f", result, float64(expected))
	}
}

func TestShapeInterface(t *testing.T) {
	// Test that both Circle and Rectangle implement Shape interface
	var _ Shape = Circle{}
	var _ Shape = Rectangle{}

	// Test polymorphism
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 5},
	}

	if len(shapes) != 2 {
		t.Errorf("Expected 2 shapes, got %d", len(shapes))
	}

	// All shapes should have Area and Perimeter
	for i, shape := range shapes {
		area := shape.Area()
		perimeter := shape.Perimeter()

		if area <= 0 {
			t.Errorf("Shape %d has invalid area: %f", i, area)
		}
		if perimeter <= 0 {
			t.Errorf("Shape %d has invalid perimeter: %f", i, perimeter)
		}
	}
}

func TestEmployeeComposition(t *testing.T) {
	employee := Employee{
		Person: Person{
			FirstName: "Alice",
			LastName:  "Smith",
			Age:       25,
		},
		EmployeeID: 1001,
		Department: "IT",
	}

	// Test field promotion from embedded Person
	if employee.FirstName != "Alice" {
		t.Errorf("FirstName = %s, want Alice", employee.FirstName)
	}

	if employee.Age != 25 {
		t.Errorf("Age = %d, want 25", employee.Age)
	}

	// Test method promotion from embedded Person
	fullName := employee.FullName()
	if fullName != "Alice Smith" {
		t.Errorf("FullName() = %s, want Alice Smith", fullName)
	}

	// Test that IsAdult works on embedded Person
	if !employee.IsAdult() {
		t.Error("Expected employee to be adult")
	}

	// Test pointer receiver method on embedded Person
	employee.HaveBirthday()
	if employee.Age != 26 {
		t.Errorf("After birthday, Age = %d, want 26", employee.Age)
	}
}

// Benchmark tests
func BenchmarkCircleArea(b *testing.B) {
	circle := Circle{Radius: 10}
	for i := 0; i < b.N; i++ {
		circle.Area()
	}
}

func BenchmarkRectangleArea(b *testing.B) {
	rect := Rectangle{Width: 10, Height: 20}
	for i := 0; i < b.N; i++ {
		rect.Area()
	}
}
