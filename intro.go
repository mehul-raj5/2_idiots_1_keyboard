package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

func takeinput() {

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
			break
		}
		outputRune = append(outputRune, temp)

	}
	fmt.Println("\nYou typed: ", outputRune)
	fmt.Println()

	//unraw
}

func main() {
	fmt.Println("Hello, World!")
	takeinput()
}
