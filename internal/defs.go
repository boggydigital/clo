package internal

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

type Definitions struct {
	Version int                 `json:"version"`
	Cmd     map[string][]string `json:"cmd"`
	Help    map[string]string   `json:"help"`
}

func LoadDefault(path string) (*Definitions, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defs, err := Load(file)
	if err != nil {
		return defs, err
	}

	return defs, nil
}

func Load(reader io.Reader) (*Definitions, error) {
	var defs *Definitions
	if e := json.NewDecoder(reader).Decode(&defs); e != nil {
		return nil, e
	}

	addInternalHelpCmd(defs)

	return defs, nil
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

func transform(arr []string, f func(string) string) []string {
	marr := make([]string, 0, len(arr))
	for _, s := range arr {
		marr = append(marr, f(s))
	}
	return marr
}

func (defs *Definitions) help(tokens []string) string {
	if defs == nil || defs.Help == nil {
		return ""
	}
	for len(tokens) > 0 {
		key := strings.Join(transform(tokens, trimAttrs), ":")
		if value, ok := defs.Help[key]; ok {
			return value
		}
		tokens = tokens[1:]
	}
	return ""
}
