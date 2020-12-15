package clidefs

import (
	"encoding/json"
	"github.com/boggydigital/clove/internal/defs"
	"io/ioutil"
)

func Load(filepath string) (*defs.Definitions, error) {
	if filepath == "" {
		filepath = "./definitions.json"
	}

	var dfs *defs.Definitions

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return dfs, err
	}

	err = json.Unmarshal(bytes, &dfs)
	if err != nil {
		return dfs, err
	}

	return dfs, nil
}
