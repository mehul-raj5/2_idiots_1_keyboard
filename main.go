package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// 1. Listen on port 8080 (Standard for Koyeb/Cloud)
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("cannot listen:", err)
	}

	log.Println("Relay server running on port 8080")

	// 2. Wait for the FIRST person (usually the Receiver/Robot)
	log.Println("Waiting for client 1 (Receiver)...")
	c1, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Client 1 connected!")

	// 3. Wait for the SECOND person (usually You/Sender)
	log.Println("Waiting for client 2 (Sender)...")
	c2, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Client 2 connected!")

	log.Println("Both clients connected. Starting relay...")

	// 4. Pipe data in both directions
	go io.Copy(c1, c2) // Data from Sender -> Receiver
	io.Copy(c2, c1)    // Data from Receiver -> Sender
}
