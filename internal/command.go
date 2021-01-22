package internal

import (
	"fmt"
	"strings"
)

type CommandDefinition struct {
	CommonDefinition
	Arguments         []string `json:"arguments,omitempty"`
	defaultArgument   string
	requiredArguments []string
}

func (cmd *CommandDefinition) ValidArg(arg string) bool {
	for _, a := range cmd.Arguments {
		if a == arg {
			return true
		}
	}
	return false
}

func (cmd *CommandDefinition) setDefaultRequired() error {
	if cmd == nil {
		return fmt.Errorf("cannot set default argument for a nil command")
	}
	for i, arg := range cmd.Arguments {
		targ := trimArgument(arg)
		if strings.HasPrefix(arg, defaultPrefix) {
			if cmd.defaultArgument != "" {
				return fmt.Errorf(
					"command %s already has default argument defined as %s",
					cmd.Token,
					cmd.defaultArgument)
			}
			cmd.defaultArgument = targ
		}
		if strings.HasSuffix(arg, requiredSuffix) {
			if cmd.requiredArguments == nil {
				cmd.requiredArguments = make([]string, 0)
			}
			cmd.requiredArguments = append(cmd.requiredArguments, targ)
		}
		if cmd.defaultArgument != "" || len(cmd.requiredArguments) > 0 {
			cmd.Arguments[i] = targ
		}
	}
	return nil
}
