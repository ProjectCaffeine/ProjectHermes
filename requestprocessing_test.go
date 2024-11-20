package main

import (
	"testing"
)

func TestGetQueryParameters(t *testing.T) {
	tt := []struct{
		Name string
		RequestTarget string 
		DesiredLength int 
	} {
		{
			"Should return 1 with 1 query param",
			"/User?id=123",
			1,
		},
		{
			"Should return 0 with no query params",
			"/User",
			0,
		},
		{
			"Should return 2 with 2 query params",
			"/User?name=test&age=23",
			2,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			params := getQueryParameters(tc.RequestTarget)

			Equal(t, len(params), tc.DesiredLength)
		})
	}
}
