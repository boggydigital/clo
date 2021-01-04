package internal

import (
	"strings"
	"testing"
)

func TestCommandDefinitionValidArg(t *testing.T) {
	for _, tt := range mockValidityTests {
		t.Run(strings.Join(tt.values, "-"), func(t *testing.T) {
			cd := mockCommandDefinition("", tt.values)
			assertEquals(t, cd.ValidArg(tt.value), tt.expected)
		})
	}
}
