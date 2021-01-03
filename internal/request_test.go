package internal

import (
	"math"
	"testing"
)

func mockParseCtx(cmd, arg string) *parseCtx {
	return &parseCtx{
		Command:  mockCommandDefinition(cmd, nil),
		Argument: mockArgumentDefinition(arg, nil),
	}
}

func TestRequestUpdate(t *testing.T) {
	tests := []struct {
		token     string
		tokenType int
		ctx       *parseCtx
		expError  bool
	}{
		{"", commandAbbr, nil, false},
		{"command", command, nil, false},
		{"command-overwrite", command, nil, true},
		{"flag", flagAbbr, nil, false},
		{"flag", flag, nil, false},
		{"arg", argument, nil, false},
		{"arg", argumentAbbr, nil, false},
		{"vd", valueDefault, mockParseCtx("", "arg"), false},
		{"v", value, mockParseCtx("", "arg"), false},
		{"vf", valueFixed, mockParseCtx("", "arg"), false},
		{"", -1, nil, true},
		{"", math.MaxInt64, nil, true},
	}
	req := Request{
		Arguments: map[string][]string{},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			err := req.update(tt.token, tt.tokenType, tt.ctx)
			assertError(t, err, tt.expError)
		})
	}
}

func TestCommandHasRequiredArgs(t *testing.T) {

}
