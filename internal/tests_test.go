package internal

type TokenTest struct {
	token  string
	expNil bool
}

type TokensTest struct {
	tokens   []string
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
