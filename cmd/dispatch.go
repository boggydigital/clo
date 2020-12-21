package cmd

import (
	"github.com/boggydigital/clove"
	"github.com/boggydigital/clove/internal"
)

func Dispatch(request *clove.Request) error {
	if request == nil {
		return internal.Dispatch(nil)
	}
	verbose := request.GetFlag("verbose")
	switch request.Command {
	case "verify":
		return Verify(request.GetValue("path"), verbose)
	default:
		return internal.Dispatch(&request.Request)
	}
}
