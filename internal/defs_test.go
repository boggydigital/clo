package internal

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

const defaultFilename = "clo.json"

func testDefs() *Definitions {
	defs := &Definitions{
		Version:   1,
		EnvPrefix: "CORRECT",
		App:       "clo",
		Hint:      "hint",
		Desc:      "desc",
		Flags: []FlagDefinition{
			{
				CommonDefinition: CommonDefinition{
					Token: "flag1",
					Abbr:  "f1",
				},
			},
			{
				CommonDefinition: CommonDefinition{
					Token: "flag2",
					Abbr:  "f2",
				},
			},
		},
		Commands: []CommandDefinition{
			{
				CommonDefinition: CommonDefinition{
					Token: "command1",
					Abbr:  "c1",
				},
				Arguments: []string{
					"argument1",
					"argument2",
				},
			},
			{
				CommonDefinition: CommonDefinition{
					Token: "command2",
				},
			},
		},
		Arguments: []ArgumentDefinition{
			{
				CommonDefinition: CommonDefinition{
					Token: "argument1",
					Abbr:  "a1",
				},
				Default:  true,
				Multiple: true,
				Required: true,
				Values:   []string{"value1", "value2"},
			},
			{
				CommonDefinition: CommonDefinition{
					Token: "argument2",
					Abbr:  "a2",
				},
				Values: []string{"value3", "value4"},
			},
		},
		Values: []ValueDefinition{
			{
				CommonDefinition: CommonDefinition{
					Token: "value1",
				},
			},
		},
	}

	return defs
}

func breakDefinitions(defs *Definitions) {
	defs.Arguments = append(defs.Arguments, ArgumentDefinition{
		CommonDefinition: CommonDefinition{
			Token: "help:command",
		},
		Values: []string{"from:nowhere"},
	})
}

func writeDefs(defs *Definitions, t *testing.T) {
	//defs := testDefs()

	//if addLoadBreaks {
	//	defs.Arguments = append(defs.Arguments, ArgumentDefinition{
	//		CommonDefinition: CommonDefinition{
	//			Token: "help:command",
	//		},
	//		Values: []string{"from:nowhere"},
	//	})
	//}

	if _, err := os.Stat(defaultFilename); os.IsNotExist(err) {
		jsonBytes, err := json.Marshal(defs)
		if err != nil {
			t.Error("cannot serialize test definitions")
			return
		}
		err = ioutil.WriteFile(defaultFilename, jsonBytes, 0644)
		if err != nil {
			t.Error("cannot write test definitions")
			return
		}
	} else {
		t.Errorf("Definitions already exist at path %s", defaultFilename)
	}
}

func writeUnreadableDefs(t *testing.T) {
	if _, err := os.Stat(defaultFilename); os.IsNotExist(err) {
		err = ioutil.WriteFile(defaultFilename, []byte{}, 0644)
		if err != nil {
			t.Error("cannot write test definitions")
		}
	} else {
		t.Errorf("Definitions already exist at path %s", defaultFilename)
	}
}

func deleteDefs(t *testing.T) {
	if os.Remove(defaultFilename) != nil {
		t.Errorf("cannot remove test definitions file at %s", defaultFilename)
	}
}

func TestLoad(t *testing.T) {
	// Load adds 'help' command
	defs := testDefs()
	writeDefs(defs, t)
	helpCmd := defs.CommandByToken("help")
	if helpCmd != nil {
		t.Error("test definitions already contain help token")
		return
	}

	defs, err := LoadDefault()
	if defs == nil || err != nil {
		t.Error("cannot load default definitions")
		return
	}
	helpCmd = defs.CommandByToken("help")
	if helpCmd == nil {
		t.Error("Load definitions didn't add 'help' command")
		return
	}
	// Load fails with a path that doesn't exist
	defs, err = Load("path/that/doesnt/exist")
	if defs != nil || err == nil {
		t.Error("loaded definitions at path that shouldn't exist")
	}
	// cleanup
	deleteDefs(t)
}

func TestLoadErrors(t *testing.T) {
	// Load fails with known breaks:
	// - Pre-existing "help:command" argument
	// - Pre-existing "from:nowhere" reference value
	defs := testDefs()
	breakDefinitions(defs)
	writeDefs(defs, t)
	if defs, err := LoadDefault(); defs != nil || err == nil {
		t.Error("Load should break with known content problems")
	}
	//cleanup
	deleteDefs(t)

	// Load empty defs
	writeUnreadableDefs(t)
	if defs, err := LoadDefault(); defs != nil || err == nil {
		t.Error("Load should not return definitions for empty file")
	}
	// cleanup
	deleteDefs(t)
}

func TestFlagBy(t *testing.T) {
	defs := testDefs()
	fd := defs.FlagByToken("flag1")
	if fd == nil {
		t.Error("cannot find an expected flag by token")
	}
	fd = defs.FlagByToken("flag-token-that-doesnt-exist")
	if fd != nil {
		t.Error("found a flag that doesn't exist by token")
	}
	fd = defs.FlagByAbbr("f1")
	if fd == nil {
		t.Error("cannot find an expected flag by abbr")
	}
	fd = defs.FlagByAbbr("flag-abbr-that-doesnt-exist")
	if fd != nil {
		t.Error("found a flag that doesn't exist by abbr")
	}
}

func TestCommandBy(t *testing.T) {
	defs := testDefs()
	cd := defs.CommandByToken("command1")
	if cd == nil {
		t.Error("cannot find an expected command by token")
	}
	cd = defs.CommandByToken("command-token-that-doesnt-exist")
	if cd != nil {
		t.Error("found a command that doesn't exist by token")
	}
	cd = defs.CommandByAbbr("c1")
	if cd == nil {
		t.Error("cannot find an expected command by abbr")
	}
	cd = defs.CommandByAbbr("command-abbr-that-doesnt-exist")
	if cd != nil {
		t.Error("found a command that doesn't exist by abbr")
	}
}

func TestArgBy(t *testing.T) {
	defs := testDefs()
	ad := defs.ArgByToken("argument1")
	if ad == nil {
		t.Error("cannot find an expected argument by token")
	}
	ad = defs.ArgByToken("argument-token-that-doesnt-exist")
	if ad != nil {
		t.Error("found an argument that doesn't exist by token")
	}
	ad = defs.ArgByAbbr("a1")
	if ad == nil {
		t.Error("cannot find an expected argument by abbr")
	}
	ad = defs.ArgByAbbr("arg-abbr-that-doesnt-exist")
	if ad != nil {
		t.Error("found an argument that doesn't exist by abbr")
	}
}

func TestValueBy(t *testing.T) {
	defs := testDefs()
	vd := defs.ValueByToken("value1")
	if vd == nil {
		t.Error("cannot find an expected value by token")
	}
	vd = defs.ValueByToken("value-token-that-doesnt-exist")
	if vd != nil {
		t.Error("found a value that doesn't exist by token")
	}
}

func TestDefinedValue(t *testing.T) {
	defs := testDefs()
	if defs.DefinedValue(nil) {
		t.Error("nil values cannot be defined")
	}
	if defs.DefinedValue([]string{}) {
		t.Error("empty values cannot be defined")
	}
	if !defs.DefinedValue([]string{"value1"}) {
		t.Error("unexpected undefined value")
	}
	if defs.DefinedValue([]string{"value-that-doesnt-exist"}) {
		t.Error("unexpected defined value")
	}
}

func TestDefaultArg(t *testing.T) {
	defs := testDefs()
	cmd := defs.CommandByToken("command1")
	if cmd == nil {
		t.Errorf("cannot find required command: command1")
		return
	}

	if defs.DefaultArg(nil) != nil {
		t.Error("nil commands shouldn't have default arg")
	}

	newArgs := make([]string, 0)
	newArgs = append(newArgs, "unknown-argument")
	for _, a := range cmd.Arguments {
		newArgs = append(newArgs, a)
	}
	cmd.Arguments = newArgs
	if defs.DefaultArg(cmd) == nil {
		t.Errorf("%s is expected to have a default arg", cmd.Token)
	}

	cmd = defs.CommandByToken("command2")
	if cmd == nil {
		t.Errorf("cannot find required command: command2")
		return
	}

	if defs.DefaultArg(cmd) != nil {
		t.Errorf("%s is NOT expected to have a default arg", cmd.Token)
	}
}

func TestRequiredArgs(t *testing.T) {
	defs := testDefs()

	if len(defs.RequiredArgs("command-that-doesnt-exist")) != 0 {
		t.Error("commands that don't exist shouldn't have required args")
	}

	defs.Commands[0].Arguments = append(defs.Commands[0].Arguments, "unknown-argument")
	if len(defs.RequiredArgs(defs.Commands[0].Token)) != 1 {
		t.Errorf("%s is expected to have 1 required argument",
			defs.Commands[0].Token)
	}

}

func TestValidArgVal(t *testing.T) {
	defs := testDefs()

	if defs.ValidArgVal("", "") {
		t.Error("empty arguments shouldn't have valid args")
	}
	if defs.ValidArgVal("argument-that-doesnt-exist", "") {
		t.Error("arguments that don't exist shouldn't have valid args")
	}
	if !defs.ValidArgVal("argument1", "value1") {
		t.Error("cannot confirm a valid argument")
	}
}
