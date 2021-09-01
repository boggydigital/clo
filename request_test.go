package clo

import (
	"math"
	"strconv"
	"testing"
)

func TestRequestHasArguments(t *testing.T) {
	tests := []struct {
		req      *request
		expected bool
	}{
		{nil, false},
		{&request{}, false},
		{&request{Arguments: map[string][]string{}}, false},
		{&request{Arguments: map[string][]string{"1": {}}}, true},
		{&request{Arguments: map[string][]string{"1": {"2"}}}, true},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertValEquals(t, tt.req.hasArguments(), tt.expected)
		})
	}
}

func TestRequestSetDefaultContext(t *testing.T) {
	defs := mockDefinitions()
	tests := []struct {
		req       *request
		tokenType int
		expError  bool
		expCmd    string
		expArg    string
	}{
		{nil, command, false, "", ""},
		{&request{}, command, false, "", ""},
		{&request{}, argument, false, "command1", ""},
		{&request{lastArgument: "argument1"}, value, false, "", "argument1"},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertError(t, tt.req.setDefaultContext(tt.tokenType, defs), tt.expError)
			if tt.req != nil {
				assertValEquals(t, tt.req.Command, tt.expCmd)
				assertValEquals(t, tt.req.lastArgument, tt.expArg)
			}
		})
	}
}

func TestRequestUpdate(t *testing.T) {

	sequentialTests := []struct {
		token     string
		tokenType int
		expError  bool
	}{
		{"", command, false},
		{"command", command, false},
		{"command-overwrite", command, true},
		{"value", value, true},
		{"arg", argument, false},
		{"value", value, false},
		{"", -1, true},
		{"", math.MaxInt64, true},
	}
	req := &request{}
	for _, tt := range sequentialTests {
		t.Run(tt.token, func(t *testing.T) {
			err := req.update(tt.token, tt.tokenType)
			assertError(t, err, tt.expError)
		})
	}
}

type RequestTest struct {
	req      *request
	expError bool
}

var mockRequestCommandTests = []RequestTest{
	{nil, true}, // will be used to test defs == nil
	{nil, true},
	{&request{Command: "command2", Arguments: nil}, false},
	{&request{Command: "command1", Arguments: nil}, true},
	{&request{Command: "command1", Arguments: map[string][]string{"argument3": {}}}, true},
	{&request{Command: "command1", Arguments: map[string][]string{"argument1": {}}}, false},
	{&request{Command: "command1", Arguments: map[string][]string{"argument1": {"1"}}}, false},
	{&request{Command: "command-that-doesnt-exist", Arguments: map[string][]string{"argument1": {"1"}}}, true},
}

func TestRequestCommandHasRequiredArgs(t *testing.T) {
	for ii, tt := range mockRequestCommandTests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			defs := mockDefinitions()
			err := tt.req.commandHasRequiredArgs(defs)
			assertError(t, err, tt.expError)
		})
	}
}

var mockRequestArgumentTests = []RequestTest{
	{nil, true}, // will be used to test defs == nil
	{nil, true},
	{&request{Command: "command2", Arguments: map[string][]string{}}, false},
	{&request{Command: "command2", Arguments: map[string][]string{"": {}}}, false},
	{&request{Command: "command1", Arguments: map[string][]string{"argument1": {"1", "2"}}}, true},
	{&request{Command: "command2", Arguments: map[string][]string{"argument2": {"1", "2"}}}, false},
	{&request{Command: "command2", Arguments: map[string][]string{"argument-that-doesnt-exist": {"1", "2"}}}, false},
}

func TestArgumentsMultipleValues(t *testing.T) {
	for ii, tt := range mockRequestArgumentTests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			defs := mockDefinitions()
			err := tt.req.argumentsMultipleValues(defs)
			assertError(t, err, tt.expError)
		})
	}
}

func TestRequestVerify(t *testing.T) {
	tests := make([]RequestTest, 0)
	tests = append(tests, mockRequestCommandTests...)
	tests = append(tests, mockRequestArgumentTests...)

	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			defs := mockDefinitions()
			err := tt.req.validate(defs)
			assertError(t, err, tt.expError)
		})
	}
}

func TestRequestArgVal(t *testing.T) {
	tests := []struct {
		req      *request
		value    string
		expected string
	}{
		{nil, "", ""},
		{&request{Arguments: map[string][]string{"1": {"3"}, "2": {}}}, "1", "3"},
		{&request{Arguments: map[string][]string{"1": {"3"}, "2": {}}}, "3", ""},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertValEquals(t, tt.req.ArgVal(tt.value), tt.expected)
		})
	}
}

func TestRequestArgValues(t *testing.T) {
	tests := []struct {
		req      *request
		value    string
		expected int
	}{
		{nil, "", 0},
		{&request{Arguments: map[string][]string{"1": {"3", "4"}, "2": {}}}, "1", 2},
		{&request{Arguments: map[string][]string{"1": {"3", "4"}, "2": {}}}, "2", 0},
		{&request{Arguments: map[string][]string{"1": {"3"}, "2": {}}}, "3", 0},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertValEquals(t, len(tt.req.ArgValues(tt.value)), tt.expected)
		})
	}
}

func TestRequestFlag(t *testing.T) {
	tests := []struct {
		req      *request
		value    string
		expected bool
	}{
		{nil, "", false},
		{&request{Arguments: map[string][]string{"1": {"3"}, "2": {}}}, "1", true},
		{&request{Arguments: map[string][]string{"1": {"3"}, "2": {}}}, "2", true},
		{&request{Arguments: map[string][]string{"1": {"3"}, "2": {}}}, "3", false},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertValEquals(t, tt.req.Flag(tt.value), tt.expected)
		})
	}
}
