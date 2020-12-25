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

func printAppIntro(defs *Definitions, verbose bool) {
	appDesc := defs.Hint
	if verbose && defs.Desc != "" {
		appDesc = defs.Desc
	}
	fmt.Printf("%s - %s\n", defs.App, appDesc)
}

func printAppUsage(defs *Definitions) {
	fmt.Printf("Usage: %s command [<--arguments [<values>]>] [<--flags>]\n",
		defs.App)
}

func printAppCommands(defs *Definitions, verbose bool) {
	fmt.Println("Commands:")
	for _, cmd := range defs.Commands {
		cmdDesc := cmd.Hint
		if verbose && cmd.Desc != "" {
			cmdDesc = cmd.Desc
		}
		fmt.Printf("  %-"+defs.CommandsPadding()+"s  %s",
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
	fmt.Println("Flags:")
	for _, flg := range defs.Flags {
		flgDesc := flg.Hint
		if verbose && flg.Desc != "" {
			flgDesc = flg.Desc
		}
		fmt.Printf("  %-"+defs.FlagsPadding()+"s  %s",
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
	verbosePrompt := ""
	if !verbose {
		verbosePrompt = fmt.Sprintf(" or '%s help --verbose' for more general information", defs.App)
	}

	fmt.Printf("Run '%s help [command]' for help on a specific command%s.\n",
		defs.App,
		verbosePrompt)
}

func printAppHelp(defs *Definitions, verbose bool) error {
	printAppIntro(defs, verbose)
	fmt.Println()
	printAppUsage(defs)
	fmt.Println()
	printAppCommands(defs, verbose)
	fmt.Println()
	if len(defs.Flags) > 0 {
		printAppFlags(defs, verbose)
		fmt.Println()
	}
	if verbose {
		printAppAttrsLegend()
		fmt.Println()
	}
	printAppMoreInfoPrompt(defs, verbose)
	return nil
}

func printExampleHelp(ex *ExampleDefinition, cmd string, defs *Definitions) {
	fmt.Printf("  '%s %s", defs.App, cmd)
	for avi, argValues := range ex.ArgumentsValues {
		for arg, values := range argValues {
			fmt.Printf(" --%s ", arg)
			if len(values) == 0 {
				fmt.Printf("<%s>", arg)
				continue
			}
			fmt.Print(strings.Join(values, " "))
		}
		if avi == len(ex.ArgumentsValues)-1 {
			fmt.Print("'")
		}
	}
	fmt.Println(":", ex.Desc)
}

func printCmdUsage(cmd string, defs *Definitions) error {
	cmdUsage := fmt.Sprintf("Usage: %s %s ", defs.App, cmd)
	cd := defs.CommandByToken(cmd)
	if cd == nil {
		return fmt.Errorf("command token '%s' is not defined", cmd)
	}
	if len(cd.Arguments) > 0 {
		cmdUsage += "[<arguments>]"
	}
	fmt.Println(cmdUsage)
	return nil
}

func printArgAttrs(cmd string, arg string, defs *Definitions) error {
	ad := defs.ArgByToken(arg)
	if ad == nil {
		return fmt.Errorf("  %s: invalid argument token\n", arg)
	}
	attrs := make([]string, 0)
	if ad.Default {
		attrs = append(attrs, "Def")
	}
	if ad.Required {
		attrs = append(attrs, "Req")
	}
	if ad.Multiple {
		attrs = append(attrs, "Mult")
	}
	if ad.Env {
		envToken := fmt.Sprintf("%s_%s", strings.ToUpper(cmd), strings.ToUpper(arg))
		if defs.EnvPrefix != "" {
			envToken = fmt.Sprintf("%s_%s", strings.ToUpper(defs.EnvPrefix), envToken)
		}
		attrs = append(attrs, fmt.Sprintf("Env:%s", envToken))
	}
	if ad.Abbr != "" {
		attrs = append(attrs, fmt.Sprintf("Abbr:%s", ad.Abbr))
	}
	if len(attrs) > 0 {
		fmt.Printf(" (%s)", strings.Join(attrs, ","))
	}
	return nil
}

func printArgValues(cmd string, arg string, defs *Definitions, verbose bool) error {
	ad := defs.ArgByToken(arg)
	if ad == nil {
		return fmt.Errorf("  %s: invalid argument token\n", arg)
	}
	if len(ad.Values) > 0 {
		ap := defs.ArgumentsPadding(cmd)
		vd := defs.DefinedValue(ad.Values)
		valuesOrNewLine := ""
		if !vd {
			valuesOrNewLine = strings.Join(ad.Values, ",")
		}
		fmt.Printf("  %-"+ap+"s  supported values: %s\n",
			"",
			valuesOrNewLine)
		if vd {
			for _, vt := range ad.Values {
				vd := defs.ValueByToken(vt)
				if vd == nil {
					return fmt.Errorf("command '%s' argument '%s' has unknown value token '%s'", cmd, arg, vt)
				}
				valDesc := vd.Hint
				if verbose && vd.Desc != "" {
					valDesc = vd.Desc
				}
				fmt.Printf("  %-"+ap+"s  - %s: %s\n", "", vt, valDesc)
			}
		}
	}
	return nil
}

func printCmdArgDesc(cmd string, arg string, defs *Definitions, verbose bool) error {
	ad := defs.ArgByToken(arg)
	if ad == nil {
		return fmt.Errorf("%s: invalid argument token\n", arg)
	}
	argDesc := ad.Hint
	if verbose && ad.Desc != "" {
		argDesc = ad.Desc
	}
	fmt.Printf("  %-"+defs.ArgumentsPadding(cmd)+"s  %s", arg, argDesc)
	return nil
}

func printCmdArgs(cmd string, defs *Definitions, verbose bool) error {
	cd := defs.CommandByToken(cmd)
	if cd == nil {
		return fmt.Errorf("command token '%s' is not defined", cmd)
	}
	fmt.Println("Arguments:")
	for _, arg := range cd.Arguments {
		if err := printCmdArgDesc(cmd, arg, defs, verbose); err != nil {
			return err
		}
		if err := printArgAttrs(cmd, arg, defs); err != nil {
			return err
		}
		fmt.Println()
		if err := printArgValues(cmd, arg, defs, verbose); err != nil {
			return err
		}
	}
	return nil
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

func printCmdExamples(cmd string, defs *Definitions) error {
	cd := defs.CommandByToken(cmd)
	if cd == nil {
		return fmt.Errorf("command token '%s' is not defined", cmd)
	}
	fmt.Println("Examples:")
	for _, ex := range cd.Examples {
		printExampleHelp(&ex, cmd, defs)
	}
	return nil
}

func printCmdMoreInfoPrompt(cmd string, defs *Definitions) {
	fmt.Printf("Run '%s help %s --verbose' for more information, "+
		"examples and arguments (attributes).\n", defs.App, cmd)
}

func printCmdHelp(cmd string, defs *Definitions, verbose bool) error {
	if err := printCmdUsage(cmd, defs); err != nil {
		return err
	}
	fmt.Println()
	if err := printCmdArgs(cmd, defs, verbose); err != nil {
		return err
	}
	fmt.Println()
	if verbose {
		printArgAttrsLegend()
		fmt.Println()
		if err := printCmdExamples(cmd, defs); err != nil {
			return err
		}
	} else {
		printCmdMoreInfoPrompt(cmd, defs)
	}
	return nil
}
