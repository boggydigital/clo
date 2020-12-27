package internal

import (
	"strconv"
	"testing"
)

func TestCmdWithInvalidArgsHasNoValidArgs(t *testing.T) {
	cd := CommandDefinition{
		Arguments: nil,
	}

	if cd.ValidArg("any") {
		t.Error("command with nil arguments shouldn't have valid args")
	}

	cd.Arguments = make([]string, 0)

	if cd.ValidArg("any") {
		t.Error("command with nil arguments shouldn't have valid args")
	}
}

func TestCmdValidArgCanBeFound(t *testing.T) {
	cVals := 3
	cd := CommandDefinition{
		Arguments: make([]string, cVals),
	}
	for i := 0; i < cVals; i++ {
		cd.Arguments[i] = strconv.Itoa(i + 1)
	}

	for i := 0; i < cVals; i++ {
		if !cd.ValidArg(strconv.Itoa(i + 1)) {
			t.Errorf("expected argument '%d' to be valid", i+1)
		}
	}
}
