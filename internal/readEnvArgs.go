package internal

import (
	"fmt"
	"os"
	"strings"
)

// readEnvArgs reads arguments values from the environmental variables
func (req *Request) readEnvArgs(def *Definitions) error {
	if req == nil {
		return fmt.Errorf("cannot fill args from env for a nil request")
	}
	if def == nil {
		return fmt.Errorf("cannot fill args from env using nil definitions")
	}

	for _, arg := range def.Arguments {
		if !arg.Env {
			continue
		}

		// using COMMAND_ARGUMENT format to allow specifying different values
		// for the same argument token for different commands
		envKey := fmt.Sprintf("%s_%s",
			strings.ToUpper(req.Command),
			strings.ToUpper(arg.Token))

		if def.EnvPrefix != "" {
			envKey = fmt.Sprintf("%s_%s", strings.ToUpper(def.EnvPrefix), envKey)
		}

		envVal := os.Getenv(envKey)

		// only add value from environmental variable if it's the only value
		if envVal != "" &&
			(req.Arguments[arg.Token] == nil ||
				len(req.Arguments[arg.Token]) == 0) {
			req.Arguments[arg.Token] = append(req.Arguments[arg.Token], envVal)
		}
	}

	return nil
}
