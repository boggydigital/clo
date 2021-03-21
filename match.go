package clo

import (
	"fmt"
)

func matchArg(
	token string,
	cmd string,
	validCmdArg func(string, string) (string, error)) (string, error) {
	if !isArg(token) {
		return "", nil
	}
	arg, err := validCmdArg(cmd, trimArgPrefix(token))
	return arg, err
}

// match a token to a given token type, in the context of existing
// command, argument (applicable to tokens.argument and tokens.value, tokens.valueFixed, tokens.valueDefault)
func match(
	token string,
	tokenType int,
	cmdCtx, argCtx string,
	def *Definitions) (string, error) {
	if def == nil {
		return "", fmt.Errorf("cannot match token with nil definitions")
	}

	switch tokenType {
	case command:
		return def.definedCmd(token)
	case argument:
		return matchArg(token, cmdCtx, def.definedArg)
	case value:
		if isArg(token) {
			break
		}
		val, err := def.definedVal(cmdCtx, argCtx, token)
		return val, err
	default:
		return "", fmt.Errorf(
			"cannot confirm match for a token '%v' of type '%v'",
			token,
			tokenString(tokenType))
	}

	return "", nil
}
