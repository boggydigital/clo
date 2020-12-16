package internal

import "fmt"

type ExampleDefinition struct {
	Arguments []string `json:"arguments"`
	Values    []string `json:"values"`
	Desc      string   `json:"desc,omitempty"`
}

type CommandDefinition struct {
	CommonDefinition
	Arguments []string            `json:"arguments,omitempty"`
	Examples  []ExampleDefinition `json:"examples,omitempty"`
}

func (cmd *CommandDefinition) ArgSupported(arg string) bool {
	for _, a := range cmd.Arguments {
		if a == arg {
			return true
		}
	}
	return false
}

func (ex *ExampleDefinition) Print() {
	fmt.Printf("Arguments:%v\nValues:%v\nDesc:%v\n",
		ex.Arguments, ex.Values, ex.Desc)
}

func (cmd *CommandDefinition) Print() {
	cmd.CommonDefinition.Print()
	fmt.Printf("Arguments:%v\n", cmd.Arguments)
	for _, ex := range cmd.Examples {
		ex.Print()
	}
}
