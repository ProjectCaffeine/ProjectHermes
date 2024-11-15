package main

import (
	"strconv"
	"bufio"
)

func buildResponse(reqData *RequestData, statusCode int, statusMessage string, rw *bufio.ReadWriter) {
	respData :=  ResponseData{}

	respData.ResponseCode = statusCode
	respData.ResponsePhrase = statusMessage
	respData.HttpVersion = reqData.HttpVersion


	processRequest(reqData, &respData)

	rw.Write([]byte("HTTP-VERSION = " + reqData.HttpVersion))
	rw.Write([]byte(" " + strconv.Itoa(statusCode)))
	rw.Write([]byte(" " + statusMessage))
	rw.Write([]byte("\n"))

	//return headers go here
	writeHeaders(rw, statusCode)

	rw.Write([]byte("\r\n"))
	rw.Write([]byte("\n"))
	
	//data goes here

	rw.Flush()
}

func processRequest(reqData *RequestData, respData *ResponseData) {
	data, headers, err := handleRequest(reqData)

	if err != nil {
		respData.ResponseCode = 404
		respData.ResponsePhrase = "Not Found"
	}
}

func writeHeaders(rw *bufio.ReadWriter, statusCode int) {
	if statusCode == 405 {
		rw.Write([]byte("Allow: GET"))
	}

}
