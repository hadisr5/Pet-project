package main

import (
	"fmt"

	"golang.org/x/net/websocket"
)

type Message struct {
	Text string `json:"text"`
}
type hub struct {
	clients         map[string]*websocket.Conn
	addClientChn    chan *websocket.Conn
	removeClientChn chan *websocket.Conn
	broadcastChn    chan Message
}

func newHub() *hub {
	return &hub{
		clients:         make(map[string]*websocket.Conn),
		addClientChn:    make(chan *websocket.Conn),
		removeClientChn: make(chan *websocket.Conn),
		broadcastChn:    make(chan Message),
	}
}

func (h *hub) run() {
	for {
		select {
		case conn := <-h.addClientChn:
			fmt.Println(conn)
		case conn := <-h.removeClientChn:
			fmt.Println(conn)
		case m := <-h.broadcastChn:
			fmt.Println(m)
		}
	}
}
func (h *hub) addClient(conn *websocket.Conn) {
	h.clients[conn.RemoteAddr().String()] = conn

}
func (h *hub) removeClient(conn *websocket.Conn) {
	delete(h.clients, conn.RemoteAddr().String())
}

func (h *hub) broadcast(m Message) {
	for _, conn := range h.clients {
		err := websocket.JSON.Send(conn, m)
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
