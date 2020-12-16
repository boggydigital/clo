package main

import (
	"fmt"
	"github.com/boggydigital/clove"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		// TODO: Results in error: argument 'path' has multiple values, supports no more than one
		args = []string{"verify", "-path", "test.json"}
	}

	dfs, err := clove.LoadDefs("./definitions.json")
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	req, err := clove.Parse(args, dfs)
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	fmt.Println("----------")
	req.Print()
}
