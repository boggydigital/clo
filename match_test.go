package clo

import (
	"testing"
)

func TestMatchArgument(t *testing.T) {
	tests := []struct {
		token  string
		cmd    string
		expArg string
	}{
		{"a", "c", ""},
		{"-a", "c", "a"},
		{"--a", "c", "a"},
	}

	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			arg, err := matchArg(tt.token, tt.cmd, mockDefinedArg)
			assertError(t, err, false)
			assertValEquals(t, arg, tt.expArg)
		})
	}
}

func TestMatch(t *testing.T) {
	tests := []struct {
		token          string
		tokenType      int
		cmdCtx, argCtx string
		expected       string
		expError       bool
	}{
		{"token", -1, "", "", "", true},
		{"token", -1, "", "", "", true},
		{"command1", command, "", "", "command1_", false},
		{"argument1", argument, "command1", "", "", false},
		{"-argument1", argument, "command1", "", "argument1_!$", false},
		{"--argument1", argument, "command1", "", "argument1_!$", false},
		{"-value1", value, "abc", "argval", "", false},
		{"value1", value, "abc", "argval", "value1", false},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			defs := mockDefinitions()
			m, err := match(tt.token, tt.tokenType, tt.cmdCtx, tt.argCtx, defs)
			assertValEquals(t, m, tt.expected)
			assertError(t, err, tt.expError)
		})
	}
}
