package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	conn, err := websocket.Dial("ws://localhost:9686", "")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	go receive(conn)
	send(conn)
}

func createDemoIp() string {
	//192.168.1.1
	var array [4]int
	for i := 0; i < len(array); i++ {
		rand.Seed(time.Now().UnixNano())
		array[i] = rand.Intn(256)
	}
	return fmt.Sprintf("http://%d.%d.%d.%d", array[0], array[1], array[2], array[3])
}

func receive(conn *websocket.Conn) {
	for {
		var m Message
		err := websocket.JSON.Receive(conn, &m)
		if err != nil {
			log.Fatalln("error in Recive Data :", err)
			continue
		}
		fmt.Println("Message from server :", m.Text)

	}
}

func send(conn *websocket.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		m := Message{
			Text: text,
		}
		err := websocket.JSON.Send(conn, m)
		if err != nil {
			fmt.Println("Error in send Data :", err)
			continue
		}
	}
}
