package clo

// this set of constants enumerates all distinct token types
const (
	command = iota
	argument
	value
)

// tokenString converts token type to a human readable string
func tokenString(tokenType int) string {
	switch tokenType {
	case command:
		return "command"
	case argument:
		return "argument"
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
		return []int{argument, value}
	case argument:
		return []int{value, argument}
	case value:
		return []int{argument, value}
	default:
		return []int{}
	}
}

func initial() []int {
	return []int{command, argument, value}
}
