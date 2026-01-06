package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for demo
	},
}

// Hub manages WebSocket connections
type Hub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan []byte
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			fmt.Printf("Client connected. Total clients: %d\n", len(h.clients))

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.Close()
				fmt.Printf("Client disconnected. Total clients: %d\n", len(h.clients))
			}

		case message := <-h.broadcast:
			fmt.Printf("Broadcasting message to %d clients\n", len(h.clients))
			for client := range h.clients {
				err := client.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					client.Close()
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	h.register <- conn

	// Read messages from client
	go func() {
		defer func() {
			h.unregister <- conn
		}()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				break
			}
			// Broadcast to all clients
			h.broadcast <- message
		}
	}()
}

func main() {
	hub := newHub()
	go hub.run()

	http.HandleFunc("/ws", hub.handleWS)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("WebSocket server starting on :8080")
	fmt.Println("Connect to: ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
