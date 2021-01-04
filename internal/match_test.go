package internal

import (
	"math"
	"testing"
)

func TestMatchArgument(t *testing.T) {
	tests := []MatchArgumentTest{
		{MatchTest{"", argument, false, false}, nil, nil},
		{MatchTest{"", argumentAbbr, false, false}, nil, nil},
		// shouldn't match
		// - tokens that don't exist
		{MatchTest{"-argument-that-doesnt-exist", argument, false, true}, nil, nil},
		{MatchTest{"-argument-that-doesnt-exist", argument, false, false}, nil, mockDefinitions()},
		{MatchTest{"--argument-that-doesnt-exist", argument, false, true}, nil, nil},
		{MatchTest{"--argument-that-doesnt-exist", argument, false, false}, nil, mockDefinitions()},
		{MatchTest{"-abbr-that-doesnt-exist", argumentAbbr, false, true}, nil, nil},
		{MatchTest{"-abbr-that-doesnt-exist", argumentAbbr, false, false}, nil, mockDefinitions()},
		{MatchTest{"--abbr-that-doesnt-exist", argumentAbbr, false, true}, nil, nil},
		{MatchTest{"--abbr-that-doesnt-exist", argumentAbbr, false, false}, nil, mockDefinitions()},
		// - tokens that exist, but are not valid
		{MatchTest{"-argument1", argument, false, true}, mockCommandDefinition("", []string{"not-argument1"}), mockDefinitions()},
		{MatchTest{"--argument1", argument, false, true}, mockCommandDefinition("", []string{"not-argument1"}), mockDefinitions()},
		{MatchTest{"-a1", argumentAbbr, false, true}, mockCommandDefinition("", []string{"not-argument1"}), mockDefinitions()},
		{MatchTest{"--a1", argumentAbbr, false, true}, mockCommandDefinition("", []string{"not-argument1"}), mockDefinitions()},
		// - nil command context, so can't validate
		{MatchTest{"-argument1", argument, false, true}, nil, mockDefinitions()},
		{MatchTest{"--argument1", argument, false, true}, nil, mockDefinitions()},
		{MatchTest{"-a1", argumentAbbr, false, true}, nil, mockDefinitions()},
		{MatchTest{"--a1", argumentAbbr, false, true}, nil, mockDefinitions()},
		// - tokenType is not argument or argumentAbbr
		{MatchTest{"-argument1", command, false, true}, nil, mockDefinitions()},
		{MatchTest{"-argument1", commandAbbr, false, true}, nil, mockDefinitions()},
		{MatchTest{"-argument1", valueDefault, false, true}, nil, mockDefinitions()},
		{MatchTest{"-argument1", valueFixed, false, true}, nil, mockDefinitions()},
		{MatchTest{"-argument1", value, false, true}, nil, mockDefinitions()},
		{MatchTest{"-argument1", flag, false, true}, nil, mockDefinitions()},
		{MatchTest{"-argument1", flagAbbr, false, true}, nil, mockDefinitions()},
		{MatchTest{"-argument1", -1, false, true}, nil, mockDefinitions()},
		{MatchTest{"-argument1", math.MaxInt64, false, true}, nil, mockDefinitions()},
		// should match
		// - valid args for command, no error
		{MatchTest{"-argument1", argument, true, false}, mockCommandDefinition("", []string{"argument1"}), mockDefinitions()},
		{MatchTest{"--argument1", argument, true, false}, mockCommandDefinition("", []string{"argument1"}), mockDefinitions()},
		{MatchTest{"-a1", argumentAbbr, true, false}, mockCommandDefinition("", []string{"argument1"}), mockDefinitions()},
		{MatchTest{"--a1", argumentAbbr, true, false}, mockCommandDefinition("", []string{"argument1"}), mockDefinitions()},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			m, err := matchArgument(tt.token, tt.tokenType, tt.cmd, tt.def)
			assertEquals(t, m, tt.expected)
			assertError(t, err, tt.expError)
		})
	}
}

func TestMatchFlag(t *testing.T) {
	tests := []MatchFlagTest{
		{MatchTest{"", flag, false, false}, nil},
		{MatchTest{"", flagAbbr, false, false}, nil},
		// shouldn't match
		// - tokens that don't exist
		{MatchTest{"-flag-that-doesnt-exist", flag, false, true}, nil},
		{MatchTest{"-flag-that-doesnt-exist", flag, false, false}, mockDefinitions()},
		{MatchTest{"--flag-that-doesnt-exist", flag, false, true}, nil},
		{MatchTest{"--flag-that-doesnt-exist", flag, false, false}, mockDefinitions()},
		{MatchTest{"-abbr-that-doesnt-exist", flagAbbr, false, true}, nil},
		{MatchTest{"-abbr-that-doesnt-exist", flagAbbr, false, false}, mockDefinitions()},
		{MatchTest{"--abbr-that-doesnt-exist", flagAbbr, false, true}, nil},
		{MatchTest{"--abbr-that-doesnt-exist", flagAbbr, false, false}, mockDefinitions()},
		// - tokenType is not flag or flagAbbr
		{MatchTest{"-flag1", command, false, true}, mockDefinitions()},
		{MatchTest{"-flag1", commandAbbr, false, true}, mockDefinitions()},
		{MatchTest{"-flag1", argument, false, true}, mockDefinitions()},
		{MatchTest{"-flag1", argumentAbbr, false, true}, mockDefinitions()},
		{MatchTest{"-flag1", valueDefault, false, true}, mockDefinitions()},
		{MatchTest{"-flag1", valueFixed, false, true}, mockDefinitions()},
		{MatchTest{"-flag1", value, false, true}, mockDefinitions()},
		{MatchTest{"-flag1", -1, false, true}, mockDefinitions()},
		{MatchTest{"-flag1", math.MaxInt64, false, true}, mockDefinitions()},
		// should match
		// - valid args for command, no error
		{MatchTest{"-flag1", flag, true, false}, mockDefinitions()},
		{MatchTest{"--flag1", flag, true, false}, mockDefinitions()},
		{MatchTest{"-f1", flagAbbr, true, false}, mockDefinitions()},
		{MatchTest{"--f1", flagAbbr, true, false}, mockDefinitions()},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			m, err := matchFlag(tt.token, tt.tokenType, tt.def)
			assertEquals(t, m, tt.expected)
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
		{MatchTest{"", flag, false, false}, nil, nil},
		{MatchTest{"", flagAbbr, false, false}, nil, nil},
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
			assertEquals(t, m, tt.expected)
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

	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			m, err := matchValue(tt.token, tt.tokenType, tt.arg)
			assertEquals(t, m, tt.expected)
			assertError(t, err, tt.expError)
		})
	}
}
