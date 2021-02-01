package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func addInternalHelpCmd(def *Definitions) {
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

	// add help topics for help command and arguments

	if _, ok := def.Help["help"]; !ok {
		def.Help["help"] = "display app help"
	}

	if _, ok := def.Help["help:command"]; !ok {
		def.Help["help:command"] = "display app command help"
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
	appHelp := defs.help([]string{appIntro})
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
		fmt.Printf("  %-"+ap+"s  supported values: %s\n",
			"",
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
		defs.help([]string{cmd, arg}))
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
