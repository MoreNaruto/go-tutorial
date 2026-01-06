package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetBooks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var responseBooks []Book
	json.Unmarshal(w.Body.Bytes(), &responseBooks)

	if len(responseBooks) < 2 {
		t.Errorf("Expected at least 2 books, got %d", len(responseBooks))
	}
}

func TestGetBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var book Book
	json.Unmarshal(w.Body.Bytes(), &book)

	if book.ID != 1 {
		t.Errorf("Expected book ID 1, got %d", book.ID)
	}
}

func TestGetBookNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/999", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestCreateBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	newBook := Book{
		Title:  "Test Book",
		Author: "Test Author",
		Year:   2024,
	}

	jsonData, _ := json.Marshal(newBook)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	var createdBook Book
	json.Unmarshal(w.Body.Bytes(), &createdBook)

	if createdBook.Title != newBook.Title {
		t.Errorf("Expected title %s, got %s", newBook.Title, createdBook.Title)
	}
	if createdBook.ID == 0 {
		t.Error("Expected book to have an ID")
	}
}

func TestCreateBookValidation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	invalidBook := map[string]interface{}{
		"title":  "",
		"author": "Author",
		"year":   2024,
	}

	jsonData, _ := json.Marshal(invalidBook)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestUpdateBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	updatedBook := Book{
		Title:  "Updated Title",
		Author: "Updated Author",
		Year:   2023,
	}

	jsonData, _ := json.Marshal(updatedBook)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/books/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var book Book
	json.Unmarshal(w.Body.Bytes(), &book)

	if book.Title != updatedBook.Title {
		t.Errorf("Expected title %s, got %s", updatedBook.Title, book.Title)
	}
}

func TestDeleteBook(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Reset books for this test
	books = []Book{
		{ID: 1, Title: "Book 1", Author: "Author 1", Year: 2020},
		{ID: 2, Title: "Book 2", Author: "Author 2", Year: 2021},
	}

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/books/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Verify book is deleted
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/books/1", nil)
	router.ServeHTTP(w2, req2)

	if w2.Code != http.StatusNotFound {
		t.Errorf("Expected book to be deleted, but got status %d", w2.Code)
	}
}
