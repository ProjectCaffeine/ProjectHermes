package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	//"strings"
)

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

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	parseRequest(rw)
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
