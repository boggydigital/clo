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
		t.Run(tt.dupe, func(t *testing.T) {
			assertValEquals(t, firstDupe(tt.slice), tt.dupe)
		})
	}
}

func TestVFail(t *testing.T) {
	vFail("", true)
}

func TestVPass(t *testing.T) {
	vPass("", true)
}

func TestCmdTokensAreNotEmpty(t *testing.T) {
	for _, tt := range mockNoEmptyTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := cmdTokensAreNotEmpty(mockCommandDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentCmdTokens(t *testing.T) {
	for _, tt := range mockDifferentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentCmdTokens(mockCommandDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentCmdAbbr(t *testing.T) {
	for _, tt := range mockDifferentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentCmdAbbr(mockCommandDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestArgTokensAreNotEmpty(t *testing.T) {
	for _, tt := range mockNoEmptyTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := argTokensAreNotEmpty(mockArgumentDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentArgTokens(t *testing.T) {
	for _, tt := range mockDifferentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentArgTokens(mockArgumentDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentArgAbbr(t *testing.T) {
	for _, tt := range mockDifferentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentArgAbbr(mockArgumentDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestFlagTokensAreNotEmpty(t *testing.T) {
	for _, tt := range mockNoEmptyTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := flagTokensAreNotEmpty(mockFlagDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentFlagTokens(t *testing.T) {
	for _, tt := range mockDifferentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentFlagTokens(mockFlagDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentFlagAbbr(t *testing.T) {
	for _, tt := range mockDifferentTokensTests() {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := differentFlagAbbr(mockFlagDefinitions(tt.tokens), false)
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
				mockArgumentDefinitions(tt.arguments),
				mockFlagDefinitions(tt.flags), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestCommandsValidArgs(t *testing.T) {
	for _, tt := range mockNoEmptyTokensTests() {
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
				mockArgumentDefinitions(tt.arguments),
				false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestDifferentArgsCmd(t *testing.T) {
	for _, tt := range mockDifferentTokensTests() {
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
			err := singleDefaultArgPerCmd(
				mockCommandDefinitions(tt.tokens),
				mockArgByToken,
				false)
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
			err := differentArgValues(mockArgumentDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestExamplesDescAreNotEmpty(t *testing.T) {
	for _, tt := range mockExamplesTests {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := examplesDescAreNotEmpty(mockCommandDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestExamplesArgumentsAreNotEmpty(t *testing.T) {
	for _, tt := range mockExamplesTests {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := examplesArgumentsAreNotEmpty(mockCommandDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestExamplesHaveArgsValues(t *testing.T) {
	for _, tt := range mockEmptyExamplesTests {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := examplesHaveArgsValues(mockCommandDefinitions(tt.tokens), false)
			assertError(t, err, tt.expError)
		})
	}
}

func TestExamplesArgumentsAreValid(t *testing.T) {
	for _, tt := range mockExamplesTests {
		t.Run(strings.Join(tt.tokens, "-"), func(t *testing.T) {
			err := examplesArgumentsAreValid(mockCommandDefinitions(tt.tokens), mockArgByToken, false)
			assertError(t, err, tt.expError)
		})
	}
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
			assertValEquals(t, len(errs), 1)
		})
	}
}

func TestDefinitionsVerify(t *testing.T) {
	// We've already verified individual error cases above
	// so running known good definitions for the coverage
	defs := mockDefinitions()
	errs := defs.Verify(false)
	assertValEquals(t, len(errs), 0)
}
