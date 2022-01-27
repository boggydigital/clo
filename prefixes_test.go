package clo

import (
	"github.com/boggydigital/testo"
	"strconv"
	"testing"
)

func TestIsArg(t *testing.T) {
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
			testo.EqualValues(t, isArg(tt.token), tt.expected)
		})
	}
}

func TestTrimArgPrefix(t *testing.T) {
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
			testo.EqualValues(t, trimArgPrefix(tt.token), tt.expected)
		})
	}
}
