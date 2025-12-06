package main

import (
	// "bufio"
	"fmt"
	"log"
	"net"
	// "os"
	// "golang.org/x/term"
)

func takeinput1(conn net.Conn) (int){
	defer conn.Close()

	for{
		buff := make([]byte, 64)
		n, err := conn.Read(buff)
		if err != nil {
			log.Fatalln("gobar hogyi")
		}

		fmt.Println(string(buff[:n]))
	}

	// return 1;
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("tatti hogyi")
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		if(takeinput1(conn) == 1){
			break;
		}

	}
}
