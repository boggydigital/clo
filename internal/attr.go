package internal

import "strings"

const (
	defaultAttr  = "_"
	requiredAttr = "!"
	multipleAttr = "..."
	envAttr      = "$"
	attrs        = defaultAttr + requiredAttr + multipleAttr + envAttr
	argValuesSep = "="
)

func isDefault(token string) bool {
	return strings.Contains(token, defaultAttr)
}

func isRequired(token string) bool {
	return strings.Contains(token, requiredAttr)
}

func isMultiple(token string) bool {
	return strings.Contains(token, multipleAttr)
}

func isEnv(token string) bool {
	return strings.Contains(token, envAttr)
}

func hasArgValues(token string) bool {
	return strings.Contains(token, argValuesSep)
}

func splitArgValues(token string) (string, []string) {
	if !hasArgValues(token) {
		return token, []string{}
	}
	argVal := strings.Split(token, argValuesSep)
	return argVal[0], strings.Split(argVal[len(argVal)-1], valuesSep)
}

func trimArgValues(token string) string {
	if hasArgValues(token) {
		token, _ = splitArgValues(token)
	}
	return token
}

func trimAttrs(token string) string {
	token = trimArgValues(token)
	return strings.Trim(token, attrs)
}
