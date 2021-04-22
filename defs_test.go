package clo

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"testing"
)

func mockDefinedArg(_, arg string) (string, error) {
	return arg, nil
}

func mockDefinitions() *Definitions {
	return &Definitions{
		Version: 1,
		Cmd: map[string][]string{
			"command1_": {"argument1_!$", "argument2...", "abbr-arg"},
			"command2":  {"argument2...", "xyz"},
			"abc":       {"argval=value1,value2_,abbr-val", "defarg_"},
		},
		Help: map[string]string{
			"command1":           "command1 help",
			"command1:argument1": "command1 argument1 help",
			"command1:argument2": "command1 argument2 help",
			"command1:abbr-arg":  "command1 abbr_arg help",
			"command2":           "command2 help",
			"command2:argument2": "command2 argument2 help",
			"command2:xyz":       "command2 xyz help",
			"abc":                "abc help",
			"abc:argval":         "abc argval help",
			"abc:defarg":         "abc defarg help",
		},
	}
}

var valueDelegates = map[string]func() []string{
	"arguments": func() []string { return []string{"a1", "a2"} },
	"values":    func() []string { return []string{"v1", "v2"} },
}

func mockDefinitionsReplace() *Definitions {
	return &Definitions{
		Version: 1,
		Cmd: map[string][]string{
			"c1": {"{arguments}"},
			"c2": {"arg={values}"},
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
			defs, err := Load(r, nil)
			assertNil(t, defs, tt.expNil)
			assertError(t, err, tt.expErr)
			// check that Load adds help command
			if defs != nil {
				helpCmd, err := defs.definedCmd("help")
				assertError(t, err, false)
				assertValNotEquals(t, helpCmd, "")
			}
		})
	}
}

type valueDelegatesTest struct {
	cmd          string
	expArgs      int
	placeholders bool
}

func TestDefinitionsLoadReplace(t *testing.T) {
	bb, err := json.Marshal(mockDefinitionsReplace())
	assertError(t, err, false)

	testsNoValuesDelegates := []valueDelegatesTest{
		{"c1", 1, true},
		{"c2", 1, true},
	}

	testsValuesDelegates := []valueDelegatesTest{
		{"c1", 2, false},
		{"c2", 1, false},
	}

	valueDelegatesTests := []struct {
		name                string
		delegates           map[string]func() []string
		valueDelegatesTests []valueDelegatesTest
	}{
		{"no-vd-", nil, testsNoValuesDelegates},
		{"vd-", valueDelegates, testsValuesDelegates},
	}

	for _, vdt := range valueDelegatesTests {
		def, err := Load(bytes.NewReader(bb), vdt.delegates)
		assertError(t, err, false)

		for _, tt := range vdt.valueDelegatesTests {
			t.Run(vdt.name+tt.cmd, func(t *testing.T) {
				assertValEquals(t, len(def.Cmd[tt.cmd]), tt.expArgs)
				assertValEquals(t, strings.Contains(def.Cmd[tt.cmd][0], placeholderPrefix), tt.placeholders)
				assertValEquals(t, strings.Contains(def.Cmd[tt.cmd][0], placeholderSuffix), tt.placeholders)
			})
		}
	}

}

func TestDefinitionsDefinedCmd(t *testing.T) {
	tests := []struct {
		cmd    string
		expCmd string
		expErr bool
	}{
		{"cmd-that-doesnt-exist", "", false}, // used to test defs == nil
		{"cmd-that-doesnt-exist", "", false},
		{"command1", "command1_", false},
		{"a", "abc", false},
		{"ab", "abc", false},
		{"abc", "abc", false},
		{"c", "", true},
		{"command", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.cmd, func(t *testing.T) {
			defs := mockDefinitions()
			dc, err := defs.definedCmd(tt.cmd)
			assertError(t, err, tt.expErr)
			assertValEquals(t, dc, tt.expCmd)
		})
	}
}

func TestDefinitionsDefinedCmdArg(t *testing.T) {
	tests := []struct {
		cmd, arg string
		expArg   string
	}{
		{"cmd-that-doesnt-exist", "arg-that-doesnt-exist", ""}, // used to test defs == nil
		{"cmd-that-doesnt-exist", "arg-that-doesnt-exist", ""},
		{"command1", "argument1", "argument1_!$"},
		{"command1", "argument-that-doesnt-exist", ""},
	}
	for _, tt := range tests {
		t.Run(tt.cmd+tt.arg, func(t *testing.T) {
			defs := mockDefinitions()
			da, err := defs.definedArg(tt.cmd, tt.arg)
			assertError(t, err, false)
			assertValEquals(t, da, tt.expArg)
		})
	}
}

func TestDefinitionsDefinedCmdArgVal(t *testing.T) {
	tests := []struct {
		cmd, arg, val string
		expVal        string
	}{
		{"cmd-that-doesnt-exist", "arg-that-doesnt-exist", "value1", ""},
		{"cmd-that-doesnt-exist", "arg-that-doesnt-exist", "value1", ""},
		{"command1", "argument1", "", ""},
		{"abc", "argval", "value1", "value1"},
	}
	for _, tt := range tests {
		t.Run(tt.cmd+tt.arg+tt.val, func(t *testing.T) {
			defs := mockDefinitions()
			dv, err := defs.definedVal(tt.cmd, tt.arg, tt.val)
			assertError(t, err, false)
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
		da, err := defs.defaultArgument(tt.cmd)
		assertError(t, err, false)
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
		{&Request{}, &Request{}, false},
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

func TestDefaultArgByNameNotValues(t *testing.T) {
	tests := []struct {
		cmd    string
		extArg string
		expErr bool
	}{
		{"abc", "defarg_", false},
	}
	defs := mockDefinitions()
	for _, tt := range tests {
		t.Run(tt.cmd+tt.extArg, func(t *testing.T) {
			defArg, err := defs.defaultArgument(tt.cmd)
			assertError(t, err, tt.expErr)
			assertValEquals(t, defArg, tt.extArg)
		})
	}
}
