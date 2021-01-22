package internal

import (
	"errors"
	"fmt"
)

type Request struct {
	Command   string
	Arguments map[string][]string
}

func (req *Request) hasArguments() bool {
	return req != nil && len(req.Arguments) > 0
}

func (req *Request) lastArgument() string {
	if req == nil {
		return ""
	}
	if len(req.Arguments) == 0 {
		return ""
	}
	keys := make([]string, 0, len(req.Arguments))
	for k := range req.Arguments {
		keys = append(keys, k)
	}
	return keys[len(keys)-1]
}

func (req *Request) update(expandedToken string, tokenType int) error {

	switch tokenType {
	case command:
		if req.Command != "" {
			return errors.New("request already has a command specified")
		}
		req.Command = expandedToken
		break
	case argument:
		arg := trimPrefix(expandedToken)
		if req.Arguments[arg] == nil {
			req.Arguments[arg] = []string{}
		}
	//case valueDefault:
	//	fallthrough
	case value:
		lastKey := req.lastArgument()
		if lastKey == "" {
			return fmt.Errorf("cannot update value for a request with no arguments")
		}
		req.Arguments[lastKey] = append(req.Arguments[lastKey], expandedToken)
	default:
		return fmt.Errorf(
			"cannot update request for a token '%v' of type '%v'",
			expandedToken,
			tokenString(tokenType))
	}
	return nil
}

func (req *Request) commandHasRequiredArgs(def *Definitions) error {
	if def == nil {
		return errors.New("cannot verify required argument using nil definitions")
	}
	if req == nil {
		return errors.New("cannot verify required argument using nil request")
	}

	// TODO: verify not nil
	cd := def.CommandByToken(req.Command)
	if cd == nil {
		return nil
	}

	for _, ra := range cd.requiredArguments {
		matched := false
		for arg, values := range req.Arguments {
			if ra == arg {
				if len(values) == 0 {
					return fmt.Errorf("required argument '%v' is missing values", ra)
				}
				matched = true
				break
			}
		}
		if !matched {
			return fmt.Errorf("required argument '%v' is missing for the command '%v'", ra, req.Command)
		}
	}
	return nil
}

func (req *Request) argumentsMultipleValues(def *Definitions) error {
	if def == nil {
		return errors.New("cannot verify required argument using nil definitions")
	}
	if req == nil {
		return errors.New("cannot verify nil request for required arguments")
	}

	for arg, values := range req.Arguments {
		if arg == "" {
			continue
		}
		ad := def.ArgByToken(arg)
		if ad == nil {
			continue
		}
		if !ad.Multiple && len(values) > 1 {
			return fmt.Errorf("argument '%v' has multiple values, supports no more than one", arg)
		}
	}

	return nil
}

func (req *Request) verify(def *Definitions) error {
	if def == nil {
		return errors.New("cannot verify required argument using nil definitions")
	}
	if req == nil {
		return errors.New("cannot verify nil request for required arguments")
	}

	err := req.commandHasRequiredArgs(def)
	if err != nil {
		return err
	}
	err = req.argumentsMultipleValues(def)
	if err != nil {
		return err
	}
	return nil
}

func (req *Request) ArgVal(arg string) string {
	if req == nil {
		return ""
	}
	vs := req.Arguments[arg]
	if len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (req *Request) ArgValues(arg string) []string {
	if req == nil {
		return []string{}
	}
	return req.Arguments[arg]
}

func (req *Request) Flag(arg string) bool {
	if req == nil {
		return false
	}
	_, ok := req.Arguments[arg]
	return ok
}
