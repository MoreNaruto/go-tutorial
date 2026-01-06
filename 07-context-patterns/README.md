# 07 - Context Patterns

**Level:** Intermediate

## What This Project Teaches

Context package for request-scoped values, cancellation, and timeouts:
- Context basics (`Background`, `TODO`)
- Timeout and deadline contexts
- Manual cancellation
- Passing values through context
- Context in goroutines and APIs

## Why This Project Matters

Context is essential for:
- Cancelling long-running operations
- Setting request timeouts
- Passing request-scoped data (user ID, trace ID)
- Graceful shutdown
- HTTP middleware chains

## Key Concepts

### Creating Contexts

```go
// Root contexts
ctx := context.Background()  // Production code
ctx := context.TODO()        // Placeholder/unsure

// Derived contexts
ctx, cancel := context.WithCancel(parent)
ctx, cancel := context.WithTimeout(parent, duration)
ctx, cancel := context.WithDeadline(parent, time)
ctx = context.WithValue(parent, key, value)
```

### Cancellation

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()  // Always call cancel

select {
case <-time.After(1 * time.Second):
    // Work completed
case <-ctx.Done():
    return ctx.Err()  // context.Canceled or context.DeadlineExceeded
}
```

### Context Values

```go
type key string
const userIDKey key = "userID"

ctx := context.WithValue(parent, userIDKey, 12345)

if userID, ok := ctx.Value(userIDKey).(int); ok {
    // Use userID
}
```

## Best Practices

1. **Always pass context as first parameter:** `func Do(ctx context.Context, ...)`
2. **Never store contexts:** They're request-scoped
3. **Always call cancel:** Use `defer cancel()`
4. **Check ctx.Done():** In long-running operations
5. **Don't put optional data in context:** Use for request-scoped only
6. **Use typed keys:** Define custom type for context keys

## Common Patterns

### HTTP Handler with Timeout

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()

    result, err := fetchData(ctx)
    // Handle response
}
```

### Concurrent Operations

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

for i := 0; i < 10; i++ {
    go worker(ctx, i)
}

// Cancel all workers
cancel()
```

## Running

```bash
go run main.go
go test ./...
go test -v
```

## What You'll Learn

1. When and why to use context
2. Different context creation patterns
3. Cancellation and timeout handling
4. Passing request-scoped values
5. Context in concurrent code

## Production Usage

- **HTTP servers:** Request cancellation
- **gRPC:** Deadline propagation
- **Database queries:** Query timeouts
- **Microservices:** Trace ID propagation
- **Background jobs:** Graceful cancellation

## Next Steps

- **05-http-server** - Context in HTTP handlers
- **12-concurrency-patterns** - Context with worker pools
- **09-clean-architecture** - Context through layers
