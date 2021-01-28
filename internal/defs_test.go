package internal

import (
	"strconv"
	"testing"
)

func mockValidCmdArg(cmd, arg string) (string, string) {
	//rCmd, rArg := cmd, arg
	//if strings.HasSuffix(cmd, "doesnt-exist") {
	//	rCmd = ""
	//}
	//if strings.HasSuffix(arg, "doesnt-exist") {
	//	rArg = ""
	//}
	return cmd, arg
}

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
			dc := defs.definedCmd(tt.addedCmd)
			assertValEquals(t, dc, "")

			defs, err := tt.load()
			assertError(t, err, !tt.validLoad)
			assertNil(t, defs, !tt.validLoad)

			if defs != nil {
				dc := defs.definedCmd(tt.addedCmd)
				assertValEquals(t, dc, tt.addedCmd)
			}
		})
	}
}

func TestDefinitionsLoadErrors(t *testing.T) {
	// Load fails with known breaks:
	tests := []struct {
		setup    func(t *testing.T)
		expNil   bool
		expError bool
	}{
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
	for ii, tt := range tests {
		t.Run(tt.cmd, func(t *testing.T) {
			defs := mockDefinitions()
			if ii == 0 {
				defs = nil
			}
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
	for ii, tt := range tests {
		t.Run(tt.cmd+tt.arg, func(t *testing.T) {
			defs := mockDefinitions()
			if ii == 0 {
				defs = nil
			}
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
	for ii, tt := range tests {
		t.Run(tt.cmd+tt.arg+tt.val, func(t *testing.T) {
			defs := mockDefinitions()
			if ii == 0 {
				defs = nil
			}
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
		{"", ""},
		{"command1", "argument1_!$"},
		{"cmd-that-doesnt-exist", ""},
		{"command2", ""},
	}
	for ii, tt := range tests {
		defs := mockDefinitions()
		if ii == 0 {
			defs = nil
		}
		da := defs.defaultArgument(tt.cmd)
		assertValEquals(t, da, tt.expArg)
	}
}
