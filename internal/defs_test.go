package internal

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestDefinitionsLoad(t *testing.T) {
	tests := []struct {
		load      func() (*Definitions, error)
		validLoad bool
		addedCmd  string
	}{
		{LoadDefault, true, "help"},
		{loadMockPathThatDoesntExist, false, "help"},
	}

	// Load adds 'help' command
	defs := mockDefinitions()
	writeMockDefs(defs, t)
	t.Cleanup(deleteMockDefs)

	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			// command shouldn't exist before we add it at load
			cmd := defs.CommandByToken(tt.addedCmd)
			assertNil(t, cmd, true)

			defs, err := tt.load()
			assertError(t, err, !tt.validLoad)
			assertNil(t, defs, !tt.validLoad)

			if defs != nil {
				cmd := defs.CommandByToken(tt.addedCmd)
				assertNil(t, cmd, false)
			}
		})
	}
}

func TestDefinitionsLoadErrors(t *testing.T) {
	// Load fails with known breaks:
	// - Pre-existing "help:command" argument
	// - Pre-existing "from:nowhere" reference value
	tests := []struct {
		setup    func(t *testing.T)
		expNil   bool
		expError bool
	}{
		{setupBrokenMockDefs, true, true},
		{setupEmptyMockDefs, true, true},
	}

	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			tt.setup(t)
			defs, err := LoadDefault()
			assertNil(t, defs, tt.expNil)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDefinitionsCommandByToken(t *testing.T) {
	defs := mockDefinitions()
	for _, tt := range mockByTokenAbbrTests("command") {
		t.Run(tt.token, func(t *testing.T) {
			cd := defs.CommandByToken(tt.token)
			assertNil(t, cd, tt.expNil)
		})
	}
}

func TestDefinitionsCommandByAbbr(t *testing.T) {
	defs := mockDefinitions()
	for _, tt := range mockByTokenAbbrTests("c") {
		t.Run(tt.token, func(t *testing.T) {
			cd := defs.CommandByAbbr(tt.token)
			assertNil(t, cd, tt.expNil)
		})
	}
}

func TestDefinitionsArgByToken(t *testing.T) {
	defs := mockDefinitions()
	for _, tt := range mockByTokenAbbrTests("argument") {
		t.Run(tt.token, func(t *testing.T) {
			cd := defs.ArgByToken(tt.token)
			assertNil(t, cd, tt.expNil)
		})
	}
}

func TestDefinitionsArgByAbbr(t *testing.T) {
	defs := mockDefinitions()
	for _, tt := range mockByTokenAbbrTests("a") {
		t.Run(tt.token, func(t *testing.T) {
			cd := defs.ArgByAbbr(tt.token)
			assertNil(t, cd, tt.expNil)
		})
	}
}

func TestDefinitionsValueByToken(t *testing.T) {
	defs := mockDefinitions()
	for _, tt := range mockByTokenAbbrTests("value") {
		t.Run(tt.token, func(t *testing.T) {
			cd := defs.ValueByToken(tt.token)
			assertNil(t, cd, tt.expNil)
		})
	}
}

func TestDefinitionsDefinedValue(t *testing.T) {
	defs := mockDefinitions()
	for _, tt := range mockValidityTests {
		t.Run(strings.Join(tt.values, "-"), func(t *testing.T) {
			assertValEquals(t, defs.DefinedValue(tt.values), tt.expected)
		})
	}
}

func TestDefinitionsDefaultArg(t *testing.T) {
	defs := mockDefinitions()
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
			assertNil(t, tt.cmd, !tt.validCmd)
			if tt.validCmd && tt.args != nil {
				tt.cmd.Arguments = tt.args
			}
			ad := defs.DefaultArg(tt.cmd)
			assertNil(t, ad, tt.expNil)
		})
	}
}

func TestDefinitionsRequiredArgs(t *testing.T) {
	defs := mockDefinitions()
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
			assertValEquals(t, len(defs.RequiredArgs(tt.cmd)), tt.requiredArgs)
		})
	}
}

func TestDefinitionsValidArgVal(t *testing.T) {
	tests := []struct {
		arg      string
		val      string
		expected bool
	}{
		{"", "", false},
		{"argument-that-doesnt-exist", "", false},
		{"argument1", "value1", true},
	}
	defs := mockDefinitions()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s-%s", tt.arg, tt.val), func(t *testing.T) {
			assertValEquals(t, defs.ValidArgVal(tt.arg, tt.val), tt.expected)
		})
	}
}
