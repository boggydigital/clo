package internal

// parseCtx stores current command that is used for determining
// if a particular argument is expected at a given moment. Same
// goes for the argument that is used to validate values
type parseCtx struct {
	Command  *CommandDefinition
	Argument *ArgumentDefinition
}

func (ctx *parseCtx) update(token string, tokenType int, def *Definitions) {
	switch tokenType {
	case command:
		ctx.Command = def.CommandByToken(token)
	//case commandAbbr:
	//	ctx.Command = def.CommandByAbbr(token)
	case argument:
		ctx.Argument = def.ArgByToken(trimPrefix(token))
		//case argumentAbbr:
		//	ctx.Argument = def.ArgByAbbr(trimPrefix(token))
	}
}
