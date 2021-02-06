package clo

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"
)

func mockValidCmdArg(cmd, arg string) (string, string) {
	return cmd, arg
}

func mockDefinitions() *Definitions {
	return &Definitions{
		Version: 1,
		Cmd: map[string][]string{
			"command1_": {"argument1_!$", "argument2..."},
			"command2":  {"argument2...", "xyz"},
			"abc":       {"argval=value1,value2_"},
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

func TestDefinitionsLoad(t *testing.T) {
	bytes, err := json.Marshal(mockDefinitions())
	assertError(t, err, false)

	tests := []struct {
		content string
		expNil  bool
		expErr  bool
	}{
		{"", true, true},
		{string(bytes), false, false},
	}

	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			r := strings.NewReader(tt.content)
			defs, err := Load(r)
			assertNil(t, defs, tt.expNil)
			assertError(t, err, tt.expErr)
			// check that Load adds help command
			if defs != nil {
				helpCmd := defs.definedCmd("help")
				assertValNotEquals(t, helpCmd, "")
			}
		})
	}
}

func TestDefinitionsDefinedCmd(t *testing.T) {
	tests := []struct {
		cmd    string
		expCmd string
	}{
		{"cmd-that-doesnt-exist", ""}, // used to test defs == nil
		{"cmd-that-doesnt-exist", ""},
		{"command1", "command1_"},
		{"a", "abc"},
		{"ab", "abc"},
		{"abc", "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.cmd, func(t *testing.T) {
			defs := mockDefinitions()
			dc := defs.definedCmd(tt.cmd)
			assertValEquals(t, dc, tt.expCmd)
		})
	}
}

func TestDefinitionsDefinedCmdArg(t *testing.T) {
	tests := []struct {
		cmd, arg       string
		expCmd, expArg string
	}{
		{"cmd-that-doesnt-exist", "arg-that-doesnt-exist", "", ""}, // used to test defs == nil
		{"cmd-that-doesnt-exist", "arg-that-doesnt-exist", "", ""},
		{"command1", "argument1", "command1_", "argument1_!$"},
		{"command1", "argument-that-doesnt-exist", "command1_", ""},
	}
	for _, tt := range tests {
		t.Run(tt.cmd+tt.arg, func(t *testing.T) {
			defs := mockDefinitions()
			dc, da := defs.definedCmdArg(tt.cmd, tt.arg)
			assertValEquals(t, dc, tt.expCmd)
			assertValEquals(t, da, tt.expArg)
		})
	}
}

func TestDefinitionsDefinedCmdArgVal(t *testing.T) {
	tests := []struct {
		cmd, arg, val          string
		expCmd, expArg, expVal string
	}{
		{"cmd-that-doesnt-exist", "arg-that-doesnt-exist", "value1", "", "", ""},
		{"cmd-that-doesnt-exist", "arg-that-doesnt-exist", "value1", "", "", ""},
		{"command1", "argument1", "", "command1_", "argument1_!$", ""},
		{"abc", "argval", "value1", "abc", "argval", "value1"},
	}
	for _, tt := range tests {
		t.Run(tt.cmd+tt.arg+tt.val, func(t *testing.T) {
			defs := mockDefinitions()
			dc, da, dv := defs.definedCmdArgVal(tt.cmd, tt.arg, tt.val)
			assertValEquals(t, dc, tt.expCmd)
			assertValEquals(t, da, tt.expArg)
			assertValEquals(t, dv, tt.expVal)
		})
	}
}

func TestDefinitionsDefaultCommand(t *testing.T) {
	tests := []struct {
		defs   *Definitions
		expCmd string
	}{
		{nil, ""},
		{mockDefinitions(), "command1_"},
		{&Definitions{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.expCmd, func(t *testing.T) {
			dc := tt.defs.defaultCommand()
			assertValEquals(t, dc, tt.expCmd)
		})
	}
}

func TestDefinitionsDefaultArgument(t *testing.T) {
	tests := []struct {
		cmd    string
		expArg string
	}{
		{"command1", "argument1_!$"},
		{"cmd-that-doesnt-exist", ""},
		{"command2", ""},
	}
	for _, tt := range tests {
		defs := mockDefinitions()
		da := defs.defaultArgument(tt.cmd)
		assertValEquals(t, da, tt.expArg)
	}
}

func TestDefinitionsDefaultArgumentValues(t *testing.T) {
	tests := []struct {
		req    *Request
		expReq *Request
		expErr bool
	}{
		{nil, nil, true},
		{&Request{}, &Request{}, true},
		{&Request{Command: "command1"}, &Request{Command: "command1", Arguments: map[string][]string{}}, false},
		{&Request{Command: "abc"}, &Request{
			Command: "abc",
			Arguments: map[string][]string{
				"argval": {"value2"},
			},
		}, false},
		{&Request{
			Command: "abc",
			Arguments: map[string][]string{
				"argval": {"value1"}},
		},
			&Request{
				Command: "abc",
				Arguments: map[string][]string{
					"argval": {"value1"}},
			}, false},
	}
	defs := mockDefinitions()
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertError(t, defs.defaultArgValues(tt.req), tt.expErr)
			assertInterfaceEquals(t, tt.req, tt.expReq)
		})
	}
}
