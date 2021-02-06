package clo

import (
	"strconv"
	"testing"
)

func TestDefinitionsParse(t *testing.T) {
	tests := []struct {
		args       []string
		expCmd     string
		expLastArg string
		expLenArgs int
		expErr     bool
	}{
		{[]string{"", "command1", "--argument1"}, "command1", "argument1", 1, false},
		{[]string{"--argument1"}, "command1", "argument1", 1, false},
		{[]string{"unknown-token"}, "command1", "argument1", 1, false},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			defs := mockDefinitions()
			req, err := defs.Parse(tt.args)
			assertError(t, err, tt.expErr)
			assertValEquals(t, req.Command, tt.expCmd)
			assertValEquals(t, req.lastArgument, tt.expLastArg)
			assertValEquals(t, len(req.Arguments), tt.expLenArgs)
		})
	}
}

func TestDefinitionsNoDefaultsParse(t *testing.T) {
	tests := []struct {
		args       []string
		expCmd     string
		expLastArg string
		expLenArgs int
		expErr     bool
	}{
		{[]string{"command2"}, "", "", 0, true},
		{[]string{"command1", "argument1", "value-that-doesnt-exist"}, "command1", "", 0, true},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			defs := mockDefinitionsNoDefaults()
			req, err := defs.Parse(tt.args)
			assertError(t, err, tt.expErr)
			assertValEquals(t, req.Command, tt.expCmd)
			assertValEquals(t, req.lastArgument, tt.expLastArg)
			assertValEquals(t, len(req.Arguments), tt.expLenArgs)
		})
	}
}
