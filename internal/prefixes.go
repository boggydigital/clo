package internal

import (
	"strings"
)

func hasPrefix(token string) bool {
	return strings.HasPrefix(token, "-") ||
		strings.HasPrefix(token, "--")
}

func trimPrefix(token string) string {
	if strings.HasPrefix(token, "--") {
		return strings.TrimPrefix(token, "--")
	} else if strings.HasPrefix(token, "-") {
		return strings.TrimPrefix(token, "-")
	} else {
		return token
	}
}
