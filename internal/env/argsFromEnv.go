package env

import (
	"fmt"
	"github.com/boggydigital/clove/internal/clireq"
	"github.com/boggydigital/clove/internal/defs"
	"os"
	"strings"
)

// EnvArgs reads arguments values from the environmental variables
func EnvArgs(req *clireq.Request, def *defs.Definitions) error {
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
		if req.Arguments[arg.Token] == nil ||
			(len(req.Arguments[arg.Token]) == 0 && envVal != "") {
			req.Arguments[arg.Token] = append(req.Arguments[arg.Token], envVal)
		}
	}

	return nil
}
