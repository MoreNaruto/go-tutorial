package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

// TestServerHealthWithHttpexpect demonstrates basic HTTP testing
func TestServerHealthWithHttpexpect(t *testing.T) {
	// Create server instance
	server := NewServer()

	// Create test server
	ts := httptest.NewServer(server)
	defer ts.Close()

	// Create httpexpect instance
	e := httpexpect.Default(t, ts.URL)

	// Test health endpoint
	e.GET("/health").
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("status").
		ValueEqual("status", "ok")
}

// TestServerCreateUser demonstrates POST request testing
func TestServerCreateUser(t *testing.T) {
	server := NewServer()
	ts := httptest.NewServer(server)
	defer ts.Close()

	e := httpexpect.Default(t, ts.URL)

	// Create a new user
	user := map[string]interface{}{
		"name":  "Alice",
		"email": "alice@example.com",
	}

	obj := e.POST("/users").
		WithJSON(user).
		Expect().
		Status(http.StatusCreated).
		JSON().Object()

	// Verify response
	obj.ContainsKey("id")
	obj.ContainsKey("name")
	obj.ContainsKey("email")
	obj.ValueEqual("name", "Alice")
	obj.ValueEqual("email", "alice@example.com")

	// Get the created user ID
	userID := obj.Value("id").Number().Raw()
	t.Logf("Created user with ID: %v", userID)
}

// TestServerGetUser demonstrates GET request with path parameter
func TestServerGetUser(t *testing.T) {
	server := NewServer()
	ts := httptest.NewServer(server)
	defer ts.Close()

	e := httpexpect.Default(t, ts.URL)

	// Create a user first
	createResp := e.POST("/users").
		WithJSON(map[string]string{
			"name":  "Bob",
			"email": "bob@example.com",
		}).
		Expect().
		Status(http.StatusCreated).
		JSON().Object()

	userID := createResp.Value("id").Number().Raw()

	// Get the user
	e.GET("/users/{id}", userID).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("id", userID).
		ValueEqual("name", "Bob").
		ValueEqual("email", "bob@example.com")
}

// TestServerListUsers demonstrates GET request returning array
func TestServerListUsers(t *testing.T) {
	server := NewServer()
	ts := httptest.NewServer(server)
	defer ts.Close()

	e := httpexpect.Default(t, ts.URL)

	// Initially empty
	e.GET("/users").
		Expect().
		Status(http.StatusOK).
		JSON().Array().
		Length().Equal(0)

	// Create some users
	users := []map[string]string{
		{"name": "Alice", "email": "alice@example.com"},
		{"name": "Bob", "email": "bob@example.com"},
		{"name": "Charlie", "email": "charlie@example.com"},
	}

	for _, user := range users {
		e.POST("/users").
			WithJSON(user).
			Expect().
			Status(http.StatusCreated)
	}

	// List should now have 3 users
	arr := e.GET("/users").
		Expect().
		Status(http.StatusOK).
		JSON().Array()

	arr.Length().Equal(3)

	// Verify first user
	arr.Element(0).Object().
		ContainsKey("id").
		ContainsKey("name").
		ContainsKey("email")
}

// TestServerUpdateUser demonstrates PUT request
func TestServerUpdateUser(t *testing.T) {
	server := NewServer()
	ts := httptest.NewServer(server)
	defer ts.Close()

	e := httpexpect.Default(t, ts.URL)

	// Create a user
	createResp := e.POST("/users").
		WithJSON(map[string]string{
			"name":  "Alice",
			"email": "alice@example.com",
		}).
		Expect().
		Status(http.StatusCreated).
		JSON().Object()

	userID := createResp.Value("id").Number().Raw()

	// Update the user
	e.PUT("/users/{id}", userID).
		WithJSON(map[string]string{
			"name":  "Alice Smith",
			"email": "alice.smith@example.com",
		}).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("name", "Alice Smith").
		ValueEqual("email", "alice.smith@example.com")

	// Verify update persisted
	e.GET("/users/{id}", userID).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("name", "Alice Smith").
		ValueEqual("email", "alice.smith@example.com")
}

// TestServerDeleteUser demonstrates DELETE request
func TestServerDeleteUser(t *testing.T) {
	server := NewServer()
	ts := httptest.NewServer(server)
	defer ts.Close()

	e := httpexpect.Default(t, ts.URL)

	// Create a user
	createResp := e.POST("/users").
		WithJSON(map[string]string{
			"name":  "Bob",
			"email": "bob@example.com",
		}).
		Expect().
		Status(http.StatusCreated).
		JSON().Object()

	userID := createResp.Value("id").Number().Raw()

	// Delete the user
	e.DELETE("/users/{id}", userID).
		Expect().
		Status(http.StatusNoContent)

	// Verify user is gone
	e.GET("/users/{id}", userID).
		Expect().
		Status(http.StatusNotFound)
}

// TestServerErrorHandling demonstrates error response testing
func TestServerErrorHandling(t *testing.T) {
	server := NewServer()
	ts := httptest.NewServer(server)
	defer ts.Close()

	e := httpexpect.Default(t, ts.URL)

	t.Run("Get non-existent user", func(t *testing.T) {
		e.GET("/users/999").
			Expect().
			Status(http.StatusNotFound).
			Body().Contains("not found")
	})

	t.Run("Create user with missing name", func(t *testing.T) {
		e.POST("/users").
			WithJSON(map[string]string{
				"email": "test@example.com",
			}).
			Expect().
			Status(http.StatusBadRequest).
			Body().Contains("Name is required")
	})

	t.Run("Create user with missing email", func(t *testing.T) {
		e.POST("/users").
			WithJSON(map[string]string{
				"name": "Test User",
			}).
			Expect().
			Status(http.StatusBadRequest).
			Body().Contains("Email is required")
	})

	t.Run("Invalid user ID", func(t *testing.T) {
		e.GET("/users/invalid").
			Expect().
			Status(http.StatusBadRequest).
			Body().Contains("Invalid user ID")
	})

	t.Run("Update non-existent user", func(t *testing.T) {
		e.PUT("/users/999").
			WithJSON(map[string]string{
				"name":  "Test",
				"email": "test@example.com",
			}).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("Delete non-existent user", func(t *testing.T) {
		e.DELETE("/users/999").
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("Method not allowed", func(t *testing.T) {
		e.PATCH("/users").
			Expect().
			Status(http.StatusMethodNotAllowed)
	})
}

// TestServerFullWorkflow demonstrates a complete CRUD workflow
func TestServerFullWorkflow(t *testing.T) {
	server := NewServer()
	ts := httptest.NewServer(server)
	defer ts.Close()

	e := httpexpect.Default(t, ts.URL)

	// Step 1: Verify empty list
	e.GET("/users").
		Expect().
		Status(http.StatusOK).
		JSON().Array().
		Length().Equal(0)

	// Step 2: Create first user
	user1 := e.POST("/users").
		WithJSON(map[string]string{
			"name":  "Alice",
			"email": "alice@example.com",
		}).
		Expect().
		Status(http.StatusCreated).
		JSON().Object()

	user1ID := user1.Value("id").Number().Raw()

	// Step 3: Create second user
	user2 := e.POST("/users").
		WithJSON(map[string]string{
			"name":  "Bob",
			"email": "bob@example.com",
		}).
		Expect().
		Status(http.StatusCreated).
		JSON().Object()

	user2ID := user2.Value("id").Number().Raw()

	// Step 4: List should have 2 users
	e.GET("/users").
		Expect().
		Status(http.StatusOK).
		JSON().Array().
		Length().Equal(2)

	// Step 5: Get individual users
	e.GET("/users/{id}", user1ID).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("name", "Alice")

	e.GET("/users/{id}", user2ID).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("name", "Bob")

	// Step 6: Update first user
	e.PUT("/users/{id}", user1ID).
		WithJSON(map[string]string{
			"name":  "Alice Updated",
			"email": "alice.updated@example.com",
		}).
		Expect().
		Status(http.StatusOK)

	// Step 7: Verify update
	e.GET("/users/{id}", user1ID).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("name", "Alice Updated")

	// Step 8: Delete second user
	e.DELETE("/users/{id}", user2ID).
		Expect().
		Status(http.StatusNoContent)

	// Step 9: Verify deletion
	e.GET("/users/{id}", user2ID).
		Expect().
		Status(http.StatusNotFound)

	// Step 10: List should have 1 user
	e.GET("/users").
		Expect().
		Status(http.StatusOK).
		JSON().Array().
		Length().Equal(1)
}

// TestServerWithCustomConfiguration demonstrates custom httpexpect configuration
func TestServerWithCustomConfiguration(t *testing.T) {
	server := NewServer()
	ts := httptest.NewServer(server)
	defer ts.Close()

	// Create httpexpect with custom configuration
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  ts.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	// Use the configured instance
	e.GET("/health").
		Expect().
		Status(http.StatusOK).
		ContentType("application/json")
}

// TestServerHeaders demonstrates header testing
func TestServerHeaders(t *testing.T) {
	server := NewServer()
	ts := httptest.NewServer(server)
	defer ts.Close()

	e := httpexpect.Default(t, ts.URL)

	// Test response headers
	e.GET("/health").
		Expect().
		Status(http.StatusOK).
		Header("Content-Type").Equal("application/json")
}

// TestServerResponseMatchers demonstrates various response matchers
func TestServerResponseMatchers(t *testing.T) {
	server := NewServer()
	ts := httptest.NewServer(server)
	defer ts.Close()

	e := httpexpect.Default(t, ts.URL)

	// Create a user for testing
	obj := e.POST("/users").
		WithJSON(map[string]string{
			"name":  "Charlie",
			"email": "charlie@example.com",
		}).
		Expect().
		Status(http.StatusCreated).
		JSON().Object()

	// Demonstrate various matchers
	obj.Keys().ContainsOnly("id", "name", "email")
	obj.Value("id").Number().Gt(0)
	obj.Value("name").String().NotEmpty()
	obj.Value("email").String().Contains("@")
}
