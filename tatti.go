package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

func main() {

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	// This GUARANTEES terminal returns to normal,
	// even if your program panics or breaks.
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	reader := bufio.NewReader(os.Stdin)
	buff := make([]rune, 1024)

	index := 0

	for {
		input, _, err := reader.ReadRune()
		if err != nil {
			log.Fatalln("tatti hogyi:", err)
		}

		if input == '=' {
			break
		}

		if index < len(buff) {
			buff[index] = input
			index++
		} else {
			log.Fatalln("buffer full")
			break
		}
	}

	fmt.Println(string(buff[:index]))
}
