package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"strconv"
)

func handleRequest(reqData *RequestData, respData *ResponseData) []byte {
	if reqData.RequestTarget == "/" ||
		reqData.RequestTarget == "" {
		data, headers := getIndex()

		respData.Headers  = headers
		return data
	}

	if reqData.RequestTarget == "/User" && reqData.HttpMethod == "POST" {
		createUser(reqData)

		respData.ResponsePhrase = "Created"
		respData.ResponseCode = 201
		return nil

	}
	
	respData.ResponseCode = 404
	respData.ResponsePhrase = "Not found"

	return nil
}

func createUser(reqData *RequestData) {
	user := User{}

	err := json.Unmarshal(reqData.body, &user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nUser name:%s\nUser email:%s\n", user.Name, user.Email)
}

func getIndex() ([]byte, map[string]string) {
	headers := make(map[string]string)
	const tpl = `<!DOCTYPE html>
		<html>
		<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
		</head>
		<body>
		Hello World
		</body>
		</html>`

	t, err := template.New("webpage").Parse(tpl)
	checkForError(err)

	var bf bytes.Buffer
	
	err = t.Execute(&bf, nil)
	checkForError(err)

	headers["Content-Type"] = "text/html; charset=utf-8"
	headers["Content-Length"] = strconv.Itoa(bf.Len())

	return bf.Bytes(), headers
}
