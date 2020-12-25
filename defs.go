package clove

import (
	"github.com/boggydigital/clove/internal"
)

// Definitions hold commands, arguments, values and flags
// constraints, descriptions and connections.
type Definitions struct {
	internal.Definitions
}

// loadEmbedded loads definitions embedded into the app.
// TODO: Will actually to that in the 1.16
func loadEmbedded() (*Definitions, error) {
	defs, err := internal.LoadEmbedded()
	if err != nil {
		return nil, err
	}
	return &Definitions{Definitions: *defs}, err
}

// LoadDefinitions loads definitions JSON at a certain path.
func LoadDefinitions(path string) (*Definitions, error) {
	defs, err := internal.Load(path)
	if err != nil {
		return nil, err
	}
	return &Definitions{Definitions: *defs}, err
}
