package cmd

import (
	"fmt"
	"github.com/boggydigital/clo"
	"io/ioutil"
)

func Verify(filepath string, verbose bool) error {

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	defs, err := clo.LoadDefinitions(bytes)
	if err != nil {
		return err
	}

	errors := defs.Verify(verbose)
	if len(errors) > 0 {
		fmt.Printf("Following errors were found in %s:\n", filepath)
		for _, err := range errors {
			fmt.Println("-", err.Error())
		}
	}

	if len(errors) == 0 {
		fmt.Printf("%s has been verified, no errors found\n", filepath)
	}

	return nil
}
