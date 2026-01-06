# 09 - Clean Architecture

**Level:** Advanced

## What This Project Teaches

Hexagonal/Clean architecture in Go:
- Domain-driven design
- Dependency inversion
- Repository pattern
- Use case layer
- Testable architecture

## Structure

```
09-clean-architecture/
├── domain/          # Business entities
├── repository/      # Data access interfaces
├── usecase/         # Business logic
├── handler/         # HTTP handlers
└── main.go          # Dependency wiring
```

## Layers

1. **Domain:** Core business entities (no dependencies)
2. **Repository:** Data access interface (domain depends on this)
3. **Use Case:** Business logic (depends on repository interface)
4. **Handler:** HTTP/external interface (depends on use case)

## Running

```bash
go run .
go test ./...
```

## Key Benefits

- Testable (mock any layer)
- Independent of frameworks
- Independent of database
- Clear separation of concerns
- Easy to understand and maintain

This structure demonstrates production-ready architecture used in large Go applications.
