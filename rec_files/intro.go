package main

//for WINDOWS sender

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/term"
)

func takeinput(conn net.Conn) {
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

		fmt.Fprintf(conn, "%s", string(char))
	}
}

func main() {
	conn, err := net.Dial("tcp", "xx:xx")
	if err != nil {
		log.Fatalln("Connection failed:", err)
	}
	defer conn.Close()

	fmt.Println("Connected! Typing here will reflect on the Mac.")
	fmt.Println("Press '=' to exit.")
	takeinput(conn)
}
