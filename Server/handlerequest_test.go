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
