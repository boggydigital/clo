package clo

import (
	"github.com/boggydigital/testo"
	"strconv"
	"testing"
)

func TestDefinitionsParse(t *testing.T) {
	tests := []struct {
		args       []string
		expCmd     string
		expLastArg string
		expLenArgs int
		expErr     bool
	}{
		{[]string{"", "command1", "--argument1"}, "command1", "argument1", 1, false},
		{[]string{"--argument1"}, "command1", "argument1", 1, false},
		{[]string{"unknown-token"}, "command1", "argument1", 1, false},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			defs := mockDefinitions()
			req, err := defs.parseRequest(tt.args)
			testo.Error(t, err, tt.expErr)
			testo.EqualValues(t, req.Command, tt.expCmd)
			testo.EqualValues(t, req.lastArgument, tt.expLastArg)
			testo.EqualValues(t, len(req.Arguments), tt.expLenArgs)
		})
	}
}

func TestDefinitionsNoDefaultsParse(t *testing.T) {
	tests := []struct {
		args       []string
		expCmd     string
		expLastArg string
		expLenArgs int
		expErr     bool
	}{
		{[]string{"command2"}, "", "", 0, true},
		{[]string{"command1", "argument1", "value-that-doesnt-exist"}, "command1", "", 0, true},
	}
	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			defs := mockDefinitionsNoDefaults()
			req, err := defs.parseRequest(tt.args)
			testo.Error(t, err, tt.expErr)
			testo.EqualValues(t, req.Command, tt.expCmd)
			testo.EqualValues(t, req.lastArgument, tt.expLastArg)
			testo.EqualValues(t, len(req.Arguments), tt.expLenArgs)
		})
	}
}
