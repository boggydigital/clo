package internal

import (
	"fmt"
	"strings"
)

const (
	requiredPrefix = "*"
	defaultPrefix  = "_"
	multipleSuffix = "..."
)

func GenDefinitions(app string, commands, arguments, flags []string) *Definitions {
	defs := &Definitions{
		Version:   1,
		EnvPrefix: strings.ToUpper(app),
		App:       app,
	}

	if app != "" {
		defs.Hint = fmt.Sprintf("%s hint", app)
		defs.Desc = fmt.Sprintf("%s description", app)
	}

	for _, c := range commands {
		defs.Commands = append(defs.Commands, *genCommand(c))
	}
	for _, a := range arguments {
		defs.Arguments = append(defs.Arguments, *genArgument(a))
	}
	for _, f := range flags {
		defs.Flags = append(defs.Flags, *genFlag(f))
	}
	return defs
}

func genCommand(cmd string) *CommandDefinition {
	return &CommandDefinition{
		CommonDefinition: CommonDefinition{
			Token: cmd,
			Hint:  fmt.Sprintf("%s hint", cmd),
			Desc:  fmt.Sprintf("%s description", cmd),
		},
		Arguments: []string{},
	}
}

func genArgument(arg string) *ArgumentDefinition {
	ad := &ArgumentDefinition{
		CommonDefinition: CommonDefinition{},
	}

	if strings.HasPrefix(arg, defaultPrefix) {
		ad.Default = true
		arg = strings.TrimPrefix(arg, defaultPrefix)
	}

	if strings.HasPrefix(arg, requiredPrefix) {
		ad.Required = true
		arg = strings.TrimPrefix(arg, requiredPrefix)
	}

	if strings.HasSuffix(arg, multipleSuffix) {
		ad.Multiple = true
		arg = strings.TrimSuffix(arg, multipleSuffix)
	}

	ad.Token = arg
	ad.Hint = fmt.Sprintf("%s hint", arg)
	ad.Desc = fmt.Sprintf("%s description", arg)

	return ad
}

func genFlag(flag string) *FlagDefinition {
	return &FlagDefinition{
		CommonDefinition: CommonDefinition{
			Token: flag,
			Hint:  fmt.Sprintf("%s hint", flag),
			Desc:  fmt.Sprintf("%s description", flag),
		},
	}
}
