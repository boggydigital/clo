package internal

import (
	"testing"
)

func TestValidArg(t *testing.T) {
	for ii, tt := range validityTests {
		t.Run(validityNames[ii], func(t *testing.T) {
			ad := CommandDefinition{
				Arguments: tt.values,
			}
			if ad.ValidArg(tt.value) != tt.expected {
				t.Error("unexpected value validity:", tt.value)
			}
		})
	}
}
