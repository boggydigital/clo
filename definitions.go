package clo

import (
	"fmt"
	"github.com/boggydigital/wits"
	"io"
	"strings"
)

type definitions struct {
	Cmd               map[string][]string `json:"cmd"`
	Help              map[string]string   `json:"help"`
	defaultsOverrides map[string][]string
}

func Load(commands, help io.Reader, valuesDelegates map[string]func() []string) (defs *definitions, err error) {

	defs = &definitions{}

	if defs.Cmd, err = wits.ReadKeyValues(commands); err != nil {
		return defs, err
	}

	if help != nil {
		if defs.Help, err = wits.ReadKeyValue(help); err != nil {
			return defs, err
		}
	}

	// post-processing definitions include the following steps:
	// - replace placeholder values using delegates (if provided)
	// - add 'help' command if not present

	if valuesDelegates != nil {
		if err := defs.replacePlaceholders(valuesDelegates); err != nil {
			return defs, err
		}
	}

	addInternalHelpCmd(defs)

	return defs, nil
}

func (defs *definitions) replaceArgValuesList(cmd, replaceArg string, replacedValues []string) {
	if replaceArg == "" {
		return
	}
	// capacity = existing arguments, plus replaced values, minus the 1 placeholder
	replacedArgs := make([]string, 0, len(defs.Cmd[cmd])+len(replacedValues)-1)

	// first, add all existing arguments, except the placeholder
	for _, arg := range defs.Cmd[cmd] {
		if arg == replaceArg {
			continue
		}
		replacedArgs = append(replacedArgs, arg)
	}
	// second, add replaced values we've got from a delegate
	replacedArgs = append(replacedArgs, replacedValues...)

	defs.Cmd[cmd] = replacedArgs
}

func (defs *definitions) replacePlaceholders(valuesDelegates map[string]func() []string) error {
	for cmd, args := range defs.Cmd {

		// replacing argument with list values can happen once per command.
		// we might populate those values as we scan arguments upon finding list placeholder
		// (placeholder is a list placeholder, when the only value is a placeholder, e.g. "{list_values}"
		// and not part of argument values, e.g. "arg={arg_values}")
		var replaceArg string
		var replacedValues []string

		for i, arg := range args {
			ph := extract(arg)

			if ph.identifier == "" {
				continue
			}

			if valuesDelegates[ph.identifier] == nil {
				return fmt.Errorf("clo: %s not present in data delegates, can't expand", ph.identifier)
			}

			values := valuesDelegates[ph.identifier]()
			if ph.multiple {
				for j := 0; j < len(values); j++ {
					values[j] = makeMultiple(values[j])
				}
			}

			// if the placeholder has been specified as "first value is default"
			if ph.defaultFirstValue {
				if len(values) == 0 {
					return fmt.Errorf("clo: replaced values are empty, can't make first value default")
				}
				values[0] = makeDefault(values[0])
			}

			// list values is the last placeholder, so if we encountered one
			// we store replaced argument, values to replace with and break from the scan
			if ph.listValues {
				replaceArg = ph.String()
				replacedValues = values
				break
			}

			// replace argument placeholder with a comma separated list of values,
			// provided by a delegate
			defs.Cmd[cmd][i] = strings.Replace(arg, ph.String(), strings.Join(values, ","), 1)
		}

		// if replaceArg, replacedValues have been filled during args scan, then
		// replace that placeholder with a list of values provided by a delegate
		defs.replaceArgValuesList(cmd, replaceArg, replacedValues)
	}
	return nil
}

func (defs *definitions) definedCmd(c string) (string, error) {
	if defs == nil {
		return "", fmt.Errorf("clo: no defined command for nil definitions")
	}
	definedToken := ""

	for cmd := range defs.Cmd {
		if strings.HasPrefix(cmd, c) {
			if definedToken != "" {
				return "", fmt.Errorf("clo: ambiguous command %s that could be %s or %s", c, definedToken, cmd)
			}
			definedToken = cmd
		}
	}

	return definedToken, nil
}

func (defs *definitions) definedArg(c, a string) (string, error) {
	if defs == nil {
		return "", fmt.Errorf("clo: no defined command argument for nil defintions")
	}

	cmd, err := defs.definedCmd(c)
	if err != nil {
		return "", err
	}
	if cmd == "" {
		return "", nil
	}

	definedArg := ""

	for _, arg := range defs.Cmd[cmd] {
		if strings.HasPrefix(arg, a) {
			if definedArg != "" {
				return arg, fmt.Errorf("clo: ambiguous argument %s that could be %s or %s", a, definedArg, arg)
			}
			definedArg = arg
		}
	}

	return definedArg, nil
}

func (defs *definitions) definedVal(c, a, v string) (string, error) {
	if defs == nil {
		return "", fmt.Errorf("clo: no defined command argument value for nil definitions")
	}

	arg, err := defs.definedArg(c, a)
	if err != nil {
		return "", err
	}
	if arg == "" {
		return "", nil
	}

	definedValue := ""
	var values []string

	// splitArgValues
	if !hasArgValues(arg) {
		return v, nil
	}

	_, values = splitArgValues(arg)
	for _, val := range values {
		if strings.HasPrefix(val, v) {
			if definedValue != "" {
				return val, fmt.Errorf("clo: ambiguous value %s that could be %s or %s", v, definedValue, val)
			}
			definedValue = trimAttrs(val)
		}
	}

	return definedValue, nil
}

func (defs *definitions) defaultCommand() string {
	if defs == nil {
		return ""
	}
	for c := range defs.Cmd {
		if isDefault(c) {
			return c
		}
	}
	return ""
}

func (defs *definitions) defaultArgument(cmd string) (string, error) {
	if defs == nil {
		return "", fmt.Errorf("clo: no default argument for nil defintions")
	}

	dc, err := defs.definedCmd(cmd)
	if err != nil {
		return dc, err
	}
	if dc == "" {
		return "", nil
	}

	for _, arg := range defs.Cmd[dc] {
		asv, _ := splitArgValues(arg)
		if isDefault(asv) {
			return asv, nil
		}
	}
	return "", nil
}

func transform(arr []string, f func(string) string) []string {
	mArr := make([]string, 0, len(arr))
	for _, s := range arr {
		mArr = append(mArr, f(s))
	}
	return mArr
}

func (defs *definitions) defaultArgValues(req *request) error {
	if req == nil {
		return fmt.Errorf("clo: can't fill default argument values for a nil request")
	}
	if req.Command == "" {
		// return if no command has been specified, nothing to fill
		return nil
	}
	if req.Arguments == nil {
		req.Arguments = make(map[string][]string)
	}

	dc, err := defs.definedCmd(req.Command)
	if err != nil {
		return err
	}
	if dc == "" {
		return fmt.Errorf("unknown request command %s", req.Command)
	}

	for _, arg := range defs.Cmd[dc] {
		a, values := splitArgValues(arg)
		ta := trimAttrs(a)

		// check if request already has some values specified for that argument
		if rv, ok := req.Arguments[ta]; ok {
			if len(rv) > 0 {
				continue
			}
		}

		// TODO: add tests for user default overrides
		// check if user has provided default overrides with SetUserDefaults
		if len(defs.defaultsOverrides) > 0 {
			dv, ok := defs.defaultsOverrides[fmt.Sprintf("%s:%s", trimAttrs(dc), ta)]
			if !ok {
				// if cmd:arg doesn't match, check generic arg
				dv, ok = defs.defaultsOverrides[ta]
			}
			if ok {
				req.Arguments[ta] = dv
				continue
			}
		}

		for _, v := range values {
			if !isDefault(v) {
				continue
			}

			if req.Arguments[ta] == nil {
				req.Arguments[ta] = make([]string, 0)
			}

			req.Arguments[ta] = append(req.Arguments[ta], trimAttrs(v))
		}
	}

	return nil
}
