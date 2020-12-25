package clove

import (
	"github.com/boggydigital/clove/internal"
)

type Request struct {
	internal.Request
}

// Parse converts args into a structured Request that has
// a command, arguments with values, flags - all according to
// definitions provided in the JSON file.
func Parse(args []string) (*Request, error) {

	if len(args) == 0 {
		return nil, nil
	}

	def, err := loadEmbedded()
	if err != nil {
		return nil, err
	}

	req, err := def.Parse(args)
	if req == nil || err != nil {
		return nil, err
	}

	return &Request{Request: *req}, err
}
