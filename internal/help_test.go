package internal

import "testing"

func TestCreateHelpCommandDefinition(t *testing.T) {
	assertNil(t, createHelpCommandDefinition(), false)
}

func TestCreateHelpArgumentDefinition(t *testing.T) {
	assertNil(t, createHelpArgumentDefinition(""), false)
}

func TestAddCommandAbbr(t *testing.T) {
	tests := []struct {
		token string
		cmd   *CommandDefinition
	}{
		{"", nil},
		{"", mockCommandDefinition("c", nil)},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			if tt.cmd != nil {
				assertNotEquals(t, tt.cmd.Abbr, tt.token)
			}
			addCommandAbbr(tt.token, tt.cmd, mockCommandByAbbr)
			if tt.cmd != nil {
				assertEquals(t, tt.cmd.Abbr, tt.token)
			}
		})
	}
}

func TestAddArgAbbr(t *testing.T) {
	tests := []struct {
		token string
		arg   *ArgumentDefinition
	}{
		{"", nil},
		{"", mockArgumentDefinition("a", nil)},
	}
	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			if tt.arg != nil {
				assertNotEquals(t, tt.arg.Abbr, tt.token)
			}
			addArgAbbr(tt.token, tt.arg, mockArgByAbbr)
			if tt.arg != nil {
				assertEquals(t, tt.arg.Abbr, tt.token)
			}
		})
	}
}

func TestAddHelpCommand(t *testing.T) {
	//tests := []struct{
	//	token string
	//	abbr string
	//	expNil bool
	//}
}
