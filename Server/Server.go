package main

import (
	"bufio"
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
	runMainServer()

}

func runMainServer() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Failed to open TCP connection")
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()

		if err != nil || conn == nil {
			log.Fatal("Could not accept connection")
		}

		go handleConnection(ln, conn)
	}
}
 
func runParrotServer() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Failed to open TCP connection")
	}

	defer ln.Close()
	for {
		conn, err := ln.Accept()

		if err != nil || conn == nil {
			log.Fatal("Could not accept connection")
		}

		go handleParrotConnection(ln, conn)
	}

	ln.Close()
}

func handleParrotConnection(ln net.Listener, conn net.Conn) {
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	for {
		req, err := rw.ReadString('\n')

		if err != nil {
			rw.WriteString("failed to read input")
			rw.Flush()
			return
		}

		rw.WriteString(fmt.Sprintf("Request received: %s", req))
		rw.Flush()
	}
}

func handleConnection(ln net.Listener, conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 2056)
	n, err := conn.Read(buffer)

	if err != nil {
		log.Fatal("Failed to read data.")
	}

	parseRequest(buffer, n)
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
