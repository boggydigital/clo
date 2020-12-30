package internal

import "testing"

func TestUpdate(t *testing.T) {
	pCtx := parseCtx{}
	defs := testDefs()

	if pCtx.Command != nil || pCtx.Argument != nil {
		t.Error("unexpected parse context initial condition")
	}

	// invalid updates
	// commands
	pCtx.update("command-that-doesnt-exist", command, defs)
	if pCtx.Command != nil {
		t.Error("unexpectedly set parse context command by command token that doesn't exist")
	}
	pCtx.update("command-abbr-that-doesnt-exist", commandAbbr, defs)
	if pCtx.Command != nil {
		t.Error("unexpectedly set parse context command  by command abbr that doesn't exist")
	}
	// arguments
	pCtx.update("--arg-that-doesnt-exist", argument, defs)
	if pCtx.Argument != nil {
		t.Error("unexpectedly set parse context command by arg token that doesn't exist")
	}
	pCtx.update("--arg-abbr-that-doesnt-exist", argumentAbbr, defs)
	if pCtx.Argument != nil {
		t.Error("unexpectedly set parse context command by arg abbr that doesn't exist")
	}

	// valid updates
	// commands
	pCtx.update("command1", command, defs)
	if pCtx.Command == nil {
		t.Error("parse context command wasn't set by a valid command token")
	}
	pCtx.update("c1", commandAbbr, defs)
	if pCtx.Command == nil {
		t.Error("parse context command wasn't set by a valid command abbr")
	}
	// arguments
	pCtx.update("--argument1", argument, defs)
	if pCtx.Argument == nil {
		t.Error("parse context command wasn't set by a valid argument token")
	}
	pCtx.update("--a1", argumentAbbr, defs)
	if pCtx.Argument == nil {
		t.Error("parse context command wasn't set by a valid argument abbr")
	}
}
