package clo

import (
	"strconv"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		input    string
		expected *placeholder
	}{
		{"", &placeholder{}},
		{"val={", &placeholder{}},
		{"val=}id{", &placeholder{}},
		{"val={}", &placeholder{}},
		{"val={_}", &placeholder{"_", false, false}},
		{"val={id_}", &placeholder{"id", true, false}},
		{"val={id}", &placeholder{"id", false, false}},
		{"{}", &placeholder{}},
		{"{_}", &placeholder{"_", false, true}},
		{"{id_}", &placeholder{"id", true, true}},
		{"{id}", &placeholder{"id", false, true}},
		{" {id}", &placeholder{"id", false, false}},
		{"{id1}={id2}", &placeholder{"id1", false, true}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			ph := extract(tt.input)
			assertInterfaceEquals(t, ph, tt.expected)
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		input    *placeholder
		expected string
	}{
		{&placeholder{}, "{}"},
		{&placeholder{"id", false, false}, "{id}"},
		{&placeholder{"id", false, true}, "{id}"},
		{&placeholder{"id", true, false}, "{id_}"},
		{&placeholder{"id", true, true}, "{id_}"},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertValEquals(t, tt.input.String(), tt.expected)
		})
	}
}
