package internal

import (
	"math"
	"testing"
)

func TestTokenString(t *testing.T) {
	tests := []struct {
		tokenType int
		tokenStr  string
	}{
		{command, "command"},
		{commandAbbr, "commandAbbr"},
		{argument, "argument"},
		{argumentAbbr, "argumentAbbr"},
		{valueDefault, "valueDefault"},
		{valueFixed, "valueFixed"},
		{value, "value"},
		{flag, "flag"},
		{flagAbbr, "flagAbbr"},
		{-1, "unknown"},
		{math.MaxInt64, "unknown"},
	}

	for _, tt := range tests {
		testName := tt.tokenStr
		t.Run(testName, func(t *testing.T) {
			act := tokenString(tt.tokenType)
			if act != tt.tokenStr {
				t.Errorf("tokenType %d string was expected to be %s and was %s",
					tt.tokenType,
					tt.tokenStr,
					act)
			}
		})
	}
}
