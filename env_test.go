package clo

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestArgEnv(t *testing.T) {
	tests := []struct {
		app, cmd, arg string
		env           string
	}{
		{"p", "c", "", ""},
		{"", "", "a", "A"},
		{"", "c", "a", "C_A"},
		{"p", "c", "a", "P_C_A"},
	}

	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertValEquals(t, argEnv(tt.app, tt.cmd, tt.arg), tt.env)
		})
	}
}

func TestReadEnvArg(t *testing.T) {
	tests := []struct {
		cmd, arg string
		expError bool
	}{
		{"cmd-that-doesnt-exist", "arg-that-doesnt-exist", true},
		{"command1", "argument1", false},
	}

	for _, tt := range tests {
		t.Run(tt.cmd+tt.arg, func(t *testing.T) {
			req := request{}
			assertError(t, req.readEnvArgs(nil), true)
			// setup - make sure argument1 expects value from env
			defs := mockDefinitions()
			// request arguments should use argument1 with 0 values
			req.Command = tt.cmd
			req.Arguments = map[string][]string{
				tt.arg: {},
			}
			// trivial validation that we're starting from an empty value
			assertValEquals(t, len(req.Arguments[tt.arg]), 0)
			// store existing value of the env. variable
			envToken := argEnv(appName(), tt.cmd, tt.arg)
			envValue := strings.ToLower(envToken)
			storedEnv := os.Getenv(envToken)
			// set the value we'll expect to see as argument1 value
			assertError(t, os.Setenv(envToken, envValue), false)
			// read empty arguments values from env
			assertError(t, req.readEnvArgs(defs), tt.expError)
			// there should be a value we got from env. variable
			expArgValLen := 1
			if tt.expError {
				expArgValLen = 0
			}
			assertValEquals(t, len(req.Arguments[tt.arg]), expArgValLen)
			// reset env. variable value to original stored value
			assertError(t, os.Setenv(envToken, storedEnv), false)
		})
	}
}
