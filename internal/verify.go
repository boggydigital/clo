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

func cmdTokensAreNotEmpty(commands []CommandDefinition, debug bool) error {
	msg := "command tokens are not empty"
	for i, c := range commands {
		if c.Token == "" {
			vFail(msg, debug)
			return fmt.Errorf("command #%d has an empty token", i+1)
		}
	}
	vPass(msg, debug)
	return nil
}

func differentCmdTokens(commands []CommandDefinition, debug bool) error {
	msg := "command tokens are different"
	dupeCmd := make([]string, 0)
	for _, c := range commands {
		dupeCmd = append(dupeCmd, c.Token)
	}
	if df := firstDupe(dupeCmd); df != "" {
		vFail(msg, debug)
		return fmt.Errorf("commands have duplicate token: '%v'", df)
	}
	vPass(msg, debug)
	return nil
}

//func differentCmdAbbr(commands []CommandDefinition, debug bool) error {
//	msg := "command abbreviations are different"
//	dupeCmd := make([]string, 0)
//	for _, c := range commands {
//		dupeCmd = append(dupeCmd, c.Abbr)
//	}
//	if df := firstDupe(dupeCmd); df != "" {
//		vFail(msg, debug)
//		return fmt.Errorf("commands have duplicate abbreviation: '%v'", df)
//	}
//	vPass(msg, debug)
//	return nil
//}

func argTokensAreNotEmpty(args []ArgumentDefinition, debug bool) error {
	msg := "argument tokens are not empty"
	for i, a := range args {
		if a.Token == "" {
			vFail(msg, debug)
			return fmt.Errorf("argument #%d has an empty token", i+1)
		}
	}
	vPass(msg, debug)
	return nil
}

func differentArgTokens(args []ArgumentDefinition, debug bool) error {
	msg := "argument tokens are different"
	dupeArgs := make([]string, 0)
	for _, a := range args {
		dupeArgs = append(dupeArgs, a.Token)
	}
	if df := firstDupe(dupeArgs); df != "" {
		vFail(msg, debug)
		return fmt.Errorf("arguments have duplicate token: '%v'", df)
	}
	vPass(msg, debug)
	return nil
}

//func differentArgAbbr(args []ArgumentDefinition, debug bool) error {
//	msg := "argument abbreviations are different"
//	dupeArgs := make([]string, 0)
//	for _, a := range args {
//		dupeArgs = append(dupeArgs, a.Abbr)
//	}
//	if df := firstDupe(dupeArgs); df != "" {
//		vFail(msg, debug)
//		return fmt.Errorf("arguments have duplicate abbreviation: '%v'", df)
//	}
//	vPass(msg, debug)
//	return nil
//}

//func differentAbbr(
//	commands []CommandDefinition,
//	args []ArgumentDefinition,
//	v bool) error {
//	msg := "all abbreviations are different"
//	abbr := make([]string, 0)
//
//	for _, c := range commands {
//		abbr = append(abbr, c.Abbr)
//	}
//	for _, a := range args {
//		abbr = append(abbr, a.Abbr)
//	}
//	if da := firstDupe(abbr); da != "" {
//		vFail(msg, v)
//		return fmt.Errorf("same abbreviation for a command, argument: '%v'", da)
//	}
//	vPass(msg, v)
//	return nil
//}

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

func appendError(errors []error, err error) []error {
	if err != nil {
		return append(errors, err)
	}
	return errors
}

func (def *Definitions) Verify(debug bool) []error {

	errors := make([]error, 0)

	// tokens and abbreviations
	errors = appendError(errors, cmdTokensAreNotEmpty(def.Commands, debug))
	errors = appendError(errors, differentCmdTokens(def.Commands, debug))
	//errors = appendError(errors, differentCmdAbbr(def.Commands, debug))
	errors = appendError(errors, argTokensAreNotEmpty(def.Arguments, debug))
	errors = appendError(errors, differentArgTokens(def.Arguments, debug))
	//errors = appendError(errors, differentArgAbbr(def.Arguments, debug))
	//errors = appendError(errors, differentAbbr(def.Commands, def.Arguments, debug))

	// arguments
	errors = appendError(errors, commandsValidArgs(def.Commands, def.ArgByToken, debug))
	errors = appendError(errors, allUsedArgs(def.Commands, def.Arguments, debug))
	errors = appendError(errors, differentArgsCmd(def.Commands, debug))
	errors = appendError(errors, differentArgValues(def.Arguments, debug))

	return errors
}
