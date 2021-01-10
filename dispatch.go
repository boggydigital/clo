package clo

import "github.com/boggydigital/clo/internal"

// Route Request to a function that can handle the command,
// with the arguments, and maybe flags.
func Route(request *Request) error {
	if request == nil {
		return internal.Route(nil)
	}
	return internal.Route(&request.Request)
}
