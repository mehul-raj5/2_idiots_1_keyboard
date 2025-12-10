package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/term"
)

func takeinput(conn net.Conn) {
	// 1. Set terminal to Raw Mode (Captures keys instantly)
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	r := bufio.NewReader(os.Stdin)

	for {
		// 2. Read exact key press
		char, _, err := r.ReadRune()
		if err != nil {
			break
		}

		// 3. Exit safely on '='
		if char == '=' {
			break
		}

		// 4. Send key to server
		fmt.Fprintf(conn, "%s", string(char))
	}
}

func main() {
	// !!! REPLACE WITH YOUR CURRENT KOYEB ADDRESS !!!
	serverAddress := "01.proxy.koyeb.app:14576"

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalln("Connection failed:", err)
	}
	defer conn.Close()

	fmt.Println("Connected! Press '=' to exit.")
	fmt.Println("Start typing...")

	takeinput(conn)
}
