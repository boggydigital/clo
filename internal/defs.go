package internal

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Definitions struct {
	Version int                 `json:"version"`
	Cmd     map[string][]string `json:"cmd"`
	Help    map[string]string   `json:"help"`
}

func LoadDefault() (*Definitions, error) {
	return Load("clo.json")
}

func Load(path string) (*Definitions, error) {

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var def *Definitions

	if e := json.Unmarshal(bytes, &def); e != nil {
		return nil, e
	}

	addHelpCmd(def)

	return def, nil
}

func (def *Definitions) definedCmd(c string) string {
	if def == nil {
		return ""
	}

	for cmd := range def.Cmd {
		if strings.HasPrefix(cmd, c) {
			return cmd
		}
	}

	return ""
}

func (def *Definitions) definedCmdArg(c, a string) (string, string) {
	if def == nil {
		return "", ""
	}

	cmd := def.definedCmd(c)
	if cmd == "" {
		return cmd, ""
	}

	for _, arg := range def.Cmd[cmd] {
		if strings.HasPrefix(arg, a) {
			return cmd, arg
		}
	}

	return cmd, ""
}

func (def *Definitions) definedCmdArgVal(c, a, v string) (string, string, string) {
	if def == nil {
		return "", "", ""
	}

	cmd, arg := def.definedCmdArg(c, a)
	if arg == "" {
		return cmd, arg, ""
	}

	// splitArgValues
	if hasArgValues(arg) {
		asv, values := splitArgValues(arg)
		for _, val := range values {
			if strings.HasPrefix(val, v) {
				return cmd, asv, val
			}
		}
	}

	return cmd, arg, v
}

func (def *Definitions) defaultCommand() string {
	if def == nil {
		return ""
	}
	for c := range def.Cmd {
		if isDefault(c) {
			return c
		}
	}
	return ""
}

func (def *Definitions) defaultArgument(cmd string) string {
	if def == nil {
		return ""
	}

	dc := def.definedCmd(cmd)
	if dc == "" {
		return ""
	}

	for _, arg := range def.Cmd[dc] {
		if isDefault(arg) {
			return arg
		}
	}
	return ""
}
