# 14 - Testing Frameworks

**Level:** Intermediate to Advanced

## What This Project Teaches

This project demonstrates popular Go testing frameworks and tools:
- **Testify** - Assertion and mocking library
- **Ginkgo** - BDD-style testing framework
- **Gomega** - Matcher/assertion library (pairs with Ginkgo)
- **Httpexpect** - End-to-end HTTP API testing
- **Benchmarks** - Go's built-in performance testing

## Why This Structure?

A **focused structure** demonstrating each framework:
- `calculator.go` - Simple calculator for testing
- `server.go` - HTTP server for API testing
- `testify_test.go` - Testify examples
- `ginkgo_test.go` - Ginkgo/Gomega examples
- `httpexpect_test.go` - HTTP API testing examples
- `benchmark_test.go` - Performance benchmarking

While Go's standard `testing` package is excellent, third-party frameworks add:
- More expressive assertions
- BDD-style test organization
- Better error messages
- Specialized testing capabilities

## Installation

Initialize the module and install dependencies:

```bash
go mod download
```

Or install individual frameworks:

```bash
# Testify
go get github.com/stretchr/testify

# Ginkgo and Gomega
go get github.com/onsi/ginkgo/v2/ginkgo
go get github.com/onsi/gomega

# Httpexpect
go get github.com/gavv/httpexpect/v2
```

## Framework Overview

### 1. Testify

**What it is:** Assertion and mocking toolkit for standard Go tests

**Why use it:**
- More readable assertions than `if/else` + `t.Error`
- Rich assertion library (`assert`, `require`)
- Mock generation for interfaces
- Suite support for setup/teardown

**When to use:**
- You want better assertions but keep standard test structure
- Need mocking capabilities
- Want minimal changes to existing tests

**Example:**
```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    assert.Equal(t, 5, result)
    assert.Greater(t, result, 0)
}
```

### 2. Ginkgo

**What it is:** BDD (Behavior-Driven Development) testing framework

**Why use it:**
- Expressive, nested test organization
- `Describe`, `Context`, `It` blocks (like RSpec/Jasmine)
- Built-in test focus and skip
- Parallel test execution
- Rich CLI with watch mode

**When to use:**
- You prefer BDD-style tests
- Complex test suites with shared setup
- Need advanced test organization
- Team has BDD background

**Example:**
```go
Describe("Calculator", func() {
    Context("when adding numbers", func() {
        It("should return correct sum", func() {
            Expect(Add(2, 3)).To(Equal(5))
        })
    })
})
```

### 3. Gomega

**What it is:** Matcher/assertion library (pairs with Ginkgo)

**Why use it:**
- Expressive matchers: `Equal`, `BeNumerically`, `ContainElement`
- Async assertions: `Eventually`, `Consistently`
- Works with Ginkgo or standalone
- Better error messages

**When to use:**
- With Ginkgo tests
- Need async/polling assertions
- Want readable matchers

**Example:**
```go
Expect(result).To(Equal(5))
Expect(slice).To(ContainElement("foo"))
Eventually(func() int { return counter.Value() }).Should(Equal(100))
```

### 4. Httpexpect

**What it is:** End-to-end HTTP API testing toolkit

**Why use it:**
- Fluent API for HTTP requests/responses
- Built-in JSON/XML parsing
- Response validation
- WebSocket support
- Cookie and session handling

**When to use:**
- Testing REST APIs
- Integration tests for HTTP services
- Need to test full request/response cycle
- Want readable API tests

**Example:**
```go
e := httpexpect.Default(t, "http://localhost:8080")
e.GET("/users/1").
    Expect().
    Status(http.StatusOK).
    JSON().Object().
    ValueEqual("name", "Alice")
```

### 5. Benchmarks

**What it is:** Go's built-in performance testing

**Why use it:**
- Measure performance and allocations
- Compare implementations
- Detect performance regressions
- No external dependencies

**When to use:**
- Optimizing critical code paths
- Comparing algorithm choices
- Ensuring performance targets
- CI performance checks

**Example:**
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

## Running Tests

### Standard Go Tests (Testify)

```bash
# Run all tests
go test ./...

# Verbose output
go test -v

# Run specific test
go test -run TestCalculator
```

### Ginkgo Tests

```bash
# Install Ginkgo CLI
go install github.com/onsi/ginkgo/v2/ginkgo

# Run Ginkgo tests
ginkgo

# Verbose output
ginkgo -v

# Watch mode (re-run on changes)
ginkgo watch

# Parallel execution
ginkgo -p
```

### Benchmarks

```bash
# Run all benchmarks
go test -bench=.

# With memory allocation stats
go test -bench=. -benchmem

# Run specific benchmark
go test -bench=BenchmarkAdd

# Compare benchmarks
go test -bench=. > old.txt
# ... make changes ...
go test -bench=. > new.txt
benchcmp old.txt new.txt  # Requires golang.org/x/tools/cmd/benchcmp
```

## Framework Comparison

| Feature | Standard | Testify | Ginkgo | Httpexpect |
|---------|----------|---------|--------|------------|
| Assertions | Manual | ✓ | ✓ (Gomega) | ✓ |
| BDD Style | ✗ | ✗ | ✓ | ✗ |
| Mocking | Manual | ✓ | Manual/3rd | ✗ |
| HTTP Testing | Manual | Manual | Manual | ✓ |
| Learning Curve | Low | Low | Medium | Low |
| Boilerplate | High | Medium | Low | Low |
| Parallel Tests | ✓ | ✓ | ✓ | ✓ |

## Best Practices

### General Testing

1. **One assertion per test** (in BDD style)
2. **Clear test names** - Describe what is being tested
3. **Arrange-Act-Assert** pattern
4. **Avoid test interdependence**
5. **Use table-driven tests** for multiple cases
6. **Test edge cases** - zero values, boundaries, errors

### Testify

1. Use `assert` for non-critical checks
2. Use `require` to stop test on failure
3. Use `suite` for setup/teardown
4. Generate mocks: `mockery`

### Ginkgo

1. Use `Describe` for components
2. Use `Context` for scenarios
3. Use `It` for specific behaviors
4. Keep `BeforeEach`/`AfterEach` minimal
5. Use `FDescribe`/`FIt` to focus tests during development
6. Avoid deep nesting (max 3-4 levels)

### Httpexpect

1. Use builders for requests
2. Chain assertions
3. Test status codes first
4. Validate response structure
5. Use matchers for flexibility

### Benchmarks

1. Use `b.ResetTimer()` after setup
2. Use `b.StopTimer()`/`b.StartTimer()` to exclude overhead
3. Use `b.RunParallel()` for concurrent scenarios
4. Use `b.ReportAllocs()` to track allocations
5. Run multiple times: `go test -bench=. -count=5`

## When to Use Which Framework

### Use Standard Testing When:
- Simple projects
- Minimal dependencies
- Standard Go idioms preferred

### Use Testify When:
- Want better assertions
- Need mocking
- Standard test structure preferred
- Existing tests to enhance

### Use Ginkgo When:
- Prefer BDD style
- Complex test organization needed
- Team has BDD experience
- Need advanced features (focus, parallel, watch)

### Use Httpexpect When:
- Testing HTTP APIs
- Integration/E2E tests
- Want fluent API test syntax

### Use Benchmarks When:
- Performance matters
- Comparing implementations
- Optimizing hot paths
- CI performance gates

## Common Patterns

### Table-Driven with Testify

```go
tests := []struct {
    name     string
    a, b     int
    expected int
}{
    {"positive", 2, 3, 5},
    {"negative", -1, -2, -3},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        assert.Equal(t, tt.expected, Add(tt.a, tt.b))
    })
}
```

### Async Assertions with Gomega

```go
Eventually(func() int {
    return counter.Value()
}, "5s", "100ms").Should(Equal(100))
```

### API Test Chain with Httpexpect

```go
e.POST("/users").
    WithJSON(user).
    Expect().
    Status(http.StatusCreated).
    JSON().Object().
    ContainsKey("id")
```

## What You'll Learn

After this project, you'll understand:

1. How to write tests with Testify assertions
2. How to organize tests with Ginkgo BDD style
3. How to use Gomega matchers
4. How to test HTTP APIs with Httpexpect
5. How to write and run benchmarks
6. When to choose each framework

## Real-World Usage

**Testify:**
- Kubernetes, Docker, Terraform
- Most popular Go testing library
- Widely used in open source

**Ginkgo:**
- Cloud Foundry, Kubernetes controllers
- Preferred for complex test suites
- Common in enterprise Go projects

**Httpexpect:**
- API integration tests
- E2E testing
- Service testing

**Benchmarks:**
- Standard library development
- Performance-critical applications
- Optimization verification

## Next Steps

After mastering testing frameworks:
- **08-testing-strategies** - Testing patterns and best practices
- **09-clean-architecture** - Testing layered architectures
- **06-rest-api-gin** - API testing in practice

## Further Reading

- [Testify Documentation](https://github.com/stretchr/testify)
- [Ginkgo Documentation](https://onsi.github.io/ginkgo/)
- [Gomega Documentation](https://onsi.github.io/gomega/)
- [Httpexpect Documentation](https://github.com/gavv/httpexpect)
- [Go Blog: Benchmarks](https://go.dev/blog/benchmark)
- [Effective Go: Testing](https://go.dev/doc/effective_go#testing)
