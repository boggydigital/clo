package internal

import "fmt"

func Dispatch(request *Request) error {
	if request == nil {
		request = &Request{
			Command: "help",
		}
	}
	verbose := request.GetFlag("verbose")
	switch request.Command {
	case "help":
		return printHelp(request.GetValue("help:command"), verbose)
	default:
		return fmt.Errorf("unknown command: '%s'", request.Command)
	}
}
