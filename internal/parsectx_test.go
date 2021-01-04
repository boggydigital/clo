package internal

import (
	"testing"
)

func TestParseCtxUpdate(t *testing.T) {
	tests := []struct {
		token     string
		tokenType int
		cmdNilExp bool
		argNilExp bool
	}{
		{"command-token-that-doesnt-exist", command, true, true},
		{"command-abbr-that-doesnt-exist", commandAbbr, true, true},
		{"--arg-token-that-doesnt-exist", argument, true, true},
		{"--arg-abbr-that-doesnt-exist", argumentAbbr, true, true},
		{"command1", command, false, true},
		{"c1", commandAbbr, false, true},
		{"--argument1", argument, false, false},
		{"--a1", argumentAbbr, false, false},
	}
	pCtx, defs := parseCtx{}, mockDefinitions()
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			pCtx.update(tt.token, tt.tokenType, defs)
			assertNil(t, pCtx.Command, tt.cmdNilExp)
			assertNil(t, pCtx.Argument, tt.argNilExp)
		})
	}
}
