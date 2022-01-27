package clo

import (
	"bytes"
	"encoding/json"
	"github.com/boggydigital/testo"
	"strconv"
	"strings"
	"testing"
)

func mockDefinedArg(_, arg string) (string, error) {
	return arg, nil
}

func mockDefinitions() *definitions {
	return &definitions{
		Version: 1,
		Cmd: map[string][]string{
			"command1" + defaultAttr: {"argument1" + defaultAttr + requiredAttr + envAttr, "argument2" + multipleAttr, "abbr-arg"},
			"command2":               {"argument2" + multipleAttr, "xyz"},
			"abc":                    {"argval" + argValuesSep + "value1,value2" + defaultAttr + ",abbr-val", "defarg" + defaultAttr},
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

func mockDefinitionsReplace() *definitions {
	return &definitions{
		Version: 1,
		Cmd: map[string][]string{
			"c1": {"{arguments}"},
			"c2": {"arg" + argValuesSep + "{values}"},
		},
	}
}

func mockDefinitionsNoDefaults() *definitions {
	return &definitions{
		Version: 1,
		Cmd: map[string][]string{
			"command1": {"argument1" + argValuesSep + "value1,value2", "argument2" + argValuesSep + "value3,value4"},
		},
	}
}

func TestDefinitionsLoad(t *testing.T) {
	bytes, err := json.Marshal(mockDefinitions())
	testo.Error(t, err, false)

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
			testo.Nil(t, defs, tt.expNil)
			testo.Error(t, err, tt.expErr)
			// check that Load adds help command
			if defs != nil {
				helpCmd, err := defs.definedCmd("help")
				testo.Error(t, err, false)
				testo.UnequalValues(t, helpCmd, "")
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
	testo.Error(t, err, false)

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
		expError            bool
	}{
		{"no-vd-", nil, testsNoValuesDelegates, false},
		{"empty-vd-", map[string]func() []string{}, nil, true},
		{"vd-", valueDelegates, testsValuesDelegates, false},
	}

	for _, vdt := range valueDelegatesTests {
		def, err := Load(bytes.NewReader(bb), vdt.delegates)
		testo.Error(t, err, vdt.expError)

		for _, tt := range vdt.valueDelegatesTests {
			t.Run(vdt.name+tt.cmd, func(t *testing.T) {
				testo.EqualValues(t, len(def.Cmd[tt.cmd]), tt.expArgs)
				testo.EqualValues(t, strings.Contains(def.Cmd[tt.cmd][0], placeholderPrefix), tt.placeholders)
				testo.EqualValues(t, strings.Contains(def.Cmd[tt.cmd][0], placeholderSuffix), tt.placeholders)
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
		{"command1", "command1" + defaultAttr, false},
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
			testo.Error(t, err, tt.expErr)
			testo.EqualValues(t, dc, tt.expCmd)
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
		{"command1", "argument1", "argument1" + defaultAttr + requiredAttr + envAttr},
		{"command1", "argument-that-doesnt-exist", ""},
	}
	for _, tt := range tests {
		t.Run(tt.cmd+tt.arg, func(t *testing.T) {
			defs := mockDefinitions()
			da, err := defs.definedArg(tt.cmd, tt.arg)
			testo.Error(t, err, false)
			testo.EqualValues(t, da, tt.expArg)
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
			testo.Error(t, err, false)
			testo.EqualValues(t, dv, tt.expVal)
		})
	}
}

func TestDefinitionsDefaultCommand(t *testing.T) {
	tests := []struct {
		defs   *definitions
		expCmd string
	}{
		{nil, ""},
		{mockDefinitions(), "command1" + defaultAttr},
		{&definitions{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.expCmd, func(t *testing.T) {
			dc := tt.defs.defaultCommand()
			testo.EqualValues(t, dc, tt.expCmd)
		})
	}
}

func TestDefinitionsDefaultArgument(t *testing.T) {
	tests := []struct {
		cmd    string
		expArg string
	}{
		{"command1", "argument1" + defaultAttr + requiredAttr + envAttr},
		{"cmd-that-doesnt-exist", ""},
		{"command2", ""},
	}
	for _, tt := range tests {
		defs := mockDefinitions()
		da, err := defs.defaultArgument(tt.cmd)
		testo.Error(t, err, false)
		testo.EqualValues(t, da, tt.expArg)
	}
}

func TestDefinitionsDefaultArgumentValues(t *testing.T) {
	tests := []struct {
		req    *request
		expReq *request
		expErr bool
	}{
		{nil, nil, true},
		{&request{}, &request{}, false},
		{&request{Command: "command1"}, &request{Command: "command1", Arguments: map[string][]string{}}, false},
		{&request{Command: "abc"}, &request{
			Command: "abc",
			Arguments: map[string][]string{
				"argval": {"value2"},
			},
		}, false},
		{&request{
			Command: "abc",
			Arguments: map[string][]string{
				"argval": {"value1"}},
		},
			&request{
				Command: "abc",
				Arguments: map[string][]string{
					"argval": {"value1"}},
			}, false},
	}
	defs := mockDefinitions()
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			testo.Error(t, defs.defaultArgValues(tt.req), tt.expErr)
			testo.DeepEqual(t, tt.req, tt.expReq)
		})
	}
}

func TestDefaultArgByNameNotValues(t *testing.T) {
	tests := []struct {
		cmd    string
		extArg string
		expErr bool
	}{
		{"abc", "defarg" + defaultAttr, false},
	}
	defs := mockDefinitions()
	for _, tt := range tests {
		t.Run(tt.cmd+tt.extArg, func(t *testing.T) {
			defArg, err := defs.defaultArgument(tt.cmd)
			testo.Error(t, err, tt.expErr)
			testo.EqualValues(t, defArg, tt.extArg)
		})
	}
}
