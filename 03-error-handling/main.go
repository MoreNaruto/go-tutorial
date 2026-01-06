package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== Error Handling Tutorial ===")
	fmt.Println()

	// Demonstrate basic error handling
	demonstrateBasicErrors()

	// Demonstrate custom errors
	demonstrateCustomErrors()

	// Demonstrate error wrapping
	demonstrateErrorWrapping()

	// Demonstrate error inspection
	demonstrateErrorInspection()
}

// divide performs division and returns an error for division by zero
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// parseAge parses and validates an age string
func parseAge(ageStr string) (int, error) {
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return 0, fmt.Errorf("invalid age format: %w", err)
	}

	if age < 0 {
		return 0, errors.New("age cannot be negative")
	}

	if age > 150 {
		return 0, errors.New("age is unrealistic")
	}

	return age, nil
}

// demonstrateBasicErrors shows basic error handling patterns
func demonstrateBasicErrors() {
	fmt.Println("--- Basic Error Handling ---")

	// Successful operation
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	// Operation with error
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error dividing 10 by 0: %v\n", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result)
	}

	// Parse valid age
	age, err := parseAge("25")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Parsed age: %d\n", age)
	}

	// Parse invalid format
	_, err = parseAge("not-a-number")
	if err != nil {
		fmt.Printf("Error parsing age: %v\n", err)
	}

	// Parse negative age
	_, err = parseAge("-5")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println()
}

// ValidationError represents a validation error with context
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s (value: %v)",
		e.Field, e.Message, e.Value)
}

// User represents a user with validated fields
type User struct {
	Username string
	Email    string
	Age      int
}

// NewUser creates a new user with validation
func NewUser(username, email string, age int) (*User, error) {
	if username == "" {
		return nil, &ValidationError{
			Field:   "username",
			Value:   username,
			Message: "cannot be empty",
		}
	}

	if len(username) < 3 {
		return nil, &ValidationError{
			Field:   "username",
			Value:   username,
			Message: "must be at least 3 characters",
		}
	}

	if email == "" {
		return nil, &ValidationError{
			Field:   "email",
			Value:   email,
			Message: "cannot be empty",
		}
	}

	if age < 13 {
		return nil, &ValidationError{
			Field:   "age",
			Value:   age,
			Message: "must be at least 13",
		}
	}

	return &User{
		Username: username,
		Email:    email,
		Age:      age,
	}, nil
}

// demonstrateCustomErrors shows custom error types
func demonstrateCustomErrors() {
	fmt.Println("--- Custom Errors ---")

	// Valid user
	user, err := NewUser("alice", "alice@example.com", 25)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Created user: %s (%s), age %d\n", user.Username, user.Email, user.Age)
	}

	// Invalid username (empty)
	_, err = NewUser("", "bob@example.com", 30)
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		// Type assertion to access custom error fields
		if valErr, ok := err.(*ValidationError); ok {
			fmt.Printf("  Failed field: %s\n", valErr.Field)
		}
	}

	// Invalid username (too short)
	_, err = NewUser("ab", "charlie@example.com", 20)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Invalid age
	_, err = NewUser("david", "david@example.com", 10)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println()
}

// DatabaseError represents a database operation error
type DatabaseError struct {
	Operation string
	Err       error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("database error during %s: %v", e.Operation, e.Err)
}

func (e *DatabaseError) Unwrap() error {
	return e.Err
}

// fetchUser simulates fetching a user from database
func fetchUser(id int) (*User, error) {
	if id <= 0 {
		return nil, &ValidationError{
			Field:   "id",
			Value:   id,
			Message: "must be positive",
		}
	}

	if id == 999 {
		// Simulate database error
		return nil, &DatabaseError{
			Operation: "SELECT",
			Err:       errors.New("connection timeout"),
		}
	}

	// Simulate user not found
	if id == 404 {
		return nil, fmt.Errorf("user with id %d not found", id)
	}

	// Return mock user
	return &User{
		Username: fmt.Sprintf("user_%d", id),
		Email:    fmt.Sprintf("user_%d@example.com", id),
		Age:      25,
	}, nil
}

// getUserInfo wraps fetchUser and adds context
func getUserInfo(id int) (*User, error) {
	user, err := fetchUser(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info for id %d: %w", id, err)
	}
	return user, nil
}

// demonstrateErrorWrapping shows error wrapping with %w
func demonstrateErrorWrapping() {
	fmt.Println("--- Error Wrapping ---")

	// Successful fetch
	user, err := getUserInfo(123)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Fetched user: %s\n", user.Username)
	}

	// Wrapped validation error
	_, err = getUserInfo(-1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Wrapped database error
	_, err = getUserInfo(999)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Not found error
	_, err = getUserInfo(404)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println()
}

// demonstrateErrorInspection shows error inspection with errors.Is and errors.As
func demonstrateErrorInspection() {
	fmt.Println("--- Error Inspection ---")

	// Define sentinel errors
	var ErrNotFound = errors.New("not found")
	var ErrUnauthorized = errors.New("unauthorized")

	// Test errors.Is
	err := fmt.Errorf("failed to fetch: %w", ErrNotFound)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Error is ErrNotFound")
	}

	if !errors.Is(err, ErrUnauthorized) {
		fmt.Println("Error is not ErrUnauthorized")
	}

	// Test errors.As with custom errors
	_, err = getUserInfo(-5)
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		fmt.Printf("Validation error detected: field=%s, message=%s\n",
			valErr.Field, valErr.Message)
	}

	// Test errors.As with database error
	_, err = getUserInfo(999)
	var dbErr *DatabaseError
	if errors.As(err, &dbErr) {
		fmt.Printf("Database error detected: operation=%s, underlying=%v\n",
			dbErr.Operation, dbErr.Err)
	}

	// Demonstrate Unwrap
	if dbErr != nil {
		unwrapped := errors.Unwrap(dbErr)
		fmt.Printf("Unwrapped error: %v\n", unwrapped)
	}
}
