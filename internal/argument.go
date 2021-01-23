package internal

type ArgumentDefinition struct {
	Env      bool     `json:"env,omitempty"`
	Multiple bool     `json:"multiple,omitempty"`
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
