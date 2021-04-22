package clo

import "strings"

type placeholder struct {
	identifier        string
	defaultFirstValue bool
	listValues        bool
}

const (
	placeholderPrefix = "{"
	placeholderSuffix = "}"
)

func extract(data string) *placeholder {
	ph := &placeholder{}
	start := strings.Index(data, placeholderPrefix)
	end := strings.Index(data, placeholderSuffix)
	if end-start > 1 && start > -1 {
		ph.identifier = data[start+1 : end]
		if isDefault(ph.identifier) && len(ph.identifier) > 1 {
			ph.defaultFirstValue = true
			ph.identifier = strings.TrimSuffix(ph.identifier, defaultAttr)
		}
		ph.listValues = start == 0
	}
	return ph
}

func (ph *placeholder) String() string {
	id := ph.identifier
	if ph.defaultFirstValue {
		id = id + "_"
	}
	return placeholderPrefix + id + placeholderSuffix
}
