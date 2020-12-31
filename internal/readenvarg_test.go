package internal

import (
	"os"
	"strings"
	"testing"
)

func TestArgEnv(t *testing.T) {
	names := []string{"empty_arg", "only_arg", "cmd_arg", "prefix_cmd_arg"}
	tests := []struct {
		prefix, cmd, arg string
		env              string
	}{
		{"p", "c", "", ""},
		{"", "", "a", "A"},
		{"", "c", "a", "C_A"},
		{"p", "c", "a", "P_C_A"},
	}

	for ii, tt := range tests {
		t.Run(names[ii], func(t *testing.T) {
			if argEnv(tt.prefix, tt.cmd, tt.arg) != tt.env {
				t.Error()
			}
		})
	}
}

func TestReadEnvArg(t *testing.T) {
	req := Request{}
	if err := req.readEnvArgs(nil); err == nil {
		t.Error()
	}
	// setup - make sure argument1 expects value from env
	defs := testDefs()
	defs.Arguments[0].Env = true

	// request arguments should use argument1 with 0 values
	req.Command = defs.Commands[0].Token
	req.Arguments = map[string][]string{
		defs.Arguments[0].Token: {},
	}
	// trivial validation that we're starting from an empty value
	if len(req.Arguments[defs.Arguments[0].Token]) > 0 {
		t.Error()
	}
	// store existing value of the env. variable
	envToken := argEnv(defs.EnvPrefix, defs.Commands[0].Token, defs.Arguments[0].Token)
	envValue := strings.ToLower(envToken)
	storedEnv := os.Getenv(envToken)
	// set the value we'll expect to see as argument1 value
	if err := os.Setenv(envToken, envValue); err != nil {
		t.Error(err.Error())
	}
	// read empty arguments values from env
	if err := req.readEnvArgs(defs); err != nil {
		t.Error(err.Error())
	}
	// there should be a value we got from env. variable
	if len(req.Arguments[defs.Arguments[0].Token]) < 1 {
		t.Error()
	}
	// reset env. variable value to original stored value
	if err := os.Setenv(envToken, storedEnv); err != nil {
		t.Error(err.Error())
	}
}
