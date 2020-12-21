package internal

import "fmt"

func Dispatch(request *Request) error {
	if request == nil {
		return help("", false)
	}
	verbose := request.GetFlag("verbose")
	switch request.Command {
	case "help":
		return help(request.GetValue("help:command"), verbose)
	default:
		return fmt.Errorf("unknown command: '%s'", request.Command)
	}
}
