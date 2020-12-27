package internal

type CommonDefinition struct {
	Token string `json:"token"`
	Abbr  string `json:"abbr,omitempty"`
	Hint  string `json:"hint,omitempty"`
	Desc  string `json:"desc,omitempty"`
}
