package internal

type ExampleDefinition struct {
	Arguments []string `json:"arguments"`
	Values    []string `json:"values"`
	Desc      string   `json:"desc,omitempty"`
}
