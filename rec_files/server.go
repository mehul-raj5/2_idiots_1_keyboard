package main

//for MAC reciever

import (
	"fmt"
	"log"
	"net"

	"github.com/go-vgo/robotgo"
)

func takeinput1(conn net.Conn) {
	defer conn.Close()

	buff := make([]byte, 8)

	for {
		n, err := conn.Read(buff)
		if err != nil {
			log.Println("Connection closed:", err)
			return
		}

		receivedData := string(buff[:n])

		switch receivedData {
		case "\r":
			robotgo.KeyTap("enter")
		case "\x7f", "\b":
			robotgo.KeyTap("backspace")
		default:
			robotgo.TypeStr(receivedData)
		}

		fmt.Printf("Simulated: %q\n", receivedData)
	}
}

func main() {
	conn, err := net.Dial("tcp", "xx:xx")
	if err != nil {
		log.Fatalln("Cannot connect:", err)
	}

	fmt.Println("Connected to stream. Waiting for input...")
	takeinput1(conn)
}
