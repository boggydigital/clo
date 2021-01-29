package main

import (
	"fmt"
	"github.com/boggydigital/clo"
	"github.com/boggydigital/clo/cmd"
	"os"
)

func main() {
	defs, err := clo.LoadDefinitions("clo.json")

	req, err := clo.Parse(os.Args[1:], defs)
	if err != nil {
		fmt.Println("error:", err.Error())
		if err = clo.Route(req, defs); err != nil {
			fmt.Println("error:", err.Error())
		}
		os.Exit(1)
	}

	if err := cmd.Route(req, defs); err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
}
