package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {

	h := newHub()
	mux := http.NewServeMux()
	mux.Handle("/", websocket.Handler(func(conn *websocket.Conn) {
		wsHandler(conn, h)
	}))

	server := http.Server{
		Addr:    ":9686",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
func wsHandler(conn *websocket.Conn, h *hub) {
	go h.run()

	h.addClientChn <- conn

	for {
		var m Message
		err := websocket.JSON.Receive(conn, &m)
		if err != nil {
			h.removeClientChn <- conn
			fmt.Println("error in Recive Message : ", err.Error())
			continue
		}
		h.broadcastChn <- m
	}
}
