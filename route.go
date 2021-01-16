package clo

import "github.com/boggydigital/clo/internal"

// Route Request to a function that can handle the command
// with the provided arguments and values.
func Route(request *Request) error {
	if request == nil {
		return internal.Route(nil)
	}
	return internal.Route(&request.Request)
}
