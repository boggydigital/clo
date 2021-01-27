package internal

import (
	"strconv"
	"testing"
)

func TestArgEnv(t *testing.T) {
	tests := []struct {
		cmd, arg string
		env      string
	}{
		//{"c", "", ""},
		//{"", "a", "A"},
		//{"c", "a", "C_A"},
		//{"c", "a", "P_C_A"},
	}

	for ii, tt := range tests {
		t.Run(strconv.Itoa(ii), func(t *testing.T) {
			assertValEquals(t, argEnv(tt.cmd, tt.arg), tt.env)
		})
	}
}

//func TestReadEnvArg(t *testing.T) {
//	req := Request{}
//	assertError(t, req.readEnvArgs(nil), true)
//	// setup - make sure argument1 expects value from env
//	defs := mockDefinitions()
//	defs.Arguments[0].Env = true
//	// request arguments should use argument1 with 0 values
//	req.Command = defs.Commands[0].Token
//	req.Arguments = map[string][]string{
//		defs.Arguments[0].Token: {},
//	}
//	// trivial validation that we're starting from an empty value
//	assertValEquals(t, len(req.Arguments[defs.Arguments[0].Token]), 0)
//	// store existing value of the env. variable
//	envToken := argEnv(defs.EnvPrefix, defs.Commands[0].Token, defs.Arguments[0].Token)
//	envValue := strings.ToLower(envToken)
//	storedEnv := os.Getenv(envToken)
//	// set the value we'll expect to see as argument1 value
//	assertError(t, os.Setenv(envToken, envValue), false)
//	// read empty arguments values from env
//	assertError(t, req.readEnvArgs(defs), false)
//	// there should be a value we got from env. variable
//	assertValEquals(t, len(req.Arguments[defs.Arguments[0].Token]), 1)
//	// reset env. variable value to original stored value
//	assertError(t, os.Setenv(envToken, storedEnv), false)
//}