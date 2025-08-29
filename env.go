package clo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func appName() string {
	if len(os.Args) > 0 {
		return filepath.Base(os.Args[0])
	}
	return "app"
}

func argEnv(app, cmd, arg string) string {
	if arg == "" {
		return ""
	}
	// using COMMAND_ARGUMENT format to allow specifying different values
	// for the same argument token for different commands
	envKey := strings.ToUpper(arg)

	if cmd != "" {
		envKey = fmt.Sprintf("%s_%s", strings.ToUpper(cmd), envKey)
	}

	if app != "" {
		envKey = fmt.Sprintf("%s_%s",
			strings.ToUpper(app),
			envKey)
	}

	return envKey
}

// readEnvArgs reads arguments values from the environmental variables
func (req *request) readEnvArgs(def *definitions) error {
	if def == nil {
		return fmt.Errorf("cannot fill args from env using nil definitions")
	}

	dc, err := def.definedCmd(req.Command)
	if err != nil {
		return err
	}
	if dc == "" {
		return fmt.Errorf("cannot fill args from env for an empty command")
	}

	for _, arg := range def.Cmd[dc] {
		if !isEnv(arg) {
			continue
		}

		tArg := trimAttrs(arg)
		envKey := argEnv(appName(), req.Command, trimAttrs(arg))
		var envVals []string
		envVal := os.Getenv(envKey)
		if envVal == "" {
			// try command-independent value APP_ARG if APP_CMD_ARG didn't work
			envKey := argEnv(appName(), "", trimAttrs(arg))
			envVal = os.Getenv(envKey)
		}

		if envVal == "" {
			continue
		}

		if isMultiple(arg) && strings.Contains(envVal, ",") {
			envVals = strings.Split(envVal, ",")
		} else {
			envVals = []string{envVal}
		}

		// only add value from environmental variable if it's the only value,
		// don't overwrite value directly provided by user
		if len(envVals) > 0 &&
			len(req.Arguments[tArg]) == 0 {
			req.Arguments[tArg] = append(req.Arguments[tArg], envVals...)
		}
	}

	return nil
}
