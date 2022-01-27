package clo

import (
	"github.com/boggydigital/testo"
	"testing"
)

func TestPrintHelp(t *testing.T) {
	tests := []struct {
		token string
	}{
		{""},
		{"command1"},
		{"command2"},
		{"command-that-doesnt-exist"},
	}
	defs := mockDefinitions()
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			err := printHelp(tt.token, defs)
			testo.Error(t, err, defs == nil)
		})
	}
}
