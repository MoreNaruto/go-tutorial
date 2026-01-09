# 13 - Sync.Mutex

**Level:** Intermediate

## What This Project Teaches

This project demonstrates Go's `sync.Mutex` for protecting shared state in concurrent programs:
- What race conditions are and why they're dangerous
- How `sync.Mutex` provides mutual exclusion
- Proper locking and unlocking patterns
- `sync.RWMutex` for read-heavy workloads
- Common pitfalls and best practices

## Why This Structure?

A **flat structure** for focused mutex learning:
- `main.go` - Race condition demonstration and mutex solutions
- `mutex_test.go` - Tests for concurrent safety

While channels are Go's preferred concurrency primitive ("Don't communicate by sharing memory, share memory by communicating"), mutexes are essential when:
- Protecting simple shared state
- Coordinating access to data structures
- Implementing caching mechanisms
- Interfacing with non-channel-safe code

## What is sync.Mutex?

`sync.Mutex` (mutual exclusion lock) ensures only one goroutine can access a critical section at a time.

**Key characteristics:**
- **Lock()** - Acquire the lock (blocks if already locked)
- **Unlock()** - Release the lock (must be called by the same goroutine that locked)
- Not reentrant (same goroutine can't lock twice)
- Zero value is ready to use
- Must not be copied after first use

## The Problem: Race Conditions

A race condition occurs when multiple goroutines access shared data concurrently and at least one modifies it:

```go
var counter int  // Shared state

for i := 0; i < 100; i++ {
    go func() {
        counter++  // UNSAFE: read-modify-write is not atomic
    }()
}

// Final counter value is unpredictable!
```

**Why it's unsafe:**
1. Goroutine A reads counter (value: 5)
2. Goroutine B reads counter (value: 5)
3. Goroutine A increments to 6, writes back
4. Goroutine B increments to 6, writes back
5. Counter is 6, not 7!

## The Solution: Mutex

```go
var (
    counter int
    mu      sync.Mutex  // Protects counter
)

for i := 0; i < 100; i++ {
    go func() {
        mu.Lock()
        counter++  // SAFE: only one goroutine at a time
        mu.Unlock()
    }()
}
```

**How it works:**
- First goroutine calls `Lock()` → succeeds, proceeds
- Other goroutines call `Lock()` → block until first calls `Unlock()`
- Only one goroutine in critical section at any time

## Basic Mutex Pattern

```go
type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()  // Ensures unlock even if panic
    c.value++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}
```

**Pattern notes:**
- Mutex embedded in struct with protected data
- Use `defer` to guarantee unlock
- Lock protects all access, reads AND writes

## RWMutex: Optimizing for Reads

`sync.RWMutex` allows multiple concurrent readers OR one exclusive writer:

```go
type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

// Multiple goroutines can read concurrently
func (c *Cache) Get(key string) string {
    c.mu.RLock()         // Reader lock
    defer c.mu.RUnlock()
    return c.data[key]
}

// Only one goroutine can write
func (c *Cache) Set(key, value string) {
    c.mu.Lock()          // Writer lock (exclusive)
    defer c.mu.Unlock()
    c.data[key] = value
}
```

**When to use RWMutex:**
- Read-heavy workloads (10+ reads per write)
- Read operations are expensive
- Contention is high

**When NOT to use:**
- Write-heavy or balanced read/write
- Critical sections are very short
- Added complexity not worth marginal gains

## Running the Code

Run to see race conditions and their fixes:

```bash
go run main.go
```

Run with race detector (IMPORTANT for concurrent code):

```bash
go run -race main.go
```

The race detector will catch data races at runtime.

## Running Tests

Execute tests:

```bash
go test ./...
```

Run with race detector:

```bash
go test -race ./...
```

Run benchmarks:

```bash
go test -bench=. -benchmem
```

## Common Patterns

### Pattern 1: Defer Unlock

```go
mu.Lock()
defer mu.Unlock()
// Critical section
// Unlock guaranteed even if panic
```

### Pattern 2: Struct Embedding

```go
type SafeMap struct {
    mu   sync.Mutex
    data map[string]int
}
```

Keep mutex and protected data together.

### Pattern 3: Mutex Per Resource

```go
type Server struct {
    usersMu   sync.RWMutex
    users     map[int]*User

    configMu  sync.RWMutex
    config    *Config
}
```

Fine-grained locking reduces contention.

### Pattern 4: Try-Lock (Advanced)

```go
type TryLocker struct {
    mu sync.Mutex
}

func (t *TryLocker) TryLock() bool {
    return t.mu.TryLock()  // Go 1.18+
}
```

Non-blocking lock attempt.

## Common Pitfalls

### 1. Forgetting to Unlock

```go
mu.Lock()
doWork()
// BUG: Never unlocked!
```

**Fix:** Always use `defer mu.Unlock()`

### 2. Copying a Mutex

```go
type Counter struct {
    mu sync.Mutex
    n  int
}

c1 := Counter{}
c2 := c1  // BUG: Copies mutex!
```

**Fix:** Pass by pointer: `func process(c *Counter)`

### 3. Locking Twice (Deadlock)

```go
mu.Lock()
helper()  // Calls mu.Lock() again!
mu.Unlock()
```

**Fix:** Document lock requirements, use RLock when possible

### 4. Holding Lock Too Long

```go
mu.Lock()
defer mu.Unlock()
slowNetworkCall()  // Blocks all other goroutines!
```

**Fix:** Minimize critical section, unlock before slow I/O

### 5. Wrong Lock for Operation

```go
func (c *Cache) Get(key string) string {
    c.mu.Lock()  // Should be RLock()!
    defer c.mu.Unlock()
    return c.data[key]
}
```

**Fix:** Use RLock() for reads with RWMutex

## Best Practices

1. **Always use defer:** `defer mu.Unlock()` after `mu.Lock()`
2. **Keep critical sections short:** Lock, modify, unlock
3. **Document lock requirements:** Which mutex protects which data
4. **Use race detector:** `go test -race` catches issues
5. **Consider channels first:** Mutexes are lower-level
6. **Don't export mutex:** Keep `sync.Mutex` unexported
7. **Pass by pointer:** Never copy structs containing mutexes
8. **One lock per resource:** Avoid global locks when possible

## When to Use Mutex vs Channels

**Use Mutex when:**
- Protecting simple state (counter, cache)
- Need direct access to shared data
- Interfacing with non-Go-safe libraries
- Very short critical sections

**Use Channels when:**
- Coordinating goroutines
- Passing ownership of data
- Distributing work
- Implementing pipelines

**Rule of thumb:** Prefer channels. Use mutexes when channels are awkward or inefficient.

## Performance Considerations

- **Contention:** High lock contention kills performance
- **Critical section size:** Smaller is faster
- **Lock granularity:** More locks = less contention (but more complexity)
- **RWMutex overhead:** Only worth it for read-heavy workloads

Benchmark to verify assumptions!

## Debugging Race Conditions

### Race Detector

```bash
go test -race
go run -race main.go
go build -race
```

**Notes:**
- Catches races at runtime (not static analysis)
- Adds significant overhead (~10x slower)
- Use in tests and development, not production
- Only catches races that actually occur during run

### Common Race Patterns

1. **Unprotected variable access**
2. **Closure variable capture**
3. **Map concurrent access**
4. **Slice concurrent modification**

## Real-World Examples

**Caching:**
```go
type UserCache struct {
    mu    sync.RWMutex
    users map[int]*User
}
```

**Connection Pool:**
```go
type Pool struct {
    mu    sync.Mutex
    conns []*Conn
}
```

**Request Counter:**
```go
type Stats struct {
    mu       sync.Mutex
    requests int
}
```

**Configuration:**
```go
type Config struct {
    mu   sync.RWMutex
    data map[string]string
}
```

## What You'll Learn

After this project, you'll understand:

1. What race conditions are and how to detect them
2. How to use `sync.Mutex` for mutual exclusion
3. Proper locking/unlocking patterns with `defer`
4. When to use `sync.RWMutex` for read-heavy workloads
5. Common pitfalls and how to avoid them
6. When to prefer channels over mutexes

## Next Steps

After mastering mutexes:
- **04-basic-concurrency** - Goroutines and channels
- **12-concurrency-patterns** - Advanced patterns (worker pools)
- **07-context-patterns** - Cancellation and timeouts

## Further Reading

- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go Memory Model](https://go.dev/ref/mem)
- [Race Detector](https://go.dev/doc/articles/race_detector)
- [sync package documentation](https://pkg.go.dev/sync)
