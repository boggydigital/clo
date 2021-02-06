package clo

import (
	"strconv"
	"testing"
)

func TestRoute(t *testing.T) {
	tests := []struct {
		req      *Request
		expError bool
	}{
		{nil, true},
		{&Request{Command: "help"}, true},
		{&Request{Command: "command-that-doesnt-exist"}, true},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			err := Route(tt.req, nil)
			assertError(t, err, tt.expError)
		})
	}
}
