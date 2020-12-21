package clove

import "github.com/boggydigital/clove/internal"

func Dispatch(request *Request) error {
	if request == nil {
		return internal.Dispatch(nil)
	}
	return internal.Dispatch(&request.Request)
}
