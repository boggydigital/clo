package clo

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

func TestInitial(t *testing.T) {
	assertValEquals(t, len(initial()), 3)
}

//func TestExpandAbbr(t *testing.T) {
//	tests := []struct {
//		token     string
//		expToken  string
//		tokenType int
//	}{
//		{"c1", "command1", command},
//		{"command-abbr-that-doesnt-exist", "command-abbr-that-doesnt-exist", command},
//		{"a1", "argument1", argument},
//		{"argument-abbr-that-doesnt-exist", "argument-abbr-that-doesnt-exist", argument},
//		{"c", "c", command},
//		{"a", "a", argument},
//		{"v", "v", value},
//	}
//	defs := mockDefinitions()
//	for _, tt := range tests {
//		t.Run(tt.token, func(t *testing.T) {
//			expToken := expandAbbr(tt.token, tt.tokenType, defs)
//			assertValEquals(t, expToken, tt.expToken)
//		})
//	}
//}
