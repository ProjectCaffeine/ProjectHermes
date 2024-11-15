package main

import (
	"strings"
	"testing"
)

func TestGetIndex(t *testing.T) {
	data := string(getIndex())

	if !strings.Contains(data, "Hello World") {
		t.Error(data)
		t.Error("Hello world not found.")
	}
}
