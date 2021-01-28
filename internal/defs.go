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

func (defs *Definitions) definedCmd(c string) string {
	if defs == nil {
		return ""
	}

	for cmd := range defs.Cmd {
		if strings.HasPrefix(cmd, c) {
			return cmd
		}
	}

	return ""
}

func (defs *Definitions) definedCmdArg(c, a string) (string, string) {
	if defs == nil {
		return "", ""
	}

	cmd := defs.definedCmd(c)
	if cmd == "" {
		return cmd, ""
	}

	for _, arg := range defs.Cmd[cmd] {
		if strings.HasPrefix(arg, a) {
			return cmd, arg
		}
	}

	return cmd, ""
}

func (defs *Definitions) definedCmdArgVal(c, a, v string) (string, string, string) {
	if defs == nil {
		return "", "", ""
	}

	cmd, arg := defs.definedCmdArg(c, a)
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

func (defs *Definitions) defaultCommand() string {
	if defs == nil {
		return ""
	}
	for c := range defs.Cmd {
		if isDefault(c) {
			return c
		}
	}
	return ""
}

func (defs *Definitions) defaultArgument(cmd string) string {
	if defs == nil {
		return ""
	}

	dc := defs.definedCmd(cmd)
	if dc == "" {
		return ""
	}

	for _, arg := range defs.Cmd[dc] {
		if isDefault(arg) {
			return arg
		}
	}
	return ""
}
