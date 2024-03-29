package clo

import (
	"github.com/boggydigital/testo"
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
			testo.Error(t, err, false)
			testo.EqualValues(t, arg, tt.expArg)
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
		{"command1", command, "", "", "command1" + defaultAttr, false},
		{"a", command, "", "", "abc", false},
		{"argument1", argument, "command1", "", "", false},
		{"-argument1", argument, "command1", "", "argument1" + defaultAttr + requiredAttr + envAttr, false},
		{"--argument1", argument, "command1", "", "argument1" + defaultAttr + requiredAttr + envAttr, false},
		{"-ab", argument, "command1", "", "abbr-arg", false},
		{"--ab", argument, "command1", "", "abbr-arg", false},
		{"-value1", value, "abc", "argval", "", false},
		{"value1", value, "abc", "argval", "value1", false},
		{"ab", value, "abc", "argval", "abbr-val", false},
		{"val", value, "command2", "argument2", "val", false},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			defs := mockDefinitions()
			m, err := match(tt.token, tt.tokenType, tt.cmdCtx, tt.argCtx, defs)
			testo.EqualValues(t, m, tt.expected)
			testo.Error(t, err, tt.expError)
		})
	}
}
