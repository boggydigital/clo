package clo

import (
	"net/url"
)

type Handler func(u *url.URL) error

var defaultHandlers = make(map[string]Handler)

func HandleFunc(pattern string, handler Handler) {
	defaultHandlers[pattern] = handler
}

func HandleFuncs(handlers map[string]Handler) {
	for p, h := range handlers {
		defaultHandlers[p] = h
	}
}
