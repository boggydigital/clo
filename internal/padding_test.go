package internal

import "testing"

func TestCommandsPadding(t *testing.T) {
	defs := mockDefinitions()
	assertEquals(t, defs.CommandsPadding(), len("command1"))
}

func TestFlagsPadding(t *testing.T) {
	defs := mockDefinitions()
	assertEquals(t, defs.FlagsPadding(), len("flag1"))
}

func TestArgumentsPadding(t *testing.T) {
	tests := []struct {
		cmd        string
		expPadding int
	}{
		{"command1", len("argument1")},
		{"command-that-doesnt-exist", 0},
	}
	defs := mockDefinitions()

	for _, tt := range tests {
		t.Run(tt.cmd, func(t *testing.T) {
			assertEquals(t, defs.ArgumentsPadding(tt.cmd), tt.expPadding)
		})
	}
}
