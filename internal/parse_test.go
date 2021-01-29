package internal

import (
	"strconv"
	"testing"
)

func TestDefinitionsParse(t *testing.T) {
	//defs := mockDefinitions()
	tests := []struct {
		args   []string
		req    *Request
		expErr bool
	}{
		{[]string{}, nil, true},
		{[]string{"--argument1"}, &Request{Command: "command1", Arguments: map[string][]string{"argument1": {}}}, false},
	}
	//{nil, []string{}, nil, true},
	//{defs, []string{""}, &Request{
	//	Command:   "",
	//	Arguments: map[string][]string{},
	//}, false},
	//{defs, []string{"c1", "-a1", "value1"},
	//	&Request{
	//		Command: "command1",
	//		Arguments: map[string][]string{
	//			"argument1": {"value1"},
	//		},
	//	},
	//	false,
	//},
	//{defs, []string{"c1", "-a1", "value-that-doesnt-exist"}, nil, true},
	//{defs, []string{"command-that-doesnt-exist"}, nil, true},
	//{defs, []string{"c1", "-a2", "value3", "value4"},
	//	&Request{
	//		Command: "command1",
	//		Arguments: map[string][]string{
	//			"argument2": {"value3", "value4"},
	//		},
	//	},
	//	true,
	//},
	//}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			defs := mockDefinitions()
			if ii == 0 {
				defs = nil
			}
			req, err := defs.Parse(tt.args)
			assertError(t, err, tt.expErr)
			assertInterfaceEquals(t, req, tt.req)
		})
	}
}
