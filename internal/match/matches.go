package match

import (
	"errors"
	"fmt"
	"github.com/boggydigital/clove/internal/defs"
	"github.com/boggydigital/clove/internal/parsectx"
	"github.com/boggydigital/clove/internal/strutil"
	"github.com/boggydigital/clove/internal/tokens"
)

// matchesArgument determines if an arg token matches argument (either in normal or
// abbreviated form) for a given command
func matchesArgument(token string, tokenType int, cmdCtx *defs.CommandDefinition, def *defs.Definitions) (bool, error) {
	// argument is expected to be prefixed with a single or double dash, if it's not -
	// there is no reason to perform further match checks
	if !strutil.HasExpectedPrefix(token, tokenType) {
		return false, nil
	}

	// first, try looking up argument by a token or an abbreviation,
	// given a token type
	var ad *defs.ArgumentDefinition
	switch tokenType {
	case tokens.Argument:
		ad = def.ArgByToken(strutil.TrimPrefix(token, tokenType))
	case tokens.ArgumentAbbr:
		ad = def.ArgByAbbr(strutil.TrimPrefix(token, tokenType))
	default:
		return false, errors.New(
			fmt.Sprintf("type '%v' cannot be used for argument matches", tokens.String(tokenType)))
	}

	// report no match if the token didn't match any of the arguments or argument abbreviations
	if ad == nil {
		return false, nil
	}

	// however if the argument matched token, we need to check if it's one of the supported values
	// for a command context
	if cmdCtx.ArgSupported(ad.Token) {
		return true, nil
	}

	// report if it's a valid argument, but is not supported for a command context
	return false, errors.New(
		fmt.Sprintf("argument '%v' is not supported by command '%v'", ad.Token, cmdCtx.Token))
}

// matchesFlag determines if a token matches flag (either in normal or abbreviated form)
func matchesFlag(token string, tokenType int, def *defs.Definitions) (bool, error) {
	// flag is expected to be prefixed with a single or double dash, if it's not -
	// there is no reason to perform further match checks
	if !strutil.HasExpectedPrefix(token, tokenType) {
		return false, nil
	}

	// check if flag matches a token or abbreviation
	switch tokenType {
	case tokens.Flag:
		return def.FlagByToken(strutil.TrimPrefix(token, tokenType)) != nil, nil
	case tokens.FlagAbbr:
		return def.FlagByAbbr(strutil.TrimPrefix(token, tokenType)) != nil, nil
	default:
		return false, errors.New(
			fmt.Sprintf("type '%v' cannot be used for flag matches", tokens.String(tokenType)))
	}
}

func matchesDefaultValue(token string, tokenType int, ctx *parsectx.Context, def *defs.Definitions) (bool, error) {

	if ctx.Command == nil ||
		(ctx.Argument != nil && !ctx.Argument.Default) {
		return false, nil
	}
	defArg := def.DefaultArg(ctx.Command)

	match, err := matchesValue(token, tokenType, defArg)

	if match && err == nil {
		ctx.Argument = defArg
	}

	return match, err
}

func matchesValue(token string, tokenType int, arg *defs.ArgumentDefinition) (bool, error) {
	if arg == nil {
		return false, errors.New("can't confirm a match for a value that is missing an argument context")
	}

	if tokenType == tokens.ValueDefault {
		if !arg.Default {
			return false, nil
		}
		tokenType = tokens.Value
		if len(arg.Values) > 0 {
			tokenType = tokens.ValueFixed
		}
	}

	switch tokenType {
	case tokens.ValueFixed:
		if arg.ValueSupported(token) {
			return true, nil
		} else {
			if len(arg.Values) > 0 {
				// check for sequence break (another argument) before verifying a fixed value
				if strutil.HasAnyPrefix(token) {
					return false, nil
				}
				return false, errors.New(fmt.Sprintf("unsupported value '%v' for an argument '%v'", token, arg.Token))
			} else {
				return false, nil
			}
		}
	case tokens.Value:
		return len(arg.Values) == 0, nil
	}
	return false, nil
}

// Matches determines if a token can be a given token type, in the context of existing
// command, argument (applicable to tokens.Argument and tokens.Value, tokens.ValueFixed, tokens.ValueDefault)
func Matches(token string, tokenType int, ctx *parsectx.Context, def *defs.Definitions) (bool, error) {
	switch tokenType {
	case tokens.Command:
		return def.CommandByToken(token) != nil, nil
	case tokens.CommandAbbr:
		return def.CommandByAbbr(token) != nil, nil
	case tokens.Flag:
		fallthrough
	case tokens.FlagAbbr:
		return matchesFlag(token, tokenType, def)
	case tokens.Argument:
		fallthrough
	case tokens.ArgumentAbbr:
		return matchesArgument(token, tokenType, ctx.Command, def)
	case tokens.ValueDefault:
		return matchesDefaultValue(token, tokenType, ctx, def)
	case tokens.ValueFixed:
		fallthrough
	case tokens.Value:
		return matchesValue(token, tokenType, ctx.Argument)
	default:
		return false, errors.New(
			fmt.Sprintf("cannot confirm match for a token '%v' of type '%v'", token, tokens.String(tokenType)))
	}
}
