package main

import (
	"bufio"
	"bytes"
	"strconv"
	"testing"
	"time"
)

func TestWriteStandardHeaders(t *testing.T) {
	tt := []struct{
		Name string
		HttpVersion string 
		ResponseCode int 
		ResponsePhrase string
	} {
		{
			"Should build 200 headers",
			"HTTP 1.1",
			200,
			"Ok",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			reqData := RequestData {
				HttpVersion: tc.HttpVersion,
			}
			respData := ResponseData {
				ResponseCode: tc.ResponseCode,
				ResponsePhrase: tc.ResponsePhrase,
			}

			
			b := new(bytes.Buffer)
			bw := bufio.NewWriter(b)
			br := bufio.NewReader(bytes.NewReader([]byte{}))
			rw := bufio.NewReadWriter(br, bw)
			testTime := time.Now().UTC()

			writeStandardHeaders(&reqData, &respData, testTime, rw)

			rw.Flush()

			expect := []byte(tc.HttpVersion + 
			" " + 
			strconv.Itoa(tc.ResponseCode) +
			" " + 
			tc.ResponsePhrase + 
			"\n" +
			"HTTP-date = " +
			testTime.Format(time.RFC1123))

			if bytes.Compare(b.Bytes(), expect) != 1 {
				t.Errorf("Response was not the expected response\nGot:\n%s\nExpected:\n%s\n", string(b.Bytes()), string(expect))

			}
		})
	}
}
