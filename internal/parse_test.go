package internal

import (
	"strings"
	"testing"
)

func TestDefinitionsParse(t *testing.T) {
	defs := mockDefinitions()
	tests := []struct {
		def      *Definitions
		args     []string
		req      *Request
		expError bool
	}{
		{nil, []string{}, nil, true},
		{defs, []string{""}, &Request{
			Command:   "",
			Arguments: map[string][]string{},
		}, false},
		{defs, []string{"c1", "-a1", "value1"},
			&Request{
				Command: "command1",
				Arguments: map[string][]string{
					"argument1": {"value1"},
				},
			},
			false,
		},
		{defs, []string{"c1", "-a1", "value-that-doesnt-exist"}, nil, true},
		{defs, []string{"command-that-doesnt-exist"}, nil, true},
		{defs, []string{"c1", "-a2", "value3", "value4"},
			&Request{
				Command: "command1",
				Arguments: map[string][]string{
					"argument2": {"value3", "value4"},
				},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, "-"), func(t *testing.T) {
			req, err := tt.def.Parse(tt.args)
			assertError(t, err, tt.expError)
			assertInterfaceEquals(t, req, tt.req)
		})
	}
}
