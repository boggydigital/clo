package internal

import (
	"strings"
	"testing"
)

func TestDefinitionsParse(t *testing.T) {
	tests := []struct {
		def      *Definitions
		args     []string
		req      *Request
		expError bool
	}{
		{nil, []string{}, nil, true},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, "-"), func(t *testing.T) {
			_, err := tt.def.Parse(tt.args)
			assertError(t, err, tt.expError)
			// TODO: verify request Equals
		})
	}
}
