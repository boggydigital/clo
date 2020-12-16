package internal

import (
	"errors"
	"fmt"
)

type Request struct {
	Flags     []string
	Command   string
	Arguments map[string][]string
}

func (req *Request) update(expandedToken string, tokenType int, ctx *parseCtx) error {

	switch tokenType {
	case commandAbbr:
		fallthrough
	case command:
		if req.Command != "" {
			return errors.New("request already has a command specified")
		}
		req.Command = expandedToken
		break
	case flagAbbr:
		req.Flags = append(req.Flags, trimPrefix(expandedToken))
	case flag:
		req.Flags = append(req.Flags, trimPrefix(expandedToken))
	case argument:
		fallthrough
	case argumentAbbr:
		arg := trimPrefix(expandedToken)
		if req.Arguments[arg] == nil {
			req.Arguments[arg] = []string{}
		}
	case valueDefault:
		fallthrough
	case value:
		fallthrough
	case valueFixed:
		argCtx := ctx.Argument.Token
		req.Arguments[argCtx] = append(req.Arguments[argCtx], expandedToken)
	default:
		return fmt.Errorf(
			"cannot update request for a token '%v' of type '%v'",
			expandedToken,
			tokenString(tokenType))
	}
	return nil
}
