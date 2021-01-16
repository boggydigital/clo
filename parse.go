package clo

import (
	"github.com/boggydigital/clo/internal"
)

type Request struct {
	internal.Request
}

// Parse converts args into a structured Request that has
// a command, arguments with values - all according to
// the definitions provided in the a JSON file.
func Parse(args []string) (*Request, error) {

	if len(args) == 0 {
		return nil, nil
	}

	defs, err := internal.LoadDefault()
	if err != nil || defs == nil {
		return nil, err
	}

	req, err := defs.Parse(args)
	if req == nil || err != nil {
		return nil, err
	}

	return &Request{Request: *req}, err
}
