package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Failed to open TCP connection")
	}

	defer ln.Close()
	fmt.Print("Opening connection on port 8080\n")
	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Fatal("Failed to accept connection\n")
		}

		buffer := make([]byte, 2056)
		n, err := conn.Read(buffer)

		if err != nil {
			log.Fatal("Failed to read data.")
		}

		fmt.Printf("Received:\n %s", buffer[:n])
		conn.Close()
		break;
	}

	fmt.Print("Closing connection\n")
	ln.Close()
}
