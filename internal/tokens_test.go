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
			assertValEquals(t, tokenString(tt.tokenType), tt.tokenStr)
		})
	}
}

func TestNext(t *testing.T) {
	tests := []struct {
		tokenType int
		nextLen   int
	}{
		{command, 5},
		{commandAbbr, 5},
		{argument, 6},
		{argumentAbbr, 6},
		{valueFixed, 5},
		{valueDefault, 5},
		{value, 5},
		{flag, 2},
		{flagAbbr, 2},
		{-1, 0},
		{math.MaxInt64, 0},
	}

	for _, tt := range tests {
		t.Run(tokenString(tt.tokenType), func(t *testing.T) {
			assertValEquals(t, len(next(tt.tokenType)), tt.nextLen)
		})
	}
}

func TestFirst(t *testing.T) {
	assertValEquals(t, len(first()), 2)
}

func TestExpandAbbr(t *testing.T) {
	tests := []struct {
		token     string
		expToken  string
		tokenType int
		expError  bool
	}{
		{"c1", "command1", commandAbbr, false},
		{"command-abbr-that-doesnt-exist", "", commandAbbr, true},
		{"a1", "argument1", argumentAbbr, false},
		{"argument-abbr-that-doesnt-exist", "", argumentAbbr, true},
		{"f1", "flag1", flagAbbr, false},
		{"flag-abbr-that-doesnt-exist", "", flagAbbr, true},
		{"c", "c", command, false},
		{"a", "a", argument, false},
		{"vd", "vd", valueDefault, false},
		{"vf", "vf", valueFixed, false},
		{"v", "v", value, false},
		{"f", "f", flag, false},
	}
	defs := mockDefinitions()
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			expToken, err := expandAbbr(tt.token, tt.tokenType, defs)
			assertValEquals(t, expToken, tt.expToken)
			assertError(t, err, tt.expError)
		})
	}
}
