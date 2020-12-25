package internal

import "fmt"

func sliceFirstDupe(sl []string) string {
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

func cmdTokensAreNotEmpty(def *Definitions, v bool) error {
	msg := "command tokens are not empty"
	for i, c := range def.Commands {
		if c.Token == "" {
			vFail(msg, v)
			return fmt.Errorf("command #%d has an empty token", i+1)
		}
	}
	vPass(msg, v)
	return nil
}

func differentCmdTokens(def *Definitions, v bool) error {
	msg := "command tokens are different"
	cmds := make([]string, 0)
	for _, c := range def.Commands {
		cmds = append(cmds, c.Token)
	}
	if df := sliceFirstDupe(cmds); df != "" {
		vFail(msg, v)
		return fmt.Errorf("commands have duplicate token: '%v'", df)
	}
	vPass(msg, v)
	return nil
}

func differentCmdAbbr(def *Definitions, v bool) error {
	msg := "command abbreviations are different"
	cmds := make([]string, 0)
	for _, c := range def.Commands {
		cmds = append(cmds, c.Abbr)
	}
	if df := sliceFirstDupe(cmds); df != "" {
		vFail(msg, v)
		return fmt.Errorf("commands have duplicate abbreviation: '%v'", df)
	}
	vPass(msg, v)
	return nil
}

func differentArgTokens(def *Definitions, v bool) error {
	msg := "argument tokens are different"
	args := make([]string, 0)
	for _, a := range def.Arguments {
		args = append(args, a.Token)
	}
	if df := sliceFirstDupe(args); df != "" {
		vFail(msg, v)
		return fmt.Errorf("arguments have duplicate token: '%v'", df)
	}
	vPass(msg, v)
	return nil
}

func argTokensAreNotEmpty(def *Definitions, v bool) error {
	msg := "argument tokens are not empty"
	for i, a := range def.Arguments {
		if a.Token == "" {
			vFail(msg, v)
			return fmt.Errorf("argument #%d has an empty token", i+1)
		}
	}
	vPass(msg, v)
	return nil
}

func differentArgAbbr(def *Definitions, v bool) error {
	msg := "argument abbreviations are different"
	args := make([]string, 0)
	for _, a := range def.Arguments {
		args = append(args, a.Abbr)
	}
	if df := sliceFirstDupe(args); df != "" {
		vFail(msg, v)
		return fmt.Errorf("arguments have duplicate abbreviation: '%v'", df)
	}
	vPass(msg, v)
	return nil
}

func differentFlagTokens(def *Definitions, v bool) error {
	msg := "flag tokens are different"
	flags := make([]string, 0)
	for _, f := range def.Flags {
		flags = append(flags, f.Token)
	}
	if df := sliceFirstDupe(flags); df != "" {
		vFail(msg, v)
		return fmt.Errorf("flags have duplicate token: '%v'", df)
	}
	vPass(msg, v)
	return nil
}

func differentFlagAbbr(def *Definitions, v bool) error {
	msg := "flag abbreviations are different"
	flags := make([]string, 0)
	for _, f := range def.Flags {
		flags = append(flags, f.Abbr)
	}
	if df := sliceFirstDupe(flags); df != "" {
		vFail(msg, v)
		return fmt.Errorf("flags have duplicate abbreviation: '%v'", df)
	}
	vPass(msg, v)
	return nil
}

func flagTokensAreNotEmpty(def *Definitions, v bool) error {
	msg := "flag tokens are not empty"
	for i, f := range def.Flags {
		if f.Token == "" {
			vFail(msg, v)
			return fmt.Errorf("flag #%d has an empty token", i+1)
		}
	}
	vPass(msg, v)
	return nil
}

func differentAbbr(def *Definitions, v bool) error {
	msg := "all abbreviations are different"
	abbr := make([]string, 0)

	for _, c := range def.Commands {
		abbr = append(abbr, c.Abbr)
	}
	for _, a := range def.Arguments {
		abbr = append(abbr, a.Abbr)
	}
	for _, f := range def.Flags {
		abbr = append(abbr, f.Abbr)
	}
	if da := sliceFirstDupe(abbr); da != "" {
		vFail(msg, v)
		return fmt.Errorf("same abbreviation for a command, argument or flag: '%v'", da)
	}
	vPass(msg, v)
	return nil
}

func commandsValidArgs(def *Definitions, v bool) error {
	msg := "commands have valid arguments"
	for _, c := range def.Commands {
		for _, a := range c.Arguments {
			da := def.ArgByToken(a)
			if da == nil {
				vFail(msg, v)
				return fmt.Errorf("command '%s' has undefined argument '%s'", c.Token, a)
			}
		}
	}
	vPass(msg, v)
	return nil
}

func allUsedArgs(def *Definitions, v bool) error {
	msg := "all arguments are used in commands"
	cas := make([]string, 0)
	for _, c := range def.Commands {
		for _, ca := range c.Arguments {
			cas = append(cas, ca)
		}
	}
	for _, a := range def.Arguments {
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

func differentArgsCmd(def *Definitions, v bool) error {
	msg := "no duplicate arguments in commands"
	for _, c := range def.Commands {
		if da := sliceFirstDupe(c.Arguments); da != "" {
			vFail(msg, v)
			return fmt.Errorf("command '%s' has duplicate argument '%s'", c.Token, da)
		}
	}
	vPass(msg, v)
	return nil
}

func singleDefaultArgPerCmd(def *Definitions, v bool) error {
	msg := "no more than one default argument"
	for _, cmd := range def.Commands {
		d := ""
		for _, at := range cmd.Arguments {
			arg := def.ArgByToken(at)
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

func differentArgValues(def *Definitions, v bool) error {
	msg := "no duplicate values in arguments"
	for _, a := range def.Arguments {
		if dv := sliceFirstDupe(a.Values); dv != "" {
			vFail(msg, v)
			return fmt.Errorf("argument '%s' has duplicate value '%s'", a.Token, dv)
		}
	}
	vPass(msg, v)
	return nil
}

func examplesDescAreNotEmpty(def *Definitions, v bool) error {
	msg := "example descriptions are not empty"
	for _, cd := range def.Commands {
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

func examplesArgumentsAreNotEmpty(def *Definitions, v bool) error {
	msg := "examples arguments are not empty"
	for _, cd := range def.Commands {
		for i, ex := range cd.Examples {
			for _, argVal := range ex.ArgumentsValues {
				for arg, _ := range argVal {
					if arg == "" {
						vFail(msg, v)
						return fmt.Errorf("command '%s' example #%d has an empty argument",
							cd.Token,
							i+1)
					}
				}
			}
		}
	}
	vPass(msg, v)
	return nil
}

func examplesHaveArgsValues(def *Definitions, v bool) error {
	msg := "examples have at least one argument (with optional value(s))"
	for _, cd := range def.Commands {
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

func examplesArgumentsAreValid(def *Definitions, v bool) error {
	msg := "examples are using valid arguments"
	for _, cd := range def.Commands {
		for i, ex := range cd.Examples {
			for _, argVal := range ex.ArgumentsValues {
				for arg, _ := range argVal {
					ad := def.ArgByToken(arg)
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
	}
	vPass(msg, v)
	return nil
}

func cmdExampleHasValidValues(cmd string, def *Definitions, ex *ExampleDefinition, i int) error {
	for _, argVal := range ex.ArgumentsValues {
		for arg, values := range argVal {
			for _, val := range values {
				if !def.ValidArgVal(val, arg) {
					return fmt.Errorf("command '%s' example #%d uses invalid "+
						"value '%s' for an argument '%s'",
						cmd,
						i+1,
						val,
						arg)
				}
			}
		}
	}
	return nil
}

func examplesHaveValidValues(def *Definitions, v bool) error {
	msg := "examples have valid values"
	for _, cd := range def.Commands {
		for i, ex := range cd.Examples {
			if err := cmdExampleHasValidValues(cd.Token, def, &ex, i); err != nil {
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
	errors = appendError(errors, cmdTokensAreNotEmpty(def, v))
	errors = appendError(errors, differentCmdTokens(def, v))
	errors = appendError(errors, differentCmdAbbr(def, v))
	errors = appendError(errors, argTokensAreNotEmpty(def, v))
	errors = appendError(errors, differentArgTokens(def, v))
	errors = appendError(errors, differentArgAbbr(def, v))
	errors = appendError(errors, flagTokensAreNotEmpty(def, v))
	errors = appendError(errors, differentFlagTokens(def, v))
	errors = appendError(errors, differentFlagAbbr(def, v))
	errors = appendError(errors, differentAbbr(def, v))

	// arguments
	errors = appendError(errors, commandsValidArgs(def, v))
	errors = appendError(errors, allUsedArgs(def, v))
	errors = appendError(errors, singleDefaultArgPerCmd(def, v))
	errors = appendError(errors, differentArgsCmd(def, v))
	errors = appendError(errors, differentArgValues(def, v))

	// examples
	errors = appendError(errors, examplesDescAreNotEmpty(def, v))
	errors = appendError(errors, examplesArgumentsAreValid(def, v))
	errors = appendError(errors, examplesArgumentsAreNotEmpty(def, v))
	errors = appendError(errors, examplesHaveArgsValues(def, v))
	errors = appendError(errors, examplesHaveValidValues(def, v))

	return errors
}
