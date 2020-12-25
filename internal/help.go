package internal

import (
	"fmt"
	"strings"
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

func printHelp(cmd string, verbose bool) error {
	defs, err := LoadEmbedded()
	if err != nil {
		return err
	}
	if cmd == "" {
		return printAppHelp(defs, verbose)
	} else {
		return printCmdHelp(cmd, defs, verbose)
	}
}

func printAppHelp(defs *Definitions, verbose bool) error {
	appDesc := defs.Hint
	if verbose && defs.Desc != "" {
		appDesc = defs.Desc
	}
	fmt.Printf("%s - %s\n", defs.App, appDesc)
	fmt.Println()
	fmt.Printf("Usage: %s command [<arguments [<values>]>] [<flags>]\n",
		defs.App)
	fmt.Println()
	fmt.Println("Commands:")
	for _, cmd := range defs.Commands {
		cmdDesc := cmd.Hint
		if verbose && cmd.Desc != "" {
			cmdDesc = cmd.Desc
		}
		fmt.Printf("  %-"+defs.CommandsPadding()+"s  %s\n",
			cmd.Token,
			cmdDesc)
	}
	fmt.Println()
	if len(defs.Flags) > 0 {
		fmt.Println("Flags:")
		for _, flg := range defs.Flags {
			flgDesc := flg.Hint
			if verbose && flg.Desc != "" {
				flgDesc = flg.Desc
			}
			fmt.Printf("  %-"+defs.FlagsPadding()+"s  %s\n",
				flg.Token,
				flgDesc)
		}
		fmt.Println()
	}
	fmt.Printf("Run '%s help [command]' for more information on a command.\n",
		defs.App)

	return nil
}

func printExampleHelp(ex *ExampleDefinition, cmd string, defs *Definitions) {
	fmt.Printf("  '%s %s", defs.App, cmd)
	for avi, argVals := range ex.ArgumentsValues {
		for arg, vals := range argVals {
			fmt.Printf(" --%s ", arg)
			if len(vals) == 0 {
				fmt.Printf("<%s>", arg)
				continue
			}
			fmt.Print(strings.Join(vals, " "))
		}
		if avi == len(ex.ArgumentsValues)-1 {
			fmt.Print("'")
		}
	}
	fmt.Println(":", ex.Desc)
}

func printCmdHelp(cmd string, defs *Definitions, verbose bool) error {
	cmdUsage := fmt.Sprintf("Usage: %s %s ", defs.App, cmd)
	cd := defs.CommandByToken(cmd)
	if cd == nil {
		return fmt.Errorf("command token '%s' is not defined", cmd)
	}
	if len(cd.Arguments) > 0 {
		cmdUsage += "[<arguments>]"
	}
	fmt.Println(cmdUsage)
	fmt.Println()
	fmt.Println("Arguments:")
	for _, arg := range cd.Arguments {
		ad := defs.ArgByToken(arg)
		if ad == nil {
			fmt.Printf("  %s: invalid argument token\n", arg)
			continue
		}
		argDesc := ad.Hint
		if verbose && ad.Desc != "" {
			argDesc = ad.Desc
		}
		fmt.Printf("  %-"+defs.ArgumentsPadding(cmd)+"s  %s ", arg, argDesc)
		if ad.Default {
			fmt.Print("[Def]")
		}
		if ad.Required {
			fmt.Print("[Req]")
		}
		if ad.Multiple {
			fmt.Print("[Mlt]")
		}
		if ad.Env {
			envToken := fmt.Sprintf("%s_%s", strings.ToUpper(cmd), strings.ToUpper(arg))
			if defs.EnvPrefix != "" {
				envToken = fmt.Sprintf("%s_%s", strings.ToUpper(defs.EnvPrefix), envToken)
			}
			fmt.Printf("[Env:%s]", envToken)
		}
		fmt.Println()

		if len(ad.Values) > 0 {
			fmt.Printf("  %-"+defs.ArgumentsPadding(cmd)+"s  supported values: ", "")
			for i, av := range ad.Values {
				if i == len(ad.Values)-1 {
					fmt.Println(av)
				} else {
					fmt.Print(av, ", ")
				}
			}
		}

	}
	fmt.Println()
	if verbose {
		fmt.Println("Arguments [attributes] explanation:")
		fmt.Printf("  %-5s  %s\n", "[Def]", "default argument - value(s) can be provided right "+
			"after a command without an argument token")
		fmt.Printf("  %-5s  %s\n", "[Mlt]", "supports multiple values, that can be provided "+
			"in sequence or each with an argument token")
		fmt.Printf("  %-5s  %s\n", "[Req]", "required argument - app cannot "+
			"meaningfully run without a value")
		fmt.Printf("  %-5s  %s\n", "[Env]", "value can be provided with an "+
			"environment variable specified above")
		fmt.Println()
		fmt.Println("Examples:")
		for _, ex := range cd.Examples {
			printExampleHelp(&ex, cmd, defs)
		}
	} else {
		fmt.Printf("Run '%s help %s --verbose' for more information, "+
			"incl. examples and arguments [attributes] explaination.\n", defs.App, cmd)
	}

	return nil
}
