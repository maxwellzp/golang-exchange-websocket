package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type Client struct {
	Conn  *websocket.Conn
	Mutex sync.Mutex
	Pairs map[string]bool
}

type Message struct {
	Action string `json:"action"`
	Pair   string `json:"pair"`
}

var (
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	clients  = make(map[*Client]bool)
	mu       sync.Mutex
)

func HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket", http.StatusBadRequest)
		return
	}

	client := &Client{Conn: conn, Pairs: make(map[string]bool)}

	mu.Lock()
	clients[client] = true
	mu.Unlock()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		switch msg.Action {
		case "subscribe":
			client.Pairs[msg.Pair] = true
		case "unsubscribe":
			delete(client.Pairs, msg.Pair)
		}
	}

	mu.Lock()
	delete(clients, client)
	mu.Unlock()
	err = conn.Close()
	if err != nil {
		log.Printf("Error closing websocket: %v", err)
	}
}
