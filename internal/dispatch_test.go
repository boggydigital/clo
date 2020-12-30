package internal

import (
	"testing"
)

func TestDispatch(t *testing.T) {
	writeDefs(testDefs(), t)
	// verify nil Request
	if err := Dispatch(nil); err != nil {
		t.Error("dispatch nil request error:", err.Error())
	}
	// verify help command
	if err := Dispatch(&Request{Command: "help"}); err != nil {
		t.Error("dispatch request with help command error:", err.Error())
	}
	// verify unknown command error
	if err := Dispatch(&Request{Command: "command-that-doesnt-exist"}); err == nil {
		t.Error("command that doesn't exist should create an error")
	}
	// cleanup
	deleteDefs(t)
}
