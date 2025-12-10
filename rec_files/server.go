package main

import (
	"fmt"
	"log"
	"net"
)

func takeinput1(conn net.Conn) {
	defer conn.Close()

	for {
		buff := make([]byte, 64)
		n, err := conn.Read(buff)
		if err != nil {
			log.Println("connection closed", err)
			return
		}

		fmt.Print(string(buff[:n]))
	}
}

func main() {
	conn, err := net.Dial("tcp", "your-app-name.koyeb.app:8080")
	if err != nil {
		log.Fatalln("cannot connect:", err)
	}

	takeinput1(conn)
}
