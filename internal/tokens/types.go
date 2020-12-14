package tokens

// this set of constants enumerates all distinct token types
const (
	Command = iota
	CommandAbbr
	Argument
	ArgumentAbbr
	ValueDefault
	ValueFixed
	Value
	Flag
	FlagAbbr
)

// String converts token type to a human readable string
func String(tokenType int) string {
	switch tokenType {
	case Command:
		return "command"
	case CommandAbbr:
		return "commandAbbr"
	case Argument:
		return "argument"
	case ArgumentAbbr:
		return "argumentAbbr"
	case ValueDefault:
		return "valueDefault"
	case ValueFixed:
		return "valueFixed"
	case Value:
		return "value"
	case Flag:
		return "flag"
	case FlagAbbr:
		return "flagAbbr"
	}
	return "unknown"
}

// Next defines possible token types that can follow a given type of token.
// Order below assumes the following sequence of token types:
// app command [<arguments> [<values>]] [<flags>]
func Next(tokenType int) []int {
	switch tokenType {
	// commands
	case Command:
		fallthrough
	case CommandAbbr:
		return []int{ArgumentAbbr, Argument, ValueDefault, FlagAbbr, Flag}
	// arguments
	case Argument:
		fallthrough
	case ArgumentAbbr:
		return []int{ValueFixed, ArgumentAbbr, Argument, FlagAbbr, Flag, Value}
	// values
	case ValueFixed:
		return []int{ValueFixed, ArgumentAbbr, Argument, FlagAbbr, Flag}
	case ValueDefault:
		return []int{ValueDefault, ArgumentAbbr, Argument, FlagAbbr, Flag}
	case Value:
		return []int{ArgumentAbbr, Argument, FlagAbbr, Flag, Value}
	// flags
	case Flag:
		fallthrough
	case FlagAbbr:
		return []int{Flag, FlagAbbr}
	default:
		return []int{}
	}
}

// First
func First() []int {
	return []int{Command, CommandAbbr}
}
