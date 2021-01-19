package internal

import (
	"fmt"
)

// this set of constants enumerates all distinct token types
const (
	command = iota
	commandAbbr
	argument
	argumentAbbr
	//valueDefault
	//valueFixed
	value
)

// tokenString converts token type to a human readable string
func tokenString(tokenType int) string {
	switch tokenType {
	case command:
		return "command"
	case commandAbbr:
		return "commandAbbr"
	case argument:
		return "argument"
	case argumentAbbr:
		return "argumentAbbr"
	//case valueDefault:
	//	return "valueDefault"
	//case valueFixed:
	//	return "valueFixed"
	case value:
		return "value"
	}
	return "unknown"
}

// next defines possible token types that can follow a given type of token.
// Order below assumes the following sequence of token types:
// app command [<arguments> [<values>]]
func next(tokenType int) []int {
	switch tokenType {
	case command:
		fallthrough
	case commandAbbr:
		return []int{argumentAbbr, argument, value}
	case argument:
		fallthrough
	case argumentAbbr:
		return []int{value, argumentAbbr, argument}
	case value:
		return []int{argumentAbbr, argument, value}
	default:
		return []int{}
	}
}

// first
func first() []int {
	return []int{command, commandAbbr}
}

func expandAbbr(token string, tokenType int, def *Definitions) (string, error) {
	switch tokenType {
	case commandAbbr:
		cd := def.CommandByAbbr(token)
		if cd == nil {
			return "", fmt.Errorf("unknown command abbreviation: '%v'", token)
		}
		return cd.Token, nil
	case argumentAbbr:
		ad := def.ArgByAbbr(trimPrefix(token))
		if ad == nil {
			return "", fmt.Errorf("unknown argument abbreviation: '%v'", token)
		}
		return ad.Token, nil
	default:
		return token, nil
	}
}
