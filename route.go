package clo

import "github.com/boggydigital/clo/internal"

// Route Request to a function that can handle the command
// with the provided arguments and values.
func Route(request *Request, defs *Definitions) error {
	if request == nil {
		return internal.Route(nil, &defs.Definitions)
	}
	return internal.Route(&request.Request, &defs.Definitions)
}
