package internal

import "testing"

func TestArgumentWithNilValuesHasNoValidValues(t *testing.T) {
	ad := ArgumentDefinition{
		Values: nil,
	}
	if ad.ValidValue("any") {
		t.Error("argument with nil values shouldn't have valid values")
	}
}

func TestArgumentWithEmptyValuesHasNoValidValues(t *testing.T) {
	ad := ArgumentDefinition{
		Values: make([]string, 0),
	}
	if ad.ValidValue("any") {
		t.Error("argument with empty values shouldn't have valid values")
	}
}
