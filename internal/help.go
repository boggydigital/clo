package internal

import (
	"fmt"
	"strings"
)

func addHelpCmd(def *Definitions) {
	if def == nil {
		return
	}

	commands := make([]string, 0, len(def.Cmd))
	for c, _ := range def.Cmd {
		commands = append(commands, trimAttrs(c))
	}

	if _, ok := def.Cmd["help"]; !ok {
		def.Cmd["help"] = []string{
			fmt.Sprintf("%s%s=%s",
				"command", defaultAttr,
				strings.Join(commands, ",")),
		}
	}
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
	//fmt.Printf("%s - %s\n", defs.App, defs.Help)
}

func printAppUsage(defs *Definitions) {
	if defs == nil {
		return
	}
	//fmt.Printf("Usage: %s command [arguments [values]]\n",
	//	defs.App)
}

func printAppCommands(defs *Definitions) {
	if defs == nil {
		return
	}
	fmt.Println("Commands:")
	//for _, cmd := range defs.Commands {
	//fmt.Printf("  %-"+strconv.Itoa(defs.CommandsPadding())+"s  %s",
	//	cmd.Token,
	//	cmd.Help)
	//fmt.Println()
	//}
}

func printAppMoreInfoPrompt(defs *Definitions) {
	if defs == nil {
		return
	}

	//fmt.Printf("Run '%s help [command]' for help on a specific command.\n",
	//	defs.App)
}

func printAppHelp(defs *Definitions) {
	printAppIntro(defs)
	fmt.Println()
	printAppUsage(defs)
	fmt.Println()
	printAppCommands(defs)
	fmt.Println()
	printAppMoreInfoPrompt(defs)
}

func printCmdUsage(cmd string, defs *Definitions) {
	//if defs == nil {
	//	return
	//}
	//cmdUsage := fmt.Sprintf("Usage: %s %s ", defs.App, cmd)
	//cd := defs.CommandByToken(cmd)
	//if cd == nil {
	//	return
	//}
	//if len(cd.Arguments) > 0 {
	//	cmdUsage += "[<arguments>]"
	//}
	//fmt.Println(cmdUsage)
}

func printArgAttrs(cmd string, arg string, defs *Definitions) {
	//if defs == nil {
	//	return
	//}
	//ad := defs.ArgByToken(arg)
	//if ad == nil {
	//	return
	//}
	//attrs := make([]string, 0)
	//if ad.Env {
	//	//envToken := argEnv(defs.EnvPrefix, cmd, arg)
	//	//attrs = append(attrs, fmt.Sprintf("Env:%s", envToken))
	//}
	//if ad.Multiple {
	//	attrs = append(attrs, "Mult")
	//}
	//if len(attrs) > 0 {
	//	fmt.Printf(" (%s)", strings.Join(attrs, ", "))
	//}
}

func printArgValues(cmd string, arg string, defs *Definitions) {
	//if defs == nil {
	//	return
	//}
	//ad := defs.ArgByToken(arg)
	//if ad == nil {
	//	return
	//}
	//if len(ad.Values) > 0 {
	//	ap := strconv.Itoa(defs.ArgumentsPadding(cmd))
	//	fmt.Printf("  %-"+ap+"s  supported values: %s\n",
	//		"",
	//		strings.Join(ad.Values, ", "))
	//}
}

func printCmdArgDesc(cmd string, arg string, defs *Definitions) {
	//if defs == nil {
	//	return
	//}
	//ad := defs.ArgByToken(arg)
	//if ad == nil {
	//	return
	//}
	//fmt.Printf("  %-"+strconv.Itoa(defs.ArgumentsPadding(cmd))+"s  %s", arg, ad.Help)
}

func printCmdArgs(cmd string, defs *Definitions) {
	//if defs == nil {
	//	return
	//}
	//cd := defs.CommandByToken(cmd)
	//if cd == nil {
	//	return
	//}
	//fmt.Println("Arguments:")
	//for _, arg := range cd.Arguments {
	//	printCmdArgDesc(cmd, arg, defs)
	//	printArgAttrs(cmd, arg, defs)
	//	fmt.Println()
	//	printArgValues(cmd, arg, defs)
	//}
}

//func printArgAttrsLegend() {
//	fmt.Println("Arguments (attributes):")
//	fmt.Printf("  %-4s  %s\n", "Def", "default argument - value(s) can be provided right "+
//		"after a command without an argument token")
//	fmt.Printf("  %-4s  %s\n", "Mult", "supports multiple values, that can be provided "+
//		"in sequence or each with an argument token")
//	fmt.Printf("  %-4s  %s\n", "Req", "required argument - app cannot "+
//		"meaningfully run without a value")
//	fmt.Printf("  %-4s  %s\n", "Env", "value can be provided with an "+
//		"environment variable specified above")
//	fmt.Printf("  %-4s  %s\n", "Abbr", "argument abbreviation that can be "+
//		"used in place of a full token")
//}

func printCmdHelp(cmd string, defs *Definitions) {
	printCmdUsage(cmd, defs)
	fmt.Println()
	printCmdArgs(cmd, defs)
	fmt.Println()
	//printArgAttrsLegend()
}
