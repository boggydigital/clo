package internal

import (
	"errors"
	"strconv"
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

func mockCommandDefinition(cmd string, args []string) *CommandDefinition {
	cd := CommandDefinition{
		CommonDefinition: CommonDefinition{Token: cmd, Abbr: cmd},
		Arguments:        args,
	}
	mockExample(&cd, args)
	return &cd
}

func mockCommandDefinitions(commands []string) []CommandDefinition {
	comDefs := make([]CommandDefinition, 0)
	for _, c := range commands {
		comDefs = append(comDefs, *mockCommandDefinition(c, commands))
	}
	return comDefs
}

func mockExample(cd *CommandDefinition, tokens []string) {
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

func mockArgumentDefinition(arg string, values []string) *ArgumentDefinition {
	ad := ArgumentDefinition{
		CommonDefinition: CommonDefinition{Token: arg, Abbr: arg},
		Values:           values,
	}
	return &ad
}

func genArgumentDefinitions(args []string) []ArgumentDefinition {
	argDefs := make([]ArgumentDefinition, 0)
	for _, a := range args {
		argDefs = append(argDefs, *mockArgumentDefinition(a, args))
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
			err := cmdTokensAreNotEmpty(mockCommandDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentCmdTokens(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentCmdTokens(mockCommandDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentCmdAbbr(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentCmdAbbr(mockCommandDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestArgTokensAreNotEmpty(t *testing.T) {
	for _, tt := range noEmptyTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := argTokensAreNotEmpty(genArgumentDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentArgTokens(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentArgTokens(genArgumentDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentArgAbbr(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentArgAbbr(genArgumentDefinitions(tt.tokens), false)
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
				mockCommandDefinitions(tt.commands),
				genArgumentDefinitions(tt.arguments),
				genFlags(tt.flags), false)
			assertError(t, err, tt.expError)
		})
	}
}

func mockArgByToken(token string) *ArgumentDefinition {
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
			err := commandsValidArgs(mockCommandDefinitions(tt.tokens), mockArgByToken, false)
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
				mockCommandDefinitions(tt.commands),
				genArgumentDefinitions(tt.arguments),
				false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentArgsCmd(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentArgsCmd(mockCommandDefinitions(tt.tokens), false)
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
			err := singleDefaultArgPerCmd(mockCommandDefinitions(tt.tokens), mockArgByToken, false)
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
			err := differentArgValues(genArgumentDefinitions(tt.tokens), false)
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
			err := examplesDescAreNotEmpty(mockCommandDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestExamplesArgumentsAreNotEmpty(t *testing.T) {
	for _, tt := range examplesTests {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := examplesArgumentsAreNotEmpty(mockCommandDefinitions(tt.tokens), false)
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
			err := examplesHaveArgsValues(mockCommandDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestExamplesArgumentsAreValid(t *testing.T) {
	for _, tt := range examplesTests {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := examplesArgumentsAreValid(mockCommandDefinitions(tt.tokens), mockArgByToken, false)
			assertError(t, err, tt.expError)
		})
	}
}

func mockValidArgVal(arg, val string) bool {
	if strings.HasPrefix(val, "invalid") {
		return false
	}
	return true
}

func TestCmdExampleHasValidValues(t *testing.T) {
	tests := []struct {
		argValues map[string][]string
		expError  bool
	}{
		{nil, false},
		{map[string][]string{}, false},
		{map[string][]string{"a": {"1", "2"}}, false},
		{map[string][]string{"a": {"invalid"}}, true},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			err := cmdExampleHasValidValues("", tt.argValues, mockValidArgVal, 0)
			assertError(t, err, tt.expError)
		})
	}
}

func TestExamplesHaveValidValues(t *testing.T) {
	tests := []TokensTest{
		{nil, false},
		{[]string{}, false},
		{[]string{"1", "2"}, false},
		{[]string{"invalid"}, true},
	}
	for _, tt := range tests {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := examplesHaveValidValues(mockCommandDefinitions(tt.tokens), mockValidArgVal, false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestAppendError(t *testing.T) {
	tests := []error{errors.New(""), nil}
	errs := make([]error, 0)
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			errs = appendError(errs, tt)
			if len(errs) != 1 {
				t.Error()
			}
		})
	}
}

func TestVerify(t *testing.T) {
	// We've already verified individual error cases above
	// so running known good definitions for the coverage
	defs := testDefs()
	errs := defs.Verify(false)
	if len(errs) > 0 {
		t.Error()
	}
}
