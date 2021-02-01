package internal

import (
	"strconv"
	"testing"
)

func mockDefinitions() *Definitions {
	return &Definitions{
		Version: 1,
		Cmd: map[string][]string{
			"command1_": {"argument1_!$", "argument2..."},
			"command2":  {"argument2...", "xyz"},
			"abc":       {"argval=value1,value2"},
		},
		Help: map[string]string{
			"command1":           "command1 help",
			"command1:argument1": "command1 argument1 help",
			"command1:argument2": "command1 argument2 help",
			"command2":           "command2 help",
			"command2:argument2": "command2 argument2 help",
			"command2:xyz":       "command2 xyz help",
			"abc":                "abc help",
			"abc:argval":         "abc argval help",
		},
	}
}

func mockDefinitionsNoDefaults() *Definitions {
	return &Definitions{
		Version: 1,
		Cmd: map[string][]string{
			"command1": {"argument1=value1,value2", "argument2=value3,value4"},
		},
	}
}

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
