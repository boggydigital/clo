package main

import (
	"encoding/json"
	"fmt"
	"github.com/boggydigital/clove/internal/defs"
	"github.com/boggydigital/clove/internal/parse"
	"io/ioutil"
	"os"
)

func main() {

	// TODO: bug: while the arg is lower case - the original definition is used as is

	filename := "definitions.json"
	args := []string{"dl", "products", "images", "--id", "1", "2", "3", "--media", "game", "movie", "--help", "-v"}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var dfs *defs.Definitions

	err = json.Unmarshal(bytes, &dfs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req, err := parse.Parse(args, dfs)
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
	fmt.Println("----------")
	req.Print()
}
