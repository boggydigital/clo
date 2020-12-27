package internal

type ExampleDefinition struct {
	ArgumentsValues []map[string][]string `json:"argumentsValues"`
	Desc            string                `json:"desc,omitempty"`
}
