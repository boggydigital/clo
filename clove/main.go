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
		args = []string{"verify", "--path", "test.json"}
	}

	req, err := clove.Parse(args)
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	if len(req.Flags) > 0 {
		fmt.Println("flags:")
		for _, f := range req.Flags {
			fmt.Println("-", f)
		}
	}
	fmt.Println("command:", req.Command)
	if len(req.Arguments) > 0 {
		fmt.Println("arguments:")
		for a, v := range req.Arguments {
			fmt.Println("-", a, v)
		}
	}
}
