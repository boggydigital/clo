package clo

import (
	"fmt"
	"strconv"
	"strings"
)

func (defs *Definitions) getHelp(topics []string) string {
	if defs == nil || defs.Help == nil {
		return ""
	}
	for ; len(topics) > 0; topics = topics[1:] {
		key := strings.Join(transform(topics, trimAttrs), ":")
		if value, ok := defs.Help[key]; ok {
			return value
		}
	}
	return ""
}

func addInternalHelpCmd(defs *Definitions) {
	if defs == nil {
		return
	}

	commands := make([]string, 0, len(defs.Cmd))
	for c, _ := range defs.Cmd {
		commands = append(commands, trimAttrs(c))
	}

	if _, ok := defs.Cmd["help"]; !ok {
		defs.Cmd["help"] = []string{
			fmt.Sprintf("%s%s=%s",
				"command", defaultAttr,
				strings.Join(commands, ",")),
		}
	}

	// add help topics for help command and arguments
	if defs.Help == nil {
		defs.Help = make(map[string]string, 0)
	}

	if _, ok := defs.Help["help"]; !ok {
		defs.Help["help"] = "display app help"
	}

	if _, ok := defs.Help["help:command"]; !ok {
		defs.Help["help:command"] = "display app command help"
	}
}

func printHelp(cmd string, defs *Definitions) error {
	if defs == nil {
		return fmt.Errorf("cannot show help for nil definitions")
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
	appIntro := appName()
	appHelp := defs.getHelp([]string{appIntro})
	if appHelp != "" {
		appIntro = fmt.Sprintf("%s - %s\n", appIntro, appHelp)
	}
	fmt.Print(appIntro)
}

func printAppUsage(defs *Definitions) {
	if defs == nil {
		return
	}
	fmt.Printf("Usage: %s command [arguments [values]]\n",
		appName())
}

func printAppCommands(defs *Definitions) {
	if defs == nil {
		return
	}
	fmt.Println("Commands:")
	for cmd := range defs.Cmd {
		tc := trimAttrs(cmd)
		cmdLine := fmt.Sprintf("  %-"+strconv.Itoa(defs.cmdPadding())+"s", tc)
		if cmdHelp, ok := defs.Help[tc]; ok {
			cmdLine = fmt.Sprintf("%s  %s", cmdLine, cmdHelp)
		}
		fmt.Println(cmdLine)
	}
}

func printAppMoreInfoPrompt(defs *Definitions) {
	if defs == nil {
		return
	}

	fmt.Printf("Run '%s help [command]' for help on a specific command.\n",
		appName())
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
	if defs == nil {
		return
	}
	cmdUsage := fmt.Sprintf("Usage: %s %s ", appName(), cmd)
	dc := defs.definedCmd(cmd)
	if dc == "" {
		return
	}
	if len(defs.Cmd[dc]) > 0 {
		// TODO: print actual arguments
		cmdUsage += "[<arguments>]"
	}
	fmt.Println(cmdUsage)
}

func printArgValues(cmd string, arg string, defs *Definitions) {
	if defs == nil {
		return
	}
	_, da := defs.definedCmdArg(cmd, arg)
	if da == "" {
		return
	}
	_, values := splitArgValues(da)
	if len(values) > 0 {
		ap := strconv.Itoa(defs.argPadding(cmd))
		singularOrPlural := "value"
		if len(values) > 1 {
			singularOrPlural = "values"
		}
		fmt.Printf("  %-"+ap+"s  supported %s: %s\n",
			"",
			singularOrPlural,
			strings.Join(transform(values, trimAttrs), ", "))
	}
}

func printCmdArgDesc(cmd string, arg string, defs *Definitions) {
	if defs == nil {
		return
	}
	_, da := defs.definedCmdArg(cmd, arg)
	if da == "" {
		return
	}
	fmt.Printf("  %-"+strconv.Itoa(defs.argPadding(cmd))+"s  %s",
		trimAttrs(da),
		defs.getHelp([]string{cmd, arg}))
}

func printCmdArgs(cmd string, defs *Definitions) {
	if defs == nil {
		return
	}
	dc := defs.definedCmd(cmd)
	if dc == "" {
		return
	}
	if len(defs.Cmd[dc]) > 0 {
		fmt.Println("Arguments:")
		for _, arg := range defs.Cmd[dc] {
			printCmdArgDesc(cmd, arg, defs)
			fmt.Println()
			printArgValues(cmd, arg, defs)
		}
	}
}

func printCmdHelp(cmd string, defs *Definitions) {
	printCmdUsage(cmd, defs)
	fmt.Println()
	printCmdArgs(cmd, defs)
	fmt.Println()
}
