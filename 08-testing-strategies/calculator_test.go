package main

import (
	"errors"
	"fmt"
	"testing"
)

// Table-driven tests
func TestCalculator_Add(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -5, -3, -8},
		{"mixed numbers", 10, -5, 5},
		{"zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Divide(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name      string
		a, b      float64
		want      float64
		wantError bool
	}{
		{"valid division", 10, 2, 5, false},
		{"division by zero", 10, 0, 0, true},
		{"negative result", -10, 2, -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Divide(tt.a, tt.b)

			if tt.wantError {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.want {
					t.Errorf("Divide(%.2f, %.2f) = %.2f, want %.2f", tt.a, tt.b, result, tt.want)
				}
			}
		})
	}
}

// Mock implementation of DataStore
type MockDataStore struct {
	data map[string]string
	err  error
}

func NewMockDataStore() *MockDataStore {
	return &MockDataStore{
		data: make(map[string]string),
	}
}

func (m *MockDataStore) Get(key string) (string, error) {
	if m.err != nil {
		return "", m.err
	}
	if val, ok := m.data[key]; ok {
		return val, nil
	}
	return "", errors.New("key not found")
}

func (m *MockDataStore) Set(key string, value string) error {
	if m.err != nil {
		return m.err
	}
	m.data[key] = value
	return nil
}

// Test with mock
func TestUserService_GetUsername(t *testing.T) {
	mock := NewMockDataStore()
	mock.data["user:123"] = "alice"

	service := NewUserService(mock)

	username, err := service.GetUsername("123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if username != "alice" {
		t.Errorf("Expected username 'alice', got '%s'", username)
	}
}

func TestUserService_GetUsername_NotFound(t *testing.T) {
	mock := NewMockDataStore()
	service := NewUserService(mock)

	_, err := service.GetUsername("999")
	if err == nil {
		t.Error("Expected error for non-existent user, got nil")
	}
}

func TestUserService_SaveUsername(t *testing.T) {
	mock := NewMockDataStore()
	service := NewUserService(mock)

	err := service.SaveUsername("456", "bob")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Verify it was saved
	if mock.data["user:456"] != "bob" {
		t.Errorf("Expected username 'bob', got '%s'", mock.data["user:456"])
	}
}

// Subtests
func TestUserService(t *testing.T) {
	t.Run("GetUsername", func(t *testing.T) {
		mock := NewMockDataStore()
		mock.data["user:1"] = "test"
		service := NewUserService(mock)

		username, _ := service.GetUsername("1")
		if username != "test" {
			t.Errorf("Expected 'test', got '%s'", username)
		}
	})

	t.Run("SaveUsername", func(t *testing.T) {
		mock := NewMockDataStore()
		service := NewUserService(mock)

		service.SaveUsername("2", "user2")
		if mock.data["user:2"] != "user2" {
			t.Error("Username not saved correctly")
		}
	})
}

// Test helpers
func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestWithHelper(t *testing.T) {
	calc := NewCalculator()
	result := calc.Add(2, 3)
	assertEqual(t, result, 5)
}

// Benchmark
func BenchmarkCalculator_Add(b *testing.B) {
	calc := NewCalculator()
	for i := 0; i < b.N; i++ {
		calc.Add(5, 10)
	}
}

// Example test (appears in documentation)
func ExampleCalculator_Add() {
	calc := NewCalculator()
	result := calc.Add(2, 3)
	fmt.Println(result)
	// Output: 5
}
