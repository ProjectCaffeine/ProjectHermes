package main

import (
	"fmt"
	"log"
	"net"
)

type RequestData struct {
	RequestType string
	Path string 
	Host string 
	UserAgent string 
	Accept string
}

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

		defer conn.Close()
		buffer := make([]byte, 2056)
		n, err := conn.Read(buffer)

		if err != nil {
			log.Fatal("Failed to read data.")
		}

		textString := string(buffer[:n])

		//sendBuffer := make([]byte, 2056)

		//conn.Write(sendBuffer)

		fmt.Printf("Received:\n %s", textString)
		conn.Close()
		break;
	}

	fmt.Print("Closing connection\n")
	ln.Close()
}

func parseRequest(buffer []byte, dataEnd int)(RequestData)  {
	data := RequestData{}

	

	return data
}
