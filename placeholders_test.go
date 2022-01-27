package clo

import (
	"github.com/boggydigital/testo"
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
		{"val={" + defaultAttr + "}", &placeholder{"", false, false, false}},
		{"val={" + multipleAttr + "}", &placeholder{"", false, false, false}},
		{"val={id" + defaultAttr + "}", &placeholder{"id", false, true, false}},
		{"val={id" + multipleAttr + "}", &placeholder{"id", true, false, false}},
		{"val={id" + defaultAttr + multipleAttr + "}", &placeholder{"id", true, true, false}},
		{"val={id}", &placeholder{"id", false, false, false}},
		{"{}", &placeholder{}},
		{"{" + defaultAttr + "}", &placeholder{"", false, false, true}},
		{"{id" + defaultAttr + "}", &placeholder{"id", false, true, true}},
		{"{id" + multipleAttr + "}", &placeholder{"id", true, false, true}},
		{"{id" + defaultAttr + multipleAttr + "}", &placeholder{"id", true, true, true}},
		{"{id}", &placeholder{"id", false, false, true}},
		{" {id}", &placeholder{"id", false, false, false}},
		{"{id1}" + argValuesSep + "{id2}", &placeholder{"id1", false, false, true}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			ph := extract(tt.input)
			testo.DeepEqual(t, ph, tt.expected)
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		input    *placeholder
		expected string
	}{
		{&placeholder{}, "{}"},
		{&placeholder{"id", false, false, false}, "{id}"},
		{&placeholder{"id", false, false, true}, "{id}"},
		{&placeholder{"id", false, true, false}, "{id" + defaultAttr + "}"},
		{&placeholder{"id", true, false, false}, "{id" + multipleAttr + "}"},
		{&placeholder{"id", true, true, false}, "{id" + defaultAttr + multipleAttr + "}"},
		{&placeholder{"id", false, true, true}, "{id" + defaultAttr + "}"},
		{&placeholder{"id", true, true, true}, "{id" + defaultAttr + multipleAttr + "}"},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			testo.EqualValues(t, tt.input.String(), tt.expected)
		})
	}
}
