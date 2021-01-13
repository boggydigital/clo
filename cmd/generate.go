package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/boggydigital/clo/internal"
)

func Generate(app string, commands, arguments, flags []string) error {
	defs := internal.GenDefinitions(app, commands, arguments, flags)

	bytes, err := json.Marshal(defs)
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))

	return nil
}
