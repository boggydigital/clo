package clo

import (
	"fmt"
	"net/url"
)

func (defs *Definitions) Serve(args []string) error {
	//1. assert that all commands have handlers
	//2. parse args into URL
	//3. route to the handler based on the path pattern

	for cmd := range defs.Cmd {
		if cmd == helpCmd {
			continue
		}
		if _, ok := defaultHandlers[cmd]; !ok {
			return fmt.Errorf("no handler registered for %s", cmd)
		}
	}

	u, err := defs.parseUrl(args)
	if err != nil {
		return err
	}

	if u == nil {
		u = &url.URL{
			Path: "help",
		}
	}

	if u.Path == "" || u.Path == helpCmd {
		q := u.Query()
		return printHelp(q.Get("command"), defs)
	}

	if handler, ok := defaultHandlers[u.Path]; ok {
		return handler(u)
	}

	return fmt.Errorf("unknown command %s", u.Path)
}
