package main

import (
	"context"
	"testing"
	"time"
)

func TestFetchDataSuccess(t *testing.T) {
	ctx := context.Background()
	data, err := fetchData(ctx, "http://example.com")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if data == "" {
		t.Error("Expected data, got empty string")
	}
}

func TestFetchDataWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	_, err := fetchData(ctx, "http://example.com")

	if err != context.DeadlineExceeded {
		t.Errorf("Expected DeadlineExceeded, got %v", err)
	}
}

func TestFetchDataWithCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// Cancel immediately
	cancel()

	_, err := fetchData(ctx, "http://example.com")

	if err != context.Canceled {
		t.Errorf("Expected Canceled error, got %v", err)
	}
}

func TestContextValues(t *testing.T) {
	ctx := context.WithValue(context.Background(), userIDKey, 999)

	userID := ctx.Value(userIDKey)
	if userID != 999 {
		t.Errorf("Expected userID 999, got %v", userID)
	}

	// Non-existent key returns nil
	missing := ctx.Value(contextKey("nonexistent"))
	if missing != nil {
		t.Errorf("Expected nil for non-existent key, got %v", missing)
	}
}

func TestContextDeadline(t *testing.T) {
	deadline := time.Now().Add(100 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// Check deadline
	dl, ok := ctx.Deadline()
	if !ok {
		t.Error("Expected deadline to be set")
	}

	if !dl.Equal(deadline) {
		t.Errorf("Expected deadline %v, got %v", deadline, dl)
	}
}

func BenchmarkContextCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_ = ctx
	}
}
