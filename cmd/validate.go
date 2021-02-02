package cmd

import (
	"fmt"
	"github.com/boggydigital/clo"
)

func Validate(path string, verbose bool) error {

	defs, err := clo.LoadDefinitions(path)
	if err != nil {
		return err
	}

	errors := defs.Validate(verbose)
	if len(errors) > 0 {
		fmt.Printf("Following errors were found in %s:\n", path)
		for _, err := range errors {
			fmt.Println("-", err.Error())
		}
	}

	if len(errors) == 0 {
		fmt.Printf("%s has been verified, no errors found\n", path)
	}

	return nil
}