package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/go-vgo/robotgo"
)

func takeinput1(conn net.Conn) {
	defer conn.Close()

	// Use a small buffer for instant typing
	buff := make([]byte, 128)

	for {
		// 1. Read data from the relay
		n, err := conn.Read(buff)
		if err != nil {
			log.Println("Connection closed:", err)
			return
		}

		// 2. Convert to string
		receivedData := string(buff[:n])

		// 3. Process each character
		for _, char := range receivedData {
			strChar := string(char)

			// Handle special keys
			switch strChar {
			case "\r", "\n": // Enter key
				robotgo.KeyTap("enter")
				fmt.Println("[Action] Enter")
			case "\x7f", "\b": // Backspace
				robotgo.KeyTap("backspace")
				fmt.Println("[Action] Backspace")
			default:
				// Avoid typing newlines as text
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
	// !!! REPLACE WITH YOUR CURRENT KOYEB ADDRESS !!!
	serverAddress := "01.proxy.koyeb.app:14576"

	fmt.Printf("Connecting to Receiver at %s...\n", serverAddress)
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalln("Cannot connect to relay:", err)
	}

	fmt.Println("Connected! Waiting for Sender commands...")
	takeinput1(conn)
}
