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

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	// This GUARANTEES terminal returns to normal,
	// even if your program panics or breaks.
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	//raw
	r := bufio.NewReader(os.Stdin)

	var outputRune []rune
	for {
		temp, _, _ := r.ReadRune()
		if temp == '=' {
			conn.Close()
			break
		}
		outputRune = append(outputRune, temp)
		fmt.Fprintf(conn, string(temp))

	}
	fmt.Println("\nYou typed: ", outputRune)
	fmt.Println()

	//unraw
}

func main() {
	//con est
	conn, err := net.Dial("tcp", "192.168.0.6:8080")
	if err != nil {
		log.Fatalln(err)
	}

	takeinput(conn)
}
