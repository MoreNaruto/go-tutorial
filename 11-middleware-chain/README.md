# 11 - Middleware Chain

**Level:** Advanced

## What This Project Teaches

HTTP middleware patterns in Go:
- Middleware function signature
- Chaining middleware
- Common middleware (logging, auth, recovery)
- Middleware composition
- Request/response wrapping

## Running

```bash
go run main.go
go test ./...
```

Test endpoints:
```bash
curl http://localhost:8080/
curl -H 'Authorization: Bearer token123' http://localhost:8080/
```

## Key Concepts

### Middleware Signature

```go
type Middleware func(http.Handler) http.Handler
```

### Chaining

```go
handler := middleware1(middleware2(middleware3(finalHandler)))
```

### Common Patterns

- **Logging:** Request/response details
- **Authentication:** Token validation
- **Recovery:** Panic handling
- **CORS:** Cross-origin requests
- **Rate Limiting:** Request throttling

## Production Usage

Middleware is essential in production HTTP servers for:
- Request logging and monitoring
- Authentication and authorization
- Error recovery and handling
- Performance tracking
- Security headers

This demonstrates patterns used in frameworks like Chi, Echo, and Gin.
