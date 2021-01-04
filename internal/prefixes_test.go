package internal

import (
	"strconv"
	"testing"
)

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		token    string
		expected bool
	}{
		{"--", true},
		{"-", true},
		{"", false},
		{"_", false},
	}

	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertEquals(t, hasPrefix(tt.token), tt.expected)
		})
	}
}

func TestTrimPrefix(t *testing.T) {
	token := "token"
	tests := []struct {
		token    string
		expected string
	}{
		{"-" + token, token},
		{"--" + token, token},
		{token, token},
	}

	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertEquals(t, trimPrefix(tt.token), tt.expected)
		})
	}
}
