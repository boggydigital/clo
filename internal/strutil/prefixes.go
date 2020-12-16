package strutil

import (
	"errors"
	"fmt"
	"github.com/boggydigital/clove/internal/defs"
	"github.com/boggydigital/clove/internal/tokens"
	"strings"
)

func HasAnyPrefix(token string) bool {
	return strings.HasPrefix(token, "-") ||
		strings.HasPrefix(token, "--")
}

func HasExpectedPrefix(token string, tokenType int) bool {
	prefix := ""
	switch tokenType {
	case tokens.Command:
		fallthrough
	case tokens.CommandAbbr:
		fallthrough
	case tokens.Value:
		fallthrough
	case tokens.ValueFixed:
		fallthrough
	case tokens.ValueDefault:
		return false
	case tokens.FlagAbbr:
		fallthrough
	case tokens.ArgumentAbbr:
		prefix = "-"
	case tokens.Flag:
		fallthrough
	case tokens.Argument:
		prefix = "--"
	}
	return strings.HasPrefix(token, prefix)
}

func TrimPrefix(token string, tokenType int) string {
	prefix := ""
	switch tokenType {
	case tokens.Command:
		fallthrough
	case tokens.CommandAbbr:
		fallthrough
	case tokens.Value:
		fallthrough
	case tokens.ValueFixed:
		fallthrough
	case tokens.ValueDefault:
		return token
	case tokens.FlagAbbr:
		fallthrough
	case tokens.ArgumentAbbr:
		prefix = "-"
	case tokens.Flag:
		fallthrough
	case tokens.Argument:
		prefix = "--"
	}
	return strings.TrimPrefix(token, prefix)
}

func ExpandAbbr(token string, tokenType int, def *defs.Definitions) (string, error) {
	switch tokenType {
	case tokens.CommandAbbr:
		cd := def.CommandByAbbr(token)
		if cd == nil {
			return "", errors.New(fmt.Sprintf("unknown command abbreviation: '%v'", token))
		}
		return cd.Token, nil
	case tokens.ArgumentAbbr:
		ad := def.ArgByAbbr(TrimPrefix(token, tokenType))
		if ad == nil {
			return "", errors.New(fmt.Sprintf("unknown argument abbreviation: '%v'", token))
		}
		return ad.Token, nil
	case tokens.FlagAbbr:
		fd := def.FlagByAbbr(TrimPrefix(token, tokenType))
		if fd == nil {
			return "", errors.New(fmt.Sprintf("unknown flag abbreviation: '%v'", token))
		}
		return fd.Token, nil
	default:
		return token, nil
	}
}
