package internal

func (def *Definitions) CommandsPadding() int {
	lToken := ""
	for _, cmd := range def.Commands {
		if len(cmd.Token) > len(lToken) {
			lToken = cmd.Token
		}
	}
	return len(lToken)
}

func (def *Definitions) FlagsPadding() int {
	lToken := ""
	for _, flg := range def.Flags {
		if len(flg.Token) > len(lToken) {
			lToken = flg.Token
		}
	}
	return len(lToken)
}

func (def *Definitions) ArgumentsPadding(cmd string) int {
	lToken := ""
	cd := def.CommandByToken(cmd)
	if cd == nil {
		return len(lToken)
	}

	for _, arg := range cd.Arguments {
		if len(arg) > len(lToken) {
			lToken = arg
		}
	}
	return len(lToken)
}
