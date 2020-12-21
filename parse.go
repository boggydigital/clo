package clove

import (
	"encoding/json"
	"fmt"
	"github.com/boggydigital/clove/internal"
	"io/ioutil"
)

type Request struct {
	internal.Request
}

type Definitions struct {
	internal.Definitions
}

func loadEmbedDefs() (*Definitions, error) {
	return LoadExtDefs("./clove.json")
}

func LoadExtDefs(path string) (*Definitions, error) {
	if path == "" {
		return nil, fmt.Errorf("cannot load definition with no path specified")
	}

	var dfs *Definitions

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return dfs, err
	}
	err = json.Unmarshal(bytes, &dfs)

	return dfs, err
}

func Parse(args []string) (*Request, error) {

	if len(args) == 0 {
		return nil, nil
	}

	// TODO: Parse should use embedded clove.json
	// in golang 1.16: https://github.com/golang/go/issues/41191
	def, err := loadEmbedDefs()
	if err != nil {
		return nil, err
	}

	req, err := def.Parse(args)
	if req == nil || err != nil {
		return nil, err
	}

	return &Request{Request: *req}, err
}
