# 10 - WebSockets

**Level:** Advanced

## What This Project Teaches

Real-time bidirectional communication with WebSockets:
- WebSocket server setup
- Connection management (Hub pattern)
- Broadcasting messages
- Client connection/disconnection
- Message handling

## Running

```bash
go mod download
go run main.go
```

Connect with a WebSocket client:
- Browser: `new WebSocket('ws://localhost:8080/ws')`
- CLI: `websocat ws://localhost:8080/ws`

## Key Concepts

- **Hub Pattern:** Central manager for connections
- **Broadcasting:** Send message to all connected clients
- **Goroutines:** Each connection handled concurrently
- **Channels:** Coordinate connection management

## Use Cases

- Chat applications
- Live notifications
- Real-time dashboards
- Collaborative editing
- Live sports scores

This demonstrates production-ready WebSocket patterns used in real-time applications.
