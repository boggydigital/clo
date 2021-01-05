package internal

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCreateHelpCommandDefinition(t *testing.T) {
	assertNil(t, createHelpCommandDefinition(), false)
}

func TestCreateHelpArgumentDefinition(t *testing.T) {
	assertNil(t, createHelpArgumentDefinition(""), false)
}

func TestAddCommandAbbr(t *testing.T) {
	tests := []struct {
		token string
		cmd   *CommandDefinition
	}{
		{"", nil},
		{"", mockCommandDefinition("c", nil)},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			if tt.cmd != nil {
				assertValNotEquals(t, tt.cmd.Abbr, tt.token)
			}
			addCommandAbbr(tt.token, tt.cmd, mockCommandByAbbr)
			if tt.cmd != nil {
				assertValEquals(t, tt.cmd.Abbr, tt.token)
			}
		})
	}
}

func TestAddArgAbbr(t *testing.T) {
	tests := []struct {
		token string
		arg   *ArgumentDefinition
	}{
		{"", nil},
		{"", mockArgumentDefinition("a", nil)},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			if tt.arg != nil {
				assertValNotEquals(t, tt.arg.Abbr, tt.token)
			}
			addArgAbbr(tt.token, tt.arg, mockArgByAbbr)
			if tt.arg != nil {
				assertValEquals(t, tt.arg.Abbr, tt.token)
			}
		})
	}
}

func TestAddHelpCommand(t *testing.T) {
	tests := []struct {
		token      string
		cmdByToken func(string) *CommandDefinition
		expNil     bool
	}{
		{"", nil, true},
		{"command", mockCommandByToken, true},
		{"help", mockCommandByTokenNoHelp, false},
	}

	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			cmd := addHelpCommand(tt.token, "", tt.cmdByToken, mockCommandByAbbr)
			assertNil(t, cmd, tt.expNil)
			if cmd != nil {
				assertValEquals(t, cmd.Token, tt.token)
			}
		})
	}
}

func TestAddHelpCommandArgument(t *testing.T) {
	tests := []struct {
		token      string
		argByToken func(string) *ArgumentDefinition
		expNil     bool
		expError   bool
	}{
		{"", nil, true, false},
		{"arg", mockArgByToken, true, true},
		{"help", mockArgByTokenNoHelp, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			arg, err := addHelpCommandArgument(tt.token, "", tt.argByToken, mockArgByAbbr)
			assertNil(t, arg, tt.expNil)
			assertError(t, err, tt.expError)
			if arg != nil {
				assertValEquals(t, arg.Token, tt.token)
			}
		})
	}
}

func TestTryAddHelpCommand(t *testing.T) {
	tests := []struct {
		defs     *Definitions
		expError bool
	}{
		{nil, true},
		{mockDefinitions(), false},
		{mockAddHelpCommand(mockDefinitions()), false},
		{mockAddHelpCommandArgument(mockDefinitions()), true},
	}

	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertError(t, tryAddHelpCommand(tt.defs), tt.expError)
		})
	}
}

func TestExpandRefValues(t *testing.T) {
	tests := []struct {
		args     []ArgumentDefinition
		commands []CommandDefinition
		expError bool
	}{
		{nil, nil, false},
		{mockArgumentDefinitions([]string{"arg1", "arg2"}), nil, false},
		{mockArgumentDefinitions([]string{"from:commands"}), nil, false},
		{mockArgumentDefinitions([]string{"from:commands"}), mockCommandDefinitions([]string{"c1", "c2"}), false},
		{mockArgumentDefinitions([]string{"from:error"}), nil, true},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			err := expandRefValues(tt.args, tt.commands)
			assertError(t, err, tt.expError)
		})
	}
}

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
	for _, tt := range mockPrintCommandHelpTests {
		t.Run(tt.token, func(t *testing.T) {
			err := printHelp(tt.token, tt.defs, false)
			assertError(t, err, tt.defs == nil)
		})
	}
}

func TestPrintAppIntro(t *testing.T) {
	for ii, dd := range mockHelpDefinitionsTests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			printAppIntro(dd, true)
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
			printAppCommands(dd, true)
		})
	}
}

func TestPrintAppFlags(t *testing.T) {
	for ii, dd := range mockHelpDefinitionsTests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			printAppFlags(dd, true)
		})
	}
}

func TestPrintAppAttrsLegend(t *testing.T) {
	printAppAttrsLegend()
}

func TestPrintAppMoreInfoPrompt(t *testing.T) {
	for ii, dd := range mockHelpDefinitionsTests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			printAppMoreInfoPrompt(dd, false)
		})
	}
}

func TestPrintAppHelp(t *testing.T) {
	for ii, dd := range mockHelpDefinitionsTests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			printAppHelp(dd, true)
		})
	}
}

func TestPrintExampleHelp(t *testing.T) {
	for di, dd := range mockHelpDefinitionsTests {
		for ti, tt := range mockEmptyExamplesTests {
			t.Run(fmt.Sprintf("%d-%d", di, ti), func(t *testing.T) {
				printExampleHelp(mockExampleDefinition("", tt.tokens), "", dd)
			})
		}
	}
}

func TestPrintCmdUsage(t *testing.T) {
	for _, tt := range mockPrintCommandHelpTests {
		t.Run(tt.token, func(t *testing.T) {
			printCmdUsage(tt.token, tt.defs)
		})
	}
}

func TestPrintArgAttrs(t *testing.T) {
	for _, tt := range mockPrintArgumentHelpTests {
		t.Run(tt.token, func(t *testing.T) {
			printArgAttrs("", tt.token, tt.defs)
		})
	}
}

func TestPrintArgValues(t *testing.T) {
	for _, tt := range mockPrintArgumentHelpTests {
		t.Run(tt.token, func(t *testing.T) {
			printArgValues("", tt.token, tt.defs, true)
		})
	}
}

func TestPrintCmdArgDesc(t *testing.T) {
	for _, tt := range mockPrintArgumentHelpTests {
		t.Run(tt.token, func(t *testing.T) {
			printCmdArgDesc("", tt.token, tt.defs, true)
		})
	}
}

func TestPrintCmdArgs(t *testing.T) {
	for _, tt := range mockPrintCommandHelpTests {
		t.Run(tt.token, func(t *testing.T) {
			printCmdArgs(tt.token, tt.defs, false)
		})
	}
}

func TestPrintArgAttrsLegend(t *testing.T) {
	printArgAttrsLegend()
}

func TestPrintCmdExamples(t *testing.T) {
	tests := make([]TokenHelpTest, 0)
	tests = append(tests, mockPrintCommandHelpTests...)
	exampleDefs := mockDefinitions()
	mockAddExample(&exampleDefs.Commands[0], []string{})
	tests = append(tests, TokenHelpTest{
		token: exampleDefs.Commands[0].Token,
		defs:  exampleDefs,
	})
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			printCmdExamples(tt.token, tt.defs)
		})
	}
}

func TestPrintCmdMoreInfoPrompt(t *testing.T) {
	for _, tt := range mockPrintCommandHelpTests {
		t.Run(tt.token, func(t *testing.T) {
			printCmdMoreInfoPrompt(tt.token, tt.defs)
		})
	}
}

func TestPrintCmdHelp(t *testing.T) {
	verbosity := []bool{true, false}
	for _, vv := range verbosity {
		for _, tt := range mockPrintCommandHelpTests {
			t.Run(fmt.Sprintf("%s-%v", tt.token, vv), func(t *testing.T) {
				printCmdHelp(tt.token, tt.defs, vv)
			})
		}
	}
}
