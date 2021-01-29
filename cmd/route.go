package cmd

import (
	"github.com/boggydigital/clo"
)

func Route(req *clo.Request, defs *clo.Definitions) error {
	if req == nil {
		return clo.Route(nil, defs)
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
		return clo.Route(req, defs)
	}
}
