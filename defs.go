package clove

import (
	"github.com/boggydigital/clove/internal"
)

type Definitions struct {
	internal.Definitions
}

func loadEmbedded() (*Definitions, error) {
	defs, err := internal.LoadEmbedded()
	if err != nil {
		return nil, err
	}
	return &Definitions{Definitions: *defs}, err
}

func LoadDefinitions(path string) (*Definitions, error) {
	defs, err := internal.Load(path)
	if err != nil {
		return nil, err
	}
	return &Definitions{Definitions: *defs}, err
}
