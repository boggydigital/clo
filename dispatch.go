package clo

import "github.com/boggydigital/clo/internal"

// Dispatch Request to a function that can handle the command,
// with the arguments, and maybe flags.
func Dispatch(request *Request) error {
	if request == nil {
		return internal.Dispatch(nil)
	}
	return internal.Dispatch(&request.Request)
}
