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

	filename := "definitions.json"
	args := []string{"dl", "products", "images", "accountProducts", "--id", "1", "2", "3", "--media", "game", "movie", "--help", "-v"}

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
