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
		return Validate(req.ArgVal("path"), req.Flag("verbose"))
	default:
		return clo.Route(req, defs)
	}
}
