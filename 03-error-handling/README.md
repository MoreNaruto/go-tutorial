# 03 - Error Handling

**Level:** Beginner

## What This Project Teaches

This project covers Go's error handling patterns:
- Basic error creation and handling
- Custom error types
- Error wrapping with `fmt.Errorf` and `%w`
- Error inspection with `errors.Is` and `errors.As`
- Sentinel errors
- Error context and debugging

## Why This Structure?

This uses a **flat structure** focused on demonstrating error patterns:
- `main.go` - Examples of error creation, wrapping, and inspection
- `errors_test.go` - Tests verifying error handling behavior

Error handling is a core Go concept that appears in every real-world application.

## Key Concepts Explained

### Basic Error Handling

Go uses explicit error returns rather than exceptions:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Always check errors
result, err := divide(10, 0)
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}
fmt.Printf("Result: %f\n", result)
```

**Key principles:**
- Errors are values (implement the `error` interface)
- Functions that can fail return `(result, error)`
- Always check `if err != nil`
- `nil` error means success

### Custom Error Types

Create custom errors for structured error information:

```go
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error on field '%s': %s",
        e.Field, e.Message)
}
```

**Benefits:**
- Add context-specific fields
- Enable type-based error handling
- Provide detailed error information

### Error Wrapping

Wrap errors to add context while preserving the original:

```go
func getUserInfo(id int) (*User, error) {
    user, err := fetchUser(id)
    if err != nil {
        return nil, fmt.Errorf("failed to get user info for id %d: %w", id, err)
    }
    return user, nil
}
```

**The `%w` verb:**
- Wraps the original error
- Allows error inspection with `errors.Is` and `errors.As`
- Preserves error chain for debugging

**Without wrapping (`%v`):**
- Converts error to string
- Loses original error type
- Cannot use `errors.Is` or `errors.As`

### Error Inspection

Go 1.13+ provides `errors.Is` and `errors.As` for error inspection:

```go
// errors.Is checks if error matches a sentinel error
var ErrNotFound = errors.New("not found")
err := fmt.Errorf("database: %w", ErrNotFound)

if errors.Is(err, ErrNotFound) {
    // Handle not found case
}

// errors.As extracts custom error type
var valErr *ValidationError
if errors.As(err, &valErr) {
    fmt.Printf("Field: %s, Message: %s\n", valErr.Field, valErr.Message)
}
```

**errors.Is:**
- Checks if any error in the chain matches
- Works with sentinel errors
- Compares by value

**errors.As:**
- Finds first error in chain of specific type
- Uses type assertion
- Assigns to target variable

### Unwrap Method

Custom errors can implement `Unwrap()` to support error chains:

```go
type DatabaseError struct {
    Operation string
    Err       error
}

func (e *DatabaseError) Unwrap() error {
    return e.Err
}
```

This enables `errors.Is` and `errors.As` to traverse the error chain.

## Running the Code

Run the program to see error handling examples:

```bash
go run main.go
```

You'll see demonstrations of:
- Basic error handling
- Custom error types
- Error wrapping
- Error inspection with `Is` and `As`

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

## What You'll Learn

After working through this project, you'll understand:

1. Go's explicit error handling philosophy
2. How to create and return errors
3. When and why to create custom error types
4. How to wrap errors for better debugging
5. How to inspect errors with `Is` and `As`
6. Best practices for error messages

## Error Handling Best Practices

1. **Always check errors:** Don't ignore error returns
2. **Add context when wrapping:** Help future debugging
3. **Use %w for wrapping:** Enables error inspection
4. **Make errors actionable:** Provide enough information to fix the issue
5. **Use custom types sparingly:** Only when you need structured data
6. **Don't panic:** Reserve panics for truly unrecoverable situations
7. **Error messages should be lowercase:** Unless starting with proper noun

## Common Patterns

### Sentinel Errors

Pre-defined errors for comparison:

```go
var (
    ErrNotFound      = errors.New("not found")
    ErrUnauthorized  = errors.New("unauthorized")
    ErrInvalidInput  = errors.New("invalid input")
)
```

Common Sentinel Errors:

```
io.EOF
sql.ErrNoRows
context.Canceled
context.DeadlineExceeded
```

### Error Variables

Package-level error variables:

```go
var ErrDivisionByZero = errors.New("division by zero")
```

### Type Assertions

Check for specific error types:

```go
if valErr, ok := err.(*ValidationError); ok {
    // Handle validation error specifically
}
```

### Multiple Error Checks

```go
result, err := operation()
if err != nil {
    if errors.Is(err, ErrNotFound) {
        // Handle not found
    } else if errors.Is(err, ErrUnauthorized) {
        // Handle unauthorized
    } else {
        // Handle other errors
    }
    return
}
```

## Real-World Usage

Error handling patterns from standard library and popular packages:

- **os package:** `os.ErrNotExist`, `os.ErrPermission`
- **io package:** `io.EOF`, `io.ErrUnexpectedEOF`
- **net package:** Custom error types with timeout information
- **database/sql:** `sql.ErrNoRows`

## Anti-Patterns to Avoid

1. **Ignoring errors:** `result, _ := operation()`
2. **Generic error messages:** `errors.New("error")`
3. **Over-wrapping:** Adding too many layers without value
4. **Panic for expected errors:** Use errors, not panics
5. **Losing error context:** Using `%v` instead of `%w`

## Next Steps

After mastering error handling, move on to:
- **04-basic-concurrency** - Error handling in concurrent code
- **05-http-server** - HTTP error responses
- **08-testing-strategies** - Testing error conditions

## Further Reading

- [Go blog: Error handling and Go](https://go.dev/blog/error-handling-and-go)
- [Go blog: Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- [Effective Go: Errors](https://go.dev/doc/effective_go#errors)
