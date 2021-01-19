package internal

import (
	"fmt"
	"strings"
)

const (
	//requiredPrefix = "*"
	//defaultPrefix  = "_"
	multipleSuffix = "..."
)

func GenDefinitions(app string, commands, arguments []string) *Definitions {
	defs := &Definitions{
		Version:   1,
		EnvPrefix: strings.ToUpper(app),
		App:       app,
	}

	if app != "" {
		defs.Help = fmt.Sprintf("%s help", app)
	}

	for _, c := range commands {
		defs.Commands = append(defs.Commands, *genCommand(c))
	}
	for _, a := range arguments {
		defs.Arguments = append(defs.Arguments, *genArgument(a))
	}
	return defs
}

func genCommand(cmd string) *CommandDefinition {
	return &CommandDefinition{
		CommonDefinition: CommonDefinition{
			Token: cmd,
			Help:  fmt.Sprintf("%s help", cmd),
		},
		Arguments: []string{},
	}
}

func genArgument(arg string) *ArgumentDefinition {
	ad := &ArgumentDefinition{
		CommonDefinition: CommonDefinition{},
	}

	//if strings.HasPrefix(arg, defaultPrefix) {
	//	ad.Default = true
	//	arg = strings.TrimPrefix(arg, defaultPrefix)
	//}
	//
	//if strings.HasPrefix(arg, requiredPrefix) {
	//	ad.Required = true
	//	arg = strings.TrimPrefix(arg, requiredPrefix)
	//}

	if strings.HasSuffix(arg, multipleSuffix) {
		ad.Multiple = true
		arg = strings.TrimSuffix(arg, multipleSuffix)
	}

	ad.Token = arg
	ad.Help = fmt.Sprintf("%s help", arg)

	return ad
}