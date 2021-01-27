package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

const defaultMockFilename = "clo.json"

func mockDefinitions() *Definitions {
	return &Definitions{
		Version: 1,
		Cmd:     map[string][]string{},
	}
}

//func mockAddHelpCommand(defs *Definitions) *Definitions {
//	defs.Commands = append(defs.Commands, CommandDefinition{
//		CommonDefinition: CommonDefinition{
//			Token: "help",
//		}})
//	return defs
//}
//
//func mockAddHelpCommandArgument(defs *Definitions) *Definitions {
//	defs.Arguments = append(defs.Arguments, ArgumentDefinition{
//		CommonDefinition: CommonDefinition{
//			Token: "help:command",
//		},
//		Values: []string{"from:nowhere"},
//	})
//
//	return defs
//}

//func mockAddArgumentThatDoesntExist(defs *Definitions) *Definitions {
//	defs.Commands[0].Arguments = append(defs.Commands[0].Arguments, "argument-that-doesnt-exist")
//	return defs
//}

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

func writeDefaultMockDefs(t *testing.T) {
	writeMockDefs(mockDefinitions(), t)
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
	//mockAddHelpCommandArgument(defs)
	writeMockDefs(defs, t)
	t.Cleanup(deleteMockDefs)
}

func setupEmptyMockDefs(t *testing.T) {
	writeEmptyMockDefs(t)
	t.Cleanup(deleteMockDefs)
}

//func mockCommandDefinition(cmd string, args []string) *CommandDefinition {
//	cd := CommandDefinition{
//		CommonDefinition: CommonDefinition{Token: cmd},
//		Arguments:        args,
//	}
//	return &cd
//}

//func mockCommandDefinitions(commands []string) []CommandDefinition {
//	comDefs := make([]CommandDefinition, 0)
//	for _, c := range commands {
//		comDefs = append(comDefs, *mockCommandDefinition(c, commands))
//	}
//	return comDefs
//}

//func mockArgumentDefinition(arg string, values []string) *ArgumentDefinition {
//	ad := ArgumentDefinition{
//		CommonDefinition: CommonDefinition{Token: arg},
//		Values:           values,
//	}
//	return &ad
//}

//func mockArgumentDefinitions(args []string) []ArgumentDefinition {
//	argDefs := make([]ArgumentDefinition, 0)
//	for _, a := range args {
//		argDefs = append(argDefs, *mockArgumentDefinition(a, args))
//	}
//	return argDefs
//}

//func mockCommandByToken(token string) *CommandDefinition {
//	switch token {
//	case "":
//		return nil
//	default:
//		return mockCommandDefinition(token, nil)
//	}
//}

//func mockCommandByTokenNoHelp(token string) *CommandDefinition {
//	if token == "help" {
//		return nil
//	} else {
//		return mockCommandByToken(token)
//	}
//}
//
//func mockCommandByAbbr(token string) *CommandDefinition {
//	return mockCommandByToken(token)
//}

//func mockArgByToken(token string) *ArgumentDefinition {
//	switch token {
//	case "":
//		return nil
//	default:
//		return mockArgumentDefinition(token, nil)
//	}
//}
//
//func mockArgByTokenNoHelp(token string) *ArgumentDefinition {
//	if token == "help" {
//		return nil
//	} else {
//		return mockArgByToken(token)
//	}
//}
//
//func mockArgByAbbr(token string) *ArgumentDefinition {
//	return mockArgByToken(token)
//}
