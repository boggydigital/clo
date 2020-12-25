package main

import (
	"fmt"
	"github.com/boggydigital/clove"
	"github.com/boggydigital/clove/cmd"
	"io/ioutil"
	"os"
)

func main() {

	defBytes, err := ioutil.ReadFile("app/clove.json")
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
	defs, err := clove.LoadDefinitions(defBytes)
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	req, err := defs.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("error:", err.Error())
		if err = clove.Dispatch(req); err != nil {
			fmt.Println("error:", err.Error())
		}
		os.Exit(1)
	}

	if err := cmd.Dispatch(req); err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
}
