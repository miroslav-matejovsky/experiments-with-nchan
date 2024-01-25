package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vcs [receiver|sender]")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "receiver":
		startReceiver()
	case "sender":
		startSender()
	default:
		fmt.Println("Usage: vcs [receiver|sender]")
	}
}

func startReceiver() {
	origin := "http://localhost/"
	url := "ws://localhost:8081/sub?id=foo-channel"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	var msg = make([]byte, 512)
	for {
		var n int
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received: %s.\n", msg[:n])
	}
}

func startSender() {
	message := fmt.Sprintf("Test message at %v", time.Now().Format("15:04:00"))
	origin := "http://localhost/"
	url := "ws://localhost:8081/pub?id=foo-channel"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	websocket.Message.Send(ws, message)
	log.Printf("Message '%s' sent\n", message)
}
