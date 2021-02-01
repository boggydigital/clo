package internal

func (defs *Definitions) cmdPadding() int {
	lToken := ""
	for cmd, _ := range defs.Cmd {
		tc := trimAttrs(cmd)
		if len(tc) > len(lToken) {
			lToken = tc
		}
	}
	return len(lToken)
}

func (defs *Definitions) argPadding(cmd string) int {
	lToken := ""
	dc := defs.definedCmd(cmd)
	if dc == "" {
		return 0
	}
	for _, arg := range defs.Cmd[dc] {
		ta := trimAttrs(arg)
		if len(ta) > len(lToken) {
			lToken = ta
		}
	}
	return len(lToken)
}
