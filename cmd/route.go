package cmd

import (
	"github.com/boggydigital/clo"
)

func Route(req *clo.Request, defs *clo.Definitions) error {
	if req == nil {
		return clo.Route(nil, defs)
	}
	switch req.Command {
	case "validate":
		verbose := req.Flag("verbose")
		return Validate(req.ArgVal("path"), verbose)
	case "generate":
		return Generate(
			req.ArgVal("app"),
			req.ArgValues("command"),
			req.ArgValues("argument"))
	default:
		return clo.Route(req, defs)
	}
}
