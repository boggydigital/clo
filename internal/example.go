package internal

type ExampleDefinition struct {
	ArgumentsValues []map[string][]string `json:"argumentsValues"`
	//Arguments []string `json:"arguments"`
	//Values    []string `json:"values"`
	Desc string `json:"desc,omitempty"`
}
