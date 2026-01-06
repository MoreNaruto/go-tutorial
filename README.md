# Go Tutorial Repository

A comprehensive Go (Golang) tutorial repository composed of **multiple isolated sub-projects**, each demonstrating specific core concepts and best practices.

## 🎯 Philosophy

This repository is a **collection of standalone Go projects**, not a single monolithic application. Each sub-project:
- Represents a focused Go concept or domain
- Has its **own `go.mod`** and is independently runnable
- Uses an appropriate structure for its scope
- Includes comprehensive tests and documentation
- Demonstrates idiomatic, production-ready Go code

## 📚 Tutorial Projects

### Beginner Level

| Project | Focus | Key Concepts |
|---------|-------|--------------|
| **[01-hello-basics](01-hello-basics/)** | Language Fundamentals | Variables, functions, control flow, basic types |
| **[02-structs-interfaces](02-structs-interfaces/)** | Type System | Structs, methods, interfaces, composition |
| **[03-error-handling](03-error-handling/)** | Error Management | Error patterns, custom errors, wrapping |
| **[04-basic-concurrency](04-basic-concurrency/)** | Concurrency Basics | Goroutines, channels, select statements |

### Intermediate Level

| Project | Focus | Key Concepts |
|---------|-------|--------------|
| **[05-http-server](05-http-server/)** | HTTP Basics | Standard library HTTP server, routing, handlers |
| **[06-rest-api-gin](06-rest-api-gin/)** | REST APIs | Gin framework, CRUD operations, validation |
| **[07-context-patterns](07-context-patterns/)** | Context Usage | Context propagation, cancellation, timeouts |
| **[08-testing-strategies](08-testing-strategies/)** | Testing | Table-driven tests, mocking, test organization |

### Advanced Level

| Project | Focus | Key Concepts |
|---------|-------|--------------|
| **[09-clean-architecture](09-clean-architecture/)** | Architecture | Hexagonal/Clean architecture, dependency injection |
| **[10-websockets](10-websockets/)** | Real-time Communication | WebSocket server, broadcasting, connection management |
| **[11-middleware-chain](11-middleware-chain/)** | HTTP Middleware | Custom middleware, logging, authentication, recovery |
| **[12-concurrency-patterns](12-concurrency-patterns/)** | Advanced Concurrency | Worker pools, fan-in/fan-out, pipelines |

## 🚀 Getting Started

Each sub-project is completely independent. Navigate to any project directory and follow its README:

```bash
cd 01-hello-basics
go run main.go
go test ./...
```

## 📖 Learning Path

**New to Go?** Start with:
1. 01-hello-basics
2. 02-structs-interfaces
3. 03-error-handling

**Have Go experience?** Jump to intermediate or advanced projects based on your needs.

**Want to build web services?** Check out:
- 05-http-server (basics)
- 06-rest-api-gin (framework)
- 11-middleware-chain (patterns)

**Interested in concurrency?** Explore:
- 04-basic-concurrency (fundamentals)
- 12-concurrency-patterns (advanced patterns)

## 🧪 Testing

Every sub-project includes comprehensive unit tests. Run tests for any project:

```bash
cd <project-directory>
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## 🎓 Best Practices Demonstrated

- Idiomatic Go code following `gofmt` standards
- Comprehensive error handling
- Interface-driven design for testability
- Context-aware APIs
- Table-driven testing
- Clear separation of concerns
- Production-ready patterns

## 📝 Code Style

All code in this repository:
- Follows Go conventions and idioms
- Includes inline comments explaining intent
- Avoids unnecessary abstractions
- Demonstrates real-world patterns
- Is ready for production use (with appropriate scaling considerations)

## 🤝 Contributing

This is an educational repository. Each sub-project is designed to be self-contained and focused on teaching specific concepts.

## 📄 License

See [LICENSE](LICENSE) file for details.

## 🔗 Additional Resources

- [Official Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Blog](https://go.dev/blog/)
