package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Definitions struct {
	Version   int                  `json:"version"`
	EnvPrefix string               `json:"env-prefix,omitempty"`
	App       string               `json:"app,omitempty"`
	Hint      string               `json:"hint,omitempty"`
	Desc      string               `json:"desc,omitempty"`
	Flags     []FlagDefinition     `json:"flags,omitempty"`
	Commands  []CommandDefinition  `json:"commands,omitempty"`
	Arguments []ArgumentDefinition `json:"arguments,omitempty"`
	Values    []ValueDefinition    `json:"values,omitempty"`
}

func loadDefault() (*Definitions, error) {
	bytes, err := ioutil.ReadFile("app/clove.json")
	if err != nil {
		return nil, err
	}
	return Load(bytes)
}

func Load(bytes []byte) (*Definitions, error) {
	var dfs *Definitions

	if err := json.Unmarshal(bytes, &dfs); err != nil {
		return nil, err
	}

	if err := dfs.addHelpCmd(); err != nil {
		// adding help is not considered fatal error, inform, continue
		fmt.Println("error adding help command:", err.Error())
	}

	if err := dfs.expandRefValues(); err != nil {
		return nil, err
	}

	return dfs, nil
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

func (def *Definitions) ValueByToken(val string) *ValueDefinition {
	for _, v := range def.Values {
		if v.Token == val {
			return &v
		}
	}
	return nil
}

func (def *Definitions) DefinedValue(values []string) bool {
	if values == nil ||
		len(values) == 0 {
		return false
	}
	for _, vt := range values {
		vd := def.ValueByToken(vt)
		if vd != nil {
			return true
		}
	}
	return false
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

func (def *Definitions) expandRefValues() error {
	for i, ad := range def.Arguments {
		if ad.Values != nil &&
			len(ad.Values) == 1 &&
			strings.HasPrefix(ad.Values[0], "from:") {
			source := strings.TrimPrefix(ad.Values[0], "from:")
			switch source {
			case "commands":
				def.Arguments[i].Values = make([]string, 0)
				for _, cd := range def.Commands {
					def.Arguments[i].Values = append(def.Arguments[i].Values, cd.Token)
				}
				return nil
			default:
				return fmt.Errorf("cannot expand values from an unknown source: '%s'", source)
			}
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

func (def *Definitions) ValidArgVal(val string, arg string) bool {
	if arg == "" {
		return false
	}
	ad := def.ArgByToken(arg)
	if ad == nil {
		return false
	}
	return ad.ValidValue(val)
}
