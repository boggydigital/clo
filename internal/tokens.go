package internal

// this set of constants enumerates all distinct token types
const (
	command = iota
	//commandAbbr
	argument
	//argumentAbbr
	//valueDefault
	//valueFixed
	value
)

// tokenString converts token type to a human readable string
func tokenString(tokenType int) string {
	switch tokenType {
	case command:
		return "command"
	//case commandAbbr:
	//	return "commandAbbr"
	case argument:
		return "argument"
	//case argumentAbbr:
	//	return "argumentAbbr"
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
	//case commandAbbr:
	//	return []int{argumentAbbr, argument, value}
	case argument:
		//	fallthrough
		//case argumentAbbr:
		return []int{argument, value}
	case value:
		return []int{argument, value}
	default:
		return []int{}
	}
}

// first
func first() []int {
	return []int{command}
}

//func expandAbbr(token string, tokenType int, def *Definitions) string {
//	switch tokenType {
//	case command:
//		cd := def.CommandByAbbr(token)
//		if cd == nil {
//			return token
//		}
//		return cd.Token
//	case argument:
//		ad := def.ArgByAbbr(trimPrefix(token))
//		if ad == nil {
//			return token
//		}
//		return ad.Token
//	default:
//		return token
//	}
//}
