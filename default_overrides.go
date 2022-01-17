package clo

import (
	"fmt"
	"strings"
)

func (defs *definitions) SetUserDefaults(defaultOverrides map[string][]string) error {
	//check for errors in cmd:args or args in overrides to help
	//user understand something might be not specified correctly
	if err := defs.validateOverrides(defaultOverrides); err != nil {
		return err
	}

	defs.defaultsOverrides = defaultOverrides

	return nil
}

func (defs *definitions) validateCmdArgValuesOverrides(cmd, arg string, values []string) error {
	dc, err := defs.definedCmd(cmd)
	if err != nil {
		return err
	}
	if dc == "" {
		return fmt.Errorf("unknown override command %s", cmd)
	}

	da, err := defs.definedArg(cmd, arg)
	if err != nil {
		return err
	}
	if da == "" {
		return fmt.Errorf("unknown override argument %s for command %s", arg, cmd)
	}

	for _, val := range values {
		dv, err := defs.definedVal(dc, da, strings.ToLower(val))
		if err != nil {
			return err
		}
		if dv == "" {
			return fmt.Errorf("unknown override value %s for argument %s, command %s", val, arg, cmd)
		}
	}

	return nil
}

func (defs *definitions) validateOverrides(do map[string][]string) error {
	for overrideArg, overrideValues := range do {
		if strings.Contains(overrideArg, ":") {
			parts := strings.Split(overrideArg, ":")
			//not checking for length given the condition before the split
			cmd, arg := parts[0], parts[1]
			return defs.validateCmdArgValuesOverrides(cmd, arg, overrideValues)
		}

		for dc, argValues := range defs.Cmd {
			for _, arg := range argValues {
				ca, _ := splitArgValues(arg)
				if trimAttrs(ca) != overrideArg {
					continue
				}
				if err := defs.validateCmdArgValuesOverrides(dc, overrideArg, overrideValues); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (defs *definitions) HasUserDefaultsFlag(flag string) bool {
	if defs.defaultsOverrides == nil {
		return false
	}

	_, ok := defs.defaultsOverrides[flag]
	return ok
}
