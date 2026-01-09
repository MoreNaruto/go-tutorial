package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// TestGinkgo is the entry point for Ginkgo tests
func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}

// Ginkgo tests use BDD-style organization with Describe, Context, and It blocks
var _ = Describe("Calculator", func() {
	var calc *Calculator

	// BeforeEach runs before each test (It block)
	BeforeEach(func() {
		calc = NewCalculator()
	})

	// Describe groups related tests
	Describe("Basic operations", func() {
		// Context adds additional context/scenarios
		Context("when adding numbers", func() {
			// It describes a specific behavior
			It("should return correct sum for positive numbers", func() {
				Expect(calc.Add(2, 3)).To(Equal(5))
			})

			It("should return correct sum for negative numbers", func() {
				Expect(calc.Add(-2, -3)).To(Equal(-5))
			})

			It("should handle zero", func() {
				Expect(calc.Add(0, 5)).To(Equal(5))
				Expect(calc.Add(5, 0)).To(Equal(5))
			})

			It("should handle large numbers", func() {
				result := calc.Add(1000000, 2000000)
				Expect(result).To(Equal(3000000))
				Expect(result).To(BeNumerically(">", 0))
			})
		})

		Context("when subtracting numbers", func() {
			It("should return correct difference", func() {
				Expect(calc.Subtract(10, 3)).To(Equal(7))
			})

			It("should handle negative results", func() {
				Expect(calc.Subtract(3, 10)).To(Equal(-7))
			})
		})

		Context("when multiplying numbers", func() {
			It("should return correct product", func() {
				Expect(calc.Multiply(4, 5)).To(Equal(20))
			})

			It("should handle zero", func() {
				Expect(calc.Multiply(5, 0)).To(Equal(0))
				Expect(calc.Multiply(0, 5)).To(Equal(0))
			})

			It("should handle negative numbers", func() {
				Expect(calc.Multiply(-2, 3)).To(Equal(-6))
				Expect(calc.Multiply(-2, -3)).To(Equal(6))
			})
		})

		Context("when dividing numbers", func() {
			It("should return correct quotient", func() {
				result, err := calc.Divide(10, 2)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(5.0))
			})

			It("should return error for division by zero", func() {
				result, err := calc.Divide(10, 0)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("division by zero"))
				Expect(result).To(Equal(0.0))
			})

			It("should handle fractional results", func() {
				result, err := calc.Divide(10, 3)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(BeNumerically("~", 3.333, 0.001))
			})
		})
	})

	Describe("Advanced operations", func() {
		Context("when calculating power", func() {
			It("should calculate positive exponents", func() {
				Expect(calc.Power(2, 3)).To(Equal(8.0))
				Expect(calc.Power(5, 2)).To(Equal(25.0))
			})

			It("should handle zero exponent", func() {
				Expect(calc.Power(5, 0)).To(Equal(1.0))
			})

			It("should handle negative exponents", func() {
				result := calc.Power(2, -2)
				Expect(result).To(BeNumerically("~", 0.25, 0.001))
			})
		})

		Context("when checking even numbers", func() {
			It("should identify even numbers", func() {
				Expect(calc.IsEven(2)).To(BeTrue())
				Expect(calc.IsEven(4)).To(BeTrue())
				Expect(calc.IsEven(100)).To(BeTrue())
			})

			It("should identify odd numbers", func() {
				Expect(calc.IsEven(1)).To(BeFalse())
				Expect(calc.IsEven(3)).To(BeFalse())
				Expect(calc.IsEven(99)).To(BeFalse())
			})

			It("should handle zero as even", func() {
				Expect(calc.IsEven(0)).To(BeTrue())
			})

			It("should handle negative numbers", func() {
				Expect(calc.IsEven(-2)).To(BeTrue())
				Expect(calc.IsEven(-3)).To(BeFalse())
			})
		})

		Context("when calculating absolute value", func() {
			It("should return same value for positive numbers", func() {
				Expect(calc.Abs(5)).To(Equal(5))
			})

			It("should return positive for negative numbers", func() {
				Expect(calc.Abs(-5)).To(Equal(5))
			})

			It("should handle zero", func() {
				Expect(calc.Abs(0)).To(Equal(0))
			})
		})
	})

	Describe("Collection operations", func() {
		Context("when summing numbers", func() {
			It("should calculate correct sum", func() {
				numbers := []int{1, 2, 3, 4, 5}
				Expect(calc.Sum(numbers)).To(Equal(15))
			})

			It("should handle empty slice", func() {
				Expect(calc.Sum([]int{})).To(Equal(0))
			})

			It("should handle single element", func() {
				Expect(calc.Sum([]int{42})).To(Equal(42))
			})

			It("should handle negative numbers", func() {
				numbers := []int{-1, -2, -3}
				Expect(calc.Sum(numbers)).To(Equal(-6))
			})
		})

		Context("when calculating average", func() {
			It("should calculate correct average", func() {
				numbers := []int{2, 4, 6, 8, 10}
				avg, err := calc.Average(numbers)
				Expect(err).NotTo(HaveOccurred())
				Expect(avg).To(Equal(6.0))
			})

			It("should return error for empty slice", func() {
				_, err := calc.Average([]int{})
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("cannot calculate average of empty slice"))
			})

			It("should handle single element", func() {
				avg, err := calc.Average([]int{42})
				Expect(err).NotTo(HaveOccurred())
				Expect(avg).To(Equal(42.0))
			})

			It("should handle fractional averages", func() {
				numbers := []int{1, 2, 3}
				avg, err := calc.Average(numbers)
				Expect(err).NotTo(HaveOccurred())
				Expect(avg).To(BeNumerically("~", 2.0, 0.001))
			})
		})

		Context("when finding maximum", func() {
			It("should find max in positive numbers", func() {
				numbers := []int{1, 5, 3, 9, 2}
				max, err := calc.Max(numbers)
				Expect(err).NotTo(HaveOccurred())
				Expect(max).To(Equal(9))
			})

			It("should return error for empty slice", func() {
				_, err := calc.Max([]int{})
				Expect(err).To(HaveOccurred())
			})

			It("should handle negative numbers", func() {
				numbers := []int{-5, -2, -10, -1}
				max, err := calc.Max(numbers)
				Expect(err).NotTo(HaveOccurred())
				Expect(max).To(Equal(-1))
			})

			It("should handle single element", func() {
				max, err := calc.Max([]int{42})
				Expect(err).NotTo(HaveOccurred())
				Expect(max).To(Equal(42))
			})
		})

		Context("when finding minimum", func() {
			It("should find min in positive numbers", func() {
				numbers := []int{5, 1, 9, 3, 2}
				min, err := calc.Min(numbers)
				Expect(err).NotTo(HaveOccurred())
				Expect(min).To(Equal(1))
			})

			It("should return error for empty slice", func() {
				_, err := calc.Min([]int{})
				Expect(err).To(HaveOccurred())
			})

			It("should handle negative numbers", func() {
				numbers := []int{-5, -2, -10, -1}
				min, err := calc.Min(numbers)
				Expect(err).NotTo(HaveOccurred())
				Expect(min).To(Equal(-10))
			})
		})
	})

	// Demonstrate table-driven tests with Ginkgo
	Describe("Table-driven tests", func() {
		DescribeTable("addition",
			func(a, b, expected int) {
				Expect(calc.Add(a, b)).To(Equal(expected))
			},
			Entry("positive numbers", 2, 3, 5),
			Entry("negative numbers", -2, -3, -5),
			Entry("mixed signs", -5, 3, -2),
			Entry("with zero", 0, 5, 5),
		)

		DescribeTable("multiplication",
			func(a, b, expected int) {
				Expect(calc.Multiply(a, b)).To(Equal(expected))
			},
			Entry("positive numbers", 2, 3, 6),
			Entry("with zero", 5, 0, 0),
			Entry("negative numbers", -2, 3, -6),
		)
	})

	// Demonstrate various Gomega matchers
	Describe("Gomega matchers", func() {
		It("should demonstrate equality matchers", func() {
			Expect(calc.Add(2, 3)).To(Equal(5))
			Expect(calc.Add(2, 3)).NotTo(Equal(6))
			Expect(calc.Add(2, 3)).To(BeEquivalentTo(5))
		})

		It("should demonstrate numeric matchers", func() {
			result := calc.Add(2, 3)
			Expect(result).To(BeNumerically("==", 5))
			Expect(result).To(BeNumerically(">", 0))
			Expect(result).To(BeNumerically(">=", 5))
			Expect(result).To(BeNumerically("<", 10))
			Expect(result).To(BeNumerically("<=", 5))
			Expect(result).To(BeNumerically("~", 5, 0.1))
		})

		It("should demonstrate boolean matchers", func() {
			Expect(calc.IsEven(4)).To(BeTrue())
			Expect(calc.IsEven(5)).To(BeFalse())
		})

		It("should demonstrate error matchers", func() {
			_, err := calc.Divide(10, 0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("division"))

			_, err = calc.Divide(10, 2)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should demonstrate collection matchers", func() {
			numbers := []int{1, 2, 3, 4, 5}
			Expect(numbers).To(ContainElement(3))
			Expect(numbers).To(HaveLen(5))
			Expect(numbers).NotTo(BeEmpty())
			Expect(numbers).To(ConsistOf(1, 2, 3, 4, 5))
		})

		It("should demonstrate type matchers", func() {
			result := calc.Add(2, 3)
			Expect(result).To(BeAssignableToTypeOf(0))
		})
	})
})

// Demonstrate focused and pending tests
var _ = Describe("Test control", func() {
	// FDescribe, FContext, FIt focus on specific tests
	// Only focused tests run when any are present
	// Use during development, remove before commit

	// PDescribe, PContext, PIt mark tests as pending
	// They show up in output but don't run

	It("should run normally", func() {
		calc := NewCalculator()
		Expect(calc.Add(1, 1)).To(Equal(2))
	})

	// Uncomment to see focused test behavior:
	// FIt("should run as focused test", func() {
	// 	calc := NewCalculator()
	// 	Expect(calc.Add(2, 2)).To(Equal(4))
	// })

	// Uncomment to see pending test behavior:
	// PIt("should be pending", func() {
	// 	// This test is marked as pending and won't run
	// })
})

// Demonstrate nested BeforeEach/AfterEach
var _ = Describe("Setup and teardown", func() {
	var calc *Calculator
	var setupOrder []string

	BeforeEach(func() {
		setupOrder = append(setupOrder, "outer-before")
		calc = NewCalculator()
	})

	AfterEach(func() {
		setupOrder = append(setupOrder, "outer-after")
	})

	It("should run outer setup/teardown", func() {
		Expect(calc).NotTo(BeNil())
	})

	Context("nested context", func() {
		BeforeEach(func() {
			setupOrder = append(setupOrder, "inner-before")
		})

		AfterEach(func() {
			setupOrder = append(setupOrder, "inner-after")
		})

		It("should run outer then inner setup", func() {
			// Order: outer-before -> inner-before -> test -> inner-after -> outer-after
			Expect(calc).NotTo(BeNil())
		})
	})
})
