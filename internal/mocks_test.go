package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

const defaultMockFilename = "clo.json"

func mockDefinitions() *Definitions {
	return &Definitions{
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
}

func breakMockDefinitions(defs *Definitions) {
	defs.Arguments = append(defs.Arguments, ArgumentDefinition{
		CommonDefinition: CommonDefinition{
			Token: "help:command",
		},
		Values: []string{"from:nowhere"},
	})
}

func writeMockDefs(defs *Definitions, t *testing.T) {
	if _, err := os.Stat(defaultMockFilename); os.IsNotExist(err) {
		jsonBytes, err := json.Marshal(defs)
		if err != nil {
			t.FailNow()
		}
		err = ioutil.WriteFile(defaultMockFilename, jsonBytes, 0644)
		if err != nil {
			t.FailNow()
		}
	} else {
		t.Error()
	}
}

func writeEmptyMockDefs(t *testing.T) {
	if _, err := os.Stat(defaultMockFilename); os.IsNotExist(err) {
		err = ioutil.WriteFile(defaultMockFilename, []byte{}, 0644)
		if err != nil {
			t.Error()
		}
	} else {
		t.Error()
	}
}

func deleteMockDefs() {
	if os.Remove(defaultMockFilename) != nil {
		log.Printf("cannot remove test definitions file at %s", defaultMockFilename)
	}
}

func loadMockPathThatDoesntExist() (*Definitions, error) {
	return Load("path/that/doesnt/exist")
}

func setupBrokenMockDefs(t *testing.T) {
	defs := mockDefinitions()
	breakMockDefinitions(defs)
	writeMockDefs(defs, t)
	t.Cleanup(deleteMockDefs)
}

func setupEmptyMockDefs(t *testing.T) {
	writeEmptyMockDefs(t)
	t.Cleanup(deleteMockDefs)
}

func mockCommandDefinition(cmd string, args []string) *CommandDefinition {
	cd := CommandDefinition{
		CommonDefinition: CommonDefinition{Token: cmd, Abbr: cmd},
		Arguments:        args,
	}
	mockExample(&cd, args)
	return &cd
}

func mockCommandDefinitions(commands []string) []CommandDefinition {
	comDefs := make([]CommandDefinition, 0)
	for _, c := range commands {
		comDefs = append(comDefs, *mockCommandDefinition(c, commands))
	}
	return comDefs
}

func mockExample(cd *CommandDefinition, tokens []string) {
	desc := strings.Join(tokens, "")
	cd.Examples = append(cd.Examples, ExampleDefinition{
		ArgumentsValues: map[string][]string{},
		Desc:            desc,
	})
	cd.Examples[0].ArgumentsValues = make(map[string][]string, 0)
	for _, tt := range tokens {
		if strings.HasPrefix(desc, "empty") {
			continue
		}
		cd.Examples[0].ArgumentsValues[tt] = tokens
	}
}

func mockArgumentDefinition(arg string, values []string) *ArgumentDefinition {
	ad := ArgumentDefinition{
		CommonDefinition: CommonDefinition{Token: arg, Abbr: arg},
		Values:           values,
	}
	return &ad
}

func mockArgumentDefinitions(args []string) []ArgumentDefinition {
	argDefs := make([]ArgumentDefinition, 0)
	for _, a := range args {
		argDefs = append(argDefs, *mockArgumentDefinition(a, args))
	}
	return argDefs
}

func mockFlagDefinitions(flags []string) []FlagDefinition {
	flagDefs := make([]FlagDefinition, 0)
	for _, f := range flags {
		fd := FlagDefinition{
			CommonDefinition: CommonDefinition{Token: f, Abbr: f},
		}
		flagDefs = append(flagDefs, fd)
	}
	return flagDefs
}

func mockCommandByToken(token string) *CommandDefinition {
	switch token {
	case "":
		return nil
	default:
		return mockCommandDefinition(token, nil)
	}
}

func mockCommandByAbbr(token string) *CommandDefinition {
	return mockCommandByToken(token)
}

func mockArgByToken(token string) *ArgumentDefinition {
	if strings.HasPrefix(token, "default") {
		// default arguments
		arg := mockArgumentDefinition(token, nil)
		arg.Default = true
		return arg
	}

	switch token {
	case "":
		return nil
	default:
		return mockArgumentDefinition(token, nil)
	}
}

func mockArgByAbbr(token string) *ArgumentDefinition {
	return mockArgByToken(token)
}

func mockValidArgVal(arg, val string) bool {
	if strings.HasPrefix(val, "invalid") {
		return false
	}
	return true
}

func mockParseCtx(cmd, arg string) *parseCtx {
	return &parseCtx{
		Command:  mockCommandDefinition(cmd, nil),
		Argument: mockArgumentDefinition(arg, nil),
	}
}
