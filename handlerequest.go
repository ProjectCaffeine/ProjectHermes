package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
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

	if reqData.RequestTarget == "/Files" && reqData.HttpMethod == "POST" {
		responseCode, responsePhrase := saveFile(reqData)

		respData.ResponsePhrase = responsePhrase
		respData.ResponseCode = responseCode

		return nil
	}

	if reqData.RequestTarget == "/User" && reqData.HttpMethod == "GET" {
		return getUser(reqData, respData)
	}
	
	respData.ResponseCode = 404
	respData.ResponsePhrase = "Not found"

	return nil
}

type JsonError struct {
	Error string
}

func saveFile(reqData *RequestData) (int, string) {
	file, err := os.Create("./test.txt")

	if err != nil {
		//handle error
		return 500, "Internal Server Error"
	}

	val, ok := reqData.Headers["Content-Type"]

	if !ok {
		return 400, "Internal Server Error"
	}

	if val == "application/json" {
		data := make([]byte, 1024)

		n, err := base64.RawStdEncoding.Decode(data, reqData.body)

		if err != nil {
			//handle error
			return 500, "Internal Server Error"
		}

		writtenLen, err := file.Write(data[:n])

		if err != nil && writtenLen != n {
			//handle error
			return 500, "Internal Server Error"
		}
	}

	file.Close()

	return 200, "Ok"
}

func getUser(reqData *RequestData, respData *ResponseData) []byte {
	id, hasKey := reqData.UrlQuerys["id"]
	if !hasKey {
		respData.ResponseCode = 400
		respData.ResponsePhrase = "Bad Request"

		jError := JsonError{Error: "Id is not present in query."}

		data, err := json.Marshal(jError)

		if err != nil {
			log.Fatal(err)
		}

		return data
	}

	user := User{Name: "John", Email: "Test@Test.com", Id: id}
	data, err := json.Marshal(user)

	if err != nil {
		log.Fatal(err)
	}

	return data
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
