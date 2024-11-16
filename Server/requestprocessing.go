package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func parseRequest(rw *bufio.ReadWriter)(RequestData)  {
	log.Print("Processing Header Line\n")
	data := processHeaderLine(rw)
	log.Print("Getting Headers\n")
	data.Headers = getHeaders(rw)

	fmt.Printf("Method: '%s'\nTarget: '%s'\nVersion: '%s'\n", data.HttpMethod, data.RequestTarget, data.HttpVersion)

	//for k := range data.Headers {
	//	fmt.Printf("Key: '%s'\nValue: '%s'\n", k, data.Headers[k])
	//}

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

		if lineText == "" || strings.ContainsRune(lineText, '\r') {
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
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}

	return results
}
