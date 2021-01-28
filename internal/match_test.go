package internal

import (
	"testing"
)

func TestMatchArgument(t *testing.T) {
	tests := []struct {
		token  string
		cmd    string
		expArg string
	}{
		{},
	}

	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			arg := matchArg(tt.token, tt.cmd, mockValidCmdArg)
			assertValEquals(t, arg, tt.expArg)
		})
	}
	//defs := mockDefinitions()
	//tests := []MatchArgumentTest{
	//	{MatchTest{"", argument, false, false}, "", nil},
	//	// shouldn't match
	//	// - tokens that don't exist
	//	{MatchTest{"-argument-that-doesnt-exist", argument, false, true}, "", nil},
	//	{MatchTest{"-argument-that-doesnt-exist", argument, false, false}, "", defs},
	//	{MatchTest{"--argument-that-doesnt-exist", argument, false, true}, "", nil},
	//	{MatchTest{"--argument-that-doesnt-exist", argument, false, false}, "", defs},
	//	{MatchTest{"-abbr-that-doesnt-exist", argument, false, true}, "", nil},
	//	{MatchTest{"-abbr-that-doesnt-exist", argument, false, false}, "", defs},
	//	{MatchTest{"--abbr-that-doesnt-exist", argument, false, true}, "", nil},
	//	{MatchTest{"--abbr-that-doesnt-exist", argument, false, false}, "", defs},
	//	// - tokens that exist, but are not valid
	//	{MatchTest{"-argument1", argument, false, true}, "", defs},
	//	{MatchTest{"--argument1", argument, false, true}, "", defs},
	//	// - nil command context, so can't validate
	//	{MatchTest{"-argument1", argument, false, true}, "", defs},
	//	{MatchTest{"--argument1", argument, false, true}, "", defs},
	//	// - tokenType is not argument or argumentAbbr
	//	{MatchTest{"-argument1", command, false, true}, "", defs},
	//	{MatchTest{"-argument1", command, false, true}, "", defs},
	//	{MatchTest{"-argument1", value, false, true}, "", defs},
	//	{MatchTest{"-argument1", -1, false, true}, "", defs},
	//	{MatchTest{"-argument1", math.MaxInt64, false, true}, "", defs},
	//	// should match
	//	// - valid args for command, no error
	//	{MatchTest{"-argument1", argument, true, false}, defs.Commands[0].Token, defs},
	//	{MatchTest{"--argument1", argument, true, false}, defs.Commands[0].Token, defs},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.token, func(t *testing.T) {
	//		m, err := matchArgument(tt.token, tt.tokenType, tt.cmd, tt.def)
	//		assertValEquals(t, m, tt.expected)
	//		assertError(t, err, tt.expError)
	//	})
	//}
}

//func TestMatch(t *testing.T) {
//	defs := mockDefinitions()
//	command1Request := &Request{Command: defs.Commands[0].Token}
//	argument1Request := &Request{Arguments: map[string][]string{defs.Arguments[0].Token: {}}}
//	argument3Request := &Request{Arguments: map[string][]string{defs.Arguments[2].Token: {}}}
//	tests := []MatchDefaultValueTest{
//		// def == nil
//		{MatchTest{"", command, false, true}, nil, nil},
//		{MatchTest{"", argument, false, true}, nil, nil},
//		{MatchTest{"", value, false, true}, nil, nil},
//		{MatchTest{"", -1, false, true}, nil, nil},
//		{MatchTest{"", math.MaxInt64, false, true}, nil, nil},
//		// command token
//		{MatchTest{"command1", command, true, false}, nil, defs},
//		{MatchTest{"command-that-doesnt-exist", command, false, false}, nil, defs},
//		// command abbr
//		{MatchTest{"c-abbr-that-doesnt-exist", command, false, false}, nil, defs},
//		// argument token
//		{MatchTest{"-argument1", argument, true, false}, command1Request, defs},
//		{MatchTest{"--argument1", argument, true, false}, command1Request, defs},
//		{MatchTest{"argument1", argument, false, false}, command1Request, defs},
//		{MatchTest{"-argument-that-doesnt-exist", argument, false, false}, command1Request, defs},
//		{MatchTest{"--argument-that-doesnt-exist", argument, false, false}, command1Request, defs},
//		// argument abbr
//		{MatchTest{"-a-abbr-that-doesnt-exist", argument, false, false}, command1Request, defs},
//		{MatchTest{"--a-abbr-that-doesnt-exist", argument, false, false}, command1Request, defs},
//		// value
//		{MatchTest{"value1", value, true, false}, argument1Request, defs},
//		{MatchTest{"any-value", value, true, false}, argument3Request, defs},
//		{MatchTest{"-value1", value, false, false}, argument1Request, defs},
//		{MatchTest{"any-value", value, false, false}, argument1Request, defs},
//		{MatchTest{"-value1", value, false, false}, argument3Request, defs},
//		// invalid token types
//		{MatchTest{"", -1, false, true}, nil, defs},
//		{MatchTest{"", math.MaxInt64, false, true}, nil, defs},
//	}
//	for _, tt := range tests {
//		t.Run(tt.token, func(t *testing.T) {
//			m, err := match(tt.token, tt.tokenType, tt.req, tt.def)
//			assertValEquals(t, m, tt.expected)
//			assertError(t, err, tt.expError)
//		})
//	}
//
//}
