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
			Hint:  "display help",
			Desc:  "display help for the app or a specific command",
		},
	}
}

func createHelpArgumentDefinition(token string) *ArgumentDefinition {
	return &ArgumentDefinition{
		CommonDefinition: CommonDefinition{
			Token: token,
			Abbr:  "",
			Hint:  "app command",
		},
		Default:  true,
		Multiple: false,
		Required: false,
		Values:   []string{"from:commands"},
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

func printHelp(cmd string, defs *Definitions, verbose bool) error {
	if defs == nil {
		return fmt.Errorf("cannot show help without definitions")
	}
	if cmd == "" {
		printAppHelp(defs, verbose)
	} else {
		printCmdHelp(cmd, defs, verbose)
	}
	return nil
}

func printAppIntro(defs *Definitions, verbose bool) {
	if defs == nil {
		return
	}
	appDesc := defs.Hint
	if verbose && defs.Desc != "" {
		appDesc = defs.Desc
	}
	fmt.Printf("%s - %s\n", defs.App, appDesc)
}

func printAppUsage(defs *Definitions) {
	if defs == nil {
		return
	}
	fmt.Printf("Usage: %s command [<arguments [<values>]>] [<flags>]\n",
		defs.App)
}

func printAppCommands(defs *Definitions, verbose bool) {
	if defs == nil {
		return
	}
	fmt.Println("Commands:")
	for _, cmd := range defs.Commands {
		cmdDesc := cmd.Hint
		if verbose && cmd.Desc != "" {
			cmdDesc = cmd.Desc
		}
		fmt.Printf("  %-"+strconv.Itoa(defs.CommandsPadding())+"s  %s",
			cmd.Token,
			cmdDesc)
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

func printAppFlags(defs *Definitions, verbose bool) {
	if defs == nil {
		return
	}
	fmt.Println("Flags:")
	for _, flg := range defs.Flags {
		flgDesc := flg.Hint
		if verbose && flg.Desc != "" {
			flgDesc = flg.Desc
		}
		fmt.Printf("  %-"+strconv.Itoa(defs.FlagsPadding())+"s  %s",
			flg.Token,
			flgDesc)
		attrs := make([]string, 0)
		if flg.Abbr != "" {
			attrs = append(attrs, fmt.Sprintf("Abbr:%s", flg.Abbr))
		}
		if len(attrs) > 0 {
			fmt.Printf(" (%s)", strings.Join(attrs, ","))
		}
		fmt.Println()
	}
}

func printAppAttrsLegend() {
	fmt.Println("Commands, flags (attributes):")
	fmt.Printf("  %-4s  %s\n", "Abbr", "abbreviation that can be used in place of a full token")
}

func printAppMoreInfoPrompt(defs *Definitions, verbose bool) {
	if defs == nil {
		return
	}
	verbosePrompt := ""
	if !verbose {
		verbosePrompt = fmt.Sprintf(" or '%s help --verbose' for more help", defs.App)
	}

	fmt.Printf("Run '%s help [command]' for help on a specific command%s.\n",
		defs.App,
		verbosePrompt)
}

func printAppHelp(defs *Definitions, verbose bool) {
	printAppIntro(defs, verbose)
	fmt.Println()
	printAppUsage(defs)
	fmt.Println()
	printAppCommands(defs, verbose)
	fmt.Println()
	if defs != nil && len(defs.Flags) > 0 {
		printAppFlags(defs, verbose)
		fmt.Println()
	}
	if verbose {
		printAppAttrsLegend()
		fmt.Println()
	}
	printAppMoreInfoPrompt(defs, verbose)
}

func printExampleHelp(ex *ExampleDefinition, cmd string, defs *Definitions) {
	if defs == nil {
		return
	}
	fmt.Printf("  '%s %s", defs.App, cmd)
	for arg, values := range ex.ArgumentsValues {
		fmt.Printf(" --%s ", arg)
		if len(values) == 0 {
			fmt.Printf("<%s>", arg)
			continue
		}
		fmt.Print(strings.Join(values, " "))
	}
	fmt.Println("':", ex.Desc)
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
	if ad.Default {
		attrs = append(attrs, "Def")
	}
	if ad.Env {
		envToken := argEnv(defs.EnvPrefix, cmd, arg)
		attrs = append(attrs, fmt.Sprintf("Env:%s", envToken))
	}
	if ad.Multiple {
		attrs = append(attrs, "Mult")
	}
	if ad.Required {
		attrs = append(attrs, "Req")
	}
	if len(attrs) > 0 {
		fmt.Printf(" (%s)", strings.Join(attrs, ", "))
	}
}

func printArgValues(cmd string, arg string, defs *Definitions, verbose bool) {
	if defs == nil {
		return
	}
	ad := defs.ArgByToken(arg)
	if ad == nil {
		return
	}
	if len(ad.Values) > 0 {
		ap := strconv.Itoa(defs.ArgumentsPadding(cmd))
		vd := defs.DefinedValue(ad.Values)
		valuesOrNewLine := ""
		if !vd {
			valuesOrNewLine = strings.Join(ad.Values, ", ")
		}
		fmt.Printf("  %-"+ap+"s  supported values: %s\n",
			"",
			valuesOrNewLine)
		if vd {
			for _, vt := range ad.Values {
				vd := defs.ValueByToken(vt)
				if vd == nil {
					continue
					//return fmt.Errorf("command '%s' argument '%s' has unknown value token '%s'", cmd, arg, vt)
				}
				valDesc := vd.Hint
				if verbose && vd.Desc != "" {
					valDesc = vd.Desc
				}
				fmt.Printf("  %-"+ap+"s  - %s: %s\n", "", vt, valDesc)
			}
		}
	}
}

func printCmdArgDesc(cmd string, arg string, defs *Definitions, verbose bool) {
	if defs == nil {
		return
	}
	ad := defs.ArgByToken(arg)
	if ad == nil {
		return
	}
	argDesc := ad.Hint
	if verbose && ad.Desc != "" {
		argDesc = ad.Desc
	}
	fmt.Printf("  %-"+strconv.Itoa(defs.ArgumentsPadding(cmd))+"s  %s", arg, argDesc)
}

func printCmdArgs(cmd string, defs *Definitions, verbose bool) {
	if defs == nil {
		return
	}
	cd := defs.CommandByToken(cmd)
	if cd == nil {
		return
	}
	fmt.Println("Arguments:")
	for _, arg := range cd.Arguments {
		printCmdArgDesc(cmd, arg, defs, verbose)
		printArgAttrs(cmd, arg, defs)
		fmt.Println()
		printArgValues(cmd, arg, defs, verbose)
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

func printCmdExamples(cmd string, defs *Definitions) {
	if defs == nil {
		return
	}
	cd := defs.CommandByToken(cmd)
	if cd == nil {
		return
	}
	if len(cd.Examples) > 0 {
		fmt.Println("Examples:")
		for _, ex := range cd.Examples {
			printExampleHelp(&ex, cmd, defs)
		}
	}
}

func printCmdMoreInfoPrompt(cmd string, defs *Definitions) {
	if defs == nil {
		return
	}
	fmt.Printf("Run '%s help %s --verbose' for more information, "+
		"examples and arguments (attributes).\n", defs.App, cmd)
}

func printCmdHelp(cmd string, defs *Definitions, verbose bool) {
	printCmdUsage(cmd, defs)
	fmt.Println()
	printCmdArgs(cmd, defs, verbose)
	fmt.Println()
	if verbose {
		printArgAttrsLegend()
		fmt.Println()
		printCmdExamples(cmd, defs)
	} else {
		printCmdMoreInfoPrompt(cmd, defs)
	}
}
