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

func (req *Request) setDefaultContext(tokenType int, def *Definitions) error {
	switch tokenType {
	case argument:
		if req.Command == "" {
			dc := def.defaultCommand()
			if dc != "" {
				return req.update(trimAttrs(dc), command)
			}
		}
	case value:
		if req.lastArgument() == "" {
			da := def.defaultArgument(req.Command)
			if da != "" {
				return req.update(trimAttrs(da), argument)
			}
		}
	}
	return nil
}

func (req *Request) update(token string, tokenType int) error {
	switch tokenType {
	case command:
		if req.Command != "" {
			return errors.New("request already has a command specified")
		}
		req.Command = token
		break
	case argument:
		arg := trimArgPrefix(token)
		if req.Arguments == nil {
			req.Arguments = map[string][]string{}
		}
		if req.Arguments[arg] == nil {
			req.Arguments[arg] = []string{}
		}
	case value:
		lastArg := req.lastArgument()
		if lastArg == "" {
			return fmt.Errorf("cannot update value for a request with no arguments")
		}
		req.Arguments[lastArg] = append(req.Arguments[lastArg], token)
	default:
		return fmt.Errorf(
			"cannot update request for a token '%v' of type '%v'",
			token,
			tokenString(tokenType))
	}
	return nil
}

func (req *Request) commandHasRequiredArgs(def *Definitions) error {
	if def == nil {
		return errors.New("cannot verify required arguments using nil definitions")
	}
	if req == nil {
		return errors.New("cannot verify required arguments using nil request")
	}

	dc := def.definedCmd(req.Command)
	if dc == "" {
		return fmt.Errorf("cannot verify required arguments without a command")
	}

	for _, arg := range def.Cmd[dc] {
		if !isRequired(arg) {
			continue
		}

		tArg := trimAttrs(arg)
		if _, ok := req.Arguments[tArg]; !ok {
			return fmt.Errorf("required argument '%v' is missing for the command '%v'", tArg, req.Command)
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

		_, da := def.definedCmdArg(req.Command, arg)
		if da == "" {
			continue
		}

		if !isMultiple(da) && len(values) > 1 {
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
