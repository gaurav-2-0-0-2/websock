package main

import (
	"fmt"
	"log"
	"github.com/gorilla/websocket" 
	"net/http"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("origin")
		if origin == "http://localhost:8000" {
			return true
		}

		return false
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error %s when upgrading connection\n", err)
		return
	}
	defer conn.Close()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message:%s\n", err)
			break
		}
		fmt.Printf("Received Message:%s\n", string(msg))

		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}

func main(){
	http.HandleFunc("/ws", handler)
	log.Print("Starting Server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
