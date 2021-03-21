package clo

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Definitions struct {
	Version int                 `json:"version"`
	Cmd     map[string][]string `json:"cmd"`
	Help    map[string]string   `json:"help"`
}

func Load(reader io.Reader) (*Definitions, error) {
	var defs *Definitions
	if e := json.NewDecoder(reader).Decode(&defs); e != nil {
		return nil, e
	}

	addInternalHelpCmd(defs)

	return defs, nil
}

func (defs *Definitions) definedCmd(c string) (string, error) {
	if defs == nil {
		return "", fmt.Errorf("clo: no defined command for nil definitions")
	}
	definedToken := ""

	for cmd := range defs.Cmd {
		if strings.HasPrefix(cmd, c) {
			if definedToken != "" {
				return "", fmt.Errorf("clo: ambiguous command %s that could be %s or %s", c, definedToken, cmd)
			}
			definedToken = cmd
		}
	}

	return definedToken, nil
}

func (defs *Definitions) definedArg(c, a string) (string, error) {
	if defs == nil {
		return "", fmt.Errorf("clo: no defined command argument for nil defintions")
	}

	cmd, err := defs.definedCmd(c)
	if err != nil {
		return "", err
	}
	if cmd == "" {
		return "", nil
	}

	definedArg := ""

	for _, arg := range defs.Cmd[cmd] {
		if strings.HasPrefix(arg, a) {
			if definedArg != "" {
				return arg, fmt.Errorf("clo: ambiguous argument %s that could be %s or %s", a, definedArg, arg)
			}
			definedArg = arg
		}
	}

	return definedArg, nil
}

func (defs *Definitions) definedVal(c, a, v string) (string, error) {
	if defs == nil {
		return "", fmt.Errorf("clo: no defined command argument value for nil definitions")
	}

	arg, err := defs.definedArg(c, a)
	if err != nil {
		return "", err
	}
	if arg == "" {
		return "", nil
	}

	definedValue := ""
	var values []string

	// splitArgValues
	if hasArgValues(arg) {
		_, values = splitArgValues(arg)
		for _, val := range values {
			if strings.HasPrefix(val, v) {
				if definedValue != "" {
					return val, fmt.Errorf("clo: ambiguous value %s that could be %s or %s", v, definedValue, val)
				}
				definedValue = val
			}
		}
	}

	return v, nil
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

func (defs *Definitions) defaultArgument(cmd string) (string, error) {
	if defs == nil {
		return "", fmt.Errorf("clo: no default argument for nil defintions")
	}

	dc, err := defs.definedCmd(cmd)
	if err != nil {
		return dc, err
	}
	if dc == "" {
		return "", nil
	}

	for _, arg := range defs.Cmd[dc] {
		if isDefault(arg) {
			return arg, nil
		}
	}
	return "", nil
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
		return fmt.Errorf("clo: can't fill default argument values for a nil request")
	}
	if req.Command == "" {
		// return if no command has been specified, nothing to fill
		return nil
	}
	if req.Arguments == nil {
		req.Arguments = make(map[string][]string, 0)
	}

	dc, err := defs.definedCmd(req.Command)
	if err != nil {
		return err
	}
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
