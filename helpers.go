package main

import (
	"errors"
	"log"
	"maps"
	"testing"
)

func MergeHeaders(source map[string]string,  target map[string]string) (map[string]string, error) {
	newMap := source

	for v := range maps.Keys(newMap) {
		_, hasKey := source[v]

		if hasKey {
			return nil, errors.New("Source map already has key.")
		}

		newMap[v] = target[v]
	}

	return newMap, nil
}

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
