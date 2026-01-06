package main

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		want      float64
		wantError bool
	}{
		{"normal division", 10, 2, 5, false},
		{"division by zero", 10, 0, 0, true},
		{"negative numbers", -10, 2, -5, false},
		{"zero numerator", 0, 5, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := divide(tt.a, tt.b)

			if tt.wantError {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.want {
					t.Errorf("divide(%f, %f) = %f, want %f", tt.a, tt.b, result, tt.want)
				}
			}
		})
	}
}

func TestParseAge(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      int
		wantError bool
	}{
		{"valid age", "25", 25, false},
		{"zero age", "0", 0, false},
		{"max valid age", "150", 150, false},
		{"invalid format", "not-a-number", 0, true},
		{"negative age", "-5", 0, true},
		{"unrealistic age", "200", 0, true},
		{"empty string", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseAge(tt.input)

			if tt.wantError {
				if err == nil {
					t.Errorf("Expected error for input %q, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for input %q: %v", tt.input, err)
				}
				if result != tt.want {
					t.Errorf("parseAge(%q) = %d, want %d", tt.input, result, tt.want)
				}
			}
		})
	}
}

func TestValidationError(t *testing.T) {
	err := &ValidationError{
		Field:   "username",
		Value:   "",
		Message: "cannot be empty",
	}

	errorMsg := err.Error()
	if errorMsg == "" {
		t.Error("ValidationError.Error() returned empty string")
	}

	// Check that error message contains field name
	if !contains(errorMsg, "username") {
		t.Errorf("Error message should contain field name 'username': %s", errorMsg)
	}
}

func TestNewUser(t *testing.T) {
	tests := []struct {
		name      string
		username  string
		email     string
		age       int
		wantError bool
		errorType *ValidationError
	}{
		{
			name:      "valid user",
			username:  "alice",
			email:     "alice@example.com",
			age:       25,
			wantError: false,
		},
		{
			name:      "empty username",
			username:  "",
			email:     "test@example.com",
			age:       25,
			wantError: true,
		},
		{
			name:      "short username",
			username:  "ab",
			email:     "test@example.com",
			age:       25,
			wantError: true,
		},
		{
			name:      "empty email",
			username:  "alice",
			email:     "",
			age:       25,
			wantError: true,
		},
		{
			name:      "age too young",
			username:  "alice",
			email:     "alice@example.com",
			age:       10,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(tt.username, tt.email, tt.age)

			if tt.wantError {
				if err == nil {
					t.Error("Expected error, got nil")
				}

				// Verify it's a ValidationError
				var valErr *ValidationError
				if !errors.As(err, &valErr) {
					t.Error("Expected ValidationError type")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if user == nil {
					t.Error("Expected user, got nil")
				}
				if user.Username != tt.username {
					t.Errorf("Username = %s, want %s", user.Username, tt.username)
				}
			}
		})
	}
}

func TestFetchUser(t *testing.T) {
	tests := []struct {
		name      string
		id        int
		wantError bool
	}{
		{"valid id", 123, false},
		{"invalid id (zero)", 0, true},
		{"invalid id (negative)", -1, true},
		{"database error", 999, true},
		{"not found", 404, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := fetchUser(tt.id)

			if tt.wantError {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if user == nil {
					t.Error("Expected user, got nil")
				}
			}
		})
	}
}

func TestGetUserInfo(t *testing.T) {
	// Test error wrapping
	_, err := getUserInfo(-1)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	// Check that error is wrapped
	var valErr *ValidationError
	if !errors.As(err, &valErr) {
		t.Error("Expected wrapped ValidationError")
	}

	// Test database error wrapping
	_, err = getUserInfo(999)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	var dbErr *DatabaseError
	if !errors.As(err, &dbErr) {
		t.Error("Expected wrapped DatabaseError")
	}
}

func TestDatabaseError(t *testing.T) {
	originalErr := errors.New("connection failed")
	dbErr := &DatabaseError{
		Operation: "INSERT",
		Err:       originalErr,
	}

	// Test Error method
	errorMsg := dbErr.Error()
	if !contains(errorMsg, "INSERT") {
		t.Errorf("Error message should contain operation: %s", errorMsg)
	}

	// Test Unwrap
	unwrapped := dbErr.Unwrap()
	if unwrapped != originalErr {
		t.Error("Unwrap should return original error")
	}

	// Test errors.Is
	if !errors.Is(dbErr, originalErr) {
		t.Error("errors.Is should find original error")
	}
}

func TestErrorInspection(t *testing.T) {
	// Test errors.Is
	ErrNotFound := errors.New("not found")
	wrappedErr := errors.New("database: not found")

	if errors.Is(wrappedErr, ErrNotFound) {
		t.Error("Should not match different error instances")
	}

	// Test with actual wrapped error
	actualWrapped := errors.New("operation failed: %w")
	if !errors.Is(actualWrapped, actualWrapped) {
		t.Error("Error should match itself")
	}

	// Test errors.As
	_, err := getUserInfo(-10)
	var valErr *ValidationError
	if !errors.As(err, &valErr) {
		t.Error("Should extract ValidationError from wrapped error")
	}

	if valErr.Field != "id" {
		t.Errorf("Expected field 'id', got %s", valErr.Field)
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) &&
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
		findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
