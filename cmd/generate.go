package cmd

import "fmt"

func Generate(commands, arguments, flags []string) error {
	fmt.Println("generate")
	fmt.Println("commands:", commands)
	fmt.Println("arguments:", arguments)
	fmt.Println("flags:", flags)
	return nil
}
