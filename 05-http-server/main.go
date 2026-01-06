package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	// Register handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/users/", userHandler) // Trailing slash for path matching

	// Static file server
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	fmt.Println("Try:")
	fmt.Println("  GET  http://localhost:8080/")
	fmt.Println("  GET  http://localhost:8080/hello?name=Alice")
	fmt.Println("  POST http://localhost:8080/echo")
	fmt.Println("  GET  http://localhost:8080/json")
	fmt.Println("  GET  http://localhost:8080/users/123")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// homeHandler handles requests to the root path
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	html := `
		<html>
		<head><title>Go HTTP Server</title></head>
		<body>
			<h1>Welcome to Go HTTP Server Tutorial</h1>
			<ul>
				<li><a href="/hello?name=World">Hello endpoint</a></li>
				<li><a href="/json">JSON endpoint</a></li>
				<li><a href="/users/123">User endpoint</a></li>
			</ul>
		</body>
		</html>
	`
	fmt.Fprint(w, html)
}

// helloHandler demonstrates query parameters
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

// echoHandler demonstrates reading request body
func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	body := make([]byte, r.ContentLength)
	_, err := r.Body.Read(body)
	if err != nil && err.Error() != "EOF" {
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Echo back
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "You sent: %s", string(body))
}

// User represents a user in our system
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// jsonHandler demonstrates JSON responses
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Username: "alice", Email: "alice@example.com"},
		{ID: 2, Username: "bob", Email: "bob@example.com"},
		{ID: 3, Username: "carol", Email: "carol@example.com"},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// userHandler demonstrates path parameters
func userHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from path: /users/{id}
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	if path == "" {
		http.Error(w, "User ID required", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Mock user lookup
	user := User{
		ID:       userID,
		Username: fmt.Sprintf("user_%d", userID),
		Email:    fmt.Sprintf("user_%d@example.com", userID),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
