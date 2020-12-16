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
	valueDefault
	valueFixed
	value
	flag
	flagAbbr
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

// next defines possible token types that can follow a given type of token.
// Order below assumes the following sequence of token types:
// app command [<arguments> [<values>]] [<flags>]
func next(tokenType int) []int {
	switch tokenType {
	// commands
	case command:
		fallthrough
	case commandAbbr:
		return []int{argumentAbbr, argument, valueDefault, flagAbbr, flag}
	// arguments
	case argument:
		fallthrough
	case argumentAbbr:
		return []int{valueFixed, argumentAbbr, argument, flagAbbr, flag, value}
	// values
	case valueFixed:
		return []int{valueFixed, argumentAbbr, argument, flagAbbr, flag}
	case valueDefault:
		return []int{valueDefault, argumentAbbr, argument, flagAbbr, flag}
	case value:
		return []int{argumentAbbr, argument, flagAbbr, flag, value}
	// flags
	case flag:
		fallthrough
	case flagAbbr:
		return []int{flag, flagAbbr}
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
		ad := def.ArgByAbbr(trimPrefix(token, tokenType))
		if ad == nil {
			return "", fmt.Errorf("unknown argument abbreviation: '%v'", token)
		}
		return ad.Token, nil
	case flagAbbr:
		fd := def.FlagByAbbr(trimPrefix(token, tokenType))
		if fd == nil {
			return "", fmt.Errorf("unknown flag abbreviation: '%v'", token)
		}
		return fd.Token, nil
	default:
		return token, nil
	}
}
