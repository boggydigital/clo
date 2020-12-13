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
	defaultValue
	fixedValue
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
	case defaultValue:
		return "defaultValue"
	case fixedValue:
		return "fixedValue"
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
		return []int{argumentAbbr, argument, defaultValue, flagAbbr, flag}
	case commandAbbr:
		return []int{argumentAbbr, argument, defaultValue, flagAbbr, flag}
	case defaultValue:
		return []int{defaultValue, argumentAbbr, argument, flagAbbr, flag}
	case argument:
		return []int{fixedValue, argumentAbbr, argument, flagAbbr, flag, value}
	case argumentAbbr:
		return []int{fixedValue, argumentAbbr, argument, flagAbbr, flag, value}
	case fixedValue:
		return []int{fixedValue, argumentAbbr, argument, flagAbbr, flag}
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

func contains(item string, slice []string) bool {
	if slice == nil {
		return false
	}
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func matchesValue(token string, tokenType int, arg *ArgumentDefinition) (bool, error) {
	if arg == nil {
		return false, errors.New("can't confirm a match for a value that is missing an argument context")
	}
	switch tokenType {
	case fixedValue:
		if contains(token, arg.Values) {
			return true, nil
		} else {
			if len(arg.Values) > 0 {
				// check for sequence break (another argument) before verifying a fixed value
				if hasDDashPrefix(token) || hasSDashPrefix(token) {
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
		if !hasDDashPrefix(token) {
			return false, nil
		}
		return def.FlagByToken(trimDDash(token)) != nil, nil
	case flagAbbr:
		if !hasSDashPrefix(token) {
			return false, nil
		}
		if def.FlagByAbbr(trimSDash(token)) != nil {
			return true, nil
		}
		return false, nil
	case argument:
		if !hasDDashPrefix(token) {
			return false, nil
		}
		ad := def.ArgByToken(trimDDash(token))
		if ad == nil {
			return false, nil
		}
		if contains(ad.Token, ctx.Command.Arguments) {
			return true, nil
		}
		return false, errors.New(
			fmt.Sprintf("argument '%v' is not supported by command '%v'", ad.Token, ctx.Command.Token))
	case argumentAbbr:
		if !hasSDashPrefix(token) {
			return false, nil
		}
		ad := def.ArgByAbbr(trimSDash(token))
		if ad == nil {
			return false, nil
		}
		if contains(ad.Token, ctx.Command.Arguments) {
			return true, nil
		}
		return false, errors.New(
			fmt.Sprintf("argument '%v' is not supported by command '%v'", ad.Token, ctx.Command.Token))
	case fixedValue:
		fallthrough
	case value:
		return matchesValue(token, tokenType, ctx.Argument)
	default:
		return false, errors.New(
			fmt.Sprintf("cannot confirm match for a token '%v' of type '%v'", token, tokenStr(tokenType)))
	}
}

func hasDDashPrefix(token string) bool {
	return strings.HasPrefix(token, "--")
}

func hasSDashPrefix(token string) bool {
	return strings.HasPrefix(token, "-") &&
		!strings.HasPrefix(token, "--")
}

func trimDDash(token string) string {
	return strings.TrimPrefix(token, "--")
}

func trimSDash(token string) string {
	return strings.TrimPrefix(token, "-")
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
		req.Flags = append(req.Flags, trimSDash(expandedToken))
	case flag:
		req.Flags = append(req.Flags, trimDDash(expandedToken))
	case argument:
		req.Arguments[trimDDash(expandedToken)] = []string{}
	case argumentAbbr:
		req.Arguments[trimSDash(expandedToken)] = []string{}
	case value:
		fallthrough
	case fixedValue:
		arg := ctx.Argument.Token
		req.Arguments[arg] = append(req.Arguments[arg], expandedToken)
	default:
		return errors.New(
			fmt.Sprintf("cannot update request for a token '%v' of type '%v'", expandedToken, tokenStr(tokenType)))
	}
	return nil
}

func updateContext(ctx *context, token string, tokenType int, def *Definitions) {
	switch tokenType {
	case command:
		ctx.Command = def.CommandByToken(token)
	case commandAbbr:
		ctx.Command = def.CommandByAbbr(token)
	case argument:
		ctx.Argument = def.ArgByToken(trimDDash(token))
	case argumentAbbr:
		ctx.Argument = def.ArgByAbbr(trimSDash(token))
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
	case defaultValue:
		fallthrough
	case fixedValue:
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
		ad := def.ArgByAbbr(trimSDash(token))
		if ad == nil {
			return "", errors.New(fmt.Sprintf("unknown argument abbreviation: '%v'", token))
		}
		return ad.Token, nil
	case flagAbbr:
		fd := def.FlagByAbbr(trimSDash(token))
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
		//fmt.Println("arg:", arg)
		matched := false
		for _, tt := range expected {
			//fmt.Println("tt:", tokenStr(tt))
			match, err := matches(arg, tt, &ctx, def)
			if err != nil {
				return &req, err
			}
			if match {
				matched = true
				//fmt.Println("match")
				expandedToken, err := expandAbbr(arg, tt, def)
				if err != nil {
					return nil, err
				}
				err = updateRequest(&req, expandedToken, tt, &ctx)
				if err != nil {
					return nil, err
				}
				updateContext(&ctx, arg, tt, def)
				expected = next(tt)
				break
			}
			//fmt.Println("no match")
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
