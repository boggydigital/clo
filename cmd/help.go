package cmd

import "fmt"

func Help(cmd string, verbose bool) error {
	fmt.Printf("help for command '%s'\n", cmd)
	return nil
}
