package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
)

const defaultFilename = "clo.json"

func testDefs() *Definitions {
	defs := &Definitions{
		Version:   1,
		EnvPrefix: "CORRECT",
		App:       "clo",
		Hint:      "hint",
		Desc:      "desc",
		Flags: []FlagDefinition{
			{
				CommonDefinition: CommonDefinition{
					Token: "flag1",
					Abbr:  "f1",
				},
			},
			{
				CommonDefinition: CommonDefinition{
					Token: "flag2",
					Abbr:  "f2",
				},
			},
		},
		Commands: []CommandDefinition{
			{
				CommonDefinition: CommonDefinition{
					Token: "command1",
					Abbr:  "c1",
				},
				Arguments: []string{
					"argument1",
					"argument2",
				},
			},
			{
				CommonDefinition: CommonDefinition{
					Token: "command2",
				},
			},
		},
		Arguments: []ArgumentDefinition{
			{
				CommonDefinition: CommonDefinition{
					Token: "argument1",
					Abbr:  "a1",
				},
				Default:  true,
				Multiple: true,
				Required: true,
				Values:   []string{"value1", "value2"},
			},
			{
				CommonDefinition: CommonDefinition{
					Token: "argument2",
					Abbr:  "a2",
				},
				Values: []string{"value3", "value4"},
			},
		},
		Values: []ValueDefinition{
			{
				CommonDefinition: CommonDefinition{
					Token: "value1",
				},
			},
		},
	}

	return defs
}

func breakDefinitions(defs *Definitions) {
	defs.Arguments = append(defs.Arguments, ArgumentDefinition{
		CommonDefinition: CommonDefinition{
			Token: "help:command",
		},
		Values: []string{"from:nowhere"},
	})
}

func writeDefs(defs *Definitions, t *testing.T) {
	if _, err := os.Stat(defaultFilename); os.IsNotExist(err) {
		jsonBytes, err := json.Marshal(defs)
		if err != nil {
			t.FailNow()
		}
		err = ioutil.WriteFile(defaultFilename, jsonBytes, 0644)
		if err != nil {
			t.FailNow()
		}
	} else {
		t.Error()
	}
}

func writeEmptyDefs(t *testing.T) {
	if _, err := os.Stat(defaultFilename); os.IsNotExist(err) {
		err = ioutil.WriteFile(defaultFilename, []byte{}, 0644)
		if err != nil {
			t.Error()
		}
	} else {
		t.Error()
	}
}

func deleteDefs() {
	if os.Remove(defaultFilename) != nil {
		log.Printf("cannot remove test definitions file at %s", defaultFilename)
	}
}

func loadPathThatDoesntExist() (*Definitions, error) {
	return Load("path/that/doesnt/exist")
}

func TestLoad(t *testing.T) {
	names := []string{"load-adds-help", "load-at-a-path-that-doesnt-exist"}
	tests := []struct {
		load      func() (*Definitions, error)
		validLoad bool
		addedCmd  string
	}{
		{LoadDefault, true, "help"},
		{loadPathThatDoesntExist, false, "help"},
	}

	// Load adds 'help' command
	defs := testDefs()
	writeDefs(defs, t)
	t.Cleanup(deleteDefs)

	for ii, tt := range tests {
		t.Run(names[ii], func(t *testing.T) {
			// command shouldn't exist before we add it at load
			cmd := defs.CommandByToken(tt.addedCmd)
			if cmd != nil {
				t.Error()
			}

			defs, err := tt.load()
			if tt.validLoad && err != nil {
				t.Error(err.Error())
			}
			if (defs == nil && tt.validLoad) || (defs != nil && !tt.validLoad) {
				t.FailNow()
			}

			if defs != nil {
				cmd := defs.CommandByToken(tt.addedCmd)
				if cmd == nil {
					t.Error()
				}
			}
		})
	}
}

func setupBrokenDefs(t *testing.T) {
	defs := testDefs()
	breakDefinitions(defs)
	writeDefs(defs, t)
	t.Cleanup(deleteDefs)
}

func setupEmptyDefs(t *testing.T) {
	writeEmptyDefs(t)
	t.Cleanup(deleteDefs)
}

func TestLoadErrors(t *testing.T) {
	// Load fails with known breaks:
	// - Pre-existing "help:command" argument
	// - Pre-existing "from:nowhere" reference value
	names := []string{"broken_defs", "empty_defs"}
	tests := []struct {
		setup func(t *testing.T)
	}{
		{setupBrokenDefs},
		{setupEmptyDefs},
	}

	for ii, tt := range tests {
		t.Run(names[ii], func(t *testing.T) {
			tt.setup(t)
			if defs, err := LoadDefault(); defs != nil || err == nil {
				t.Error()
			}
		})
	}
}

type TokenTest struct {
	token  string
	expNil bool
}

func genByTokenAbbrTests(prefix string) []TokenTest {
	return []TokenTest{
		// valid token/abbr
		{prefix + "1", false},
		// invalid token/abbr
		{prefix + "-token-that-doesnt-exist", true},
	}
}

func assertNil(t *testing.T, v interface{}, expNil bool) {
	val := reflect.ValueOf(v)
	if (val.IsNil() && !expNil) || (!val.IsNil() && expNil) {
		t.Error()
	}
}

func TestFlagByToken(t *testing.T) {
	defs := testDefs()
	for _, tt := range genByTokenAbbrTests("flag") {
		t.Run(tt.token, func(t *testing.T) {
			fd := defs.FlagByToken(tt.token)
			assertNil(t, fd, tt.expNil)
		})
	}
}

func TestFlagByAbbr(t *testing.T) {
	defs := testDefs()
	for _, tt := range genByTokenAbbrTests("f") {
		t.Run(tt.token, func(t *testing.T) {
			fd := defs.FlagByAbbr(tt.token)
			assertNil(t, fd, tt.expNil)
		})
	}
}

func TestCommandByToken(t *testing.T) {
	defs := testDefs()
	for _, tt := range genByTokenAbbrTests("command") {
		t.Run(tt.token, func(t *testing.T) {
			cd := defs.CommandByToken(tt.token)
			assertNil(t, cd, tt.expNil)
		})
	}
}

func TestCommandByAbbr(t *testing.T) {
	defs := testDefs()
	for _, tt := range genByTokenAbbrTests("c") {
		t.Run(tt.token, func(t *testing.T) {
			cd := defs.CommandByAbbr(tt.token)
			assertNil(t, cd, tt.expNil)
		})
	}
}

func TestArgByToken(t *testing.T) {
	defs := testDefs()
	for _, tt := range genByTokenAbbrTests("argument") {
		t.Run(tt.token, func(t *testing.T) {
			cd := defs.ArgByToken(tt.token)
			assertNil(t, cd, tt.expNil)
		})
	}
}

func TestArgByAbbr(t *testing.T) {
	defs := testDefs()
	for _, tt := range genByTokenAbbrTests("a") {
		t.Run(tt.token, func(t *testing.T) {
			cd := defs.ArgByAbbr(tt.token)
			assertNil(t, cd, tt.expNil)
		})
	}
}

func TestValueBy(t *testing.T) {
	defs := testDefs()
	for _, tt := range genByTokenAbbrTests("value") {
		t.Run(tt.token, func(t *testing.T) {
			cd := defs.ValueByToken(tt.token)
			assertNil(t, cd, tt.expNil)
		})
	}
}

func TestDefinedValue(t *testing.T) {
	defs := testDefs()
	for ii, tt := range validityTests {
		t.Run(validityNames[ii], func(t *testing.T) {
			if defs.DefinedValue(tt.values) != tt.expected {
				t.Error()
			}
		})
	}
}

func TestDefaultArg(t *testing.T) {
	defs := testDefs()
	tests := []struct {
		cmd      *CommandDefinition
		validCmd bool
		args     []string
		expNil   bool
	}{
		{nil, false, nil, true},
		{
			defs.CommandByToken("command1"),
			true,
			[]string{"argument-that-doesnt-exist", "argument1", "argument2"},
			false,
		},
		{defs.CommandByToken("command2"), true, nil, true},
	}

	for _, tt := range tests {
		name := "nil"
		if tt.cmd != nil {
			name = tt.cmd.Token
		}
		t.Run(name, func(t *testing.T) {
			if tt.validCmd && tt.cmd == nil {
				t.Error()
			}
			if tt.validCmd && tt.args != nil {
				tt.cmd.Arguments = tt.args
			}
			ad := defs.DefaultArg(tt.cmd)
			assertNil(t, ad, tt.expNil)
		})
	}
}

func TestRequiredArgs(t *testing.T) {
	defs := testDefs()
	tests := []struct {
		cmd          string
		requiredArgs int
	}{
		{"command-that-doesnt-exist", 0},
		{defs.Commands[0].Token, 1},
	}
	// this is required to hit a "if arg == nil {" condition
	defs.Commands[0].Arguments = append(defs.Commands[0].Arguments, "argument-that-doesnt-exist")
	for _, tt := range tests {
		t.Run(tt.cmd, func(t *testing.T) {
			if len(defs.RequiredArgs(tt.cmd)) != tt.requiredArgs {
				t.Error()
			}
		})
	}
}

func TestValidArgVal(t *testing.T) {
	tests := []struct {
		arg      string
		val      string
		expected bool
	}{
		{"", "", false},
		{"argument-that-doesnt-exist", "", false},
		{"argument1", "value1", true},
	}
	defs := testDefs()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s-%s", tt.arg, tt.val), func(t *testing.T) {
			if defs.ValidArgVal(tt.arg, tt.val) != tt.expected {
				t.Error()
			}
		})
	}
}
