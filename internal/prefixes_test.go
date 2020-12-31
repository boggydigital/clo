package internal

import "testing"

func TestHasPrefix(t *testing.T) {
	names := []string{"double", "single", "empty", "underscore"}
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
		t.Run(names[ii], func(t *testing.T) {
			if hasPrefix(tt.token) != tt.expected {
				t.Error()
			}
		})
	}
}

func TestTrimPrefix(t *testing.T) {
	token := "token"
	names := []string{"single", "double", "no_prefix"}
	tests := []struct {
		token    string
		expected string
	}{
		{"-" + token, token},
		{"--" + token, token},
		{token, token},
	}

	for ii, tt := range tests {
		t.Run(names[ii], func(t *testing.T) {
			if trimPrefix(tt.token) != tt.expected {
				t.Error()
			}
		})
	}
}
