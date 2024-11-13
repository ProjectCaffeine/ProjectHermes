package main

import "testing"

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