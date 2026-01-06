package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Create handler
	finalHandler := http.HandlerFunc(handleRequest)

	// Chain middleware
	handler := loggingMiddleware(
		authMiddleware(
			recoveryMiddleware(finalHandler),
		),
	)

	fmt.Println("Server starting on :8080")
	fmt.Println("Try: curl http://localhost:8080/")
	fmt.Println("Try with auth: curl -H 'Authorization: Bearer token123' http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! Your request was processed successfully.\n")
}

// loggingMiddleware logs request details
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		fmt.Printf("[%s] %s %s\n", start.Format("15:04:05"), r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		fmt.Printf("[%s] Request completed in %v\n", time.Now().Format("15:04:05"), duration)
	})
}

// authMiddleware checks authorization
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			fmt.Println("No auth token provided")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		fmt.Println("Auth token validated:", token)
		next.ServeHTTP(w, r)
	})
}

// recoveryMiddleware recovers from panics
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Recovered from panic: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// Middleware builder pattern
type Middleware func(http.Handler) http.Handler

func Chain(middlewares ...Middleware) Middleware {
	return func(final http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}
