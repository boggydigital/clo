package clo

import (
	"fmt"
	"net/url"
	"strings"
)

// ParseRequest converts args to a structured request or returns an error if there are unexpected values,
// order or if any of the defined constraints are not met: fixed values, required,
// multiple values, etc.
func (defs *definitions) parseRequest(args []string) (*request, error) {

	if len(args) == 0 {
		return nil, nil
	}

	if defs == nil {
		return nil, fmt.Errorf("cannot parse using nil definitions")
	}

	var req = &request{
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
			err := req.setDefaultContext(tt, defs)
			if err != nil {
				return req, err
			}
			//
			definedToken, err := match(
				arg,
				tt,
				req.Command,
				req.lastArgument,
				defs)
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
	if err := req.readEnvArgs(defs); err != nil {
		return req, fmt.Errorf("failed to fill arguments from the environment")
	}

	// check is any arguments have default values that can be used
	// instead of leaving those arguments empty
	if err := defs.defaultArgValues(req); err != nil {
		return req, fmt.Errorf("failed to fill arguments default values")
	}

	// validate request using definition constraints
	if err := req.validate(defs); err != nil {
		return req, err
	}

	return req, nil
}

func (defs *definitions) parseUrl(args []string) (*url.URL, error) {
	req, err := defs.parseRequest(args)
	if err != nil {
		return nil, err
	}

	if req == nil {
		req = &request{Command: helpCmd}
	}

	u := &url.URL{
		Scheme: "cli",
		Host:   "clo",
		Path:   req.Command,
	}

	q := u.Query()
	for arg, values := range req.Arguments {
		val := ""
		if len(values) == 0 {
			val = "true"
		} else {
			val = strings.Join(values, ",")
		}
		q.Add(arg, val)
	}

	u.RawQuery = q.Encode()

	return u, nil
}
