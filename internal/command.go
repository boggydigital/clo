package internal

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
