package clo

import (
	"github.com/boggydigital/clo/internal"
)

// Definitions hold commands, arguments, values and flags
// constraints, descriptions and connections.
type Definitions struct {
	internal.Definitions
}

// LoadDefinitions loads definitions JSON from bytes.
func LoadDefinitions(path string, verbose bool) (*Definitions, error) {
	defs, err := internal.Load(path)
	if err != nil {
		return nil, err
	}
	return &Definitions{Definitions: *defs}, err
}
