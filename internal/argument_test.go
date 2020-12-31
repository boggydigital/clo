package internal

import (
	"testing"
)

var validityNames = []string{"nil", "empty", "valid", "value-that-doesnt-exist"}
var validityTests = []struct {
	values   []string
	value    string
	expected bool
}{
	{nil, "any", false},
	{[]string{}, "any", false},
	{[]string{"value1"}, "value1", true},
	{[]string{"value-that-doesnt-exist"}, "value1", false},
}

func assertEquals(t *testing.T, v1, v2 interface{}) {
	if v1 != v2 {
		t.Error()
	}
}

func TestValidValue(t *testing.T) {
	for ii, tt := range validityTests {
		t.Run(validityNames[ii], func(t *testing.T) {
			ad := ArgumentDefinition{Values: tt.values}
			assertEquals(t, ad.ValidValue(tt.value), tt.expected)
		})
	}
}
