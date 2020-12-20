package internal

import (
	"strings"
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
