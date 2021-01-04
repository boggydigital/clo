package internal

import (
	"strings"
	"testing"
)

func TestValidValue(t *testing.T) {
	for _, tt := range mockValidityTests {
		t.Run(strings.Join(tt.values, "-"), func(t *testing.T) {
			ad := mockArgumentDefinition("", tt.values)
			assertEquals(t, ad.ValidValue(tt.value), tt.expected)
		})
	}
}
