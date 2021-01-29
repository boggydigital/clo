package internal

func mockDefinitions() *Definitions {
	return &Definitions{
		Version: 1,
		Cmd: map[string][]string{
			"command1_": {"argument1_!$", "argument2..."},
			"command2":  {"argument2...", "xyz"},
			"abc":       {"argval=value1,value2"},
		},
	}
}
