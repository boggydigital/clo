package main

import (
	"fmt"
	"github.com/boggydigital/clo"
	"github.com/boggydigital/clo/cmd"
	"io/ioutil"
	"os"
)

func main() {

	defBytes, err := ioutil.ReadFile("app/clo.json")
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
	defs, err := clo.LoadDefinitions(defBytes)
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	req, err := defs.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("error:", err.Error())
		if err = clo.Dispatch(req); err != nil {
			fmt.Println("error:", err.Error())
		}
		os.Exit(1)
	}

	if err := cmd.Dispatch(req); err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
}
