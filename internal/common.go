package internal

type CommonDefinition struct {
	Token string `json:"token"`
	Abbr  string `json:"abbr,omitempty"`
	Help  string `json:"help,omitempty"`
}
