package main

import (
	"bytes"
	"errors"
	"html/template"
	"strconv"
)

func handleRequest(reqData *RequestData) ([]byte, map[string]string, error) {
	if reqData.RequestTarget == "/" ||
		reqData.RequestTarget == "" {
		data, headers := getIndex()
		return data, headers, nil
	}
	
	return nil, nil, errors.New("Request Target not found")
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
