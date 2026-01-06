# 08 - Testing Strategies

**Level:** Intermediate

## What This Project Teaches

Go testing patterns and best practices:
- Table-driven tests
- Subtests with `t.Run`
- Mocking with interfaces
- Test helpers
- Benchmarks
- Example tests
- Test coverage

## Key Patterns

### Table-Driven Tests

```go
tests := []struct {
    name     string
    input    int
    expected int
}{
    {"case 1", 5, 10},
    {"case 2", 3, 6},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        result := function(tt.input)
        if result != tt.expected {
            t.Errorf("got %d, want %d", result, tt.expected)
        }
    })
}
```

### Mocking with Interfaces

```go
type DataStore interface {
    Get(key string) (string, error)
}

type MockDataStore struct {
    data map[string]string
}

func (m *MockDataStore) Get(key string) (string, error) {
    return m.data[key], nil
}
```

### Test Helpers

```go
func assertEqual(t *testing.T, got, want interface{}) {
    t.Helper()  // Marks this as helper function
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}
```

## Running Tests

```bash
go test ./...                    # Run all tests
go test -v                       # Verbose output
go test -run TestName            # Run specific test
go test -cover                   # Show coverage
go test -coverprofile=cover.out  # Generate coverage report
go tool cover -html=cover.out    # View coverage in browser
go test -bench=.                 # Run benchmarks
```

## Best Practices

1. **Use table-driven tests:** Test multiple cases systematically
2. **Mock dependencies:** Use interfaces for testability
3. **Test one thing:** Each test should verify one behavior
4. **Clear test names:** Describe what is being tested
5. **Use subtests:** Group related tests
6. **Test edge cases:** Zero values, boundaries, errors
7. **Use test helpers:** Reduce duplication with `t.Helper()`

## Test Organization

```
package/
  ├── service.go          # Implementation
  ├── service_test.go     # Tests
  └── mock_store_test.go  # Test mocks
```

## What You'll Learn

1. Writing comprehensive tests
2. Table-driven test pattern
3. Mocking dependencies
4. Test organization
5. Coverage analysis
6. Benchmarking

## Next Steps

- **09-clean-architecture** - Testing layered architecture
- **06-rest-api-gin** - HTTP handler testing
- **12-concurrency-patterns** - Testing concurrent code
