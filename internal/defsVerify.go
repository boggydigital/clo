package internal

import "fmt"

func (def *Definitions) haveCommands() error {
	return nil
}

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

func cmdTokensAreNotEmpty(def *Definitions) error {
	for i, c := range def.Commands {
		if c.Token == "" {
			return fmt.Errorf("command #%d has an empty token", i+1)
		}
	}
	return nil
}

func differentCmdTokens(def *Definitions) error {
	cmds := make([]string, 0)
	for _, c := range def.Commands {
		cmds = append(cmds, c.Token)
	}
	if df := sliceFirstDupe(cmds); df != "" {
		return fmt.Errorf("commands have duplicate token: '%v'", df)
	}
	return nil
}

func differentCmdAbbr(def *Definitions) error {
	cmds := make([]string, 0)
	for _, c := range def.Commands {
		cmds = append(cmds, c.Abbr)
	}
	if df := sliceFirstDupe(cmds); df != "" {
		return fmt.Errorf("commands have duplicate abbreviation: '%v'", df)
	}
	return nil
}

func differentArgTokens(def *Definitions) error {
	args := make([]string, 0)
	for _, a := range def.Arguments {
		args = append(args, a.Token)
	}
	if df := sliceFirstDupe(args); df != "" {
		return fmt.Errorf("arguments have duplicate token: '%v'", df)
	}
	return nil
}

func argTokensAreNotEmpty(def *Definitions) error {
	for i, a := range def.Arguments {
		if a.Token == "" {
			return fmt.Errorf("argument #%d has an empty token", i+1)
		}
	}
	return nil
}

func differentArgAbbr(def *Definitions) error {
	args := make([]string, 0)
	for _, a := range def.Arguments {
		args = append(args, a.Abbr)
	}
	if df := sliceFirstDupe(args); df != "" {
		return fmt.Errorf("arguments have duplicate abbreviation: '%v'", df)
	}
	return nil
}

func differentFlagTokens(def *Definitions) error {
	flags := make([]string, 0)
	for _, f := range def.Flags {
		flags = append(flags, f.Token)
	}
	if df := sliceFirstDupe(flags); df != "" {
		return fmt.Errorf("flags have duplicate token: '%v'", df)
	}
	return nil
}

func differentFlagAbbr(def *Definitions) error {
	flags := make([]string, 0)
	for _, f := range def.Flags {
		flags = append(flags, f.Abbr)
	}
	if df := sliceFirstDupe(flags); df != "" {
		return fmt.Errorf("flags have duplicate abbreviation: '%v'", df)
	}
	return nil
}

func flagTokensAreNotEmpty(def *Definitions) error {
	for i, f := range def.Flags {
		if f.Token == "" {
			return fmt.Errorf("flag #%d has an empty token", i+1)
		}
	}
	return nil
}

func differentAbbr(def *Definitions) error {
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
		return fmt.Errorf("same abbreviation for a command, argument or flag: '%v'", da)
	}
	return nil
}

func commandsValidArgs(def *Definitions) error {
	for _, c := range def.Commands {
		for _, a := range c.Arguments {
			da := def.ArgByToken(a)
			if da == nil {
				return fmt.Errorf("command '%s' has undefined argument '%s'", c.Token, a)
			}
		}
	}
	return nil
}

func allUsedArgs(def *Definitions) error {
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
			return fmt.Errorf("argument '%s' is not used in any command", a.Token)
		}
	}
	return nil
}

func differentArgsCmd(def *Definitions) error {
	for _, c := range def.Commands {
		if da := sliceFirstDupe(c.Arguments); da != "" {
			return fmt.Errorf("command '%s' has duplicate argument '%s'", c.Token, da)
		}
	}
	return nil
}

func singleDefaultArg(def *Definitions) error {
	d := ""
	for _, a := range def.Arguments {
		if a.Default {
			if d != "" {
				return fmt.Errorf("argument '%s' redefines default from '%s'", a.Token, d)
			}
			d = a.Token
		}
	}
	return nil
}

func differentArgValues(def *Definitions) error {
	for _, a := range def.Arguments {
		if dv := sliceFirstDupe(a.Values); dv != "" {
			return fmt.Errorf("argument '%s' has duplicate value '%s'", a.Token, dv)
		}
	}
	return nil
}

func appendError(errors []error, err error) []error {
	if err != nil {
		return append(errors, err)
	}
	return errors
}

func (def *Definitions) Verify() []error {

	errors := make([]error, 0)

	// tokens and abbreviations
	errors = appendError(errors, cmdTokensAreNotEmpty(def))
	errors = appendError(errors, differentCmdTokens(def))
	errors = appendError(errors, differentCmdAbbr(def))
	errors = appendError(errors, argTokensAreNotEmpty(def))
	errors = appendError(errors, differentArgTokens(def))
	errors = appendError(errors, differentArgAbbr(def))
	errors = appendError(errors, flagTokensAreNotEmpty(def))
	errors = appendError(errors, differentFlagTokens(def))
	errors = appendError(errors, differentFlagAbbr(def))
	errors = appendError(errors, differentAbbr(def))

	// arguments
	errors = appendError(errors, commandsValidArgs(def))
	errors = appendError(errors, allUsedArgs(def))
	errors = appendError(errors, singleDefaultArg(def))
	errors = appendError(errors, differentArgsCmd(def))
	errors = appendError(errors, differentArgValues(def))

	return errors
}
