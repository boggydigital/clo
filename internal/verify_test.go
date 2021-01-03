package internal

import (
	"strings"
	"testing"
)

func TestFirstDupe(t *testing.T) {
	tests := []struct {
		slice []string
		dupe  string
	}{
		{[]string{}, ""},
		{[]string{"1"}, ""},
		{[]string{"1", "2", "3"}, ""},
		{[]string{"1", "2", "3", "1"}, "1"},
		{[]string{"1", "2", "3", "2"}, "2"},
		{[]string{"1", "2", "3", "3"}, "3"},
	}

	for _, tt := range tests {
		name := tt.dupe
		if name == "" {
			name = "empty"
		}
		t.Run(name, func(t *testing.T) {
			dupe := firstDupe(tt.slice)
			if dupe != tt.dupe {
				t.Error()
			}
		})
	}
}

func TestVFail(t *testing.T) {
	vFail("", true)
}

func TestVPass(t *testing.T) {
	vPass("", true)
}

func genCommands(commands []string) []CommandDefinition {
	comDefs := make([]CommandDefinition, 0)
	for _, c := range commands {
		cd := CommandDefinition{
			CommonDefinition: CommonDefinition{Token: c, Abbr: c},
			Arguments:        commands,
		}
		genExample(&cd, commands)
		comDefs = append(comDefs, cd)
	}
	return comDefs
}

func genExample(cd *CommandDefinition, tokens []string) {
	desc := strings.Join(tokens, "")
	cd.Examples = append(cd.Examples, ExampleDefinition{
		ArgumentsValues: map[string][]string{},
		Desc:            desc,
	})
	cd.Examples[0].ArgumentsValues = make(map[string][]string, 0)
	for _, tt := range tokens {
		if strings.HasPrefix(desc, "empty") {
			continue
		}
		cd.Examples[0].ArgumentsValues[tt] = tokens
	}
}

func genArguments(args []string) []ArgumentDefinition {
	argDefs := make([]ArgumentDefinition, 0)
	for _, a := range args {
		ad := ArgumentDefinition{
			CommonDefinition: CommonDefinition{Token: a, Abbr: a},
			Values:           args,
		}
		argDefs = append(argDefs, ad)
	}
	return argDefs
}

func genFlags(flags []string) []FlagDefinition {
	flagDefs := make([]FlagDefinition, 0)
	for _, f := range flags {
		fd := FlagDefinition{
			CommonDefinition: CommonDefinition{Token: f, Abbr: f},
		}
		flagDefs = append(flagDefs, fd)
	}
	return flagDefs
}

type TokensTest struct {
	tokens   []string
	expError bool
}

func noEmptyTokensTests() []TokensTest {
	return []TokensTest{
		{nil, false},
		{[]string{}, false},
		{[]string{"1", "2"}, false},
		{[]string{"", "x"}, true},
		{[]string{"x", ""}, true},
	}
}

func differentTokensTests() []TokensTest {
	return []TokensTest{
		{nil, false},
		{[]string{}, false},
		{[]string{"1"}, false},
		{[]string{"1", "2"}, false},
		{[]string{"1", "2", "2"}, true},
	}
}

func assertError(t *testing.T, err error, expError bool) {
	if (err != nil && !expError) || (err == nil && expError) {
		t.Error()
	}
}

func TestCmdTokensAreNotEmpty(t *testing.T) {
	for _, tt := range noEmptyTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := cmdTokensAreNotEmpty(genCommands(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentCmdTokens(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentCmdTokens(genCommands(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentCmdAbbr(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentCmdAbbr(genCommands(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestArgTokensAreNotEmpty(t *testing.T) {
	for _, tt := range noEmptyTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := argTokensAreNotEmpty(genArguments(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentArgTokens(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentArgTokens(genArguments(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentArgAbbr(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentArgAbbr(genArguments(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestFlagTokensAreNotEmpty(t *testing.T) {
	for _, tt := range noEmptyTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := flagTokensAreNotEmpty(genFlags(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentFlagTokens(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentFlagTokens(genFlags(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentFlagAbbr(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentFlagAbbr(genFlags(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentAbbr(t *testing.T) {
	tests := []struct {
		commands  []string
		arguments []string
		flags     []string
		expError  bool
	}{
		{nil, nil, nil, false},
		{[]string{}, []string{}, []string{}, false},
		{[]string{"1"}, []string{"2"}, []string{"3"}, false},
		{[]string{"1"}, []string{"2"}, []string{"3"}, false},
		{[]string{"1", "2"}, []string{"2"}, []string{"3"}, true},
		{[]string{"1"}, []string{"2", "1"}, []string{"3"}, true},
		{[]string{"1"}, []string{"2"}, []string{"1", "3"}, true},
	}
	for _, tt := range tests {
		name := strings.Join(tt.commands, "-") + "-" +
			strings.Join(tt.arguments, "-") + "-" +
			strings.Join(tt.flags, "-")
		t.Run(name, func(t *testing.T) {
			err := differentAbbr(
				genCommands(tt.commands),
				genArguments(tt.arguments),
				genFlags(tt.flags), false)
			assertError(t, err, tt.expError)
		})
	}
}

func argByToken(token string) *ArgumentDefinition {
	if strings.HasPrefix(token, "default") {
		// default arguments
		return &ArgumentDefinition{
			CommonDefinition: CommonDefinition{
				Token: token,
			},
			Default: true,
		}
	}

	switch token {
	case "":
		return nil
	default:
		return &ArgumentDefinition{
			CommonDefinition: CommonDefinition{
				Token: token,
			},
		}
	}
}

func TestCommandsValidArgs(t *testing.T) {
	for _, tt := range noEmptyTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := commandsValidArgs(genCommands(tt.tokens), argByToken, false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestAllUsedArgs(t *testing.T) {
	tests := []struct {
		commands  []string
		arguments []string
		expError  bool
	}{
		{nil, nil, false},
		{[]string{}, []string{}, false},
		{[]string{"1"}, []string{"1"}, false},
		{[]string{"1"}, []string{"2"}, true},
	}

	for _, tt := range tests {
		name := strings.Join(tt.commands, "-") + "-" +
			strings.Join(tt.arguments, "-")
		t.Run(name, func(t *testing.T) {
			err := allUsedArgs(
				genCommands(tt.commands),
				genArguments(tt.arguments),
				false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentArgsCmd(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentArgsCmd(genCommands(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestSingleDefaultArgPerCmd(t *testing.T) {
	tests := []TokensTest{
		{nil, false},
		{[]string{}, false},
		{[]string{"1", "2", "3"}, false},
		{[]string{"default1", "default2"}, true},
		{[]string{"default1", "", "default2"}, true},
	}
	for _, tt := range tests {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := singleDefaultArgPerCmd(genCommands(tt.tokens), argByToken, false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentArgValues(t *testing.T) {
	tests := []TokensTest{
		{nil, false},
		{[]string{}, false},
		{[]string{"1", "2", "3"}, false},
		{[]string{"1", "1"}, true},
	}
	for _, tt := range tests {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentArgValues(genArguments(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

var examplesTests = []TokensTest{
	{nil, false},
	{[]string{}, false},
	{[]string{"1", "2"}, false},
	{[]string{"", ""}, true},
}

func TestExamplesDescAreNotEmpty(t *testing.T) {
	for _, tt := range examplesTests {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := examplesDescAreNotEmpty(genCommands(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestExamplesArgumentsAreNotEmpty(t *testing.T) {
	for _, tt := range examplesTests {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := examplesArgumentsAreNotEmpty(genCommands(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestExamplesHaveArgsValues(t *testing.T) {
	tests := []TokensTest{
		{nil, false},
		{[]string{}, false},
		{[]string{"empty"}, true},
	}
	for _, tt := range tests {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := examplesHaveArgsValues(genCommands(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestExamplesArgumentsAreValid(t *testing.T) {
	for _, tt := range examplesTests {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := examplesArgumentsAreValid(genCommands(tt.tokens), argByToken, false)
			assertError(t, err, tt.expError)
		})
	}
}
