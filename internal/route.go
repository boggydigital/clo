package internal

import "fmt"

const (
	helpCommand         = "help"
	helpCommandArgument = "help:command"
)

func Route(request *Request) error {
	if request == nil {
		request = &Request{
			Command: helpCommand,
		}
	}
	switch request.Command {
	case helpCommand:
		defs, err := LoadDefault()
		if err != nil {
			return err
		}
		return printHelp(request.ArgVal(helpCommandArgument), defs)
	default:
		return fmt.Errorf("unknown command: '%s'", request.Command)
	}
}
