package main

import (
	"bufio"
	"fmt"
	"strconv"
	"log"
	"net"
	"strings"
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

func buildResponse(reqData *RequestData, statusCode int, statusMessage string, rw *bufio.ReadWriter) {

	rw.Write([]byte("HTTP-VERSION = " + reqData.HttpVersion))
	rw.Write([]byte(" " + strconv.Itoa(statusCode)))
	rw.Write([]byte(" " + statusMessage))
	rw.Write([]byte("\n"))

	//return headers go here

	rw.Write([]byte("\r\n"))
	rw.Write([]byte("\n"))
	
	//data goes here

	rw.Flush()
}

func parseRequest(rw *bufio.ReadWriter)(RequestData)  {
	data := processHeaderLine(rw)
	data.Headers = getHeaders(rw)

	fmt.Printf("Method: '%s'\nTarget: '%s'\nVersion: '%s'\n", data.HttpMethod, data.RequestTarget, data.HttpVersion)

	for k := range data.Headers {
		fmt.Printf("Key: '%s'\nValue: '%s'\n", k, data.Headers[k])
	}

	return data
}

func processHeaderLine(rw *bufio.ReadWriter)(RequestData) {
	data := RequestData{}
	firstLine, err := rw.ReadBytes('\n')

	if err != nil {
		log.Fatal(err)
	}

	headerLineProperties := strings.Split(string(firstLine), " ")
	cleanedVersion := strings.TrimSuffix(headerLineProperties[2], "\r\n")
	cleanedVersion = strings.TrimSuffix(cleanedVersion, "\n")

	data.HttpMethod = headerLineProperties[0]
	data.RequestTarget = headerLineProperties[1]
	data.HttpVersion = cleanedVersion

	return data
}

func getHeaders(rw *bufio.ReadWriter)(map[string]string)  {
	results := make(map[string]string)
	scanner := bufio.NewScanner(rw.Reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lineText := scanner.Text()

		if lineText == "" {
			break
		}

		var delim string
		colonAndSpaceIdx := strings.Index(lineText, ": ")
		colonIdx := strings.Index(lineText, ":")
		
		if colonAndSpaceIdx == colonIdx {
			delim = ": "
		} else {
			delim = ":"
		}

		splitHeaders := strings.SplitN(lineText, delim, 2)
		results[splitHeaders[0]] = splitHeaders[1]
	}

	return results
}

func validateHeaderRecords(data *RequestData) (int, string)  {
	if data.HttpMethod != "GET" {
		return 405, "Method Not Allowed"
	}

	return 200, "Ok"
}
