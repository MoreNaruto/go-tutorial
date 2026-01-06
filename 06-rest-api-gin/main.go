package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Book represents a book in our library
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required,min=1000,max=2100"`
}

// In-memory storage
var books = []Book{
	{ID: 1, Title: "The Go Programming Language", Author: "Donovan & Kernighan", Year: 2015},
	{ID: 2, Title: "Clean Code", Author: "Robert Martin", Year: 2008},
}
var nextID = 3

func main() {
	router := setupRouter()
	router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Routes
	router.GET("/", homeHandler)
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.POST("/books", createBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)

	return router
}

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Book API",
		"endpoints": gin.H{
			"GET /books":      "List all books",
			"GET /books/:id":  "Get book by ID",
			"POST /books":     "Create new book",
			"PUT /books/:id":  "Update book",
			"DELETE /books/:id": "Delete book",
		},
	})
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func createBook(c *gin.Context) {
	var newBook Book

	// Bind and validate
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook.ID = nextID
	nextID++

	books = append(books, newBook)

	c.JSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
