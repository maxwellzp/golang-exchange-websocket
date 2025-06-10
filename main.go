package main

import (
	"golang-exchange-websocket/redis"
	"golang-exchange-websocket/websocket"
	"log"
	"net/http"
)

func main() {
	redis.Init()
	go websocket.BroadCastLoop()

	http.HandleFunc("/ws", websocket.HandleWebsocket)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Printf("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
