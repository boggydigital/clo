package clo

import (
	"github.com/boggydigital/testo"
	"testing"
)

func TestCommandsPadding(t *testing.T) {
	defs := mockDefinitions()
	testo.EqualValues(t, defs.cmdPadding(), len("command1"))
}

func TestDefinitionsArgumentsPadding(t *testing.T) {
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
			argPadding, err := defs.argPadding(tt.cmd)
			testo.Error(t, err, false)
			testo.EqualValues(t, argPadding, tt.expPadding)
		})
	}
}
