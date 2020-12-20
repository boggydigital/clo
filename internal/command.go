package internal

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
