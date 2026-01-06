# 05 - HTTP Server

**Level:** Intermediate

## What This Project Teaches

Building HTTP servers with Go's standard library:
- HTTP handlers and routing
- Query parameters and path parameters
- Request/response handling
- JSON encoding/decoding
- HTTP testing with `httptest`
- Static file serving

## Why This Structure?

Simple **flat structure** for an HTTP server:
- `main.go` - HTTP server with multiple endpoints
- `server_test.go` - Handler tests using `httptest`
- `static/` - Directory for static files

Go's `net/http` package provides everything needed for production HTTP servers.

## Key Concepts Explained

### Basic HTTP Server

```go
http.HandleFunc("/", handler)
http.ListenAndServe(":8080", nil)
```

### Handler Function

Handlers have signature: `func(w http.ResponseWriter, r *http.Request)`

```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}
```

### Query Parameters

```go
name := r.URL.Query().Get("name")
```

Access: `http://localhost:8080/hello?name=Alice`

### Path Parameters

Extract from URL path:

```go
path := strings.TrimPrefix(r.URL.Path, "/users/")
userID, _ := strconv.Atoi(path)
```

Access: `http://localhost:8080/users/123`

### JSON Responses

```go
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(data)
```

### Error Responses

```go
http.Error(w, "Not Found", http.StatusNotFound)
```

### Static Files

```go
fs := http.FileServer(http.Dir("./static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

## Running the Server

Start the server:

```bash
go run main.go
```

Server runs on `http://localhost:8080`

Test endpoints:

```bash
curl http://localhost:8080/
curl http://localhost:8080/hello?name=Alice
curl -X POST -d "test message" http://localhost:8080/echo
curl http://localhost:8080/json
curl http://localhost:8080/users/123
```

## Running Tests

```bash
go test ./...
go test -v
go test -cover
```

## What You'll Learn

1. Building HTTP servers with standard library
2. Routing and handling different endpoints
3. Processing query and path parameters
4. Working with JSON
5. Testing HTTP handlers
6. Serving static files

## HTTP Testing

Use `httptest` for testing without starting a server:

```go
req := httptest.NewRequest(http.MethodGet, "/path", nil)
w := httptest.NewRecorder()

handler(w, req)

// Assert on w.Code, w.Body, w.Header
```

## Best Practices

1. **Set Content-Type:** Always set appropriate headers
2. **Validate input:** Check method, parameters, body
3. **Handle errors:** Return appropriate status codes
4. **Close bodies:** Use `defer r.Body.Close()`
5. **Use constants:** `http.StatusOK` instead of `200`
6. **Test handlers:** Use `httptest` for unit tests

## Common HTTP Status Codes

- `200 OK` - Success
- `201 Created` - Resource created
- `400 Bad Request` - Invalid input
- `404 Not Found` - Resource not found
- `405 Method Not Allowed` - Wrong HTTP method
- `500 Internal Server Error` - Server error

## Next Steps

- **06-rest-api-gin** - Full REST API with framework
- **11-middleware-chain** - Middleware patterns
- **07-context-patterns** - Context in HTTP handlers

## Real-World Usage

Production considerations:
- Structured routing (Chi, Gorilla Mux)
- Middleware (logging, auth, recovery)
- Graceful shutdown
- TLS/HTTPS
- Rate limiting
- Request validation
