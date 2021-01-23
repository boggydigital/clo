package internal

func (def *Definitions) CommandsPadding() int {
	lToken := ""
	for cmd, _ := range def.Cmd {
		if len(cmd) > len(lToken) {
			lToken = cmd
		}
	}
	return len(lToken)
}

func (def *Definitions) ArgumentsPadding(cmd string) int {
	lToken := ""
	//cd := def.CommandByToken(cmd)
	//if cd == nil {
	//	return len(lToken)
	//}
	//
	//for _, arg := range cd.Arguments {
	//	if len(arg) > len(lToken) {
	//		lToken = arg
	//	}
	//}
	return len(lToken)
}
