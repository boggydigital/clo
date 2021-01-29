package internal

type TokenHelpTest struct {
	token string
}

var mockHelpDefinitionsTests = []*Definitions{nil, mockDefinitions()}

var mockPrintArgumentHelpTests = []TokenHelpTest{
	{""},
	{"argument1"},
	{"argument2"},
	{"argument3"},
	{"argument-that-doesnt-exist"},
}

var mockPrintCommandHelpTests = []TokenHelpTest{
	{""},
	{"command1"},
	{"command2"},
	{"command-that-doesnt-exist"},
}
