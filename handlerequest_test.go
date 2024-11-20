package main

import (
	"strings"
	"testing"
)

func TestGetIndex(t *testing.T) {
	indexData, _ := getIndex()
	data := string(indexData)

	if !strings.Contains(data, "Hello World") {
		t.Error(data)
		t.Error("Hello world not found.")
	}
}


func TestHandleRequest(t *testing.T) {
	tt := []struct {
		Name string
		Route string
		WantErr bool
	}{
		{
			"Handle Request Should error on foo",
			"/foo",
			true,
		},
		{
			"Handle Request Should not error on slash",
			"/",
			false,
		},
		{
			"Handle Request Should not error on empty string",
			"",
			false,
		},
	}

	for _, tc := range tt {

		t.Run(tc.Name, func(t *testing.T) {
			data := RequestData{
				RequestTarget: tc.Route,
			}
			respData := ResponseData{}

			handleRequest(&data, &respData)

			Equal(t, (respData.ResponseCode == 404), tc.WantErr)
		})
	}
}
