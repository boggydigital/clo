package defs

import "fmt"

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
