package internal

import (
	"strings"
)

func hasPrefix(token string) bool {
	return strings.HasPrefix(token, "-") ||
		strings.HasPrefix(token, "--")
}

//func hasExpectedPrefix(token string, tokenType int) bool {
//	prefix := ""
//	switch tokenType {
//	case command:
//		fallthrough
//	case commandAbbr:
//		fallthrough
//	case value:
//		fallthrough
//	case valueFixed:
//		fallthrough
//	case valueDefault:
//		return false
//	case flagAbbr:
//		fallthrough
//	case argumentAbbr:
//		prefix = "-"
//	case flag:
//		fallthrough
//	case argument:
//		prefix = "--"
//	}
//	return strings.HasPrefix(token, prefix)
//}

func trimPrefix(token string) string {
	if strings.HasPrefix(token, "--") {
		return strings.TrimPrefix(token, "--")
	} else if strings.HasPrefix(token, "-") {
		return strings.TrimPrefix(token, "-")
	} else {
		return token
	}
	//prefix := ""
	//switch tokenType {
	//case command:
	//	fallthrough
	//case commandAbbr:
	//	fallthrough
	//case value:
	//	fallthrough
	//case valueFixed:
	//	fallthrough
	//case valueDefault:
	//	return token
	//case flagAbbr:
	//	fallthrough
	//case argumentAbbr:
	//	prefix = "-"
	//case flag:
	//	fallthrough
	//case argument:
	//	prefix = "--"
	//}
	//return strings.TrimPrefix(token, prefix)
}
