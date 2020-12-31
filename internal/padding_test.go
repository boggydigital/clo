package internal

import "testing"

func TestCommandsPadding(t *testing.T) {
	defs := testDefs()
	if defs.CommandsPadding() != len("command1") {
		t.Error("unexpected commands padding")
	}
}

func TestFlagsPadding(t *testing.T) {
	defs := testDefs()
	if defs.FlagsPadding() != len("flag1") {
		t.Error("unexpected flags padding")
	}
}

func TestArgumentsPadding(t *testing.T) {
	tests := []struct {
		cmd        string
		expPadding int
	}{
		{"command1", len("argument1")},
		{"command-that-doesnt-exist", 0},
	}
	defs := testDefs()

	for _, tt := range tests {
		t.Run(tt.cmd, func(t *testing.T) {
			if defs.ArgumentsPadding(tt.cmd) != tt.expPadding {
				t.Errorf("unexpected arguments padding for command '%s'", tt.cmd)
			}
		})
	}
}
