package cmd

import (
	"github.com/boggydigital/clo"
)

func Dispatch(request *clo.Request) error {
	if request == nil {
		return clo.Route(nil)
	}
	verbose := request.GetFlag("verbose")
	switch request.Command {
	case "verify":
		return Verify(request.GetValue("path"), verbose)
	default:
		return clo.Route(request)
	}
}
