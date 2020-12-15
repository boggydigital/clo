package main

import (
	"fmt"
	"github.com/boggydigital/clove/pkg/cliargs"
	"github.com/boggydigital/clove/pkg/clidefs"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"dl", "products", "images", "accountProducts", "--id", "1", "2", "3", "--media", "game", "movie", "--help", "-v"}
	}

	dfs, err := clidefs.Load("")
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	req, err := cliargs.Parse(args, dfs)
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	fmt.Println("----------")
	req.Print()
}
