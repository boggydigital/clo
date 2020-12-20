package cmd

import (
	"fmt"
	"github.com/boggydigital/clove"
)

func Dispatch(request *clove.Request) error {
	verbose := request.GetFlag("verbose")
	switch request.Command {
	case "verify":
		return Verify(request.GetValue("path"), verbose)
	case "help":
		return Help(request.GetValue("command"), verbose)
	default:
		return fmt.Errorf("unknown command: '%s'", request.Command)
	}
}
