package internal

import (
	"testing"
)

func TestDispatch(t *testing.T) {
	tests := []struct {
		request  *Request
		expError bool
	}{
		{nil, false},
		{&Request{Command: "help"}, false},
		{&Request{Command: "command-that-doesnt-exist"}, true},
	}
	writeDefs(testDefs(), t)
	t.Cleanup(deleteDefs)

	for _, tt := range tests {
		name := "nil"
		if tt.request != nil {
			name = tt.request.Command
		}
		t.Run(name, func(t *testing.T) {
			err := Dispatch(tt.request)
			assertError(t, err, tt.expError)
		})
	}
}
