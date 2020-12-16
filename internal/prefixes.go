package internal

import (
	"strings"
)

func hasAnyPrefix(token string) bool {
	return strings.HasPrefix(token, "-") ||
		strings.HasPrefix(token, "--")
}

func hasExpectedPrefix(token string, tokenType int) bool {
	prefix := ""
	switch tokenType {
	case command:
		fallthrough
	case commandAbbr:
		fallthrough
	case value:
		fallthrough
	case valueFixed:
		fallthrough
	case valueDefault:
		return false
	case flagAbbr:
		fallthrough
	case argumentAbbr:
		prefix = "-"
	case flag:
		fallthrough
	case argument:
		prefix = "--"
	}
	return strings.HasPrefix(token, prefix)
}

func trimPrefix(token string, tokenType int) string {
	prefix := ""
	switch tokenType {
	case command:
		fallthrough
	case commandAbbr:
		fallthrough
	case value:
		fallthrough
	case valueFixed:
		fallthrough
	case valueDefault:
		return token
	case flagAbbr:
		fallthrough
	case argumentAbbr:
		prefix = "-"
	case flag:
		fallthrough
	case argument:
		prefix = "--"
	}
	return strings.TrimPrefix(token, prefix)
}
