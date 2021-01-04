package internal

type TokenTest struct {
	token  string
	expNil bool
}

type TokensTest struct {
	tokens   []string
	expError bool
}

type RequestTest struct {
	req      *Request
	defs     *Definitions
	expError bool
}

var mockValidityTests = []struct {
	values   []string
	value    string
	expected bool
}{
	{nil, "any", false},
	{[]string{}, "any", false},
	{[]string{"value1"}, "value1", true},
	{[]string{"value-that-doesnt-exist"}, "value1", false},
}

var mockExamplesTests = []TokensTest{
	{nil, false},
	{[]string{}, false},
	{[]string{"1", "2"}, false},
	{[]string{"", ""}, true},
}

func mockNoEmptyTokensTests() []TokensTest {
	return []TokensTest{
		{nil, false},
		{[]string{}, false},
		{[]string{"1", "2"}, false},
		{[]string{"", "x"}, true},
		{[]string{"x", ""}, true},
	}
}

func mockDifferentTokensTests() []TokensTest {
	return []TokensTest{
		{nil, false},
		{[]string{}, false},
		{[]string{"1"}, false},
		{[]string{"1", "2"}, false},
		{[]string{"1", "2", "2"}, true},
	}
}

func mockByTokenAbbrTests(prefix string) []TokenTest {
	return []TokenTest{
		// valid token/abbr
		{prefix + "1", false},
		// invalid token/abbr
		{prefix + "-token-that-doesnt-exist", true},
	}
}

var mockRequestCommandTests = []RequestTest{
	{nil, nil, true},
	{nil, mockDefinitions(), true},
	{&Request{Command: "command2", Arguments: nil}, mockDefinitions(), false},
	{&Request{Command: "command1", Arguments: nil}, mockDefinitions(), true},
	{&Request{Command: "command1", Arguments: map[string][]string{"argument3": {}}}, mockDefinitions(), true},
	{&Request{Command: "command1", Arguments: map[string][]string{"argument1": {}}}, mockDefinitions(), true},
	{&Request{Command: "command1", Arguments: map[string][]string{"argument1": {"1"}}}, mockDefinitions(), false},
}

var mockRequestArgumentTests = []RequestTest{
	{nil, nil, true},
	{nil, mockDefinitions(), true},
	{&Request{Arguments: map[string][]string{}}, mockDefinitions(), false},
	{&Request{Arguments: map[string][]string{"": {}}}, mockDefinitions(), false},
	{&Request{Arguments: map[string][]string{"argument1": {"1", "2"}}}, mockDefinitions(), false},
	{&Request{Arguments: map[string][]string{"argument2": {"1", "2"}}}, mockDefinitions(), true},
	{&Request{Arguments: map[string][]string{"argument3": {"1", "2"}}}, mockDefinitions(), false},
}

var mockHelpDefinitionsTests = []*Definitions{nil, mockDefinitions()}
