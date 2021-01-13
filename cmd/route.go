package cmd

import (
	"github.com/boggydigital/clo"
)

func Dispatch(req *clo.Request) error {
	if req == nil {
		return clo.Route(nil)
	}
	verbose := req.GetFlag("verbose")
	switch req.Command {
	case "verify":
		return Verify(req.GetValue("path"), verbose)
	case "generate":
		return Generate(
			req.GetValue("app"),
			req.GetValues("command"),
			req.GetValues("argument"),
			req.GetValues("flag"))
	default:
		return clo.Route(req)
	}
}
