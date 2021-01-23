package internal

import (
	"errors"
	"fmt"
)

func matchCommand(token string, tokenType int, def *Definitions) (bool, error) {
	//if def == nil {
	//	return false, fmt.Errorf("cannot match argument with nil definitions")
	//}
	//var cd *CommandDefinition
	//switch tokenType {
	//case command:
	//	cd = def.CommandByToken(token)
	//default:
	//	return false, fmt.Errorf("type '%v' cannot be used for command matches", tokenString(tokenType))
	//}
	//
	//if cd == nil {
	//	return false, nil
	//}

	return false, nil
}

// matchesArgument determines if an arg token matches argument (either in normal or
// abbreviated form) for a given command
func matchArgument(token string, tokenType int, cmd string, def *Definitions) (bool, error) {
	//// argument is expected to be prefixed with a single or double dash, if it's not -
	//// there is no reason to perform further match checks
	//if !isArg(token) {
	//	return false, nil
	//}
	//
	//if def == nil {
	//	return false, fmt.Errorf("cannot match argument with nil definitions")
	//}
	//// first, try looking up argument by a token or an abbreviation,
	//// given a token type
	//var ad *ArgumentDefinition
	//switch tokenType {
	//case argument:
	//	ad = def.ArgByToken(trimArgPrefix(token))
	////case argumentAbbr:
	//default:
	//	return false, fmt.Errorf("type '%v' cannot be used for argument matches", tokenString(tokenType))
	//}
	//
	//// report no match if the token didn't match any of the arguments or argument abbreviations
	//if ad == nil {
	//	return false, nil
	//}
	//
	//cd := def.CommandByToken(cmd)
	//if cd == nil {
	//	return false, fmt.Errorf("cannot validate argument in the context of nil command")
	//}
	//
	//// however if the argument matched token, we need to check if it's one of the supported values
	//// for a command context
	//if cd.ValidArg(ad.Token) {
	//	return true, nil
	//}
	//
	//// report if it's a valid argument, but is not supported for a command context
	//return false, fmt.Errorf("argument '%v' is not supported by command '%v'", ad.Token, cmd)
	return false, nil
}

//func matchDefaultValue(token string, tokenType int, ctx *parseCtx, def *Definitions) (bool, error) {
//
//	if tokenType != valueDefault {
//		return false, nil
//	}
//
//	if ctx == nil {
//		return false, fmt.Errorf("cannot match default value with nil context")
//	}
//	if ctx.Command == nil {
//		return false, fmt.Errorf("cannot match default value with nil ctx.Command")
//	}
//
//	if def == nil {
//		return false, fmt.Errorf("cannot match default value with nil definitions")
//	}
//	// TODO: verify not nil
//	defArg := def.ArgByToken(ctx.Command.DefaultArgument)
//
//	m, err := matchValue(token, tokenType, defArg)
//
//	if m && err == nil {
//		ctx.Argument = defArg
//	}
//
//	return m, err
//}

func matchValue(token string, tokenType int, arg *ArgumentDefinition) (bool, error) {
	if arg == nil {
		return false, errors.New("can't confirm a match for a value that is missing argument context")
	}

	if isArg(token) {
		return false, nil
	}

	//if tokenType == valueDefault {
	//	tokenType = value
	//	if len(arg.Values) > 0 {
	//		tokenType = valueFixed
	//	}
	//}

	switch tokenType {
	//case valueFixed:
	//	if arg.ValidValue(token) {
	//		return true, nil
	//	} else {
	//		if len(arg.Values) > 0 {
	//			// check for sequence break (another argument) before verifying a fixed value
	//			return false, fmt.Errorf("unsupported value '%v' for an argument '%v'", token, arg.Token)
	//		} else {
	//			return false, nil
	//		}
	//	}
	case value:
		if len(arg.Values) == 0 {
			return true, nil
		} else {
			return arg.ValidValue(token), nil
		}
	default:
		return false, nil
	}
}

// match a token to a given token type, in the context of existing
// command, argument (applicable to tokens.argument and tokens.value, tokens.valueFixed, tokens.valueDefault)
func match(token string, tokenType int, req *Request, def *Definitions) (bool, error) {
	if def == nil {
		return false, fmt.Errorf("cannot match token with nil definitions")
	}
	switch tokenType {
	case command:
		return matchCommand(token, tokenType, def)
	case argument:
		return matchArgument(token, tokenType, req.Command, def)
	case value:
		//if !req.hasArguments() {
		//	if req.Command != "" {
		//		cd := def.CommandByToken(req.Command)
		//		if cd != nil && cd.defaultArgument != "" {
		//			err := req.update(cd.defaultArgument, argument)
		//			if err != nil {
		//				return false, err
		//			}
		//		} else {
		//			return false, fmt.Errorf("cannot match value for missing argument")
		//		}
		//	} else {
		//		return false, fmt.Errorf("cannot match value for missing command")
		//	}
		//}

		//arg := req.lastArgument()
		//if arg == "" {
		//	return false, fmt.Errorf("cannot match value for missing argument")
		//}
		//ad := def.ArgByToken(arg)
		//if ad == nil {
		//	return false, fmt.Errorf("cannot match vlaue for nil argument")
		//}

		//return matchValue(token, tokenType, ad)
		return false, nil
	default:
		return false, fmt.Errorf(
			"cannot confirm match for a token '%v' of type '%v'",
			token,
			tokenString(tokenType))
	}
}
