package clo

import (
	"fmt"
	"github.com/boggydigital/nod"
	"net/url"
)

func (defs *definitions) AssertCommandsHaveHandlers() error {
	for cmd := range defs.Cmd {
		if cmd == helpCmd {
			continue
		}
		if _, ok := defaultHandlers[cmd]; !ok {
			return fmt.Errorf("no handler registered for %s", cmd)
		}
	}

	return nil
}

func (defs *definitions) Serve(args []string) error {

	//1. parse args into URL
	u, err := defs.parseUrl(args)
	if err != nil {
		return err
	}

	nod.Log("clo: serving URL %s", u)

	if u == nil {
		u = &url.URL{
			Path: "help",
		}
	}

	if u.Path == "" || u.Path == helpCmd {
		q := u.Query()
		return printHelp(q.Get("command"), defs)
	}

	//2. route to the handler based on the path pattern
	if handler, ok := defaultHandlers[u.Path]; ok {
		return handler(u)
	}

	return fmt.Errorf("unknown command %s", u.Path)
}
