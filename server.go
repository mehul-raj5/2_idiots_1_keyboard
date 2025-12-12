package main

import (
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("cannot listen:", err)
	}

	log.Println("Relay server running on port 8080")

	log.Println("Waiting for client 1 (Receiver)...")
	c1, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Client 1 connected!")

	log.Println("Waiting for client 2 (Sender)...")
	c2, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Client 2 connected!")

	log.Println("Both clients connected. Starting relay...")

	go io.Copy(c1, c2)
	io.Copy(c2, c1)
}
