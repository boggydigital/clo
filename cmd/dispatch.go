package cmd

import (
	"github.com/boggydigital/clove"
)

func Dispatch(request *clove.Request) error {
	if request == nil {
		return clove.Dispatch(nil)
	}
	verbose := request.GetFlag("verbose")
	switch request.Command {
	case "verify":
		return Verify(request.GetValue("path"), verbose)
	default:
		return clove.Dispatch(request)
	}
}
