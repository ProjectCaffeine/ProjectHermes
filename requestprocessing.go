package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func parseRequest(rw *bufio.ReadWriter)(RequestData)  {
	fmt.Print("Processing Header Line\n")
	data := processHeaderLine(rw)
	fmt.Print("Getting Headers\n")
	//data.Headers = getHeaders(rw)
	data.Headers = getHeadersCustom(rw)

	fmt.Printf("\n\nMethod: '%s'\nTarget: '%s'\nVersion: '%s'\n\n", data.HttpMethod, data.RequestTarget, data.HttpVersion)


	//for k := range data.Headers {
	//	fmt.Printf("Key: '%s'\nValue: '%s'\n", k, data.Headers[k])
	//}

	getBody(rw, data.Headers["Content-Length"])

	return data
}

func getBody(rw *bufio.ReadWriter, cl string) {
	_, err := rw.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	contentLength, err := strconv.Atoi(cl)

	if err != nil {
		log.Fatal(err)
	}

	body, err := rw.Peek(contentLength)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Body: \n%s\n\n", string(body))
	user := User{}

	err = json.Unmarshal(body, &user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User name:%s\nUser email:%s\n", user.Name, user.Email)
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

func getHeadersCustom(rw *bufio.ReadWriter)(map[string]string) {
	results := make(map[string]string)

	for {
		lineText, err := rw.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		nextPreview, err := rw.Peek(2)

		if err != nil {
			log.Fatal(err)
		}

		//fmt.Printf("Prev: %q", string(nextPreview))

		if string(nextPreview) == "\r\n" {
			break
		}
		lineText = strings.Replace(lineText, "\r\n", "", -1)
		
		//log.Printf("\nScannerText:\n%q\n", lineText)
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

	return results
}

func getHeaders(rw *bufio.ReadWriter)(map[string]string)  {
	results := make(map[string]string)
	scanner := bufio.NewScanner(rw.Reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lineText := scanner.Text()

		log.Printf("\nScannerText:\n%q\n", lineText)
		log.Printf("\nBytes:\n%q\n", scanner.Bytes())
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
