package internal

import (
	"strconv"
	"testing"
)

func TestArgWithInvalidValuesHasNoValidValues(t *testing.T) {
	ad := ArgumentDefinition{
		Values: nil,
	}

	if ad.ValidValue("any") {
		t.Error("argument with nil values shouldn't have valid values")
	}

	ad.Values = make([]string, 0)

	if ad.ValidValue("any") {
		t.Error("argument with empty values shouldn't have valid values")
	}
}

func TestArgValidValueCanBeFound(t *testing.T) {
	cVals := 3
	ad := ArgumentDefinition{
		Values: make([]string, cVals),
	}
	for i := 0; i < cVals; i++ {
		ad.Values[i] = strconv.Itoa(i + 1)
	}

	for i := 0; i < cVals; i++ {
		if !ad.ValidValue(strconv.Itoa(i + 1)) {
			t.Errorf("expected value '%d' to be valid", i+1)
		}
	}
}
