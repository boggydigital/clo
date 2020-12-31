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
				t.Error("unexpected first dupe")
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
		}
		comDefs = append(comDefs, cd)
	}
	return comDefs
}

func genArguments(args []string) []ArgumentDefinition {
	argDefs := make([]ArgumentDefinition, 0)
	for _, a := range args {
		ad := ArgumentDefinition{
			CommonDefinition: CommonDefinition{Token: a, Abbr: a},
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

func TestCmdTokensAreNotEmpty(t *testing.T) {
	for _, tt := range noEmptyTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := cmdTokensAreNotEmpty(genCommands(tt.tokens), false)
			if (err != nil && !tt.expError) || (err == nil && tt.expError) {
				t.Error("unexpected command tokens not empty error state")
			}
		})
	}
}

func TestDifferentCmdTokens(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentCmdTokens(genCommands(tt.tokens), false)
			if (err != nil && !tt.expError) || (err == nil && tt.expError) {
				t.Error("unexpected different command tokens error state")
			}
		})
	}
}

func TestDifferentCmdAbbr(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentCmdAbbr(genCommands(tt.tokens), false)
			if (err != nil && !tt.expError) || (err == nil && tt.expError) {
				t.Error("unexpected different command abbr error state")
			}
		})
	}
}

func TestArgTokensAreNotEmpty(t *testing.T) {
	for _, tt := range noEmptyTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := argTokensAreNotEmpty(genArguments(tt.tokens), false)
			if (err != nil && !tt.expError) || (err == nil && tt.expError) {
				t.Error("unexpected argument tokens not empty error state")
			}
		})
	}
}

func TestDifferentArgTokens(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentArgTokens(genArguments(tt.tokens), false)
			if (err != nil && !tt.expError) || (err == nil && tt.expError) {
				t.Error("unexpected different argument tokens error state")
			}
		})
	}
}

func TestDifferentArgAbbr(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentArgAbbr(genArguments(tt.tokens), false)
			if (err != nil && !tt.expError) || (err == nil && tt.expError) {
				t.Error("unexpected different arguments abbr error state")
			}
		})
	}
}

func TestFlagTokensAreNotEmpty(t *testing.T) {
	for _, tt := range noEmptyTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := flagTokensAreNotEmpty(genFlags(tt.tokens), false)
			if (err != nil && !tt.expError) || (err == nil && tt.expError) {
				t.Error("unexpected flag tokens not empty error state")
			}
		})
	}
}

func TestDifferentFlagTokens(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentFlagTokens(genFlags(tt.tokens), false)
			if (err != nil && !tt.expError) || (err == nil && tt.expError) {
				t.Error("unexpected different flag tokens error state")
			}
		})
	}
}

func TestDifferentFlagAbbr(t *testing.T) {
	for _, tt := range differentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentFlagAbbr(genFlags(tt.tokens), false)
			if (err != nil && !tt.expError) || (err == nil && tt.expError) {
				t.Error("unexpected different flag abbr error state")
			}
		})
	}
}
