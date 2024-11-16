package main

import (
	"bufio"
	"errors"
	"fmt"
	"maps"
	"strconv"
	"time"
)

func buildResponse(reqData *RequestData, statusCode int, statusMessage string, rw *bufio.ReadWriter) {
	respData :=  ResponseData{}

	respData.ResponseCode = statusCode
	respData.ResponsePhrase = statusMessage
	respData.HttpVersion = reqData.HttpVersion

	processTime := time.Now().UTC()
	data, err := processRequest(reqData, &respData)

	writeStandardHeaders(reqData, &respData, processTime, rw)
	writeResponseDataHeaders(rw, &respData)

	rw.Write([]byte("\r\n"))
	rw.Write([]byte("\n"))
	
	if err == nil {
		rw.Write(data)
	}

	rw.Flush()
}

func processRequest(reqData *RequestData, respData *ResponseData) ([]byte, error) {
	data, headers, err := handleRequest(reqData)

	if err != nil {
		respData.ResponseCode = 404
		respData.ResponsePhrase = "Not Found"

		return nil, errors.New("Request Target not found")
	}

	respData.Headers = headers

	return data, nil
}

func writeStandardHeaders(reqData *RequestData, respData *ResponseData, processTime time.Time, rw *bufio.ReadWriter) {
	rw.Write([]byte(reqData.HttpVersion))
	rw.Write([]byte(" " + strconv.Itoa(respData.ResponseCode)))
	rw.Write([]byte(" " + respData.ResponsePhrase))
	rw.Write([]byte("\n"))
	rw.Write([]byte("HTTP-date = " + processTime.Format(time.RFC1123)))
	rw.Write([]byte("\n"))
}

func writeResponseDataHeaders(rw *bufio.ReadWriter, respData *ResponseData) {
	if respData.ResponseCode == 405 {
		rw.Write([]byte("Allow: GET"))
	}

	for v := range maps.Keys(respData.Headers) {
		rw.Write([]byte(fmt.Sprintf("%s: %s", v, respData.Headers[v])))
	}
}
