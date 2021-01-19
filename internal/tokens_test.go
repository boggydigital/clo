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
		{argument, "argument"},
		{value, "value"},
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
		{command, 2},
		{argument, 2},
		{value, 2},
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
	assertValEquals(t, len(first()), 1)
}

func TestExpandAbbr(t *testing.T) {
	tests := []struct {
		token     string
		expToken  string
		tokenType int
		expError  bool
	}{
		{"c1", "command1", command, false},
		{"command-abbr-that-doesnt-exist", "", command, true},
		{"a1", "argument1", argument, false},
		{"argument-abbr-that-doesnt-exist", "", argument, true},
		{"c", "c", command, true},
		{"a", "a", argument, true},
		{"v", "v", value, false},
	}
	defs := mockDefinitions()
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			expToken, err := expandAbbr(tt.token, tt.tokenType, defs)
			if err == nil {
				assertValEquals(t, expToken, tt.expToken)
			}
			assertError(t, err, tt.expError)
		})
	}
}
