package internal

import (
	"testing"
)

func TestDispatch(t *testing.T) {
	tests := []struct {
		writeDefs func(*testing.T)
		request   *Request
		expError  bool
	}{
		{writeDefaultMockDefs, nil, false},
		{writeDefaultMockDefs, &Request{Command: "help"}, false},
		{writeDefaultMockDefs, &Request{Command: "command-that-doesnt-exist"}, true},
		{writeEmptyMockDefs, nil, true},
	}

	for _, tt := range tests {
		name := "nil"
		if tt.request != nil {
			name = tt.request.Command
		}
		t.Run(name, func(t *testing.T) {
			tt.writeDefs(t)
			//writeMockDefs(tt.defs, t)
			t.Cleanup(deleteMockDefs)
			err := Route(tt.request)
			assertError(t, err, tt.expError)
		})
	}
}
