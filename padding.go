package clo

func (defs *Definitions) cmdPadding() int {
	lToken := ""
	for cmd := range defs.Cmd {
		tc := trimAttrs(cmd)
		if len(tc) > len(lToken) {
			lToken = tc
		}
	}
	return len(lToken)
}

func (defs *Definitions) argPadding(cmd string) (int, error) {
	lToken := ""
	dc, err := defs.definedCmd(cmd)
	if err != nil {
		return 0, err
	}
	if dc == "" {
		return 0, nil
	}
	for _, arg := range defs.Cmd[dc] {
		ta := trimAttrs(arg)
		if len(ta) > len(lToken) {
			lToken = ta
		}
	}
	return len(lToken), nil
}
