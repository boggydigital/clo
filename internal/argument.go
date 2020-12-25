package internal

type ArgumentDefinition struct {
	CommonDefinition
	Env      bool     `json:"env,omitempty"`
	Default  bool     `json:"default,omitempty"`
	Multiple bool     `json:"multiple,omitempty"`
	Required bool     `json:"required,omitempty"`
	Values   []string `json:"values,omitempty"`
}

func (arg *ArgumentDefinition) ValidValue(val string) bool {
	for _, v := range arg.Values {
		if v == val {
			return true
		}
	}
	return false
}
