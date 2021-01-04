package internal

import "fmt"

const (
	helpCommand         = "help"
	helpCommandArgument = "help:command"
	verboseFlag         = "verbose"
)

func Dispatch(request *Request) error {
	if request == nil {
		request = &Request{
			Command: helpCommand,
		}
	}
	verbose := request.GetFlag(verboseFlag)
	switch request.Command {
	case helpCommand:
		defs, err := LoadDefault()
		if err != nil {
			return err
		}
		return printHelp(request.GetValue(helpCommandArgument), defs, verbose)
	default:
		return fmt.Errorf("unknown command: '%s'", request.Command)
	}
}
