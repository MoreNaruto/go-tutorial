# 04 - Basic Concurrency

**Level:** Beginner

## What This Project Teaches

This project introduces Go's concurrency primitives:
- Goroutines (lightweight threads)
- Channels (communication between goroutines)
- Buffered vs unbuffered channels
- Select statement (multiplexing channels)
- WaitGroups (synchronization)
- Channel closing and range

## Why This Structure?

A **flat structure** for focused concurrency learning:
- `main.go` - Demonstrations of goroutines, channels, and synchronization
- `concurrency_test.go` - Tests for concurrent behavior

Concurrency is one of Go's most powerful features and a key reason for its popularity.

## Key Concepts Explained

### Goroutines

Goroutines are lightweight threads managed by the Go runtime:

```go
// Sequential
doWork()

// Concurrent
go doWork()  // Runs in separate goroutine
```

**Key points:**
- Very cheap (thousands can run concurrently)
- Managed by Go runtime, not OS
- Use `go` keyword to launch
- Don't return values directly (use channels)

### Channels

Channels are typed conduits for communication:

```go
ch := make(chan int)  // Create channel

// Send
ch <- 42

// Receive
value := <-ch
```

**Channel behavior:**
- **Unbuffered:** Send blocks until receive (synchronous)
- **Buffered:** Send blocks only when full
- Receive blocks until value available
- Both send and receive are thread-safe

### Buffered Channels

Channels with capacity:

```go
ch := make(chan int, 5)  // Buffer of 5

// Can send 5 values without blocking
ch <- 1
ch <- 2
// ...
ch <- 5
// ch <- 6  // Would block until someone receives
```

**Use cases:**
- Producer faster than consumer
- Limiting concurrency
- Reducing blocking

### Channel Direction

Function parameters can specify direction:

```go
// Send-only channel
func send(ch chan<- int) {
    ch <- 42
}

// Receive-only channel
func receive(ch <-chan int) {
    value := <-ch
}
```

This provides compile-time safety.

### Closing Channels

Sender can close channel to signal no more values:

```go
ch := make(chan int)
ch <- 1
ch <- 2
close(ch)

// Range over closed channel
for val := range ch {
    fmt.Println(val)  // Prints 1, 2, then stops
}

// Check if closed
val, ok := <-ch
if !ok {
    // Channel is closed
}
```

**Important:**
- Only sender should close
- Receiving from closed channel returns zero value
- Sending to closed channel panics

### Select Statement

Handle multiple channel operations:

```go
select {
case msg := <-ch1:
    fmt.Println("From ch1:", msg)
case msg := <-ch2:
    fmt.Println("From ch2:", msg)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout")
default:
    fmt.Println("No channel ready")
}
```

**Select behavior:**
- Waits until one case can proceed
- If multiple ready, chooses randomly
- `default` makes it non-blocking
- Common for timeouts and cancellation

### WaitGroup

Synchronize completion of multiple goroutines:

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        // Do work
    }(i)
}

wg.Wait()  // Block until all Done() called
```

**Pattern:**
- `Add(n)` before launching goroutine
- `Done()` when goroutine completes (use `defer`)
- `Wait()` blocks until counter reaches zero

## Running the Code

Run to see concurrency demonstrations:

```bash
go run main.go
```

Outputs will be interleaved due to concurrent execution.

## Running Tests

Execute tests:

```bash
go test ./...
```

Run with race detector (important for concurrent code):

```bash
go test -race ./...
```

Run benchmarks:

```bash
go test -bench=.
```

## What You'll Learn

After this project, you'll understand:

1. How to launch concurrent operations with goroutines
2. Channel-based communication between goroutines
3. Difference between buffered and unbuffered channels
4. Using select for channel multiplexing
5. Synchronizing goroutines with WaitGroups
6. Common concurrency patterns

## Common Patterns

### Fan-Out (distributing work)

```go
jobs := make(chan int, 100)
results := make(chan int, 100)

// Launch multiple workers
for w := 1; w <= 3; w++ {
    go worker(jobs, results)
}
```

### Pipeline

```go
ch1 := generate()
ch2 := square(ch1)
ch3 := print(ch2)
```

### Timeout

```go
select {
case result := <-ch:
    // Use result
case <-time.After(1 * time.Second):
    return errors.New("timeout")
}
```

## Common Pitfalls

1. **Forgetting to receive:** Sends on unbuffered channels block forever
2. **Closing channels:** Only sender should close, never the receiver
3. **Range over open channel:** Will block forever if channel never closes
4. **Closure variables:** Pass values to goroutines to avoid race conditions
5. **Not using `-race` flag:** Always test concurrent code with race detector

## Best Practices

1. **Pass channels:** Don't share memory, communicate via channels
2. **Close channels:** Sender's responsibility
3. **Use `defer wg.Done()`:** Ensures completion even with panics
4. **Buffered channels:** Size based on actual requirements
5. **Select default:** Only when you want non-blocking behavior
6. **Race detector:** Always use in tests

## Next Steps

After mastering basic concurrency:
- **07-context-patterns** - Cancellation and timeouts
- **12-concurrency-patterns** - Advanced patterns (worker pools, fan-in/out)
- **05-http-server** - Concurrent HTTP request handling

## Real-World Usage

Concurrency in production Go:
- **HTTP servers:** Each request in a goroutine
- **Background processing:** Worker pools
- **Data pipelines:** Producer-consumer chains
- **Timeouts:** Context with deadlines
- **Graceful shutdown:** Coordinating goroutine cleanup

## Further Reading

- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go blog: Concurrency is not parallelism](https://go.dev/blog/waza-talk)
- [Go blog: Go Concurrency Patterns](https://go.dev/blog/pipelines)
