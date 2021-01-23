package internal

import (
	"strings"
)

const (
	argPfx       = "-"
	argDblPfx    = "--"
	argValuesSep = "="
	valuesSep    = ","
)

func isArg(token string) bool {
	return strings.HasPrefix(token, argPfx)
}

func hasArgValues(token string) bool {
	return strings.Contains(token, argValuesSep)
}

func argValues(token string) (string, []string) {
	if !hasArgValues(token) {
		return token, []string{}
	}
	argVal := strings.Split(token, argValuesSep)
	if len(argVal) > 1 {
		return argVal[0], strings.Split(argVal[len(argVal)-1], valuesSep)
	}
	return token, []string{}
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
