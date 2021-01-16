package clo

import (
	"github.com/boggydigital/clo/internal"
)

// Definitions hold commands, arguments, values.
type Definitions struct {
	internal.Definitions
}

// LoadDefinitions loads definitions JSON from a path.
func LoadDefinitions(path string) (*Definitions, error) {
	defs, err := internal.Load(path)
	if err != nil {
		return nil, err
	}
	return &Definitions{Definitions: *defs}, err
}
