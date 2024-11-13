package main

import (
	"bufio"
	"bytes"
	"net"
	"os"
	"testing"
)

type Writer int

func TestGetHeaders(t *testing.T) {
	testData := []byte(`GET /test HTTP/1.1
Host: 127.0.0.1:8080
User-Agent: curl/7.81.0
Accept: */*`)
	bw := bufio.NewWriter(os.Stdout)
	br := bufio.NewReader(bytes.NewReader(testData))
	rw := bufio.NewReadWriter(br, bw)
	reqData := processHeaderLine(rw)

	Equal(t, reqData.HttpMethod, "GET" )
	//if reqData.HttpMethod != "GET" {
	//	t.Errorf("Expected HTTP Method: GET, got: %s", reqData.HttpMethod)
	//}

	headers := getHeaders(rw)

	Equal(t, len(headers), 4)
}

func TestValidateHeaderFile(t *testing.T) {
	tt := []struct {
		test string
		data RequestData
		want int
	} {
		{
			"Sending a GET",
			RequestData{HttpMethod: "GET"},
			200,
		},
		{
			"Sending a POST",
			RequestData{HttpMethod: "POST"},
			405,
		},
	}

	for _, tc := range tt {

		t.Run(tc.test, func(t *testing.T) {
			Equal(t, validateHeaderRecords(&tc.data), tc.want)
		})
	}
}

func TestMainServerConnection(t *testing.T) {
	tt := []struct {
		test string 
		payload []byte 
		want []byte
	} {
		{
			"Sending a GET Request displays tokens",
			[]byte(`GET /test HTTP/1.1
Host: 127.0.0.1:8080
User-Agent: curl/7.81.0
Accept: */*`),
			[]byte("Request received: hello world"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.test, func(t *testing.T) {
			conn, err := net.Dial("tcp", ":8080")
			
			if err != nil {
				t.Error("Could not connect to TCP server: ", err)
			}

			defer conn.Close()

			if _, err := conn.Write(tc.payload); err != nil {
				t.Error("could not write payload to TCP server:", err)
			}

			out := make([]byte, 1024)

			if _, err := conn.Read(out); err == nil {
				if bytes.Compare(out, tc.want) == 0 {
					t.Error("response did match expected output")
				}
			} else {
				t.Error("could not read from connection")
			}
		})
	}
}

func TestParrotServerConnection(t *testing.T) {
	tt := []struct {
		test string 
		payload []byte 
		want []byte
	} {
		{
			"Sending a simple request returns result",
			[]byte("Hello world\n"),
			[]byte("Request received: hello world"),
		},
		{
			"Sending another simple request works",
			[]byte("goodbye world\n"),
			[]byte("Request received: goodbyte world"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.test, func(t *testing.T) {
			conn, err := net.Dial("tcp", ":8080")
			
			if err != nil {
				t.Error("Could not connect to TCP server: ", err)
			}

			defer conn.Close()

			if _, err := conn.Write(tc.payload); err != nil {
				t.Error("could not write payload to TCP server:", err)
			}

			out := make([]byte, 1024)

			if _, err := conn.Read(out); err == nil {
				if bytes.Compare(out, tc.want) == 0 {
					t.Error("response did match expected output")
				}
			} else {
				t.Error("could not read from connection")
			}
		})
	}
}
