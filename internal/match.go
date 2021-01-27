package internal

import (
	"fmt"
)

func matchArg(
	token string,
	cmd string,
	validCmdArg func(string, string) (string, string)) string {
	if !isArg(token) {
		return ""
	}
	_, arg := validCmdArg(cmd, trimArgPrefix(token))
	return arg
}

// match a token to a given token type, in the context of existing
// command, argument (applicable to tokens.argument and tokens.value, tokens.valueFixed, tokens.valueDefault)
func match(
	token string,
	tokenType int,
	req *Request,
	def *Definitions) (string, error) {
	if def == nil {
		return "", fmt.Errorf("cannot match token with nil definitions")
	}

	switch tokenType {
	case command:
		return def.definedCmd(token), nil
	case argument:
		return matchArg(token, req.Command, def.definedCmdArg), nil
	case value:
		if isArg(token) {
			break
		}
		_, _, val := def.definedCmdArgVal(req.Command, req.lastArgument(), token)
		return val, nil
	default:
		return "", fmt.Errorf(
			"cannot confirm match for a token '%v' of type '%v'",
			token,
			tokenString(tokenType))
	}

	return "", nil
}
