package internal

import (
	"math"
	"strconv"
	"testing"
)

func TestRequestUpdate(t *testing.T) {
	tests := []struct {
		token     string
		tokenType int
		expError  bool
	}{
		{"", command, false},
		{"command", command, false},
		{"command-overwrite", command, true},
		{"arg", argument, false},
		//{"v", value, mockParseCtx("", "arg"), false},
		{"", -1, true},
		{"", math.MaxInt64, true},
	}
	req := Request{
		Arguments: map[string][]string{},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			err := req.update(tt.token, tt.tokenType)
			assertError(t, err, tt.expError)
		})
	}
}

func TestCommandHasRequiredArgs(t *testing.T) {
	for ii, tt := range mockRequestCommandTests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			err := tt.req.commandHasRequiredArgs(tt.defs)
			assertError(t, err, tt.expError)
		})
	}
}

func TestArgumentsMultipleValues(t *testing.T) {
	for ii, tt := range mockRequestArgumentTests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			err := tt.req.argumentsMultipleValues(tt.defs)
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
			err := tt.req.verify(tt.defs)
			assertError(t, err, tt.expError)
		})
	}
}

func TestRequestFirstValue(t *testing.T) {
	tests := []struct {
		req      *Request
		value    string
		expected string
	}{
		{nil, "", ""},
		{&Request{Arguments: map[string][]string{"1": {"3"}, "2": {}}}, "1", "3"},
		{&Request{Arguments: map[string][]string{"1": {"3"}, "2": {}}}, "3", ""},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertValEquals(t, tt.req.ArgVal(tt.value), tt.expected)
		})
	}
}

func TestRequestGetValues(t *testing.T) {
	tests := []struct {
		req      *Request
		value    string
		expected int
	}{
		{nil, "", 0},
		{&Request{Arguments: map[string][]string{"1": {"3", "4"}, "2": {}}}, "1", 2},
		{&Request{Arguments: map[string][]string{"1": {"3"}, "2": {}}}, "3", 0},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertValEquals(t, len(tt.req.ArgValues(tt.value)), tt.expected)
		})
	}
}

func TestRequestGetFlag(t *testing.T) {
	tests := []struct {
		req      *Request
		value    string
		expected bool
	}{
		{nil, "", false},
		{&Request{Arguments: map[string][]string{"1": {"3"}, "2": {}}}, "1", true},
		{&Request{Arguments: map[string][]string{"1": {"3"}, "2": {}}}, "2", true},
		{&Request{Arguments: map[string][]string{"1": {"3"}, "2": {}}}, "3", false},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertValEquals(t, tt.req.Flag(tt.value), tt.expected)
		})
	}
}
