package clo

import (
	"github.com/boggydigital/testo"
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
			testo.EqualValues(t, len(defs.defaultsOverrides), 0)

			err := defs.SetUserDefaults(tt.overrides)
			testo.Error(t, err, tt.expError)

			testo.EqualValues(t, len(defs.defaultsOverrides), tt.expLen)
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
			testo.Error(t, err, false)

			testo.EqualValues(t, defs.HasUserDefaultsFlag(tt.flag), tt.expVal)
		})
	}
}
