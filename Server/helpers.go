package main

import (
	"testing"
	"log"
)

func Equal[V comparable](t *testing.T, got, expected V) {
	t.Helper()

	if expected != got {
		t.Errorf(`assert.Equal(
			t,
			got: 
			%v
			,
			expected:
			%v
			)`, got, expected)
	}
}

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
