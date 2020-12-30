package internal

import "testing"

func TestCommandsPadding(t *testing.T) {
	defs := Definitions{
		Commands: []CommandDefinition{
			{
				CommonDefinition: CommonDefinition{
					Token: "a",
				},
			},
			{
				CommonDefinition: CommonDefinition{
					Token: "bc",
				},
			},
			{
				CommonDefinition: CommonDefinition{
					Token: "def",
				},
			},
		},
	}

	if defs.CommandsPadding() != 3 {
		t.Error("unexpected commands padding")
	}
}

func TestFlagsPadding(t *testing.T) {
	defs := Definitions{
		Flags: []FlagDefinition{
			{
				CommonDefinition: CommonDefinition{
					Token: "a",
				},
			},
			{
				CommonDefinition: CommonDefinition{
					Token: "bc",
				},
			},
			{
				CommonDefinition: CommonDefinition{
					Token: "def",
				},
			},
		},
	}

	if defs.FlagsPadding() != 3 {
		t.Error("unexpected flags padding")
	}
}

func TestArgumentsPadding(t *testing.T) {
	defs := Definitions{
		Commands: []CommandDefinition{
			{
				CommonDefinition: CommonDefinition{
					Token: "cmd",
				},
				Arguments: []string{
					"a", "bc", "def",
				},
			},
		},
	}
	if defs.ArgumentsPadding("cmd") != 3 {
		t.Error("unexpected arguments padding for valid command")
	}

	if defs.ArgumentsPadding("command-that-doesnt-exist") > 0 {
		t.Error("unexpected arguments padding for command that doesn't exist")
	}
}
