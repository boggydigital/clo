package parsectx

import (
	"github.com/boggydigital/clove/internal/defs"
	"github.com/boggydigital/clove/internal/strutil"
	"github.com/boggydigital/clove/internal/tokens"
)

// context stores current command that is used for determining
// if a particular argument is expected at a given moment. Same
// goes for the argument that is used to validate values
type Context struct {
	Command  *defs.CommandDefinition
	Argument *defs.ArgumentDefinition
}

func Update(token string, tokenType int, ctx *Context, def *defs.Definitions) {
	switch tokenType {
	case tokens.Command:
		ctx.Command = def.CommandByToken(token)
	case tokens.CommandAbbr:
		ctx.Command = def.CommandByAbbr(token)
	case tokens.Argument:
		ctx.Argument = def.ArgByToken(strutil.TrimPrefix(token, tokenType))
	case tokens.ArgumentAbbr:
		ctx.Argument = def.ArgByAbbr(strutil.TrimPrefix(token, tokenType))
	}
}
