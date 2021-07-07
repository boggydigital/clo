package clo

import "strings"

type placeholder struct {
	identifier        string
	multiple          bool
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
		}
		if isMultiple(ph.identifier) && len(ph.identifier) > 1 {
			ph.multiple = true
		}
		ph.identifier = trimAttrs(ph.identifier)
		ph.listValues = start == 0
	}
	return ph
}

func (ph *placeholder) String() string {
	id := ph.identifier
	if ph.defaultFirstValue {
		id = makeDefault(id)
	}
	if ph.multiple {
		id = makeMultiple(id)
	}
	return placeholderPrefix + id + placeholderSuffix
}
