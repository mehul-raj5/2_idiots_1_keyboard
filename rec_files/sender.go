package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/term"
)

func SendKeystrokes(conn net.Conn) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	r := bufio.NewReader(os.Stdin)

	for {
		char, _, err := r.ReadRune()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		if char == '=' {
			conn.Close()
			break
		}

		_, err = conn.Write([]byte(string(char)))
		if err != nil {
			log.Println("Connection lost:", err)
			break
		}
	}
}

func main() {
	serverAddress := "01.proxy.koyeb.app:13714"

	fmt.Printf("Connecting to Relay at %s...\n", serverAddress)
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalln("Connection failed:", err)
	}
	defer conn.Close()

	fmt.Println("Connected! Start typing (Press '=' to exit)...")
	SendKeystrokes(conn)
}
