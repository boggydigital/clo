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

func TestGenArgument(t *testing.T) {
	tests := []struct {
		input string
		token string
		mult  bool
	}{
		{"", "", false},
		{"a1", "a1", false},
		{"a1...", "a1", true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			arg := genArgument(tt.input)
			assertValEquals(t, arg.Token, tt.token)
			assertValEquals(t, arg.Multiple, tt.mult)
		})
	}
}

func TestGenDefinitions(t *testing.T) {
	tests := []struct {
		app      string
		cmd, arg []string
	}{
		{"app", []string{"c1"}, []string{"a1", "a2"}},
		{"app", []string{"c1", "c2", "c3"}, []string{"a1"}},
	}
	for _, tt := range tests {
		t.Run(tt.app, func(t *testing.T) {
			defs := GenDefinitions(tt.app, tt.cmd, tt.arg)
			assertValEquals(t, defs.App, tt.app)
			assertValEquals(t, len(defs.Commands), len(tt.cmd))
			assertValEquals(t, len(defs.Arguments), len(tt.arg))
		})
	}
}
