package main

import (
	"log"
	"net"

	"github.com/storyofhis/tcp-chat/httpserver/controllers"
)

func main() {
	// creating server
	server := controllers.NewServer()
	go server.Run()

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Unable to start server: %s", err.Error())
	}
	defer listener.Close()
	log.Println("Starting Server on :9090")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %s", err)
			continue
		}
		c := server.NewClient(conn)
		go c.ReadInput()
	}
}
