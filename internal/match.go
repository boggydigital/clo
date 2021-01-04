package internal

import (
	"errors"
	"fmt"
)

// matchesArgument determines if an arg token matches argument (either in normal or
// abbreviated form) for a given command
func matchArgument(token string, tokenType int, cmdCtx *CommandDefinition, def *Definitions) (bool, error) {
	// argument is expected to be prefixed with a single or double dash, if it's not -
	// there is no reason to perform further match checks
	if !hasPrefix(token) {
		return false, nil
	}

	if def == nil {
		return false, fmt.Errorf("cannot match argument with nil definitions")
	}
	// first, try looking up argument by a token or an abbreviation,
	// given a token type
	var ad *ArgumentDefinition
	switch tokenType {
	case argument:
		ad = def.ArgByToken(trimPrefix(token))
	case argumentAbbr:
		ad = def.ArgByAbbr(trimPrefix(token))
	default:
		return false, fmt.Errorf("type '%v' cannot be used for argument matches", tokenString(tokenType))
	}

	// report no match if the token didn't match any of the arguments or argument abbreviations
	if ad == nil {
		return false, nil
	}

	if cmdCtx == nil {
		return false, fmt.Errorf("cannot validate argument in the context of nil command")
	}

	// however if the argument matched token, we need to check if it's one of the supported values
	// for a command context
	if cmdCtx.ValidArg(ad.Token) {
		return true, nil
	}

	// report if it's a valid argument, but is not supported for a command context
	return false, fmt.Errorf("argument '%v' is not supported by command '%v'", ad.Token, cmdCtx.Token)
}

// matchesFlag determines if a token matches flag (either in normal or abbreviated form)
func matchFlag(token string, tokenType int, def *Definitions) (bool, error) {
	// flag is expected to be prefixed with a single or double dash, if it's not -
	// there is no reason to perform further match checks
	if !hasPrefix(token) {
		return false, nil
	}

	if def == nil {
		return false, fmt.Errorf("cannot match flag with nil definitions")
	}
	// check if flag matches a token or abbreviation
	switch tokenType {
	case flag:
		return def.FlagByToken(trimPrefix(token)) != nil, nil
	case flagAbbr:
		return def.FlagByAbbr(trimPrefix(token)) != nil, nil
	default:
		return false, fmt.Errorf("type '%v' cannot be used for flag matches", tokenString(tokenType))
	}
}

func matchDefaultValue(token string, tokenType int, ctx *parseCtx, def *Definitions) (bool, error) {

	if tokenType != valueDefault {
		return false, nil
	}

	if ctx == nil {
		return false, fmt.Errorf("cannot match default value with nil context")
	}
	if ctx.Command == nil ||
		(ctx.Argument != nil && !ctx.Argument.Default) {
		return false, nil
	}

	if def == nil {
		return false, fmt.Errorf("cannot match default value with nil definitions")
	}
	defArg := def.DefaultArg(ctx.Command)

	m, err := matchValue(token, tokenType, defArg)

	if m && err == nil {
		ctx.Argument = defArg
	}

	return m, err
}

func matchValue(token string, tokenType int, arg *ArgumentDefinition) (bool, error) {
	if arg == nil {
		return false, errors.New("can't confirm a match for a value that is missing an argument context")
	}

	if hasPrefix(token) {
		return false, nil
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
		if arg.ValidValue(token) {
			return true, nil
		} else {
			if len(arg.Values) > 0 {
				// check for sequence break (another argument) before verifying a fixed value
				if hasPrefix(token) {
					return false, nil
				}
				return false, fmt.Errorf("unsupported value '%v' for an argument '%v'", token, arg.Token)
			} else {
				return false, nil
			}
		}
	case value:
		return len(arg.Values) == 0, nil
	}
	return false, nil
}

// match a token to a given token type, in the context of existing
// command, argument (applicable to tokens.argument and tokens.value, tokens.valueFixed, tokens.valueDefault)
func match(token string, tokenType int, ctx *parseCtx, def *Definitions) (bool, error) {
	switch tokenType {
	case command:
		return def.CommandByToken(token) != nil, nil
	case commandAbbr:
		return def.CommandByAbbr(token) != nil, nil
	case flag:
		fallthrough
	case flagAbbr:
		return matchFlag(token, tokenType, def)
	case argument:
		fallthrough
	case argumentAbbr:
		return matchArgument(token, tokenType, ctx.Command, def)
	case valueDefault:
		return matchDefaultValue(token, tokenType, ctx, def)
	case valueFixed:
		fallthrough
	case value:
		return matchValue(token, tokenType, ctx.Argument)
	default:
		return false, fmt.Errorf(
			"cannot confirm match for a token '%v' of type '%v'",
			token,
			tokenString(tokenType))
	}
}
