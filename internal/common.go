package internal

import "fmt"

type CommonDefinition struct {
	Token string `json:"token"`
	Abbr  string `json:"abbr,omitempty"`
	Hint  string `json:"hint,omitempty"`
	Desc  string `json:"desc,omitempty"`
}

func (cd *CommonDefinition) Print() {
	fmt.Printf("Token:%v\nAbbr:%v\nHint:%v\nDesc:%v\n",
		cd.Token, cd.Abbr, cd.Hint, cd.Desc)
}

func commonByToken(commons []CommonDefinition, token string) *CommonDefinition {
	for _, c := range commons {
		if c.Token == token {
			return &c
		}
	}
	return nil
}
