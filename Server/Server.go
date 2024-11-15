package main

import (
	"bufio"
	"log"
	"net"
)

type ResponseData struct {
	HttpVersion string
	ResponseCode int 
	ResponsePhrase string
}

type RequestData struct {
	HttpMethod string
	RequestTarget string 
	HttpVersion string
	Headers map[string]string
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
 

func handleConnection(ln net.Listener, conn net.Conn) {
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	reqData := parseRequest(rw)
	statusCode, statusMessage := validateHeaderRecords(&reqData)

	buildResponse(&reqData, statusCode, statusMessage, rw)
}



func validateHeaderRecords(data *RequestData) (int, string)  {
	if data.HttpMethod != "GET" {
		return 405, "Method Not Allowed"
	}

	return 200, "Ok"
}
