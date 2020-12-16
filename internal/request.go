package internal

import (
	"errors"
	"fmt"
)

type Request struct {
	Flags     []string
	Command   string
	Arguments map[string][]string
}

func (req *Request) update(expandedToken string, tokenType int, ctx *parseCtx) error {

	switch tokenType {
	case commandAbbr:
		fallthrough
	case command:
		if req.Command != "" {
			return errors.New("request already has a command specified")
		}
		req.Command = expandedToken
		break
	case flagAbbr:
		req.Flags = append(req.Flags, trimPrefix(expandedToken))
	case flag:
		req.Flags = append(req.Flags, trimPrefix(expandedToken))
	case argument:
		fallthrough
	case argumentAbbr:
		arg := trimPrefix(expandedToken)
		if req.Arguments[arg] == nil {
			req.Arguments[arg] = []string{}
		}
	case valueDefault:
		fallthrough
	case value:
		fallthrough
	case valueFixed:
		argCtx := ctx.Argument.Token
		req.Arguments[argCtx] = append(req.Arguments[argCtx], expandedToken)
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
		return errors.New("cannot verify nil request for required arguments")
	}

	requiredArgs := def.RequiredArgs(req.Command)
	for _, ra := range requiredArgs {
		matched := false
		for at, avs := range req.Arguments {
			if ra == at {
				if len(avs) == 0 {
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

	for at, avs := range req.Arguments {
		if at == "" {
			continue
		}
		arg := def.ArgByToken(at)
		if arg == nil {
			continue
		}
		if !arg.Multiple && len(avs) > 1 {
			return fmt.Errorf("argument '%v' has multiple values, supports no more than one", at)
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
