package internal

import (
	"fmt"
)

// Parse converts args to a structured Request or returns an error if there are unexpected values,
// order or if any of the defined constraints are not met: fixed values, required,
// multiple values, etc.
func (def *Definitions) Parse(args []string) (*Request, error) {

	// TODO: consider adding built-in help command if one has not been provided

	var req = Request{
		Flags:     []string{},
		Command:   "",
		Arguments: make(map[string][]string),
	}

	var expected = first()
	var ctx parseCtx

	for _, arg := range args {
		if arg == "" {
			continue
		}
		matched := false
		for _, tt := range expected {
			success, err := match(arg, tt, &ctx, def)
			if err != nil {
				return &req, err
			}
			if success {
				matched = true
				expandedArg, err := expandAbbr(arg, tt, def)
				if err != nil {
					return nil, err
				}
				err = req.update(expandedArg, tt, &ctx)
				if err != nil {
					return nil, err
				}
				ctx.update(arg, tt, def)
				expected = next(tt)
				break
			}
		}
		if !matched {
			return nil, fmt.Errorf("unknown argument: '%v'", arg)
		}
	}

	// read arguments that are specified as supporting env
	// if the value has not been provided as a CLI flag
	err := req.readEnvArgs(def)
	if err != nil {
		return &req, err
	}

	err = req.verify(def)
	if err != nil {
		return &req, err
	}

	return &req, nil
}
