package cmd

import (
	"fmt"
	"github.com/boggydigital/clove"
)

func Dispatch(request *clove.Request) error {
	switch request.Command {
	case "verify":
		return Verify(request.Arguments, request.Flags)
	default:
		return fmt.Errorf("unknown command: '%s'", request.Command)
	}
}
