# 06 - REST API with Gin

**Level:** Intermediate

## What This Project Teaches

Building a REST API with the Gin framework:
- Gin router and HTTP methods
- CRUD operations (Create, Read, Update, Delete)
- JSON binding and validation
- Path parameters
- Error handling in APIs
- Testing Gin handlers

## Why This Structure?

**Flat structure** for focused API learning:
- `main.go` - Complete REST API with all CRUD operations
- `main_test.go` - Comprehensive API tests
- `go.mod` - Gin dependency management

Demonstrates real-world API patterns with a popular Go framework.

## Key Concepts

### Gin Setup

```go
router := gin.Default()  // With logger and recovery middleware
router.GET("/path", handler)
router.POST("/path", handler)
router.Run(":8080")
```

### Handler Function

```go
func handler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "success"})
}
```

### Path Parameters

```go
router.GET("/books/:id", getBook)

// In handler
id := c.Param("id")
```

### JSON Binding & Validation

```go
type Book struct {
    Title  string `json:"title" binding:"required"`
    Author string `json:"author" binding:"required"`
    Year   int    `json:"year" binding:"required,min=1000,max=2100"`
}

var book Book
if err := c.ShouldBindJSON(&book); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}
```

### Response Formats

```go
// Success
c.JSON(http.StatusOK, data)

// Created
c.JSON(http.StatusCreated, newResource)

// Error
c.JSON(http.StatusBadRequest, gin.H{"error": "message"})
```

## Running the API

Install dependencies:

```bash
go mod download
```

Run the server:

```bash
go run main.go
```

## Testing the API

With curl:

```bash
# List all books
curl http://localhost:8080/books

# Get specific book
curl http://localhost:8080/books/1

# Create book
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title":"New Book","author":"Author Name","year":2024}'

# Update book
curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Title","author":"Author","year":2024}'

# Delete book
curl -X DELETE http://localhost:8080/books/1
```

## Running Tests

```bash
go test ./...
go test -v
go test -cover
```

## What You'll Learn

1. Building REST APIs with Gin
2. Implementing CRUD operations
3. Request validation
4. Error handling in APIs
5. Testing HTTP APIs
6. JSON serialization

## REST Principles

### HTTP Methods

- **GET** - Retrieve resources (idempotent, safe)
- **POST** - Create resources
- **PUT** - Update entire resource (idempotent)
- **PATCH** - Partial update
- **DELETE** - Remove resource (idempotent)

### Status Codes

- **2xx Success**
  - 200 OK - General success
  - 201 Created - Resource created
  - 204 No Content - Success with no body

- **4xx Client Errors**
  - 400 Bad Request - Invalid input
  - 404 Not Found - Resource doesn't exist
  - 422 Unprocessable Entity - Validation failed

- **5xx Server Errors**
  - 500 Internal Server Error

### Resource Naming

- Use nouns, not verbs: `/books` not `/getBooks`
- Plural for collections: `/books`
- ID for specific resource: `/books/123`
- Nested resources: `/authors/1/books`

## Gin Features

### Middleware

```go
router.Use(gin.Logger())
router.Use(gin.Recovery())
router.Use(customMiddleware())
```

### Route Groups

```go
api := router.Group("/api/v1")
{
    api.GET("/books", getBooks)
    api.POST("/books", createBook)
}
```

### Query Parameters

```go
search := c.Query("search")
limit := c.DefaultQuery("limit", "10")
```

## Best Practices

1. **Validate input:** Use binding tags
2. **Consistent responses:** Standard error format
3. **Proper status codes:** Match HTTP semantics
4. **Versioning:** `/api/v1/` prefix
5. **Documentation:** Document endpoints
6. **Testing:** Test all CRUD operations

## Next Steps

- **09-clean-architecture** - Layered API architecture
- **11-middleware-chain** - Custom middleware
- **08-testing-strategies** - Advanced testing patterns

## Production Considerations

- Database integration (PostgreSQL, MongoDB)
- Authentication & authorization
- Rate limiting
- CORS configuration
- Logging and monitoring
- Graceful shutdown
- Environment configuration
- API documentation (Swagger)
