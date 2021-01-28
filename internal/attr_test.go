package internal

import (
	"fmt"
	"testing"
)

type AttrTest struct {
	token     string
	trimToken string
	expected  bool
}

func testAttr(t *testing.T, tests []AttrTest, isAttr func(string) bool) {
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			assertValEquals(t, isAttr(tt.token), tt.expected)
		})
	}
}

func genTests(attr string) []AttrTest {
	const val = "value"
	return []AttrTest{
		{token: val, trimToken: val, expected: false},
		{token: val + attr, trimToken: val, expected: true},
		{token: attr + val, trimToken: val, expected: true},
		{token: val + attr + val, trimToken: val + attr + val, expected: true},
	}
}

func TestAttrs(t *testing.T) {
	attrs := []string{defaultAttr, requiredAttr, multipleAttr, envAttr, argValuesSep}
	names := []string{"isDefault", "isRequired", "isMultiple", "isEnv", "hasArgValues"}
	attrFunc := []func(string) bool{isDefault, isRequired, isMultiple, isEnv, hasArgValues}

	for i, attr := range attrs {
		t.Run(names[i], func(t *testing.T) {
			testAttr(t, genTests(attr), attrFunc[i])
		})
	}
}

var argValuesTests = []struct {
	token  string
	arg    string
	values []string
}{
	{"", "", []string{}},
	{"abc", "abc", []string{}},
	{"=a", "", []string{"a"}},
	{"a=", "a", []string{""}},
	{"a=b", "a", []string{"b"}},
	{"a=b,c", "a", []string{"b", "c"}},
	{"a,b=c", "a,b", []string{"c"}},
	{"a=b=c", "a", []string{"c"}},
}

func TestSplitArgValues(t *testing.T) {
	for _, tt := range argValuesTests {
		t.Run(tt.token, func(t *testing.T) {
			a, v := splitArgValues(tt.token)
			fmt.Println(a, v)
			assertValEquals(t, a, tt.arg)
			assertInterfaceEquals(t, v, tt.values)
		})
	}
}

func TestTrimArgValue(t *testing.T) {
	for _, tt := range argValuesTests {
		t.Run(tt.token, func(t *testing.T) {
			a := trimArgValues(tt.token)
			assertValEquals(t, a, tt.arg)
		})
	}
}

func testTrimAttr(t *testing.T, tests []AttrTest) {
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			assertValEquals(t, trimAttrs(tt.token), tt.trimToken)
		})
	}
}

func TestTrimAttrs(t *testing.T) {
	attrs := []string{defaultAttr, requiredAttr, multipleAttr, envAttr}
	names := []string{"isDefault", "isRequired", "isMultiple", "isEnv"}

	for i, attr := range attrs {
		t.Run(names[i], func(t *testing.T) {
			testTrimAttr(t, genTests(attr))
		})
	}
}
