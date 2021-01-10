package main

import (
	"fmt"
	"github.com/boggydigital/clo"
	"github.com/boggydigital/clo/cmd"
	"os"
)

func main() {
	req, err := clo.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("error:", err.Error())
		if err = clo.Route(req); err != nil {
			fmt.Println("error:", err.Error())
		}
		os.Exit(1)
	}

	if err := cmd.Dispatch(req); err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
}
