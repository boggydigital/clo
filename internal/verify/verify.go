package verify

import (
	"errors"
	"fmt"
	"github.com/boggydigital/clove/internal/defs"
	"github.com/boggydigital/clove/internal/request"
)

func CommandHasRequiredArgs(req *request.Request, def *defs.Definitions) error {
	if def == nil {
		return errors.New("cannot verify required argument using nil definitions")
	}
	if req == nil {
		return errors.New("cannot verify nil request for required arguments")
	}

	requiredArgs := def.RequiredArgs(req.Command)
	for _, ra := range requiredArgs {
		matched := false
		for at, _ := range req.Arguments {
			if ra == at {
				matched = true
				break
			}
		}
		if !matched {
			return errors.New(
				fmt.Sprintf("required argument '%v' is missing for the command '%v'", ra, req.Command))
		}
	}
	return nil
}

func ArgumentsMultipleValues(req *request.Request, def *defs.Definitions) error {
	if def == nil {
		return errors.New("cannot verify required argument using nil definitions")
	}
	if req == nil {
		return errors.New("cannot verify nil request for required arguments")
	}

	for at, avs := range req.Arguments {
		if at == "" {
			continue
		}
		arg := def.ArgByToken(at)
		if arg == nil {
			continue
		}
		if !arg.Multiple && len(avs) > 1 {
			return errors.New(
				fmt.Sprintf("argument '%v' has multiple values, supports no more than one", at))
		}
	}

	return nil
}

func Constraints(req *request.Request, def *defs.Definitions) error {
	err := CommandHasRequiredArgs(req, def)
	if err != nil {
		return err
	}
	err = ArgumentsMultipleValues(req, def)
	if err != nil {
		return err
	}
	return nil
}
