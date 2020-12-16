package clove

import (
	"encoding/json"
	"github.com/boggydigital/clove/internal"
	"io/ioutil"
	"path/filepath"
)

type Request struct {
	internal.Request
}

type Definitions struct {
	internal.Definitions
}

func requestFromInternal(request *internal.Request) *Request {
	if request == nil {
		return nil
	}
	var req Request
	req.Request = *request
	return &req
}

func lookupPaths() []string {
	return []string{
		".",
	}
}

func embeddedDefs() *Definitions {
	return nil
}

func LoadDefs(path string) (*Definitions, error) {

	dfs := embeddedDefs()
	if dfs != nil {
		return dfs, nil
	}

	defFilename := "clove.json"

	for _, p := range lookupPaths() {
		path := filepath.Join(p, defFilename)

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

	return nil, nil
}

func Parse(args []string) (*Request, error) {

	def, err := LoadDefs("")
	if err != nil {
		return nil, err
	}

	req, err := def.Parse(args)

	return requestFromInternal(req), err
}
