package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Definitions struct {
	Version        int                  `json:"version"`
	EnvPrefix      string               `json:"env-prefix,omitempty"`
	App            string               `json:"app,omitempty"`
	Help           string               `json:"help,omitempty"`
	DefaultCommand string               `json:"default_command,omitempty"`
	Commands       []CommandDefinition  `json:"commands,omitempty"`
	Arguments      []ArgumentDefinition `json:"arguments,omitempty"`
}

func LoadDefault() (*Definitions, error) {
	return Load("clo.json")
}

func Load(path string) (*Definitions, error) {

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var dfs *Definitions

	if err := json.Unmarshal(bytes, &dfs); err != nil {
		return nil, err
	}

	if err := tryAddHelpCommand(dfs); err != nil {
		// adding help is not considered fatal error, inform, continue
		log.Println("error adding help command:", err.Error())
	}

	for i, _ := range dfs.Commands {
		e := dfs.Commands[i].setDefaultRequired()
		if e != nil {
			return dfs, e
		}
	}

	if e := expandRefValues(dfs.Arguments, dfs.Commands); e != nil {
		return nil, e
	}

	return dfs, nil
}

func (def *Definitions) CommandByToken(token string) *CommandDefinition {
	for _, c := range def.Commands {
		if c.Token == token {
			return &c
		}
	}
	return nil
}

func (def *Definitions) ArgByToken(token string) *ArgumentDefinition {
	for _, a := range def.Arguments {
		if a.Token == token {
			return &a
		}
	}
	return nil
}

func (def *Definitions) ValidArgVal(arg string, val string) bool {
	if arg == "" {
		return false
	}
	ad := def.ArgByToken(arg)
	if ad == nil {
		return false
	}
	return ad.ValidValue(val)
}
