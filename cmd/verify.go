package cmd

import (
	"fmt"
	"github.com/boggydigital/clove"
)

func Verify(args map[string][]string, flags []string) error {

	file := args["path"][0]
	defs, err := clove.LoadDefs(file)
	if err != nil {
		return err
	}

	errors := defs.Verify()
	if len(errors) > 0 {
		fmt.Printf("Following errors were found in %s:\n", file)
		for _, err := range errors {
			fmt.Println("-", err.Error())
		}
	}

	if len(errors) == 0 {
		fmt.Printf("%s has been verified, no errors found\n", file)
	}

	return nil
}
