package clo

import (
	"errors"
	"fmt"
)

func firstDupe(sl []string) string {
	if len(sl) < 2 {
		return ""
	}
	for i := 0; i < len(sl)-1; i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[i] == sl[j] {
				return sl[i]
			}
		}
	}
	return ""
}

func vFail(msg string, debug bool) {
	if debug {
		fmt.Printf("FAIL: %s\n", msg)
	}
}

func vPass(msg string, debug bool) {
	if debug {
		fmt.Printf("PASS: %s\n", msg)
	}
}

func noEmptyCmd(cmd map[string][]string, debug bool) error {
	msg := "no empty commands"
	for c := range cmd {
		if c == "" {
			vFail(msg, debug)
			return fmt.Errorf("found an empty command")
		}
	}
	vPass(msg, debug)
	return nil
}

func differentCmd(cmd map[string][]string, debug bool) error {
	msg := "commands are different"
	cmdKeys := make([]string, 0)
	for c := range cmd {
		cmdKeys = append(cmdKeys, c)
	}
	if df := firstDupe(cmdKeys); df != "" {
		vFail(msg, debug)
		return fmt.Errorf("duplicate commands: %s", df)
	}
	vPass(msg, debug)
	return nil
}

func differentCmdArgs(cmd map[string][]string, v bool) error {
	msg := "no duplicate command arguments"
	for c := range cmd {
		if da := firstDupe(cmd[c]); da != "" {
			vFail(msg, v)
			return fmt.Errorf("command %s has duplicate arguments: %s", c, da)
		}
	}
	vPass(msg, v)
	return nil
}

func noEmptyArgs(cmd string, args []string, debug bool) error {
	msg := fmt.Sprintf("no empty arguments for command %s", cmd)
	for _, a := range args {
		if a == "" {
			vFail(msg, debug)
			return fmt.Errorf("found an empty argument for command %s", cmd)
		}
	}
	vPass(msg, debug)
	return nil
}

func differentCmdArgValues(cmd string, args []string, v bool) error {
	msg := fmt.Sprintf("no duplicate arguments values for command %s", cmd)
	for _, a := range args {
		_, values := splitArgValues(a)
		if dv := firstDupe(values); dv != "" {
			vFail(msg, v)
			return fmt.Errorf("command %s argument %s has duplicate values: %s", cmd, a, dv)
		}
	}
	vPass(msg, v)
	return nil
}

func noEmptyHelpTopic(help map[string]string, v bool) error {
	msg := "no empty help topics"
	for t := range help {
		if t == "" {
			vFail(msg, v)
			return fmt.Errorf("help contains an empty topic")
		}
	}
	vPass(msg, v)
	return nil
}

func cmdHaveHelp(cmd map[string][]string, getHelp func([]string) string, v bool) error {
	msg := "commands have help topics"
	for c := range cmd {
		cHelp := getHelp([]string{c})
		if cHelp == "" {
			vFail(msg, v)
			return fmt.Errorf("command %s doesn't have help topic", c)
		}
	}
	vPass(msg, v)
	return nil
}

func noEmptyHelpMessages(help map[string]string, v bool) error {
	msg := "no empty help messages"
	for t := range help {
		if help[t] == "" {
			vFail(msg, v)
			return fmt.Errorf("help message for topic %s is empty", t)
		}
	}
	vPass(msg, v)
	return nil
}

func cmdArgsHaveHelp(cmd string, args []string, getHelp func([]string) string, v bool) error {
	msg := fmt.Sprintf("arguments have help topics for command %s", cmd)
	for _, arg := range args {
		ta := trimAttrs(arg)
		argHelp := getHelp([]string{cmd, ta})
		if argHelp == "" {
			vFail(msg, v)
			return fmt.Errorf("command %s argument %s doesn't have help topic", cmd, ta)
		}
	}
	vPass(msg, v)
	return nil
}

func appendErr(errors []error, err error) []error {
	if err != nil {
		return append(errors, err)
	}
	return errors
}

func (defs *Definitions) Validate(verbose bool) []error {

	if defs == nil {
		return []error{
			errors.New("can't validate nil definitions"),
		}
	}

	errs := make([]error, 0)

	// commands
	errs = appendErr(errs, noEmptyCmd(defs.Cmd, verbose))
	errs = appendErr(errs, differentCmd(defs.Cmd, verbose))
	errs = appendErr(errs, differentCmdArgs(defs.Cmd, verbose))
	errs = appendErr(errs, cmdHaveHelp(defs.Cmd, defs.getHelp, verbose))

	// arguments
	for c := range defs.Cmd {
		tc, args := trimAttrs(c), defs.Cmd[c]
		errs = appendErr(errs, noEmptyArgs(tc, args, verbose))
		errs = appendErr(errs, differentCmdArgValues(tc, args, verbose))
		errs = appendErr(errs, cmdArgsHaveHelp(tc, args, defs.getHelp, verbose))
	}

	// help
	errs = appendErr(errs, noEmptyHelpTopic(defs.Help, verbose))
	errs = appendErr(errs, noEmptyHelpMessages(defs.Help, verbose))

	return errs
}
