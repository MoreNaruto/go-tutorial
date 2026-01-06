package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Context Patterns Tutorial ===")
	fmt.Println()

	// Basic context
	demonstrateBasicContext()

	// Context with timeout
	demonstrateTimeout()

	// Context with deadline
	demonstrateDeadline()

	// Context with cancellation
	demonstrateCancellation()

	// Context with values
	demonstrateValues()
}

func demonstrateBasicContext() {
	fmt.Println("--- Basic Context ---")

	// Background context - parent of all contexts
	ctx := context.Background()
	fmt.Printf("Background context: %v\n", ctx)

	// TODO context - placeholder
	todoCtx := context.TODO()
	fmt.Printf("TODO context: %v\n\n", todoCtx)
}

func doWork(ctx context.Context, name string, duration time.Duration) {
	select {
	case <-time.After(duration):
		fmt.Printf("%s: completed work\n", name)
	case <-ctx.Done():
		fmt.Printf("%s: cancelled - %v\n", name, ctx.Err())
	}
}

func demonstrateTimeout() {
	fmt.Println("--- Context with Timeout ---")

	// Create context that cancels after 100ms
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel() // Always call cancel to release resources

	// This will complete (50ms < 100ms timeout)
	doWork(ctx, "Fast task", 50*time.Millisecond)

	// This will be cancelled (200ms > 100ms timeout)
	doWork(ctx, "Slow task", 200*time.Millisecond)

	fmt.Println()
}

func demonstrateDeadline() {
	fmt.Println("--- Context with Deadline ---")

	// Create context that cancels at specific time
	deadline := time.Now().Add(150 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fmt.Printf("Deadline set for: %v\n", deadline.Format("15:04:05.000"))

	doWork(ctx, "Task with deadline", 200*time.Millisecond)

	fmt.Println()
}

func demonstrateCancellation() {
	fmt.Println("--- Manual Cancellation ---")

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Manually cancelling context")
		cancel() // Trigger cancellation
	}()

	doWork(ctx, "Cancellable task", 300*time.Millisecond)

	fmt.Println()
}

type contextKey string

const (
	userIDKey    contextKey = "userID"
	requestIDKey contextKey = "requestID"
)

func demonstrateValues() {
	fmt.Println("--- Context with Values ---")

	// Add values to context
	ctx := context.WithValue(context.Background(), userIDKey, 12345)
	ctx = context.WithValue(ctx, requestIDKey, "req-abc-123")

	// Retrieve values
	processRequest(ctx)

	fmt.Println()
}

func processRequest(ctx context.Context) {
	userID := ctx.Value(userIDKey)
	requestID := ctx.Value(requestIDKey)

	fmt.Printf("Processing request: %v for user: %v\n", requestID, userID)

	// Pass context down the call chain
	authenticateUser(ctx)
}

func authenticateUser(ctx context.Context) {
	userID := ctx.Value(userIDKey)
	fmt.Printf("Authenticating user: %v\n", userID)
}

// Simulated API call with context
func fetchData(ctx context.Context, url string) (string, error) {
	// Create a channel to receive result
	result := make(chan string, 1)
	errChan := make(chan error, 1)

	go func() {
		// Simulate slow API call
		time.Sleep(100 * time.Millisecond)
		result <- fmt.Sprintf("Data from %s", url)
	}()

	// Wait for result or cancellation
	select {
	case data := <-result:
		return data, nil
	case err := <-errChan:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
