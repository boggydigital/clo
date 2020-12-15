package clireq

import (
	"errors"
	"fmt"
	"github.com/boggydigital/clove/internal/parsectx"
	"github.com/boggydigital/clove/internal/strutil"
	"github.com/boggydigital/clove/internal/tokens"
)

type Request struct {
	Flags     []string
	Command   string
	Arguments map[string][]string
}

func (req *Request) Print() {
	if len(req.Flags) > 0 {
		fmt.Printf("Flags:%v\n", req.Flags)
	}
	fmt.Printf("Command:%v\n", req.Command)
	if len(req.Arguments) > 0 {
		fmt.Println("Arguments:")
		for k, v := range req.Arguments {
			fmt.Printf("- %v=%v\n", k, v)
		}
	}
}

func Update(req *Request, expandedToken string, tokenType int, ctx *parsectx.Context) error {

	switch tokenType {
	case tokens.CommandAbbr:
		fallthrough
	case tokens.Command:
		if req.Command != "" {
			return errors.New("request already has a command specified")
		}
		req.Command = expandedToken
		break
	case tokens.FlagAbbr:
		req.Flags = append(req.Flags, strutil.TrimPrefix(expandedToken, tokenType))
	case tokens.Flag:
		req.Flags = append(req.Flags, strutil.TrimPrefix(expandedToken, tokenType))
	case tokens.Argument:
		fallthrough
	case tokens.ArgumentAbbr:
		arg := strutil.TrimPrefix(expandedToken, tokenType)
		if req.Arguments[arg] == nil {
			req.Arguments[arg] = []string{}
		}
	case tokens.ValueDefault:
		fallthrough
	case tokens.Value:
		fallthrough
	case tokens.ValueFixed:
		arg := ctx.Argument.Token
		req.Arguments[arg] = append(req.Arguments[arg], expandedToken)
	default:
		return errors.New(
			fmt.Sprintf(
				"cannot update request for a token '%v' of type '%v'",
				expandedToken,
				tokens.String(tokenType)))
	}
	return nil
}
