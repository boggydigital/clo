package internal

import "testing"

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
		name := tt.dupe
		if name == "" {
			name = "empty"
		}
		t.Run(name, func(t *testing.T) {
			dupe := firstDupe(tt.slice)
			if dupe != tt.dupe {
				t.Error("unexpected first dupe")
			}
		})
	}
}

func TestVFail(t *testing.T) {
	vFail("", true)
}

func TestVPass(t *testing.T) {
	vPass("", true)
}
