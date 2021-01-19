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
				return req, err
			}
			if success {
				matched = true
				expandedArg := expandAbbr(arg, tt, def)
				// it's ok to ignore the error below, since we'd only return
				// error in two cases:
				// 1) req.Command is already set - this shouldn't be
				// possible in this flow since after matching command
				// token or abbreviation we would progress to another type
				// 2) if tokenType is an unsupported value, however this is
				// not possible in this flow given the next() function
				_ = req.update(expandedArg, tt, &ctx)
				ctx.update(expandedArg, tt, def)
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
