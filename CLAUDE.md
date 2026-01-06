# Go / Golang Tutorial Generation Guidelines

You are an AI assistant helping generate a Go (Golang) tutorial repository intended for **Beginner**, **Intermediate**, and **Advanced** developers.

The output must be **GitHub-ready**, idiomatic Go, and production-oriented.

---

## 🎯 Goals

This repository should:
- Teach Go progressively (Beginner → Intermediate → Advanced)
- Follow Go best practices and idioms
- Include clear explanations and runnable examples
- Include unit tests for all non-trivial logic
- Demonstrate commonly used Go libraries and patterns

---

## 🧠 Project Structure Guidance (IMPORTANT)

**Do NOT assume or reuse a predefined repository structure.**

Instead:
- Choose an idiomatic Go project structure appropriate for the example
- Prefer simplicity over excessive layering
- Follow Go community conventions (`cmd/`, `internal/`, flat packages when appropriate)
- Let the **problem domain drive the structure**, not frameworks
- Explain the chosen structure in the README

Examples of acceptable approaches (non-exhaustive):
- Single-module with `main.go` for beginner examples
- `cmd/<app>` + `internal/` for intermediate/advanced services
- Flat packages for libraries
- Feature-based organization when appropriate

⚠️ **Do not hardcode a structure across all tutorials**  
Each tutorial may choose a different structure if it improves clarity.

---

## 🧠 Educational Levels

### Beginner
- Go syntax and structure
- `main` package
- Variables, structs, interfaces
- Error handling
- Basic concurrency (`goroutines`, `channels`)
- Standard library only

### Intermediate
- HTTP servers
- REST APIs
- Routing (Gin or Chi)
- Context usage
- Dependency injection
- Table-driven tests
- Mocking with interfaces

### Advanced
- Clean Architecture / Hexagonal patterns
- WebSockets
- Middleware
- Graceful shutdown
- Observability (logging, metrics)
- Performance considerations
- Concurrency patterns (worker pools, fan-in/fan-out)

---

## 🧪 Testing Requirements

All non-trivial logic **must include unit tests**.

Testing rules:
- Use Go’s standard `testing` package
- Prefer table-driven tests
- Avoid global state
- Mock dependencies using interfaces (not monkey patching)
- Tests must be runnable with:

```bash
go test ./...
