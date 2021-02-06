package clo

import (
	"errors"
	"strconv"
	"testing"
)

func TestFirstDupe(t *testing.T) {
	tests := []struct {
		slice []string
		dupe  string
	}{
		{[]string{}, ""},
		{[]string{"1"}, ""},
		{[]string{"1", "2", "3"}, ""},
		{[]string{"1", "2", "3", "1"}, "1"},
		{[]string{"1", "2", "3", "2"}, "2"},
		{[]string{"1", "2", "3", "3"}, "3"},
	}

	for _, tt := range tests {
		t.Run(tt.dupe, func(t *testing.T) {
			assertValEquals(t, firstDupe(tt.slice), tt.dupe)
		})
	}
}

func TestAppendError(t *testing.T) {
	tests := []error{errors.New(""), nil}
	errs := make([]error, 0)
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			errs = appendErr(errs, tt)
			assertValEquals(t, len(errs), 1)
		})
	}
}

func TestDefinitionsVerify(t *testing.T) {
	// We've already verified individual error cases above
	// so running known good definitions for the coverage
	defs := mockDefinitions()
	errs := defs.Validate(false)
	assertValEquals(t, len(errs), 0)
}
