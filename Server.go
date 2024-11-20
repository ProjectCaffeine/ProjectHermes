package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type User struct {
	Name string 
	Email string
}

type ResponseData struct {
	HttpVersion string
	ResponseCode int 
	ResponsePhrase string
	Headers map[string]string
}

type RequestData struct {
	HttpMethod string
	RequestTarget string 
	HttpVersion string
	Headers map[string]string
	body []byte
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
 
func printEntireRequest(conn net.Conn) {

	// Create a buffer to read data into
	buffer := make([]byte, 4096)

	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Process and use the data (here, we'll just print it)
		fmt.Printf("Received: \n%q\n", buffer[:n])
	}
}

func handleConnection(ln net.Listener, conn net.Conn) {
	defer conn.Close()
	//printEntireRequest(conn);
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	log.Print("Parsing Request\n")
	reqData := parseRequest(rw)
	log.Print("Validating Headers\n")
	statusCode, statusMessage := validateHeaderRecords(&reqData)

	log.Print("Building Response\n")
	buildResponse(&reqData, statusCode, statusMessage, rw)
}



func validateHeaderRecords(data *RequestData) (int, string)  {
	if data.HttpMethod != "GET" && data.HttpMethod != "POST" {
		return 405, "Method Not Allowed"
	}

	return 200, "Ok"
}
