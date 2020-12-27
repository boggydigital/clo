package clo

import (
	"github.com/boggydigital/clo/internal"
	"io/ioutil"
)

// Definitions hold commands, arguments, values and flags
// constraints, descriptions and connections.
type Definitions struct {
	internal.Definitions
}

// LoadDefinitions loads definitions JSON from bytes.
func LoadDefinitions(path string) (*Definitions, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	defs, err := internal.Load(bytes)
	if err != nil {
		return nil, err
	}
	return &Definitions{Definitions: *defs}, err
}
