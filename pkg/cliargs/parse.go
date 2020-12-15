package cliargs

import (
	"github.com/boggydigital/clove/internal/cliparse"
	"github.com/boggydigital/clove/internal/clireq"
	"github.com/boggydigital/clove/internal/defs"
)

func Parse(args []string, def *defs.Definitions) (*clireq.Request, error) {
	return cliparse.Parse(args, def)
}
