package main

import (
	"fmt"
	"github.com/boggydigital/clove"
	"github.com/boggydigital/clove/cmd"
	"os"
)

func main() {

	req, err := clove.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	if err := cmd.Dispatch(req); err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
}
