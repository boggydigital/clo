package main

import (
	"encoding/json"
	"fmt"
	"github.com/boggydigital/clove/internal/defs"
	"io/ioutil"
)

func main() {

	filename := "definitions.json"
	args := []string{"download", "--type", "products", "--media", "game"}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var dfs defs.Definitions

	err = json.Unmarshal(bytes, &dfs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req := dfs.Parse(args)
	req.Print()
}
