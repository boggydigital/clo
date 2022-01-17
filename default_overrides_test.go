package clo

import (
	"strconv"
	"testing"
)

func TestSetValidUserDefaults(t *testing.T) {

	tests := []struct {
		overrides map[string][]string
		expError  bool
		expLen    int
	}{
		{
			map[string][]string{"command2" + cmdArgValDefSep + "xyz": {"userdefaultval"}},
			false,
			1,
		},
		{
			map[string][]string{"xyz": {"userdefaultval"}},
			false,
			1,
		},
		{
			map[string][]string{"invalidcommand" + cmdArgValDefSep + "invalidarg": {""}},
			true,
			0,
		},
		{
			map[string][]string{"invalidcommand" + cmdArgValDefSep + "invalidarg": {""}},
			true,
			0,
		},
	}

	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			defs := mockDefinitions()
			assertValEquals(t, len(defs.defaultsOverrides), 0)

			err := defs.SetUserDefaults(tt.overrides)
			assertError(t, err, tt.expError)

			assertValEquals(t, len(defs.defaultsOverrides), tt.expLen)
		})
	}
}

func TestHasUserDefaultsFlag(t *testing.T) {
	tests := []struct {
		overrides map[string][]string
		flag      string
		expVal    bool
	}{
		{
			nil,
			"xyz",
			false,
		},
		{
			map[string][]string{"xyz": {"userdefaultval"}},
			"xyz",
			true,
		},
	}

	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			defs := mockDefinitions()

			err := defs.SetUserDefaults(tt.overrides)
			assertError(t, err, false)

			assertValEquals(t, defs.HasUserDefaultsFlag(tt.flag), tt.expVal)
		})
	}
}
