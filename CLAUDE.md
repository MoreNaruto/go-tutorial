# Go / Golang Tutorial Generation Guidelines

You are an AI assistant helping generate a Go (Golang) tutorial repository composed of **multiple isolated sub-Go projects**.

Each sub-project focuses on a **specific core area of Go** and must be runnable and testable **independently**.

The output must be **GitHub-ready**, idiomatic Go, and production-oriented.

---

## 🎯 Repository Philosophy

This repository is a **collection of standalone Go projects**, not a single monolithic application.

Each sub-project:
- Represents a focused Go concept or domain
- Has its **own `go.mod`**
- Can be run and tested in isolation
- Does **not** depend on other sub-projects
- Uses a project structure appropriate to its scope

Examples of focus areas (non-exhaustive):
- Language fundamentals
- Concurrency patterns
- HTTP servers
- REST APIs
- Middleware
- WebSockets
- Testing strategies
- Clean Architecture
- Performance & profiling

---

## 🧠 Project Structure Guidance (IMPORTANT)

**Do NOT enforce a global repository structure.**

Instead:
- Treat each sub-project as its own Go module
- Choose an idiomatic structure per sub-project
- Prefer simplicity and clarity
- Use `cmd/`, `internal/`, or flat layouts only when appropriate
- Avoid over-structuring small examples
- Explain structural decisions in each sub-project’s README

⚠️ **Do not share packages or code between sub-projects**

Isolation is intentional and required.

---

## 🧩 Sub-Project Requirements

Each sub-Go project **must** include:

- Its own `go.mod`
- A clear entry point (when applicable)
- A README.md explaining:
    - The concept being taught
    - Why this project exists
    - Why the chosen structure makes sense
    - How to run the project
    - How to run tests
- Unit tests for all non-trivial logic
- Inline code comments explaining intent and design decisions

---

## 🧠 Educational Levels

Each sub-project should be tagged or described as:

### Beginner
- Core language constructs
- Error handling
- Structs and interfaces
- Basic concurrency
- Standard library focus

### Intermediate
- HTTP servers
- Routing (Gin or Chi)
- Context propagation
- Dependency injection
- Testing with interfaces
- Table-driven tests

### Advanced
- Clean / Hexagonal Architecture
- WebSockets
- Middleware chains
- Graceful shutdown
- Observability (logging, metrics)
- Concurrency patterns (fan-in/fan-out, worker pools)
- Performance considerations

---

## 🧪 Testing Requirements

All non-trivial logic **must include unit tests**.

Testing rules:
- Use Go’s standard `testing` package
- Prefer table-driven tests
- Mock dependencies via interfaces
- Avoid global state
- Tests must pass with:

```bash
go test ./...
```

## 📘 README Requirements

Each tutorial must include a README.md that explains:

- What the example demonstrates
- Why the chosen structure makes sense 
- Key Go concepts being taught 
- How the code is organized 
- How to run the example 
- How to run tests

README tone:

- Clear 
- Beginner-friendly 
- Avoid excessive theory 
- Focus on practical understanding

## 🧑‍💻 Code Style & Best Practices

All Go code must:

- Follow gofmt 
- Use idiomatic naming (camelCase, PascalCase)
- Avoid unnecessary abstractions 
- Return errors explicitly 
- Avoid panics in application code 
- Prefer composition over inheritance 
- Use context-aware APIs (context.Context)

## 🧩 Dependency Usage Examples

Demonstrate realistic, minimal usage of commonly used libraries:

- Commonly Used Libraries (when applicable)
- net/http 
- github.com/gin-gonic/gin 
- github.com/go-chi/chi/v5 
- golang.org/x/net/websocket or github.com/gorilla/websocket 
- context 
- testing

Explain why a dependency is used, not just how.

## 📝 Code Commenting Rules

- Comments should explain intent, not obvious syntax 
- Avoid redundant comments 
- Explain concurrency and architectural decisions 
- Use doc comments for exported functions/types

Example:

```
// UserService handles business logic related to users.
// It depends on a UserRepository interface to allow
// easy mocking during tests.
type UserService struct {
    repo UserRepository
}
```

## 🚫 What to Avoid

- Overengineering 
- Magic constants 
- Large untested functions 
- Framework-driven design 
- Outdated Go patterns 
- Repeating the same structure for all examples

## ✅ Output Expectations

When generating content:

- Produce complete files, not snippets 
- Ensure examples compile 
- Ensure tests pass 
- Ensure README instructions match the code 
- Prefer clarity over cleverness

You are generating educational, production-grade Go code.
