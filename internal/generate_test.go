package internal

import "testing"

func TestGenCommand(t *testing.T) {
	tests := []string{"", "c1"}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			cmd := genCommand(tt)
			assertValEquals(t, cmd.Token, tt)
		})
	}
}

func TestGenFlags(t *testing.T) {
	tests := []string{"", "f1"}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			flg := genFlag(tt)
			assertValEquals(t, flg.Token, tt)
		})
	}
}

func TestGenArgument(t *testing.T) {
	tests := []struct {
		input string
		token string
		def   bool
		req   bool
		mult  bool
	}{
		{"", "", false, false, false},
		{"a1", "a1", false, false, false},
		{"_a1", "a1", true, false, false},
		{"_*a1", "a1", true, true, false},
		{"*a1", "a1", false, true, false},
		{"*_a1", "_a1", false, true, false},
		{"a1...", "a1", false, false, true},
		{"_a1...", "a1", true, false, true},
		{"_*a1...", "a1", true, true, true},
		{"*a1...", "a1", false, true, true},
		{"*_a1...", "_a1", false, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			arg := genArgument(tt.input)
			assertValEquals(t, arg.Token, tt.token)
			assertValEquals(t, arg.Default, tt.def)
			assertValEquals(t, arg.Required, tt.req)
			assertValEquals(t, arg.Multiple, tt.mult)
		})
	}
}

func TestGenDefinitions(t *testing.T) {
	tests := []struct {
		app           string
		cmd, arg, flg []string
	}{
		{"app", []string{"c1"}, []string{"a1", "a2"}, []string{"f1", "f2", "f3"}},
		{"app", []string{"c1", "c2", "c3"}, []string{"a1"}, []string{"f1", "f2"}},
	}
	for _, tt := range tests {
		t.Run(tt.app, func(t *testing.T) {
			defs := GenDefinitions(tt.app, tt.cmd, tt.arg, tt.flg)
			assertValEquals(t, defs.App, tt.app)
			assertValEquals(t, len(defs.Commands), len(tt.cmd))
			assertValEquals(t, len(defs.Arguments), len(tt.arg))
			assertValEquals(t, len(defs.Flags), len(tt.flg))
		})
	}
}
