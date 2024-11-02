package main

import (
	"fmt"
	"net"
	"log"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Failed to open TCP connection")
	}

	fmt.Print("Opening connection on port 8080")
	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Fatal("Failed to accept connection")
		}

		conn.Close()
		break;
	}

	fmt.Print("Closing connection")
	ln.Close()
}
