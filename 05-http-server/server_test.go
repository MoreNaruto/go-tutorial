package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	homeHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "text/html" {
		t.Errorf("Expected Content-Type text/html, got %s", contentType)
	}

	body := w.Body.String()
	if !strings.Contains(body, "Welcome") {
		t.Error("Expected body to contain 'Welcome'")
	}
}

func TestHomeHandler404(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/notfound", nil)
	w := httptest.NewRecorder()

	homeHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected string
	}{
		{"with name", "?name=Alice", "Hello, Alice!"},
		{"without name", "", "Hello, Guest!"},
		{"with empty name", "?name=", "Hello, Guest!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/hello"+tt.query, nil)
			w := httptest.NewRecorder()

			helloHandler(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status 200, got %d", w.Code)
			}

			body := w.Body.String()
			if body != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, body)
			}
		})
	}
}

func TestEchoHandler(t *testing.T) {
	body := "test message"
	req := httptest.NewRequest(http.MethodPost, "/echo", strings.NewReader(body))
	w := httptest.NewRecorder()

	echoHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	response := w.Body.String()
	if !strings.Contains(response, body) {
		t.Errorf("Expected response to contain %q, got %q", body, response)
	}
}

func TestEchoHandlerMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/echo", nil)
	w := httptest.NewRecorder()

	echoHandler(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", w.Code)
	}
}

func TestJSONHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/json", nil)
	w := httptest.NewRecorder()

	jsonHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}

	var users []User
	if err := json.NewDecoder(w.Body).Decode(&users); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	if len(users) != 3 {
		t.Errorf("Expected 3 users, got %d", len(users))
	}

	if users[0].Username != "alice" {
		t.Errorf("Expected first user to be alice, got %s", users[0].Username)
	}
}

func TestUserHandler(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		wantStatus int
		wantID     int
	}{
		{"valid user", "/users/123", http.StatusOK, 123},
		{"another valid user", "/users/456", http.StatusOK, 456},
		{"invalid ID", "/users/abc", http.StatusBadRequest, 0},
		{"missing ID", "/users/", http.StatusBadRequest, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			userHandler(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var user User
				if err := json.NewDecoder(w.Body).Decode(&user); err != nil {
					t.Fatalf("Failed to decode JSON: %v", err)
				}

				if user.ID != tt.wantID {
					t.Errorf("Expected user ID %d, got %d", tt.wantID, user.ID)
				}
			}
		})
	}
}

func BenchmarkJSONHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/json", nil)

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		jsonHandler(w, req)
	}
}
