package clo

import (
	"fmt"
	"sort"
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
	for c := range defs.Cmd {
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
		defs.Help = make(map[string]string)
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
		if err := printCmdHelp(cmd, defs); err != nil {
			return err
		}
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

	// print cmds as a sorted list
	sortedCmds := make([]string, 0, len(defs.Cmd))
	for cmd := range defs.Cmd {
		sortedCmds = append(sortedCmds, cmd)
	}
	sort.Strings(sortedCmds)

	for _, cmd := range sortedCmds {
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

func printCmdUsage(cmd string, defs *Definitions) error {
	if defs == nil {
		return fmt.Errorf("clo: can't print command usage for nil defintions")
	}
	cmdUsage := fmt.Sprintf("Usage: %s %s ", appName(), cmd)
	dc, err := defs.definedCmd(cmd)
	if err != nil {
		return err
	}
	if dc == "" {
		return nil
	}
	if len(defs.Cmd[dc]) > 0 {
		// TODO: print actual arguments
		cmdUsage += "[<arguments>]"
	}
	fmt.Println(cmdUsage)
	return nil
}

func printArgValues(cmd string, arg string, defs *Definitions) error {
	if defs == nil {
		return fmt.Errorf("clo: can't print argument values for nil definitions")
	}
	da, err := defs.definedArg(cmd, arg)
	if err != nil {
		return err
	}
	if da == "" {
		return nil
	}
	_, values := splitArgValues(da)
	if len(values) > 0 {
		argPadding, err := defs.argPadding(cmd)
		if err != nil {
			return err
		}
		ap := strconv.Itoa(argPadding)
		singularOrPlural := "value"
		if len(values) > 1 {
			singularOrPlural = "values"
		}
		fmt.Printf("  %-"+ap+"s  supported %s: %s\n",
			"",
			singularOrPlural,
			strings.Join(transform(values, trimAttrs), ", "))
	}
	return nil
}

func printCmdArgDesc(cmd string, arg string, defs *Definitions) error {
	if defs == nil {
		return fmt.Errorf("clo: can't print command argument description for nil defintions")
	}
	da, err := defs.definedArg(cmd, arg)
	if err != nil {
		return err
	}
	if da == "" {
		return nil
	}
	argPadding, err := defs.argPadding(cmd)
	if err != nil {
		return err
	}
	fmt.Printf("  %-"+strconv.Itoa(argPadding)+"s  %s",
		trimAttrs(da),
		defs.getHelp([]string{cmd, arg}))
	return nil
}

func printCmdArgs(cmd string, defs *Definitions) error {
	if defs == nil {
		return fmt.Errorf("clo: can't print command args for nil defintions")
	}
	dc, err := defs.definedCmd(cmd)
	if err != nil {
		return err
	}
	if dc == "" {
		return nil
	}
	if len(defs.Cmd[dc]) > 0 {
		fmt.Println("Arguments:")
		// print cmd args as a sorted list
		sort.Strings(defs.Cmd[dc])
		for _, arg := range defs.Cmd[dc] {
			if err := printCmdArgDesc(cmd, arg, defs); err != nil {
				return err
			}
			fmt.Println()
			if err := printArgValues(cmd, arg, defs); err != nil {
				return err
			}
		}
	}
	return nil
}

func printCmdHelp(cmd string, defs *Definitions) error {
	if err := printCmdUsage(cmd, defs); err != nil {
		return err
	}
	fmt.Println()
	if err := printCmdArgs(cmd, defs); err != nil {
		return err
	}
	fmt.Println()
	return nil
}
