package parse

import (
	"errors"
	"fmt"
	"github.com/boggydigital/clove/internal/defs"
	"github.com/boggydigital/clove/internal/match"
	"github.com/boggydigital/clove/internal/parsectx"
	"github.com/boggydigital/clove/internal/request"
	"github.com/boggydigital/clove/internal/strutil"
	"github.com/boggydigital/clove/internal/tokens"
	"github.com/boggydigital/clove/internal/verify"
	"strings"
)

// Parse converts args to a structured Request or returns an error if there are unexpected values,
// order or if any of the defined constraints are not met: fixed values, required,
// multiple values, etc.
func Parse(args []string, def *defs.Definitions) (*request.Request, error) {

	var req = request.Request{
		Flags:     []string{},
		Command:   "",
		Arguments: make(map[string][]string),
	}

	var expected = tokens.First()
	var ctx parsectx.Context

	for _, arg := range args {
		if arg == "" {
			continue
		}
		arg = strings.ToLower(arg)
		matched := false
		for _, tt := range expected {
			success, err := match.Matches(arg, tt, &ctx, def)
			if err != nil {
				return &req, err
			}
			if success {
				matched = true
				expandedArg, err := strutil.ExpandAbbr(arg, tt, def)
				if err != nil {
					return nil, err
				}
				err = request.Update(&req, expandedArg, tt, &ctx)
				if err != nil {
					return nil, err
				}
				parsectx.Update(arg, tt, &ctx, def)
				expected = tokens.Next(tt)
				break
			}
		}
		if !matched {
			return nil, errors.New(fmt.Sprintf("unknown argument: '%v'", arg))
		}
	}

	err := verify.Constraints(&req, def)
	if err != nil {
		return &req, err
	}

	return &req, nil
}
