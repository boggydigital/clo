package internal

import (
	"fmt"
	"strconv"
	"testing"
)

//func TestCreateHelpCommandDefinition(t *testing.T) {
//	assertNil(t, createHelpCommandDefinition(), false)
//}
//
//func TestCreateHelpArgumentDefinition(t *testing.T) {
//	assertNil(t, createHelpArgumentDefinition(""), false)
//}

//func TestAddCommandAbbr(t *testing.T) {
//	tests := []struct {
//		token string
//		cmd   *CommandDefinition
//	}{
//		{"", nil},
//		{"", mockCommandDefinition("c", nil)},
//	}
//	for _, tt := range tests {
//		t.Run(tt.token, func(t *testing.T) {
//			if tt.cmd != nil {
//				assertValNotEquals(t, tt.cmd.Abbr, tt.token)
//			}
//			addCommandAbbr(tt.token, tt.cmd, mockCommandByAbbr)
//			if tt.cmd != nil {
//				assertValEquals(t, tt.cmd.Abbr, tt.token)
//			}
//		})
//	}
//}

//func TestAddArgAbbr(t *testing.T) {
//	tests := []struct {
//		token string
//		arg   *ArgumentDefinition
//	}{
//		{"", nil},
//		{"", mockArgumentDefinition("a", nil)},
//	}
//	for _, tt := range tests {
//		t.Run(tt.token, func(t *testing.T) {
//			if tt.arg != nil {
//				assertValNotEquals(t, tt.arg.Abbr, tt.token)
//			}
//			addArgAbbr(tt.token, tt.arg, mockArgByAbbr)
//			if tt.arg != nil {
//				assertValEquals(t, tt.arg.Abbr, tt.token)
//			}
//		})
//	}
//}

//func TestAddHelpCommand(t *testing.T) {
//	tests := []struct {
//		token      string
//		cmdByToken func(string) *CommandDefinition
//		expNil     bool
//	}{
//		{"", nil, true},
//		{"command", mockCommandByToken, true},
//		{"help", mockCommandByTokenNoHelp, false},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.token, func(t *testing.T) {
//			cmd := addHelpCommand(tt.token, tt.cmdByToken)
//			assertNil(t, cmd, tt.expNil)
//			if cmd != nil {
//				assertValEquals(t, cmd.Token, tt.token)
//			}
//		})
//	}
//}

//func TestAddHelpCommandArgument(t *testing.T) {
//	tests := []struct {
//		token      string
//		argByToken func(string) *ArgumentDefinition
//		expNil     bool
//		expError   bool
//	}{
//		{"", nil, true, false},
//		{"arg", mockArgByToken, true, true},
//		{"help", mockArgByTokenNoHelp, false, false},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.token, func(t *testing.T) {
//			arg, err := addHelpCommandArgument(tt.token, tt.argByToken)
//			assertNil(t, arg, tt.expNil)
//			assertError(t, err, tt.expError)
//			if arg != nil {
//				assertValEquals(t, arg.Token, tt.token)
//			}
//		})
//	}
//}

//func TestTryAddHelpCommand(t *testing.T) {
//	tests := []struct {
//		defs     *Definitions
//		expError bool
//	}{
//		{nil, true},
//		{mockDefinitions(), false},
//		{mockAddHelpCommand(mockDefinitions()), false},
//		{mockAddHelpCommandArgument(mockDefinitions()), true},
//	}
//
//	for ii, tt := range tests {
//		t.Run(strconv.Itoa(ii), func(t *testing.T) {
//			assertError(t, tryAddHelpCommand(tt.defs), tt.expError)
//		})
//	}
//}

//func TestExpandRefValues(t *testing.T) {
//	tests := []struct {
//		args     []ArgumentDefinition
//		commands []CommandDefinition
//		expError bool
//	}{
//		{nil, nil, false},
//		{mockArgumentDefinitions([]string{"arg1", "arg2"}), nil, false},
//		{mockArgumentDefinitions([]string{"from:commands"}), nil, false},
//		{mockArgumentDefinitions([]string{"from:commands"}), mockCommandDefinitions([]string{"c1", "c2"}), false},
//		{mockArgumentDefinitions([]string{"from:error"}), nil, true},
//	}
//	for ii, tt := range tests {
//		t.Run(strconv.Itoa(ii), func(t *testing.T) {
//			err := expandRefValues(tt.args, tt.commands)
//			assertError(t, err, tt.expError)
//		})
//	}
//}

func TestPrintHelp(t *testing.T) {
	//tests := []struct {
	//	cmd      string
	//	defs     *Definitions
	//	expError bool
	//}{
	//	{"", nil, true},
	//	{"", mockDefinitions(), false},
	//	{"command1", mockDefinitions(), false},
	//	{"command-that-doesnt-exist", mockDefinitions(), false},
	//}
	defs := mockDefinitions()
	for _, tt := range mockPrintCommandHelpTests {
		t.Run(tt.token, func(t *testing.T) {
			err := printHelp(tt.token, defs)
			assertError(t, err, defs == nil)
		})
	}
}

func TestPrintAppIntro(t *testing.T) {
	for ii, dd := range mockHelpDefinitionsTests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			printAppIntro(dd)
		})
	}
}

func TestPrintAppUsage(t *testing.T) {
	for ii, dd := range mockHelpDefinitionsTests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			printAppUsage(dd)
		})
	}
}

func TestPrintAppCommands(t *testing.T) {
	for ii, dd := range mockHelpDefinitionsTests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			printAppCommands(dd)
		})
	}
}

//func TestPrintAppAttrsLegend(t *testing.T) {
//	printAppAttrsLegend()
//}

func TestPrintAppMoreInfoPrompt(t *testing.T) {
	for ii, dd := range mockHelpDefinitionsTests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			printAppMoreInfoPrompt(dd)
		})
	}
}

func TestPrintAppHelp(t *testing.T) {
	for ii, dd := range mockHelpDefinitionsTests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			printAppHelp(dd)
		})
	}
}

func TestPrintCmdUsage(t *testing.T) {
	defs := mockDefinitions()
	for _, tt := range mockPrintCommandHelpTests {
		t.Run(tt.token, func(t *testing.T) {
			printCmdUsage(tt.token, defs)
		})
	}
}

func TestPrintArgAttrs(t *testing.T) {
	defs := mockDefinitions()
	for _, tt := range mockPrintArgumentHelpTests {
		t.Run(tt.token, func(t *testing.T) {
			printArgAttrs("", tt.token, defs)
		})
	}
}

func TestPrintArgValues(t *testing.T) {
	defs := mockDefinitions()
	for _, tt := range mockPrintArgumentHelpTests {
		t.Run(tt.token, func(t *testing.T) {
			printArgValues("", tt.token, defs)
		})
	}
}

func TestPrintCmdArgDesc(t *testing.T) {
	defs := mockDefinitions()
	for _, tt := range mockPrintArgumentHelpTests {
		t.Run(tt.token, func(t *testing.T) {
			printCmdArgDesc("", tt.token, defs)
		})
	}
}

func TestPrintCmdArgs(t *testing.T) {
	defs := mockDefinitions()
	for _, tt := range mockPrintCommandHelpTests {
		t.Run(tt.token, func(t *testing.T) {
			printCmdArgs(tt.token, defs)
		})
	}
}

//func TestPrintArgAttrsLegend(t *testing.T) {
//	printArgAttrsLegend()
//}

func TestPrintCmdHelp(t *testing.T) {
	defs := mockDefinitions()
	for _, tt := range mockPrintCommandHelpTests {
		t.Run(fmt.Sprintf("%s", tt.token), func(t *testing.T) {
			printCmdHelp(tt.token, defs)
		})
	}
}
