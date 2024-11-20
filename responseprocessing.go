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
	var data []byte

	if statusCode != 405 {
		data, _ = processRequest(reqData, &respData)
	}

	fmt.Print("Writing standard headers\n")
	writeStandardHeaders(reqData, &respData, processTime, rw)
	fmt.Print("Writing response headers\n")
	writeResponseDataHeaders(rw, &respData)

	rw.Write([]byte("\r\n"))
	
	//if err == nil && statusCode != 405 && len(data) > 0{
	if data != nil && len(data) > 0{
		rw.Write(data)
	}

	rw.Flush()
}

func processRequest(reqData *RequestData, respData *ResponseData) ([]byte, error) {
	data := handleRequest(reqData, respData)

	if respData.ResponseCode == 404 {
		return nil, errors.New("Request Target not found")
	}

	return data, nil
}

func writeStandardHeaders(reqData *RequestData, respData *ResponseData, processTime time.Time, rw *bufio.ReadWriter) {
	rw.Write([]byte(reqData.HttpVersion))
	rw.Write([]byte(" " + strconv.Itoa(respData.ResponseCode)))
	rw.Write([]byte(" " + respData.ResponsePhrase))
	rw.Write([]byte("\n"))
	rw.Write([]byte("Date: " + processTime.Format(time.RFC1123)))
	rw.Write([]byte("\n"))
}

func writeResponseDataHeaders(rw *bufio.ReadWriter, respData *ResponseData) {
	if respData.ResponseCode == 405 {
		rw.Write([]byte("Allow: GET\n"))
		rw.Write([]byte("Content-Length: 0\n"))
	}

	if respData.Headers != nil {
		for v := range maps.Keys(respData.Headers) {
			rw.Write([]byte(fmt.Sprintf("%s: %s\n", v, respData.Headers[v])))
		}
	}
}
