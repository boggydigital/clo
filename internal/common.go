package internal

type CommonDefinition struct {
	Token string `json:"token"`
	Abbr  string `json:"abbr,omitempty"`
	Hint  string `json:"hint,omitempty"`
	Desc  string `json:"desc,omitempty"`
}

func commonByToken(commons []CommonDefinition, token string) *CommonDefinition {
	for _, c := range commons {
		if c.Token == token {
			return &c
		}
	}
	return nil
}
