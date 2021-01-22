package internal

import (
	"fmt"
	"strings"
)

const (
	defaultPrefix  = "_"
	requiredSuffix = "!"
)

func hasPrefix(token string) bool {
	// not testing for strings.HasPrefix(token, "--"), since it'll match this case as well
	return strings.HasPrefix(token, "-")

}

func trimPrefix(token string) string {
	pfx := ""
	if strings.HasPrefix(token, "--") {
		pfx = "--"
	} else if strings.HasPrefix(token, "-") {
		pfx = "-"
	}
	return strings.TrimPrefix(token, pfx)
}

func trimArgument(arg string) string {
	return strings.Trim(arg, defaultPrefix+requiredSuffix)
}

func decorateDefault(arg string) string {
	return fmt.Sprintf("%s%s", defaultPrefix, arg)
}
