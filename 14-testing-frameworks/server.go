package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Server provides HTTP API endpoints
type Server struct {
	users   map[int]*User
	nextID  int
	mu      sync.RWMutex
	handler http.Handler
}

// NewServer creates a new Server instance
func NewServer() *Server {
	s := &Server{
		users:  make(map[int]*User),
		nextID: 1,
	}
	s.setupRoutes()
	return s
}

// setupRoutes configures HTTP routes
func (s *Server) setupRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", s.handleUsers)
	mux.HandleFunc("/users/", s.handleUser)
	mux.HandleFunc("/health", s.handleHealth)
	s.handler = mux
}

// ServeHTTP implements http.Handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

// handleHealth returns health status
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

// handleUsers handles requests to /users
func (s *Server) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.listUsers(w, r)
	case http.MethodPost:
		s.createUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleUser handles requests to /users/{id}
func (s *Server) handleUser(w http.ResponseWriter, r *http.Request) {
	// Extract ID from path
	idStr := r.URL.Path[len("/users/"):]
	if idStr == "" {
		http.Error(w, "User ID required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		s.getUser(w, r, id)
	case http.MethodPut:
		s.updateUser(w, r, id)
	case http.MethodDelete:
		s.deleteUser(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// listUsers returns all users
func (s *Server) listUsers(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]*User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// createUser creates a new user
func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate
	if user.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	if user.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	user.ID = s.nextID
	s.nextID++
	s.users[user.ID] = &user
	s.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// getUser returns a single user by ID
func (s *Server) getUser(w http.ResponseWriter, r *http.Request, id int) {
	s.mu.RLock()
	user, exists := s.users[id]
	s.mu.RUnlock()

	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// updateUser updates an existing user
func (s *Server) updateUser(w http.ResponseWriter, r *http.Request, id int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user.ID = id
	s.users[id] = &user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// deleteUser deletes a user by ID
func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	delete(s.users, id)

	w.WriteHeader(http.StatusNoContent)
}

// Start starts the HTTP server
func (s *Server) Start(port int) error {
	addr := fmt.Sprintf(":%d", port)
	return http.ListenAndServe(addr, s)
}
