package internal

import (
	"errors"
	"fmt"
)

func commandHasRequiredArgs(req *Request, def *Definitions) error {
	if def == nil {
		return errors.New("cannot verify required argument using nil definitions")
	}
	if req == nil {
		return errors.New("cannot verify nil request for required arguments")
	}

	requiredArgs := def.RequiredArgs(req.Command)
	for _, ra := range requiredArgs {
		matched := false
		for at, avs := range req.Arguments {
			if ra == at {
				if len(avs) == 0 {
					return fmt.Errorf("required argument '%v' is missing values", ra)
				}
				matched = true
				break
			}
		}
		if !matched {
			return fmt.Errorf("required argument '%v' is missing for the command '%v'", ra, req.Command)
		}
	}
	return nil
}

func argumentsMultipleValues(req *Request, def *Definitions) error {
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
			return fmt.Errorf("argument '%v' has multiple values, supports no more than one", at)
		}
	}

	return nil
}

func (req *Request) verify(def *Definitions) error {
	err := commandHasRequiredArgs(req, def)
	if err != nil {
		return err
	}
	err = argumentsMultipleValues(req, def)
	if err != nil {
		return err
	}
	return nil
}
