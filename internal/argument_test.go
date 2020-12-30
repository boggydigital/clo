package internal

import (
	"testing"
)

var validityNames = []string{"nil", "empty", "valid"}
var validityTests = []struct {
	values   []string
	value    string
	expected bool
}{
	{nil, "any", false},
	{[]string{}, "any", false},
	{[]string{"value1"}, "value1", true},
}

func TestValidValue(t *testing.T) {
	for ii, tt := range validityTests {
		t.Run(validityNames[ii], func(t *testing.T) {
			ad := ArgumentDefinition{
				Values: tt.values,
			}
			if ad.ValidValue(tt.value) != tt.expected {
				t.Error("unexpected value validity:", tt.value)
			}
		})
	}
}
