package clo

import (
	"fmt"
	"net/url"

	"github.com/boggydigital/nod"
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

func (defs *definitions) Serve(u *url.URL) error {
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

func (defs *definitions) ParseServe(args []string) error {

	//1. parse args into URL
	u, err := defs.Parse(args)
	if err != nil {
		return err
	}

	return defs.Serve(u)
}
