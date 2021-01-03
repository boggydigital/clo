package internal

import "fmt"

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

func vFail(msg string, verbose bool) {
	if verbose {
		fmt.Printf("FAIL: %s\n", msg)
	}
}

func vPass(msg string, verbose bool) {
	if verbose {
		fmt.Printf("PASS: %s\n", msg)
	}
}

func cmdTokensAreNotEmpty(commands []CommandDefinition, v bool) error {
	msg := "command tokens are not empty"
	for i, c := range commands {
		if c.Token == "" {
			vFail(msg, v)
			return fmt.Errorf("command #%d has an empty token", i+1)
		}
	}
	vPass(msg, v)
	return nil
}

func differentCmdTokens(commands []CommandDefinition, v bool) error {
	msg := "command tokens are different"
	dupeCmd := make([]string, 0)
	for _, c := range commands {
		dupeCmd = append(dupeCmd, c.Token)
	}
	if df := firstDupe(dupeCmd); df != "" {
		vFail(msg, v)
		return fmt.Errorf("commands have duplicate token: '%v'", df)
	}
	vPass(msg, v)
	return nil
}

func differentCmdAbbr(commands []CommandDefinition, v bool) error {
	msg := "command abbreviations are different"
	dupeCmd := make([]string, 0)
	for _, c := range commands {
		dupeCmd = append(dupeCmd, c.Abbr)
	}
	if df := firstDupe(dupeCmd); df != "" {
		vFail(msg, v)
		return fmt.Errorf("commands have duplicate abbreviation: '%v'", df)
	}
	vPass(msg, v)
	return nil
}

func argTokensAreNotEmpty(args []ArgumentDefinition, v bool) error {
	msg := "argument tokens are not empty"
	for i, a := range args {
		if a.Token == "" {
			vFail(msg, v)
			return fmt.Errorf("argument #%d has an empty token", i+1)
		}
	}
	vPass(msg, v)
	return nil
}

func differentArgTokens(args []ArgumentDefinition, v bool) error {
	msg := "argument tokens are different"
	dupeArgs := make([]string, 0)
	for _, a := range args {
		dupeArgs = append(dupeArgs, a.Token)
	}
	if df := firstDupe(dupeArgs); df != "" {
		vFail(msg, v)
		return fmt.Errorf("arguments have duplicate token: '%v'", df)
	}
	vPass(msg, v)
	return nil
}

func differentArgAbbr(args []ArgumentDefinition, v bool) error {
	msg := "argument abbreviations are different"
	dupeArgs := make([]string, 0)
	for _, a := range args {
		dupeArgs = append(dupeArgs, a.Abbr)
	}
	if df := firstDupe(dupeArgs); df != "" {
		vFail(msg, v)
		return fmt.Errorf("arguments have duplicate abbreviation: '%v'", df)
	}
	vPass(msg, v)
	return nil
}

func flagTokensAreNotEmpty(flags []FlagDefinition, v bool) error {
	msg := "flag tokens are not empty"
	for i, f := range flags {
		if f.Token == "" {
			vFail(msg, v)
			return fmt.Errorf("flag #%d has an empty token", i+1)
		}
	}
	vPass(msg, v)
	return nil
}

func differentFlagTokens(flags []FlagDefinition, v bool) error {
	msg := "flag tokens are different"
	dupeFlags := make([]string, 0)
	for _, f := range flags {
		dupeFlags = append(dupeFlags, f.Token)
	}
	if df := firstDupe(dupeFlags); df != "" {
		vFail(msg, v)
		return fmt.Errorf("flags have duplicate token: '%v'", df)
	}
	vPass(msg, v)
	return nil
}

func differentFlagAbbr(flags []FlagDefinition, v bool) error {
	msg := "flag abbreviations are different"
	dupeFlags := make([]string, 0)
	for _, f := range flags {
		dupeFlags = append(dupeFlags, f.Abbr)
	}
	if df := firstDupe(dupeFlags); df != "" {
		vFail(msg, v)
		return fmt.Errorf("flags have duplicate abbreviation: '%v'", df)
	}
	vPass(msg, v)
	return nil
}

func differentAbbr(
	commands []CommandDefinition,
	args []ArgumentDefinition,
	flags []FlagDefinition,
	v bool) error {
	msg := "all abbreviations are different"
	abbr := make([]string, 0)

	for _, c := range commands {
		abbr = append(abbr, c.Abbr)
	}
	for _, a := range args {
		abbr = append(abbr, a.Abbr)
	}
	for _, f := range flags {
		abbr = append(abbr, f.Abbr)
	}
	if da := firstDupe(abbr); da != "" {
		vFail(msg, v)
		return fmt.Errorf("same abbreviation for a command, argument or flag: '%v'", da)
	}
	vPass(msg, v)
	return nil
}

func commandsValidArgs(
	commands []CommandDefinition,
	argByToken func(string) *ArgumentDefinition,
	v bool) error {
	msg := "commands have valid arguments"
	for _, c := range commands {
		for _, a := range c.Arguments {
			da := argByToken(a)
			if da == nil {
				vFail(msg, v)
				return fmt.Errorf("command '%s' has undefined argument '%s'", c.Token, a)
			}
		}
	}
	vPass(msg, v)
	return nil
}

func allUsedArgs(commands []CommandDefinition, args []ArgumentDefinition, v bool) error {
	msg := "all arguments are used in commands"
	cas := make([]string, 0)
	for _, c := range commands {
		for _, ca := range c.Arguments {
			cas = append(cas, ca)
		}
	}
	for _, a := range args {
		match := false
		for _, da := range cas {
			if a.Token == da {
				match = true
				break
			}
		}
		if !match {
			vFail(msg, v)
			return fmt.Errorf("argument '%s' is not used in any command", a.Token)
		}
	}
	vPass(msg, v)
	return nil
}

func differentArgsCmd(commands []CommandDefinition, v bool) error {
	msg := "no duplicate arguments in commands"
	for _, c := range commands {
		if da := firstDupe(c.Arguments); da != "" {
			vFail(msg, v)
			return fmt.Errorf("command '%s' has duplicate argument '%s'", c.Token, da)
		}
	}
	vPass(msg, v)
	return nil
}

func singleDefaultArgPerCmd(
	commands []CommandDefinition,
	argByToken func(string) *ArgumentDefinition,
	v bool) error {
	msg := "no more than one default argument"
	for _, cmd := range commands {
		d := ""
		for _, at := range cmd.Arguments {
			arg := argByToken(at)
			if arg == nil {
				continue
			}
			if !arg.Default {
				continue
			}
			if d != "" {
				vFail(msg, v)
				return fmt.Errorf("'%s' has more than one default argument: '%s' and '%s'", cmd.Token, arg.Token, d)
			}
			d = arg.Token
		}
	}
	vPass(msg, v)
	return nil
}

func differentArgValues(args []ArgumentDefinition, v bool) error {
	msg := "no duplicate values in arguments"
	for _, a := range args {
		if dv := firstDupe(a.Values); dv != "" {
			vFail(msg, v)
			return fmt.Errorf("argument '%s' has duplicate value '%s'", a.Token, dv)
		}
	}
	vPass(msg, v)
	return nil
}

func examplesDescAreNotEmpty(commands []CommandDefinition, v bool) error {
	msg := "example descriptions are not empty"
	for _, cd := range commands {
		for i, ex := range cd.Examples {
			if ex.Desc == "" {
				vFail(msg, v)
				return fmt.Errorf("command '%s' example #%d doesn't have a description", cd.Token, i+1)
			}
		}
	}
	vPass(msg, v)
	return nil
}

func examplesArgumentsAreNotEmpty(commands []CommandDefinition, v bool) error {
	msg := "examples arguments are not empty"
	for _, cd := range commands {
		for i, ex := range cd.Examples {
			for arg := range ex.ArgumentsValues {
				if arg == "" {
					vFail(msg, v)
					return fmt.Errorf("command '%s' example #%d has an empty argument",
						cd.Token,
						i+1)
				}
			}
		}
	}
	vPass(msg, v)
	return nil
}

func examplesHaveArgsValues(commands []CommandDefinition, v bool) error {
	msg := "examples have at least one argument (with optional value(s))"
	for _, cd := range commands {
		for i, ex := range cd.Examples {
			if len(ex.ArgumentsValues) == 0 {
				vFail(msg, v)
				return fmt.Errorf("command '%s' example #%d doesn't have arguments defined",
					cd.Token,
					i+1)
			}
		}
	}
	vPass(msg, v)
	return nil
}

func examplesArgumentsAreValid(
	commands []CommandDefinition,
	argByToken func(string) *ArgumentDefinition,
	v bool) error {
	msg := "examples are using valid arguments"
	for _, cd := range commands {
		for i, ex := range cd.Examples {
			for arg := range ex.ArgumentsValues {
				ad := argByToken(arg)
				if ad == nil {
					vFail(msg, v)
					return fmt.Errorf("command '%s' example #%d uses undefined argument '%s'",
						cd.Token,
						i+1,
						arg)
				}
			}
		}
	}
	vPass(msg, v)
	return nil
}

func cmdExampleHasValidValues(
	cmd string,
	argumentsValues map[string][]string,
	validArgVal func(string, string) bool,
	i int) error {
	for arg, values := range argumentsValues {
		for _, val := range values {
			if !validArgVal(arg, val) {
				return fmt.Errorf("command '%s' example #%d uses invalid "+
					"value '%s' for an argument '%s'",
					cmd,
					i+1,
					val,
					arg)
			}
		}
	}
	return nil
}

func examplesHaveValidValues(
	commands []CommandDefinition,
	validArgVal func(string, string) bool,
	v bool) error {
	msg := "examples have valid values"
	for _, cd := range commands {
		for i, ex := range cd.Examples {
			if err := cmdExampleHasValidValues(
				cd.Token,
				ex.ArgumentsValues,
				validArgVal,
				i); err != nil {
				vFail(msg, v)
				return err
			}
		}
	}
	vPass(msg, v)
	return nil
}

func appendError(errors []error, err error) []error {
	if err != nil {
		return append(errors, err)
	}
	return errors
}

func (def *Definitions) Verify(v bool) []error {

	errors := make([]error, 0)

	// tokens and abbreviations
	errors = appendError(errors, cmdTokensAreNotEmpty(def.Commands, v))
	errors = appendError(errors, differentCmdTokens(def.Commands, v))
	errors = appendError(errors, differentCmdAbbr(def.Commands, v))
	errors = appendError(errors, argTokensAreNotEmpty(def.Arguments, v))
	errors = appendError(errors, differentArgTokens(def.Arguments, v))
	errors = appendError(errors, differentArgAbbr(def.Arguments, v))
	errors = appendError(errors, flagTokensAreNotEmpty(def.Flags, v))
	errors = appendError(errors, differentFlagTokens(def.Flags, v))
	errors = appendError(errors, differentFlagAbbr(def.Flags, v))
	errors = appendError(errors, differentAbbr(def.Commands, def.Arguments, def.Flags, v))

	// arguments
	errors = appendError(errors, commandsValidArgs(def.Commands, def.ArgByToken, v))
	errors = appendError(errors, allUsedArgs(def.Commands, def.Arguments, v))
	errors = appendError(errors, singleDefaultArgPerCmd(def.Commands, def.ArgByToken, v))
	errors = appendError(errors, differentArgsCmd(def.Commands, v))
	errors = appendError(errors, differentArgValues(def.Arguments, v))

	// examples
	errors = appendError(errors, examplesDescAreNotEmpty(def.Commands, v))
	errors = appendError(errors, examplesArgumentsAreValid(def.Commands, def.ArgByToken, v))
	errors = appendError(errors, examplesArgumentsAreNotEmpty(def.Commands, v))
	errors = appendError(errors, examplesHaveArgsValues(def.Commands, v))
	errors = appendError(errors, examplesHaveValidValues(def.Commands, def.ValidArgVal, v))

	return errors
}
