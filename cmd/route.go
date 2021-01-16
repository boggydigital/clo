package cmd

import (
	"github.com/boggydigital/clo"
)

func Dispatch(req *clo.Request) error {
	if req == nil {
		return clo.Route(nil)
	}
	switch req.Command {
	case "verify":
		debug := req.Flag("debug")
		return Verify(req.ArgVal("path"), debug)
	case "generate":
		return Generate(
			req.ArgVal("app"),
			req.ArgValues("command"),
			req.ArgValues("argument"))
	default:
		return clo.Route(req)
	}
}
