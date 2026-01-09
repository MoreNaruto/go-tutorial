package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Sync.Mutex Tutorial ===")
	fmt.Println()

	// Demonstrate the race condition problem
	demonstrateRaceCondition()

	// Show the fix with Mutex
	demonstrateMutexSolution()

	// Show RWMutex for read-heavy workloads
	demonstrateRWMutex()

	// Show proper patterns
	demonstrateBestPractices()
}

// demonstrateRaceCondition shows what happens without proper synchronization
func demonstrateRaceCondition() {
	fmt.Println("--- Race Condition (UNSAFE) ---")
	fmt.Println("Running 1000 concurrent increments without mutex...")

	var counter int
	var wg sync.WaitGroup

	// Launch 1000 goroutines that increment counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// UNSAFE: This is a race condition!
			// counter++ is actually three operations:
			// 1. Read current value
			// 2. Add 1
			// 3. Write new value
			// Multiple goroutines can interleave these steps
			counter++
		}()
	}

	wg.Wait()

	// Expected: 1000, Actual: varies (race condition!)
	fmt.Printf("Expected: 1000, Got: %d\n", counter)
	if counter != 1000 {
		fmt.Println("⚠️  Race condition detected! Counter is incorrect.")
		fmt.Println("Run with: go run -race main.go to see race detector output")
	}
	fmt.Println()
}

// SafeCounter protects an integer with a mutex
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Increment safely increments the counter
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock() // Ensures unlock even if panic occurs
	c.value++
}

// Value safely reads the counter value
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// demonstrateMutexSolution shows the fix using sync.Mutex
func demonstrateMutexSolution() {
	fmt.Println("--- Mutex Solution (SAFE) ---")
	fmt.Println("Running 1000 concurrent increments with mutex...")

	counter := SafeCounter{}
	var wg sync.WaitGroup

	// Launch 1000 goroutines that safely increment counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// SAFE: Mutex ensures only one goroutine modifies at a time
			counter.Increment()
		}()
	}

	wg.Wait()

	result := counter.Value()
	fmt.Printf("Expected: 1000, Got: %d\n", result)
	if result == 1000 {
		fmt.Println("✓ Counter is correct! Mutex prevented race condition.")
	}
	fmt.Println()
}

// Cache demonstrates RWMutex for read-heavy workloads
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewCache creates a new cache
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

// Get retrieves a value (read lock)
// Multiple readers can access simultaneously
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock() // Read lock - multiple goroutines can hold this
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

// Set stores a value (write lock)
// Only one writer, blocks all readers
func (c *Cache) Set(key, value string) {
	c.mu.Lock() // Write lock - exclusive access
	defer c.mu.Unlock()
	c.data[key] = value
}

// demonstrateRWMutex shows RWMutex for read-heavy workloads
func demonstrateRWMutex() {
	fmt.Println("--- RWMutex (Read-Heavy Optimization) ---")

	cache := NewCache()

	// Pre-populate cache
	cache.Set("user:1", "Alice")
	cache.Set("user:2", "Bob")
	cache.Set("user:3", "Charlie")

	var wg sync.WaitGroup

	// Launch many readers (can run concurrently)
	fmt.Println("Launching 50 concurrent readers...")
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("user:%d", (id%3)+1)
			if val, ok := cache.Get(key); ok {
				// Read operations can happen concurrently
				_ = val // Use value
			}
		}(i)
	}

	// Launch occasional writers (exclusive access)
	fmt.Println("Launching 5 concurrent writers...")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("user:%d", id+10)
			cache.Set(key, fmt.Sprintf("User-%d", id+10))
			time.Sleep(10 * time.Millisecond) // Simulate slow write
		}(i)
	}

	wg.Wait()
	fmt.Println("✓ All operations completed safely")
	fmt.Println("RWMutex allowed multiple concurrent readers while ensuring safe writes")
	fmt.Println()
}

// BankAccount demonstrates proper mutex patterns
type BankAccount struct {
	mu      sync.Mutex
	balance int
}

// Deposit adds money to account
func (a *BankAccount) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

// Withdraw removes money if sufficient balance
func (a *BankAccount) Withdraw(amount int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.balance >= amount {
		a.balance -= amount
		return true
	}
	return false
}

// Balance returns current balance
func (a *BankAccount) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

// demonstrateBestPractices shows proper mutex usage patterns
func demonstrateBestPractices() {
	fmt.Println("--- Best Practices ---")

	account := &BankAccount{balance: 1000}

	var wg sync.WaitGroup

	// Multiple concurrent deposits
	fmt.Println("Processing 10 concurrent deposits of $100...")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.Deposit(100)
		}()
	}

	// Multiple concurrent withdrawals
	fmt.Println("Processing 5 concurrent withdrawals of $200...")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if account.Withdraw(200) {
				// Withdrawal successful
			} else {
				// Insufficient funds
			}
		}(i)
	}

	wg.Wait()

	finalBalance := account.Balance()
	fmt.Printf("Final balance: $%d\n", finalBalance)
	fmt.Println("✓ All transactions processed safely")
	fmt.Println()

	fmt.Println("Key takeaways:")
	fmt.Println("1. Always use 'defer mu.Unlock()' after 'mu.Lock()'")
	fmt.Println("2. Keep critical sections short")
	fmt.Println("3. Use RWMutex for read-heavy workloads")
	fmt.Println("4. Embed mutex with protected data in same struct")
	fmt.Println("5. Test with 'go run -race' to catch race conditions")
}
