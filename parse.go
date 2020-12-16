package clove

import (
	"encoding/json"
	"github.com/boggydigital/clove/internal"
	"io/ioutil"
)

type Request struct {
	internal.Request
}

func requestFromInternal(request *internal.Request) *Request {
	if request == nil {
		return nil
	}
	var req Request
	req.Request = *request
	return &req
}

//func lookupPaths() []string {
//	return []string {
//		"~/Library/Application Support/" + os.Args[0],
//	}
//}

func loadDefs(path string) (*internal.Definitions, error) {

	if path == "" {
		path = "./clove.json"
	}

	var dfs *internal.Definitions

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

func Parse(args []string) (*Request, error) {

	def, err := loadDefs("")
	if err != nil {
		return nil, err
	}

	req, err := def.Parse(args)
	return requestFromInternal(req), err
}
