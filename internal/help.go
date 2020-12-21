package internal

import (
	"fmt"
)

func (def *Definitions) addHelpCmd() error {
	if def == nil {
		return fmt.Errorf("cannot add 'help' command to nil definitions")
	}

	// check if definitions already have 'help' token specified
	helpCmd := def.CommandByToken("help")
	if helpCmd != nil {
		// don't return error as we're assuming author knows what they're doing
		return nil
	}

	helpCmd = &CommandDefinition{
		CommonDefinition: CommonDefinition{
			Token: "help",
			Abbr:  "",
			Hint:  "display help",
			Desc:  "display help for the app or a specific command",
		},
	}

	// set abbreviation if not used by other commands
	hAbbr := def.CommandByAbbr("h")
	if hAbbr == nil {
		helpCmd.Abbr = "h"
	}

	// create argument
	argToken := "help:command"
	commandArg := def.ArgByToken(argToken)
	if commandArg != nil {
		return fmt.Errorf("cannot add 'help' command as argument '%s' already exists", argToken)
	}

	commandArg = &ArgumentDefinition{
		CommonDefinition: CommonDefinition{
			Token: argToken,
			Abbr:  "",
			Hint:  "app command",
		},
		Default:  true,
		Multiple: false,
		Required: false,
		Values:   []string{"from:commands"},
	}

	cAbbr := def.ArgByAbbr("c")
	if cAbbr == nil {
		commandArg.Abbr = "c"
	}

	def.Arguments = append(def.Arguments, *commandArg)

	helpCmd.Arguments = []string{argToken}

	def.Commands = append(def.Commands, *helpCmd)

	return nil
}

func help(cmd string, verbose bool) error {
	if cmd == "" {
		fmt.Println("help for the app")
	} else {
		fmt.Printf("help for command '%s'\n", cmd)
	}
	return nil
}
