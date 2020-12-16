package internal

type Definitions struct {
	Version   int                  `json:"version"`
	EnvPrefix string               `json:"env-prefix"`
	Flags     []FlagDefinition     `json:"flags,omitempty"`
	Commands  []CommandDefinition  `json:"commands,omitempty"`
	Arguments []ArgumentDefinition `json:"arguments,omitempty"`
	Values    []ValueDefinition    `json:"values,omitempty"`
}

func (def *Definitions) FlagByToken(token string) *FlagDefinition {
	for _, f := range def.Flags {
		if f.Token == token {
			return &f
		}
	}
	return nil
}

func (def *Definitions) FlagByAbbr(abbr string) *FlagDefinition {
	for _, f := range def.Flags {
		if f.Abbr == abbr {
			return &f
		}
	}
	return nil
}

func (def *Definitions) CommandByToken(token string) *CommandDefinition {
	for _, c := range def.Commands {
		if c.Token == token {
			return &c
		}
	}
	return nil
}

func (def *Definitions) CommandByAbbr(abbr string) *CommandDefinition {
	for _, c := range def.Commands {
		if c.Abbr == abbr {
			return &c
		}
	}
	return nil
}

func (def *Definitions) ArgByToken(token string) *ArgumentDefinition {
	for _, a := range def.Arguments {
		if a.Token == token {
			return &a
		}
	}
	return nil
}

func (def *Definitions) ArgByAbbr(abbr string) *ArgumentDefinition {
	for _, a := range def.Arguments {
		if a.Abbr == abbr {
			return &a
		}
	}
	return nil
}

func (def *Definitions) DefaultArg(cmd *CommandDefinition) *ArgumentDefinition {
	if cmd == nil {
		return nil
	}
	for _, arg := range cmd.Arguments {
		ad := def.ArgByToken(arg)
		if ad == nil {
			continue
		}
		if ad.Default {
			return ad
		}
	}
	return nil
}

func (def *Definitions) RequiredArgs(cmd string) []string {
	required := make([]string, 0)

	command := def.CommandByToken(cmd)
	if command == nil {
		return required
	}

	for _, at := range command.Arguments {
		if at == "" {
			continue
		}
		arg := def.ArgByToken(at)
		if arg == nil {
			continue
		}
		if arg.Required {
			required = append(required, arg.Token)
		}
	}

	return required
}
