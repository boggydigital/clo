package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func createHelpCommandDefinition() *CommandDefinition {
	return &CommandDefinition{
		CommonDefinition: CommonDefinition{
			Token: "help",
			Abbr:  "",
			Help:  "display help",
		},
	}
}

func createHelpArgumentDefinition(token string) *ArgumentDefinition {
	return &ArgumentDefinition{
		CommonDefinition: CommonDefinition{
			Token: token,
			Help:  "app command",
		},
		Values: []string{"from:commands"},
	}
}

func addCommandAbbr(abbr string, cd *CommandDefinition, cmdByAbbr func(string) *CommandDefinition) {
	cmd := cmdByAbbr(abbr)
	if cmd == nil && cd != nil {
		cd.Abbr = abbr
	}
}

func addArgAbbr(abbr string, ad *ArgumentDefinition, argByAbbr func(string) *ArgumentDefinition) {
	arg := argByAbbr(abbr)
	if arg == nil && ad != nil {
		ad.Abbr = abbr
	}
}

func addHelpCommand(token, abbr string, cmdByToken, cmdByArg func(string) *CommandDefinition) *CommandDefinition {
	if cmdByToken == nil {
		return nil
	}
	// check if definitions already have 'help' token specified
	helpCmd := cmdByToken(token)
	if helpCmd != nil {
		return nil
	}

	helpCmd = createHelpCommandDefinition()
	addCommandAbbr(abbr, helpCmd, cmdByArg)

	return helpCmd
}

func addHelpCommandArgument(token, abbr string, argByToken, argByAbbr func(string) *ArgumentDefinition) (*ArgumentDefinition, error) {
	if argByToken == nil {
		return nil, nil
	}
	// create argument
	helpCmdArg := argByToken(token)
	if helpCmdArg != nil {
		return nil, fmt.Errorf("cannot add 'help' command as argument '%s' already exists", token)
	}

	helpCmdArg = createHelpArgumentDefinition(token)
	addArgAbbr(abbr, helpCmdArg, argByAbbr)

	return helpCmdArg, nil
}

func tryAddHelpCommand(def *Definitions) error {
	if def == nil {
		return fmt.Errorf("cannot add 'help' command to nil definitions")
	}

	const (
		cmdToken = "help"
		cmdAbbr  = "h"
		argToken = "help:command"
		argAbbr  = "c"
	)

	addedHelpCmd := addHelpCommand(cmdToken, cmdAbbr, def.CommandByToken, def.CommandByAbbr)
	if addedHelpCmd == nil {
		// don't return error as we're assuming author knows what they're doing
		return nil
	}

	helpCmdArg, err := addHelpCommandArgument(argToken, argAbbr, def.ArgByToken, def.ArgByAbbr)
	if err != nil {
		return err
	}

	if helpCmdArg != nil {
		def.Arguments = append(def.Arguments, *helpCmdArg)
	}

	addedHelpCmd.Arguments = []string{argToken}
	def.Commands = append(def.Commands, *addedHelpCmd)

	return nil
}

func expandRefValues(args []ArgumentDefinition, commands []CommandDefinition) error {
	for i, ad := range args {
		if ad.Values != nil &&
			len(ad.Values) == 1 &&
			strings.HasPrefix(ad.Values[0], "from:") {
			source := strings.TrimPrefix(ad.Values[0], "from:")
			switch source {
			case "commands":
				args[i].Values = make([]string, 0)
				for _, cd := range commands {
					args[i].Values = append(args[i].Values, cd.Token)
				}
				return nil
			default:
				return fmt.Errorf("cannot expand values from an unknown source: '%s'", source)
			}
		}
	}
	return nil
}

func printHelp(cmd string, defs *Definitions) error {
	if defs == nil {
		return fmt.Errorf("cannot show help without definitions")
	}
	if cmd == "" {
		printAppHelp(defs)
	} else {
		printCmdHelp(cmd, defs)
	}
	return nil
}

func printAppIntro(defs *Definitions) {
	if defs == nil {
		return
	}
	fmt.Printf("%s - %s\n", defs.App, defs.Help)
}

func printAppUsage(defs *Definitions) {
	if defs == nil {
		return
	}
	fmt.Printf("Usage: %s command [arguments [values]]\n",
		defs.App)
}

func printAppCommands(defs *Definitions) {
	if defs == nil {
		return
	}
	fmt.Println("Commands:")
	for _, cmd := range defs.Commands {
		fmt.Printf("  %-"+strconv.Itoa(defs.CommandsPadding())+"s  %s",
			cmd.Token,
			cmd.Help)
		attrs := make([]string, 0)
		if cmd.Abbr != "" {
			attrs = append(attrs, fmt.Sprintf("Abbr:%s", cmd.Abbr))
		}
		if len(attrs) > 0 {
			fmt.Printf(" (%s)", strings.Join(attrs, ","))
		}
		fmt.Println()
	}
}

func printAppAttrsLegend() {
	fmt.Println("Commands:")
	fmt.Printf("  %-4s  %s\n", "Abbr", "abbreviation that can be used in place of a full token")
}

func printAppMoreInfoPrompt(defs *Definitions) {
	if defs == nil {
		return
	}

	fmt.Printf("Run '%s help [command]' for help on a specific command.\n",
		defs.App)
}

func printAppHelp(defs *Definitions) {
	printAppIntro(defs)
	fmt.Println()
	printAppUsage(defs)
	fmt.Println()
	printAppCommands(defs)
	fmt.Println()
	printAppAttrsLegend()
	fmt.Println()
	printAppMoreInfoPrompt(defs)
}

func printCmdUsage(cmd string, defs *Definitions) {
	if defs == nil {
		return
	}
	cmdUsage := fmt.Sprintf("Usage: %s %s ", defs.App, cmd)
	cd := defs.CommandByToken(cmd)
	if cd == nil {
		return
	}
	if len(cd.Arguments) > 0 {
		cmdUsage += "[<arguments>]"
	}
	fmt.Println(cmdUsage)
}

func printArgAttrs(cmd string, arg string, defs *Definitions) {
	if defs == nil {
		return
	}
	ad := defs.ArgByToken(arg)
	if ad == nil {
		return
	}
	attrs := make([]string, 0)
	if ad.Abbr != "" {
		attrs = append(attrs, fmt.Sprintf("Abbr:%s", ad.Abbr))
	}
	if ad.Env {
		envToken := argEnv(defs.EnvPrefix, cmd, arg)
		attrs = append(attrs, fmt.Sprintf("Env:%s", envToken))
	}
	if ad.Multiple {
		attrs = append(attrs, "Mult")
	}
	if len(attrs) > 0 {
		fmt.Printf(" (%s)", strings.Join(attrs, ", "))
	}
}

func printArgValues(cmd string, arg string, defs *Definitions) {
	if defs == nil {
		return
	}
	ad := defs.ArgByToken(arg)
	if ad == nil {
		return
	}
	if len(ad.Values) > 0 {
		ap := strconv.Itoa(defs.ArgumentsPadding(cmd))
		fmt.Printf("  %-"+ap+"s  supported values: %s\n",
			"",
			strings.Join(ad.Values, ", "))
	}
}

func printCmdArgDesc(cmd string, arg string, defs *Definitions) {
	if defs == nil {
		return
	}
	ad := defs.ArgByToken(arg)
	if ad == nil {
		return
	}
	fmt.Printf("  %-"+strconv.Itoa(defs.ArgumentsPadding(cmd))+"s  %s", arg, ad.Help)
}

func printCmdArgs(cmd string, defs *Definitions) {
	if defs == nil {
		return
	}
	cd := defs.CommandByToken(cmd)
	if cd == nil {
		return
	}
	fmt.Println("Arguments:")
	for _, arg := range cd.Arguments {
		printCmdArgDesc(cmd, arg, defs)
		printArgAttrs(cmd, arg, defs)
		fmt.Println()
		printArgValues(cmd, arg, defs)
	}
}

func printArgAttrsLegend() {
	fmt.Println("Arguments (attributes):")
	fmt.Printf("  %-4s  %s\n", "Def", "default argument - value(s) can be provided right "+
		"after a command without an argument token")
	fmt.Printf("  %-4s  %s\n", "Mult", "supports multiple values, that can be provided "+
		"in sequence or each with an argument token")
	fmt.Printf("  %-4s  %s\n", "Req", "required argument - app cannot "+
		"meaningfully run without a value")
	fmt.Printf("  %-4s  %s\n", "Env", "value can be provided with an "+
		"environment variable specified above")
	fmt.Printf("  %-4s  %s\n", "Abbr", "argument abbreviation that can be "+
		"used in place of a full token")
}

func printCmdHelp(cmd string, defs *Definitions) {
	printCmdUsage(cmd, defs)
	fmt.Println()
	printCmdArgs(cmd, defs)
	fmt.Println()
	printArgAttrsLegend()
}
