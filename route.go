package clo

import "fmt"

func Route(request *Request, defs *Definitions) error {
	if request == nil {
		request = &Request{
			Command: "help",
		}
	}

	if request.Command == "" {
		request.Command = "help"
	}

	switch request.Command {
	case "help":
		return printHelp(request.ArgVal("command"), defs)
	default:
		return fmt.Errorf("unknown command: '%s'", request.Command)
	}
}
