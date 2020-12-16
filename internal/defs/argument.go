package defs

import "fmt"

type ArgumentDefinition struct {
	CommonDefinition
	Env      bool     `json:"env,omitempty"`
	Default  bool     `json:"default,omitempty"`
	Multiple bool     `json:"multiple,omitempty"`
	Required bool     `json:"required,omitempty"`
	Values   []string `json:"values,omitempty"`
}

func (arg *ArgumentDefinition) ValueSupported(val string) bool {
	for _, v := range arg.Values {
		if v == val {
			return true
		}
	}
	return false
}

func (arg *ArgumentDefinition) Print() {
	arg.CommonDefinition.Print()
	fmt.Printf("Default:%v\nMultiple:%v\nRequired:%v\nValues:%v\n",
		arg.Default, arg.Multiple, arg.Required, arg.Values)
}
