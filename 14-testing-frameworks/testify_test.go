package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// TestCalculatorWithTestify demonstrates Testify assertions
func TestCalculatorWithTestify(t *testing.T) {
	calc := NewCalculator()

	// Basic assertions with assert
	// assert continues test execution even after failure
	t.Run("Add", func(t *testing.T) {
		result := calc.Add(2, 3)
		assert.Equal(t, 5, result, "2 + 3 should equal 5")
		assert.Greater(t, result, 0, "result should be positive")
	})

	t.Run("Subtract", func(t *testing.T) {
		result := calc.Subtract(10, 3)
		assert.Equal(t, 7, result)
		assert.NotEqual(t, 0, result)
	})

	t.Run("Multiply", func(t *testing.T) {
		result := calc.Multiply(4, 5)
		assert.Equal(t, 20, result)
		assert.IsType(t, 0, result)
	})

	// Using require - stops test execution on failure
	t.Run("Divide", func(t *testing.T) {
		result, err := calc.Divide(10, 2)
		require.NoError(t, err, "should not return error for valid division")
		assert.Equal(t, 5.0, result)
	})

	t.Run("Divide by zero", func(t *testing.T) {
		result, err := calc.Divide(10, 0)
		require.Error(t, err, "should return error for division by zero")
		assert.Equal(t, 0.0, result)
		assert.Contains(t, err.Error(), "division by zero")
	})
}

// TestCalculatorTableDrivenWithTestify demonstrates table-driven tests with Testify
func TestCalculatorTableDrivenWithTestify(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -5, 3, -2},
		{"zero", 0, 5, 5},
		{"large numbers", 1000, 2000, 3000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			assert.Equal(t, tt.expected, result,
				"Add(%d, %d) should equal %d", tt.a, tt.b, tt.expected)
		})
	}
}

// TestCalculatorAdvancedAssertions demonstrates various Testify assertions
func TestCalculatorAdvancedAssertions(t *testing.T) {
	calc := NewCalculator()

	t.Run("Collection assertions", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		sum := calc.Sum(numbers)

		assert.Equal(t, 15, sum)
		assert.Contains(t, numbers, 3, "slice should contain 3")
		assert.Len(t, numbers, 5, "slice should have 5 elements")
		assert.NotEmpty(t, numbers, "slice should not be empty")
	})

	t.Run("Boolean assertions", func(t *testing.T) {
		assert.True(t, calc.IsEven(4), "4 should be even")
		assert.False(t, calc.IsEven(5), "5 should not be even")
	})

	t.Run("Numeric assertions", func(t *testing.T) {
		result := calc.Power(2, 3)
		assert.InDelta(t, 8.0, result, 0.001, "2^3 should be approximately 8")
		assert.Greater(t, result, 0.0)
		assert.GreaterOrEqual(t, result, 8.0)
		assert.Less(t, result, 10.0)
		assert.LessOrEqual(t, result, 8.0)
	})

	t.Run("Nil assertions", func(t *testing.T) {
		var nilCalc *Calculator
		assert.Nil(t, nilCalc)
		assert.NotNil(t, calc)
	})
}

// TestCalculatorErrorHandling demonstrates error assertion patterns
func TestCalculatorErrorHandling(t *testing.T) {
	calc := NewCalculator()

	t.Run("Average of empty slice", func(t *testing.T) {
		_, err := calc.Average([]int{})
		require.Error(t, err)
		assert.EqualError(t, err, "cannot calculate average of empty slice")
	})

	t.Run("Max of empty slice", func(t *testing.T) {
		_, err := calc.Max([]int{})
		assert.Error(t, err)
		assert.ErrorContains(t, err, "empty slice")
	})

	t.Run("Min of empty slice", func(t *testing.T) {
		_, err := calc.Min([]int{})
		assert.Error(t, err)
	})

	t.Run("Valid operations", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		avg, err := calc.Average(numbers)
		require.NoError(t, err)
		assert.Equal(t, 3.0, avg)

		max, err := calc.Max(numbers)
		require.NoError(t, err)
		assert.Equal(t, 5, max)

		min, err := calc.Min(numbers)
		require.NoError(t, err)
		assert.Equal(t, 1, min)
	})
}

// CalculatorTestSuite demonstrates Testify test suites
// Test suites provide setup/teardown functionality
type CalculatorTestSuite struct {
	suite.Suite
	calc *Calculator
}

// SetupSuite runs once before all tests in the suite
func (suite *CalculatorTestSuite) SetupSuite() {
	// Setup that runs once for entire suite
}

// TearDownSuite runs once after all tests in the suite
func (suite *CalculatorTestSuite) TearDownSuite() {
	// Cleanup that runs once for entire suite
}

// SetupTest runs before each test
func (suite *CalculatorTestSuite) SetupTest() {
	suite.calc = NewCalculator()
}

// TearDownTest runs after each test
func (suite *CalculatorTestSuite) TearDownTest() {
	// Cleanup after each test
}

// TestAdd tests addition in the suite
func (suite *CalculatorTestSuite) TestAdd() {
	result := suite.calc.Add(2, 3)
	suite.Equal(5, result)
}

// TestSubtract tests subtraction in the suite
func (suite *CalculatorTestSuite) TestSubtract() {
	result := suite.calc.Subtract(10, 3)
	suite.Equal(7, result)
}

// TestMultiply tests multiplication in the suite
func (suite *CalculatorTestSuite) TestMultiply() {
	result := suite.calc.Multiply(4, 5)
	suite.Equal(20, result)
}

// TestCalculatorSuite runs the test suite
func TestCalculatorSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}

// TestAssertVsRequire demonstrates the difference between assert and require
// Note: In real scenarios, assert allows multiple checks to run even if some fail,
// while require stops immediately on failure. This test uses passing assertions
// to keep the test suite green while demonstrating the concept.
func TestAssertVsRequire(t *testing.T) {
	calc := NewCalculator()

	t.Run("Using assert - would continue after failure", func(t *testing.T) {
		result := calc.Add(2, 3)
		assert.Equal(t, 5, result, "assert checks continue even if previous ones failed")
		// In a real scenario where the above failed, the test would continue
		assert.NotNil(t, calc, "assert allows multiple checks in sequence")
		// All assertions are evaluated
		assert.Greater(t, result, 0, "assert is useful for multiple related checks")
	})

	t.Run("Using require - stops on failure", func(t *testing.T) {
		result, err := calc.Divide(10, 2)
		// If this fails, test stops immediately - use for critical preconditions
		require.NoError(t, err, "require stops test immediately on failure")
		// This only runs if above passes - safe to use result
		require.Equal(t, 5.0, result, "require is ideal for critical assertions")
		// Using require ensures we don't proceed with invalid state
	})
}

// TestTestifyBestPractices demonstrates Testify best practices
func TestTestifyBestPractices(t *testing.T) {
	calc := NewCalculator()

	t.Run("Use require for critical assertions", func(t *testing.T) {
		// Use require when subsequent code depends on the assertion
		numbers := []int{1, 2, 3}
		require.NotEmpty(t, numbers, "numbers must not be empty")

		// Safe to proceed - we know numbers is not empty
		sum := calc.Sum(numbers)
		assert.Equal(t, 6, sum)
	})

	t.Run("Use assert for multiple checks", func(t *testing.T) {
		// Use assert when you want to see all failures
		result := calc.Add(2, 3)
		assert.Equal(t, 5, result)
		assert.Greater(t, result, 0)
		assert.Less(t, result, 10)
		assert.IsType(t, 0, result)
	})

	t.Run("Provide helpful messages", func(t *testing.T) {
		result := calc.Multiply(3, 4)
		assert.Equal(t, 12, result, "Expected 3 * 4 = 12, got %d", result)
	})

	t.Run("Test edge cases", func(t *testing.T) {
		// Zero values
		assert.Equal(t, 0, calc.Add(0, 0))
		assert.Equal(t, 0, calc.Multiply(5, 0))

		// Negative numbers
		assert.Equal(t, -5, calc.Add(-2, -3))

		// Large numbers
		assert.Equal(t, 2000000000, calc.Add(1000000000, 1000000000))
	})
}
