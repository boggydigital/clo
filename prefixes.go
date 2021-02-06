package clo

import (
	"strings"
)

const (
	argPfx    = "-"
	argDblPfx = "--"
	valuesSep = ","
)

func isArg(token string) bool {
	return strings.HasPrefix(token, argPfx)
}

func trimArgPrefix(token string) string {
	pfx := ""
	if strings.HasPrefix(token, argDblPfx) {
		pfx = argDblPfx
	} else if strings.HasPrefix(token, argPfx) {
		pfx = argPfx
	}
	return strings.TrimPrefix(token, pfx)
}
