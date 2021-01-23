package internal

import "strings"

const (
	defaultAttr  = "_"
	requiredAttr = "!"
	multipleAttr = "..."
	envAttr      = "$"
	attrs        = defaultAttr + requiredAttr + multipleAttr + envAttr
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

func trimAttrs(token string) string {
	return strings.Trim(token, attrs)
}
