package internal

import (
	"fmt"
)

// Parse converts args to a structured Request or returns an error if there are unexpected values,
// order or if any of the defined constraints are not met: fixed values, required,
// multiple values, etc.
func (def *Definitions) Parse(args []string) (*Request, error) {

	if def == nil {
		return nil, fmt.Errorf("cannot parse using nil definitions")
	}

	var req = &Request{
		Command:   "",
		Arguments: make(map[string][]string),
	}

	var expected = initial()

	for _, arg := range args {
		if arg == "" {
			continue
		}
		matched := false
		for _, tt := range expected {
			// set default context for certain token types
			// based on defaults defined in clo.json
			err := req.setDefaultContext(tt, def)
			if err != nil {
				return req, err
			}
			//
			definedToken, err := match(arg, tt, req, def)
			if err != nil {
				return req, err
			}

			if definedToken != "" {
				matched = true
				err = req.update(trimAttrs(definedToken), tt)
				if err != nil {
					return req, err
				}
				expected = next(tt)
				break
			}
		}
		if !matched {
			return nil, fmt.Errorf("unknown argument: '%v'", arg)
		}
	}

	// read arguments that are specified as supporting env
	// if the value has not been provided as a CLI arg.
	// Safely ignoring error here as well, since the only condition
	// that would lead to an error is a nil definitions,
	// and we've already tested that above
	_ = req.readEnvArgs(def)

	if err := req.verify(def); err != nil {
		return req, err
	}

	return req, nil
}
