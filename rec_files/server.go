package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/go-vgo/robotgo"
)

func takeinput1(conn net.Conn) {

	buff := make([]byte, 128)

	for {
		n, err := conn.Read(buff)
		if err != nil {
			continue
		}

		receivedData := string(buff[:n])

		for _, char := range receivedData {
			strChar := string(char)

			switch strChar {
			case "\r", "\n":
				robotgo.KeyTap("enter")
				fmt.Println("[Action] Enter")
			case "\x7f", "\b":
				robotgo.KeyTap("backspace")
				fmt.Println("[Action] Backspace")
			default:
				if strings.TrimSpace(strChar) == "" && strChar != " " {
					continue
				}
				robotgo.TypeStr(strChar)
				fmt.Printf("[Typed] %s\n", strChar)
			}
		}
	}
}

func main() {
	serverAddress := "xx:xx"

	fmt.Printf("Connecting to Receiver at %s...\n", serverAddress)
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalln("Cannot connect to relay:", err)
	}

	fmt.Println("Connected! Waiting for Sender commands...")
	takeinput1(conn)
}
