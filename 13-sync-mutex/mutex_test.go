package main

import (
	"sync"
	"testing"
)

// TestSafeCounter verifies counter is thread-safe
func TestSafeCounter(t *testing.T) {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	// Launch 1000 goroutines incrementing concurrently
	iterations := 1000
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	if counter.Value() != iterations {
		t.Errorf("Expected %d, got %d", iterations, counter.Value())
	}
}

// TestSafeCounterConcurrentReads verifies concurrent reads are safe
func TestSafeCounterConcurrentReads(t *testing.T) {
	counter := SafeCounter{}
	counter.Increment()
	counter.Increment()
	counter.Increment()

	var wg sync.WaitGroup

	// Multiple goroutines reading concurrently
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val := counter.Value()
			if val < 0 {
				t.Errorf("Unexpected negative value: %d", val)
			}
		}()
	}

	wg.Wait()
}

// TestCache verifies cache operations are thread-safe
func TestCache(t *testing.T) {
	cache := NewCache()
	var wg sync.WaitGroup

	// Concurrent writes
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cache.Set(string(rune('A'+id%26)), string(rune('a'+id%26)))
		}(i)
	}

	wg.Wait()

	// Verify data exists
	val, ok := cache.Get("A")
	if !ok {
		t.Error("Expected key 'A' to exist")
	}
	if val == "" {
		t.Error("Expected non-empty value")
	}
}

// TestCacheConcurrentReadWrite verifies concurrent reads and writes are safe
func TestCacheConcurrentReadWrite(t *testing.T) {
	cache := NewCache()
	cache.Set("key", "initial")

	var wg sync.WaitGroup

	// Launch readers
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				cache.Get("key")
			}
		}()
	}

	// Launch writers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				cache.Set("key", string(rune('A'+id)))
			}
		}(i)
	}

	wg.Wait()

	// Verify cache still works
	_, ok := cache.Get("key")
	if !ok {
		t.Error("Expected key to exist after concurrent operations")
	}
}

// TestBankAccount verifies bank account operations are thread-safe
func TestBankAccount(t *testing.T) {
	account := &BankAccount{balance: 1000}
	var wg sync.WaitGroup

	// Concurrent deposits
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.Deposit(10)
		}()
	}

	wg.Wait()

	expected := 1000 + (100 * 10) // Initial + 100 deposits of $10
	if account.Balance() != expected {
		t.Errorf("Expected balance %d, got %d", expected, account.Balance())
	}
}

// TestBankAccountWithdraw verifies withdrawal logic is correct
func TestBankAccountWithdraw(t *testing.T) {
	tests := []struct {
		name           string
		initialBalance int
		withdrawAmount int
		shouldSucceed  bool
		finalBalance   int
	}{
		{"sufficient funds", 100, 50, true, 50},
		{"exact funds", 100, 100, true, 0},
		{"insufficient funds", 100, 150, false, 100},
		{"zero withdrawal", 100, 0, true, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := &BankAccount{balance: tt.initialBalance}
			success := account.Withdraw(tt.withdrawAmount)

			if success != tt.shouldSucceed {
				t.Errorf("Expected success=%v, got %v", tt.shouldSucceed, success)
			}

			if account.Balance() != tt.finalBalance {
				t.Errorf("Expected balance %d, got %d", tt.finalBalance, account.Balance())
			}
		})
	}
}

// TestBankAccountConcurrentOperations verifies mixed operations are safe
func TestBankAccountConcurrentOperations(t *testing.T) {
	account := &BankAccount{balance: 10000}
	var wg sync.WaitGroup

	// Track successful withdrawals
	successfulWithdrawals := &SafeCounter{}

	// Concurrent deposits
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.Deposit(50)
		}()
	}

	// Concurrent withdrawals
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if account.Withdraw(30) {
				successfulWithdrawals.Increment()
			}
		}()
	}

	wg.Wait()

	// Verify accounting is consistent
	// Starting: 10000
	// Deposits: 100 * 50 = +5000
	// Withdrawals: successfulWithdrawals * 30
	expected := 10000 + (100 * 50) - (successfulWithdrawals.Value() * 30)
	actual := account.Balance()

	if actual != expected {
		t.Errorf("Balance inconsistent: expected %d, got %d", expected, actual)
	}

	// Balance should never be negative
	if actual < 0 {
		t.Errorf("Balance is negative: %d", actual)
	}
}

// BenchmarkSafeCounterIncrement measures mutex overhead
func BenchmarkSafeCounterIncrement(b *testing.B) {
	counter := SafeCounter{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Increment()
		}
	})
}

// BenchmarkSafeCounterValue measures read performance
func BenchmarkSafeCounterValue(b *testing.B) {
	counter := SafeCounter{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Value()
		}
	})
}

// BenchmarkCacheGet measures cache read performance with RWMutex
func BenchmarkCacheGet(b *testing.B) {
	cache := NewCache()
	cache.Set("key", "value")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.Get("key")
		}
	})
}

// BenchmarkCacheSet measures cache write performance with RWMutex
func BenchmarkCacheSet(b *testing.B) {
	cache := NewCache()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.Set("key", "value")
		}
	})
}

// BenchmarkCacheMixed measures mixed read/write performance (90% reads)
func BenchmarkCacheMixed(b *testing.B) {
	cache := NewCache()
	cache.Set("key", "value")

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%10 == 0 {
				cache.Set("key", "value")
			} else {
				cache.Get("key")
			}
			i++
		}
	})
}
