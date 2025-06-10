package main

import (
	"golang-exchange-websocket/redis"
	"golang-exchange-websocket/websocket"
	"log"
	"net/http"
)

func main() {
	redisService := redis.NewService("redis:6379")
	go websocket.BroadCastLoop(redisService)

	//http.HandleFunc("/ws", websocket.HandleWebsocket)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.HandleWebsocket(w, r, redisService)
	})
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Printf("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
