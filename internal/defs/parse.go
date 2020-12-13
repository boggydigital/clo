package defs

import (
	"errors"
	"fmt"
	"strings"
)

const (
	command = iota
	commandAbbr
	argument
	argumentAbbr
	valueDefault
	valueFixed
	value
	flag
	flagAbbr
)

func tokenStr(tokenType int) string {
	switch tokenType {
	case command:
		return "command"
	case commandAbbr:
		return "commandAbbr"
	case argument:
		return "argument"
	case argumentAbbr:
		return "argumentAbbr"
	case valueDefault:
		return "valueDefault"
	case valueFixed:
		return "valueFixed"
	case value:
		return "value"
	case flag:
		return "flag"
	case flagAbbr:
		return "flagAbbr"
	}
	return "unknown"
}

// next defines token sequencing expectations and
// what types of tokens can follow a given type of token
func next(token int) []int {
	switch token {
	case command:
		return []int{argumentAbbr, argument, valueDefault, flagAbbr, flag}
	case commandAbbr:
		return []int{argumentAbbr, argument, valueDefault, flagAbbr, flag}
	case valueDefault:
		return []int{valueDefault, argumentAbbr, argument, flagAbbr, flag}
	case argument:
		return []int{valueFixed, argumentAbbr, argument, flagAbbr, flag, value}
	case argumentAbbr:
		return []int{valueFixed, argumentAbbr, argument, flagAbbr, flag, value}
	case valueFixed:
		return []int{valueFixed, argumentAbbr, argument, flagAbbr, flag}
	case value:
		return []int{argumentAbbr, argument, flagAbbr, flag, value}
	case flag:
		return []int{flag, flagAbbr}
	case flagAbbr:
		return []int{flag, flagAbbr}
	default:
		return []int{}
	}
}

type context struct {
	Command  *CommandDefinition
	Argument *ArgumentDefinition
}

func matchesArgument(token string, tokenType int, cmd *CommandDefinition, def *Definitions) (bool, error) {
	if !hasPrefix(token, tokenType) {
		return false, nil
	}
	var ad *ArgumentDefinition
	switch tokenType {
	case argument:
		ad = def.ArgByToken(trimPrefix(token, tokenType))
	case argumentAbbr:
		ad = def.ArgByAbbr(trimPrefix(token, tokenType))
	default:
		return false, errors.New(
			fmt.Sprintf("type '%v' cannot be used for argument matches", tokenStr(tokenType)))
	}

	if ad == nil {
		// TODO: test if this need to return error
		return false, nil
	}
	if cmd.ArgSupported(ad.Token) {
		return true, nil
	}
	return false, errors.New(
		fmt.Sprintf("argument '%v' is not supported by command '%v'", ad.Token, cmd.Token))
}

func matchesFlag(token string, tokenType int, def *Definitions) (bool, error) {
	if !hasPrefix(token, tokenType) {
		return false, nil
	}
	switch tokenType {
	case flag:
		return def.FlagByToken(trimPrefix(token, tokenType)) != nil, nil
	case flagAbbr:
		return def.FlagByAbbr(trimPrefix(token, tokenType)) != nil, nil
	default:
		return false, errors.New(
			fmt.Sprintf("type '%v' cannot be used for flag matches", tokenStr(tokenType)))
	}
}

func matchesValue(token string, tokenType int, arg *ArgumentDefinition) (bool, error) {
	if arg == nil {
		return false, errors.New("can't confirm a match for a value that is missing an argument context")
	}

	if tokenType == valueDefault {
		if !arg.Default {
			return false, nil
		}
		tokenType = value
		if len(arg.Values) > 0 {
			tokenType = valueFixed
		}
	}

	switch tokenType {
	case valueFixed:
		if arg.ValueSupported(token) {
			return true, nil
		} else {
			if len(arg.Values) > 0 {
				// check for sequence break (another argument) before verifying a fixed value
				if hasAnyPrefix(token) {
					return false, nil
				}
				return false, errors.New(fmt.Sprintf("unsupported value '%v' for an argument '%v'", token, arg.Token))
			} else {
				return false, nil
			}
		}
	case value:
		return len(arg.Values) == 0, nil
	}
	return false, nil
}

func defaultArgument(cmd string, def *Definitions) (*ArgumentDefinition, error) {
	return nil, nil
}

func matches(token string, tokenType int, ctx *context, def *Definitions) (bool, error) {
	switch tokenType {
	case command:
		return def.CommandByToken(token) != nil, nil
	case commandAbbr:
		return def.CommandByAbbr(token) != nil, nil
	case flag:
		fallthrough
	case flagAbbr:
		return matchesFlag(token, tokenType, def)
	case argument:
		fallthrough
	case argumentAbbr:
		return matchesArgument(token, tokenType, ctx.Command, def)
	case valueDefault:
		if ctx.Command == nil ||
			ctx.Argument != nil {
			return false, nil
		}
		ctx.Argument = def.DefaultArg(ctx.Command)
		fallthrough
	case valueFixed:
		fallthrough
	case value:
		return matchesValue(token, tokenType, ctx.Argument)
	default:
		return false, errors.New(
			fmt.Sprintf("cannot confirm match for a token '%v' of type '%v'", token, tokenStr(tokenType)))
	}
}

func hasAnyPrefix(token string) bool {
	return strings.HasPrefix(token, "-") ||
		strings.HasPrefix(token, "--")
}

func hasPrefix(token string, tokenType int) bool {
	prefix := ""
	switch tokenType {
	case command:
		fallthrough
	case commandAbbr:
		fallthrough
	case value:
		fallthrough
	case valueFixed:
		fallthrough
	case valueDefault:
		return false
	case flagAbbr:
		fallthrough
	case argumentAbbr:
		prefix = "-"
	case flag:
		fallthrough
	case argument:
		prefix = "--"
	}
	return strings.HasPrefix(token, prefix)
}

func trimPrefix(token string, tokenType int) string {
	prefix := ""
	switch tokenType {
	case command:
		fallthrough
	case commandAbbr:
		fallthrough
	case value:
		fallthrough
	case valueFixed:
		fallthrough
	case valueDefault:
		return token
	case flagAbbr:
		fallthrough
	case argumentAbbr:
		prefix = "-"
	case flag:
		fallthrough
	case argument:
		prefix = "--"
	}
	return strings.TrimPrefix(token, prefix)
}

func updateRequest(req *Request, expandedToken string, tokenType int, ctx *context) error {

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
		req.Flags = append(req.Flags, trimPrefix(expandedToken, tokenType))
	case flag:
		req.Flags = append(req.Flags, trimPrefix(expandedToken, tokenType))
	case argument:
		fallthrough
	case argumentAbbr:
		arg := trimPrefix(expandedToken, tokenType)
		if req.Arguments[arg] == nil {
			req.Arguments[arg] = []string{}
		}
	case valueDefault:
		fallthrough
	case value:
		fallthrough
	case valueFixed:
		arg := ctx.Argument.Token
		req.Arguments[arg] = append(req.Arguments[arg], expandedToken)
	default:
		return errors.New(
			fmt.Sprintf("cannot update request for a token '%v' of type '%v'", expandedToken, tokenStr(tokenType)))
	}
	return nil
}

func setContext(token string, tokenType int, ctx *context, def *Definitions) {
	switch tokenType {
	case command:
		ctx.Command = def.CommandByToken(token)
	case commandAbbr:
		ctx.Command = def.CommandByAbbr(token)
	case argument:
		ctx.Argument = def.ArgByToken(trimPrefix(token, tokenType))
	case argumentAbbr:
		ctx.Argument = def.ArgByAbbr(trimPrefix(token, tokenType))
	}
}

func expandAbbr(token string, tokenType int, def *Definitions) (string, error) {
	switch tokenType {
	case command:
		fallthrough
	case argument:
		fallthrough
	case flag:
		fallthrough
	case valueDefault:
		fallthrough
	case valueFixed:
		fallthrough
	case value:
		return token, nil
	case commandAbbr:
		cd := def.CommandByAbbr(token)
		if cd == nil {
			return "", errors.New(fmt.Sprintf("unknown command abbreviation: '%v'", token))
		}
		return cd.Token, nil
	case argumentAbbr:
		ad := def.ArgByAbbr(trimPrefix(token, tokenType))
		if ad == nil {
			return "", errors.New(fmt.Sprintf("unknown argument abbreviation: '%v'", token))
		}
		return ad.Token, nil
	case flagAbbr:
		fd := def.FlagByAbbr(trimPrefix(token, tokenType))
		if fd == nil {
			return "", errors.New(fmt.Sprintf("unknown flag abbreviation: '%v'", token))
		}
		return fd.Token, nil
	default:
		return "", errors.New(
			fmt.Sprintf("cannot expand token '%v' of unknown type '%v'", token, tokenStr(tokenType)))
	}
}

func verifyCommandHasRequiredArgs(req *Request, def *Definitions) error {
	return nil
}

func verifyConstraints(req *Request, def *Definitions) error {
	err := verifyCommandHasRequiredArgs(req, def)
	if err != nil {
		return err
	}
	return nil
}

// Parse converts args to a structured Request or returns an error if there are unexpected values,
// order or if any of the defined constraints are not met: fixed values, required,
// multiple values, etc.
func (def *Definitions) Parse(args []string) (*Request, error) {

	var req = Request{
		Flags:     []string{},
		Command:   "",
		Arguments: make(map[string][]string),
	}

	var expected = []int{command, commandAbbr}
	var ctx context

	for _, arg := range args {
		if arg == "" {
			continue
		}
		arg = strings.ToLower(arg)
		fmt.Println("----- arg:", arg)
		matched := false
		for _, tt := range expected {
			fmt.Println("tt:", tokenStr(tt))
			match, err := matches(arg, tt, &ctx, def)
			if err != nil {
				return &req, err
			}
			if match {
				matched = true
				fmt.Println("match")
				expandedArg, err := expandAbbr(arg, tt, def)
				if err != nil {
					return nil, err
				}
				err = updateRequest(&req, expandedArg, tt, &ctx)
				if err != nil {
					return nil, err
				}
				setContext(arg, tt, &ctx, def)
				expected = next(tt)
				break
			}
			fmt.Println("no match")
		}
		if !matched {
			return nil, errors.New(fmt.Sprintf("unknown argument: '%v'", arg))
		}
	}

	err := verifyConstraints(&req, def)
	if err != nil {
		return &req, err
	}

	return &req, nil
}
