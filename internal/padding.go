package internal

func (defs *Definitions) cmdPadding() int {
	lToken := ""
	for cmd, _ := range defs.Cmd {
		if len(cmd) > len(lToken) {
			lToken = cmd
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
		if len(arg) > len(lToken) {
			lToken = arg
		}
	}
	return len(lToken)
}
