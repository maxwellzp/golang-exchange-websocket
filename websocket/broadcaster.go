package websocket

import (
	"encoding/json"
	"golang-exchange-websocket/redis"
	"log"
	"time"
)

func BroadCastLoop(redisService *redis.Service) {
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	for range ticker.C {
		mu.Lock()
		for client := range clients {
			for pair := range client.Pairs {
				rate, err := redisService.GetRate(pair)
				if err != nil {
					log.Printf("Error getting rate for %s: %v", pair, err)
					continue
				}
				msg := map[string]string{
					"pair": pair,
					"rate": rate,
				}

				dataMsg, err := json.Marshal(msg)
				if err != nil {
					log.Printf("Error marshalling message for %s: %v", pair, err)
					continue
				}

				// Only one goroutine writes to the WebSocket connection at a time.
				client.Mutex.Lock()
				err = client.Conn.WriteMessage(1, dataMsg)
				client.Mutex.Unlock()
				if err != nil {
					client.Conn.Close()
					delete(clients, client)
					log.Printf("Error writing message for %s: %v", pair, err)
					continue
				}
			}
		}
		mu.Unlock()
	}
}
