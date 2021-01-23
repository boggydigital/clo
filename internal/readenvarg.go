package internal

import (
	"fmt"
	"strings"
)

func argEnv(prefix, cmd, arg string) string {
	if arg == "" {
		return ""
	}
	// using COMMAND_ARGUMENT format to allow specifying different values
	// for the same argument token for different commands
	envKey := strings.ToUpper(arg)

	if cmd != "" {
		envKey = fmt.Sprintf("%s_%s", strings.ToUpper(cmd), envKey)
	}

	if prefix != "" {
		envKey = fmt.Sprintf("%s_%s", strings.ToUpper(prefix), envKey)
	}

	return envKey
}

// readEnvArgs reads arguments values from the environmental variables
func (req *Request) readEnvArgs(def *Definitions) error {
	if def == nil {
		return fmt.Errorf("cannot fill args from env using nil definitions")
	}

	//for _, arg := range def.Arguments {
	//	if !arg.Env {
	//		continue
	//	}
	//
	//	envKey := argEnv(def.EnvPrefix, req.Command, arg.Token)
	//	envVal := os.Getenv(envKey)
	//
	//	// only add value from environmental variable if it's the only value
	//	if envVal != "" &&
	//		(req.Arguments[arg.Token] == nil ||
	//			len(req.Arguments[arg.Token]) == 0) {
	//		req.Arguments[arg.Token] = append(req.Arguments[arg.Token], envVal)
	//	}
	//}

	return nil
}
