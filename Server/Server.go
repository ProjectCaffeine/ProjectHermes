package main

import (
	"fmt"
	"log"
	"net"
	//"strings"
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

		//sendBuffer := make([]byte, 2056)
		parseRequest(buffer, n)

		//fmt.Printf("Received:\n%s", textString)
		conn.Close()
		break;
	}

	fmt.Print("Closing connection\n")
	ln.Close()
}

func parseRequest(buffer []byte, dataEnd int)(RequestData)  {
	data := RequestData{}
	getTokens(buffer, dataEnd)

	return data
}

func getTokens(buffer []byte, dataEnd int)([]string)  {
	textString := string(buffer[:dataEnd])
	tokens := []string{}
	startIndex := 0

	for i, v := range(textString) {
		if v == ' ' {
			foundToken := textString[startIndex:i]
			tokens = append(tokens, foundToken)

			if i < len(textString) - 1 {
				startIndex = i + 1
			}
		} else if i == len(textString) - 1 {
			foundToken := textString[startIndex:i]
			tokens = append(tokens, foundToken)
		}
	}

	for _, v := range(tokens) {
		fmt.Printf("%s\n", v)
	}

	fmt.Printf("\n\n\n%s", textString)

	return tokens
}
