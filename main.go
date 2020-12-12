package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CommonDefinition struct {
	Token string `json:"token"`
	Abbr  string `json:"abbr,omitempty"`
	Hint  string `json:"hint,omitempty"`
	Desc  string `json:"desc,omitempty"`
}

type FlagDefinition struct {
	CommonDefinition
}

type ExampleDefinition struct {
	Arguments []string `json:"arguments"`
	Values    []string `json:"values"`
	Desc      string   `json:"desc,omitempty"`
}

type CommandDefinition struct {
	CommonDefinition
	Arguments []string            `json:"arguments,omitempty"`
	Examples  []ExampleDefinition `json:"examples,omitempty"`
}

type ArgumentDefinition struct {
	CommonDefinition
	Default  bool     `json:"default,omitempty"`
	Multiple bool     `json:"multiple,omitempty"`
	Required bool     `json:"required,omitempty"`
	Values   []string `json:"values,omitempty"`
}

type ValueDefinition struct {
	CommonDefinition
	Hidden bool `json:"hidden,omitempty"`
}

type Definitions struct {
	Version   int                  `json:"version"`
	Flags     []FlagDefinition     `json:"flags,omitempty"`
	Commands  []CommandDefinition  `json:"commands,omitempty"`
	Arguments []ArgumentDefinition `json:"arguments,omitempty"`
	Values    []ValueDefinition    `json:"values,omitempty"`
}

func main() {

	filename := "definitions.json"

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var defs Definitions

	err = json.Unmarshal(bytes, &defs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(defs.Version)
	fmt.Println(defs.Flags)
}
