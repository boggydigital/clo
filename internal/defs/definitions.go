package defs

import (
	"fmt"
)

type Definitions struct {
	Version   int                  `json:"version"`
	Flags     []FlagDefinition     `json:"flags,omitempty"`
	Commands  []CommandDefinition  `json:"commands,omitempty"`
	Arguments []ArgumentDefinition `json:"arguments,omitempty"`
	Values    []ValueDefinition    `json:"values,omitempty"`
}

func (def *Definitions) Print() {
	fmt.Printf("Version:%v\n", def.Version)
	fmt.Println("----- Flags -----")
	for _, f := range def.Flags {
		fmt.Println("----------")
		f.Print()
	}
	fmt.Println("----- Commands -----")
	for _, c := range def.Commands {
		fmt.Println("----------")
		c.Print()
	}
	fmt.Println("----- Arguments -----")
	for _, a := range def.Arguments {
		fmt.Println("----------")
		a.Print()
	}
	fmt.Println("----- Values -----")
	for _, v := range def.Values {
		fmt.Println("----------")
		v.Print()
	}
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

//func (def *Definitions) FlagByToA(tokenOrAbbr string) *FlagDefinition{
//	fd := def.FlagByToken(tokenOrAbbr)
//	if fd == nil {
//		return def.FlagByAbbr(tokenOrAbbr)
//	}
//	return fd
//}

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

//func (def *Definitions) CommandByToA(tokenOrAbbr string) *CommandDefinition {
//	cd := def.CommandByToken(tokenOrAbbr)
//	if cd == nil {
//		return def.CommandByAbbr(tokenOrAbbr)
//	}
//	return cd
//}

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

//func (def *Definitions) ArgByToA(tokenOrAbbr string) *ArgumentDefinition {
//	ad := def.ArgByToken(tokenOrAbbr)
//	if ad == nil {
//		return def.ArgByAbbr(tokenOrAbbr)
//	}
//	return ad
//}

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
