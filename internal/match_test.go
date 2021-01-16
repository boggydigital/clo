package internal

import (
	"math"
	"testing"
)

func TestMatchArgument(t *testing.T) {
	defs := mockDefinitions()
	tests := []MatchArgumentTest{
		{MatchTest{"", argument, false, false}, nil, nil},
		{MatchTest{"", argumentAbbr, false, false}, nil, nil},
		// shouldn't match
		// - tokens that don't exist
		{MatchTest{"-argument-that-doesnt-exist", argument, false, true}, nil, nil},
		{MatchTest{"-argument-that-doesnt-exist", argument, false, false}, nil, defs},
		{MatchTest{"--argument-that-doesnt-exist", argument, false, true}, nil, nil},
		{MatchTest{"--argument-that-doesnt-exist", argument, false, false}, nil, defs},
		{MatchTest{"-abbr-that-doesnt-exist", argumentAbbr, false, true}, nil, nil},
		{MatchTest{"-abbr-that-doesnt-exist", argumentAbbr, false, false}, nil, defs},
		{MatchTest{"--abbr-that-doesnt-exist", argumentAbbr, false, true}, nil, nil},
		{MatchTest{"--abbr-that-doesnt-exist", argumentAbbr, false, false}, nil, defs},
		// - tokens that exist, but are not valid
		{MatchTest{"-argument1", argument, false, true}, mockCommandDefinition("", []string{"not-argument1"}), defs},
		{MatchTest{"--argument1", argument, false, true}, mockCommandDefinition("", []string{"not-argument1"}), defs},
		{MatchTest{"-a1", argumentAbbr, false, true}, mockCommandDefinition("", []string{"not-argument1"}), defs},
		{MatchTest{"--a1", argumentAbbr, false, true}, mockCommandDefinition("", []string{"not-argument1"}), defs},
		// - nil command context, so can't validate
		{MatchTest{"-argument1", argument, false, true}, nil, defs},
		{MatchTest{"--argument1", argument, false, true}, nil, defs},
		{MatchTest{"-a1", argumentAbbr, false, true}, nil, defs},
		{MatchTest{"--a1", argumentAbbr, false, true}, nil, defs},
		// - tokenType is not argument or argumentAbbr
		{MatchTest{"-argument1", command, false, true}, nil, defs},
		{MatchTest{"-argument1", commandAbbr, false, true}, nil, defs},
		{MatchTest{"-argument1", valueDefault, false, true}, nil, defs},
		{MatchTest{"-argument1", valueFixed, false, true}, nil, defs},
		{MatchTest{"-argument1", value, false, true}, nil, defs},
		{MatchTest{"-argument1", -1, false, true}, nil, defs},
		{MatchTest{"-argument1", math.MaxInt64, false, true}, nil, defs},
		// should match
		// - valid args for command, no error
		{MatchTest{"-argument1", argument, true, false}, mockCommandDefinition("", []string{"argument1"}), defs},
		{MatchTest{"--argument1", argument, true, false}, mockCommandDefinition("", []string{"argument1"}), defs},
		{MatchTest{"-a1", argumentAbbr, true, false}, mockCommandDefinition("", []string{"argument1"}), defs},
		{MatchTest{"--a1", argumentAbbr, true, false}, mockCommandDefinition("", []string{"argument1"}), defs},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			m, err := matchArgument(tt.token, tt.tokenType, tt.cmd, tt.def)
			assertValEquals(t, m, tt.expected)
			assertError(t, err, tt.expError)
		})
	}
}

func TestMatchDefaultValue(t *testing.T) {
	tests := []MatchDefaultValueTest{
		{MatchTest{"", valueDefault, false, true}, nil, nil},
		// shouldn't match
		// - tokenType is not valueDefault
		{MatchTest{"", command, false, false}, nil, nil},
		{MatchTest{"", commandAbbr, false, false}, nil, nil},
		{MatchTest{"", argument, false, false}, nil, nil},
		{MatchTest{"", argumentAbbr, false, false}, nil, nil},
		{MatchTest{"", valueFixed, false, false}, nil, nil},
		{MatchTest{"", value, false, false}, nil, nil},
		// - ctx.Command is nil || !ctx.Argument.Default
		{MatchTest{"", valueDefault, false, false}, &parseCtx{
			Command:  nil,
			Argument: nil,
		}, nil},
		{MatchTest{"", valueDefault, false, false}, &parseCtx{
			Command:  mockCommandDefinition("command1", []string{}),
			Argument: mockArgumentDefinition("not-default", []string{}),
		}, nil},
		// - nil definitions
		{MatchTest{"", valueDefault, false, true}, &parseCtx{
			Command:  mockCommandDefinition("command1", []string{}),
			Argument: mockArgumentDefinition("default1", []string{}),
		}, nil},
		// - not a matching value for the argument
		{MatchTest{"value-that-doesnt-exist", valueDefault, false, true}, &parseCtx{
			Command:  mockCommandDefinition("command1", []string{"argument1"}),
			Argument: nil,
		}, mockDefinitions()},
		// should match
		{MatchTest{"value1", valueDefault, true, false}, &parseCtx{
			Command:  mockCommandDefinition("command1", []string{"argument1"}),
			Argument: nil,
		}, mockDefinitions()},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			m, err := matchDefaultValue(tt.token, tt.tokenType, tt.ctx, tt.def)
			assertValEquals(t, m, tt.expected)
			assertError(t, err, tt.expError)
		})
	}
}

func TestMatchValue(t *testing.T) {
	tests := []MatchValueTest{
		// arg == nil
		{MatchTest{"", value, false, true}, nil},
		// hasPrefix
		{MatchTest{"-", value, false, false}, mockArgumentDefinition("", []string{})},
		// tokenType == valueDefault
		{MatchTest{"value-that-doesnt-exist", valueDefault, false, false}, mockArgumentDefinition("", []string{"value1", "value2"})},
		{MatchTest{"any-value", valueDefault, false, false}, mockArgumentDefinition("", []string{""})},
		{MatchTest{"value1", valueDefault, true, false}, mockArgumentDefinition("default", []string{"value1", "value2"})},
		{MatchTest{"any-value", valueDefault, true, false}, mockArgumentDefinition("default", []string{})},
		// tokenType == valueFixed
		{MatchTest{"value1", valueFixed, true, false}, mockArgumentDefinition("", []string{"value1", "value2"})},
		{MatchTest{"value-that-doesnt-exist", valueFixed, false, true}, mockArgumentDefinition("", []string{"value1", "value2"})},
		{MatchTest{"any-value", valueFixed, false, false}, mockArgumentDefinition("", []string{})},
		// tokenType == value
		{MatchTest{"value1", value, true, false}, mockArgumentDefinition("", []string{"value1", "value2"})},
		{MatchTest{"value-that-doesnt-exist", value, false, false}, mockArgumentDefinition("", []string{"value1", "value2"})},
		{MatchTest{"any-value", value, true, false}, mockArgumentDefinition("", []string{})},
		// other token types
		{MatchTest{"value", command, false, false}, mockArgumentDefinition("", []string{})},
		{MatchTest{"value", commandAbbr, false, false}, mockArgumentDefinition("", []string{})},
		{MatchTest{"value", argument, false, false}, mockArgumentDefinition("", []string{})},
		{MatchTest{"value", argumentAbbr, false, false}, mockArgumentDefinition("", []string{})},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			m, err := matchValue(tt.token, tt.tokenType, tt.arg)
			assertValEquals(t, m, tt.expected)
			assertError(t, err, tt.expError)
		})
	}
}

func TestMatch(t *testing.T) {
	defs := mockDefinitions()
	command1ParseCtx := &parseCtx{Command: &defs.Commands[0]}
	argument1ParseCtx := &parseCtx{Argument: &defs.Arguments[0]}
	argument3ParseCtx := &parseCtx{Argument: &defs.Arguments[2]}
	tests := []MatchDefaultValueTest{
		// def == nil
		{MatchTest{"", command, false, true}, nil, nil},
		{MatchTest{"", commandAbbr, false, true}, nil, nil},
		{MatchTest{"", argument, false, true}, nil, nil},
		{MatchTest{"", argumentAbbr, false, true}, nil, nil},
		{MatchTest{"", valueDefault, false, true}, nil, nil},
		{MatchTest{"", valueFixed, false, true}, nil, nil},
		{MatchTest{"", value, false, true}, nil, nil},
		{MatchTest{"", -1, false, true}, nil, nil},
		{MatchTest{"", math.MaxInt64, false, true}, nil, nil},
		// command token
		{MatchTest{"command1", command, true, false}, nil, defs},
		{MatchTest{"command-that-doesnt-exist", command, false, false}, nil, defs},
		// command abbr
		{MatchTest{"c1", commandAbbr, true, false}, nil, defs},
		{MatchTest{"c-abbr-that-doesnt-exist", commandAbbr, false, false}, nil, defs},
		// argument token
		{MatchTest{"-argument1", argument, true, false}, command1ParseCtx, defs},
		{MatchTest{"--argument1", argument, true, false}, command1ParseCtx, defs},
		{MatchTest{"argument1", argument, false, false}, command1ParseCtx, defs},
		{MatchTest{"-argument-that-doesnt-exist", argument, false, false}, command1ParseCtx, defs},
		{MatchTest{"--argument-that-doesnt-exist", argument, false, false}, command1ParseCtx, defs},
		// argument abbr
		{MatchTest{"-a1", argumentAbbr, true, false}, command1ParseCtx, defs},
		{MatchTest{"--a1", argumentAbbr, true, false}, command1ParseCtx, defs},
		{MatchTest{"a1", argumentAbbr, false, false}, command1ParseCtx, defs},
		{MatchTest{"-a-abbr-that-doesnt-exist", argumentAbbr, false, false}, command1ParseCtx, defs},
		{MatchTest{"--a-abbr-that-doesnt-exist", argumentAbbr, false, false}, command1ParseCtx, defs},
		// value
		{MatchTest{"value1", value, true, false}, argument1ParseCtx, defs},
		{MatchTest{"any-value", value, true, false}, argument3ParseCtx, defs},
		{MatchTest{"-value1", value, false, false}, argument1ParseCtx, defs},
		{MatchTest{"any-value", value, false, false}, argument1ParseCtx, defs},
		{MatchTest{"-value1", value, false, false}, argument3ParseCtx, defs},
		// valueFixed
		{MatchTest{"value1", valueFixed, true, false}, argument1ParseCtx, defs},
		{MatchTest{"any-value", valueFixed, false, false}, argument3ParseCtx, defs},
		{MatchTest{"-value1", valueFixed, false, false}, argument1ParseCtx, defs},
		{MatchTest{"any-value", valueFixed, false, true}, argument1ParseCtx, defs},
		{MatchTest{"-value1", valueFixed, false, false}, argument3ParseCtx, defs},
		// valueDefault
		{MatchTest{"value1", valueDefault, true, false}, command1ParseCtx, defs},
		{MatchTest{"any-value", valueDefault, false, true}, command1ParseCtx, defs},
		{MatchTest{"-value1", valueDefault, false, false}, command1ParseCtx, defs},
		// invalid token types
		{MatchTest{"", -1, false, true}, nil, defs},
		{MatchTest{"", math.MaxInt64, false, true}, nil, defs},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			m, err := match(tt.token, tt.tokenType, tt.ctx, tt.def)
			assertValEquals(t, m, tt.expected)
			assertError(t, err, tt.expError)
		})
	}

}
