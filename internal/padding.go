package internal

import (
	"strconv"
)

func (def *Definitions) CommandsPadding() string {
	lToken := ""
	for _, cmd := range def.Commands {
		if len(cmd.Token) > len(lToken) {
			lToken = cmd.Token
		}
	}
	return strconv.Itoa(len(lToken))
}

func (def *Definitions) FlagsPadding() string {
	lToken := ""
	for _, flg := range def.Flags {
		if len(flg.Token) > len(lToken) {
			lToken = flg.Token
		}
	}
	return strconv.Itoa(len(lToken))
}

func (def *Definitions) ArgumentsPadding(cmd string) string {
	lToken := ""
	cd := def.CommandByToken(cmd)
	if cd == nil {
		return "0"
	}

	for _, arg := range cd.Arguments {
		if len(arg) > len(lToken) {
			lToken = arg
		}
	}
	return strconv.Itoa(len(lToken))
}
