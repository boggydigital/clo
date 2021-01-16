package internal

type CommandDefinition struct {
	CommonDefinition
	Arguments []string `json:"arguments,omitempty"`
}

func (cmd *CommandDefinition) ValidArg(arg string) bool {
	for _, a := range cmd.Arguments {
		if a == arg {
			return true
		}
	}
	return false
}
