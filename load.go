package clove

import (
	"encoding/json"
	"github.com/boggydigital/clove/internal/defs"
	"io/ioutil"
)

//func lookupPaths() []string {
//	return []string {
//		"~/Library/Application Support/" + os.Args[0],
//	}
//}

func LoadDefs(path string) (*defs.Definitions, error) {
	if path == "" {
		path = "./definitions.json"
	}

	var dfs *defs.Definitions

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return dfs, err
	}

	err = json.Unmarshal(bytes, &dfs)
	if err != nil {
		return dfs, err
	}

	return dfs, nil
}
