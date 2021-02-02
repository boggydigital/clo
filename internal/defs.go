package internal

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (defs *Definitions) defaultArgValues(req *Request) error {
	if req == nil {
		return errors.New("cannot fill default argument values for a nil request")
	}
	if req.Command == "" {
		return errors.New("cannot fill default argument values for a request without a command")
	}
	if req.Arguments == nil {
		req.Arguments = make(map[string][]string, 0)
	}

	dc := defs.definedCmd(req.Command)
	if dc == "" {
		return fmt.Errorf("unknown request command %s", req.Command)
	}

	for _, arg := range defs.Cmd[dc] {
		a, values := splitArgValues(arg)
		ta := trimAttrs(a)

		// check if request already has some values specified for that argument
		if rv, ok := req.Arguments[ta]; ok {
			if len(rv) > 0 {
				continue
			}
		}

		for _, v := range values {
			if !isDefault(v) {
				continue
			}

			if req.Arguments[ta] == nil {
				req.Arguments[ta] = make([]string, 0)
			}

			req.Arguments[ta] = append(req.Arguments[ta], trimAttrs(v))
		}
	}

	return nil
}
