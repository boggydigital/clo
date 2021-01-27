package internal

import "fmt"

func Route(request *Request) error {
	if request == nil {
		request = &Request{
			Command: "help",
		}
	}
	switch request.Command {
	case "help":
		defs, err := LoadDefault()
		if err != nil {
			return err
		}
		return printHelp(request.ArgVal("command"), defs)
	default:
		return fmt.Errorf("unknown command: '%s'", request.Command)
	}
}
