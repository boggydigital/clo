package clove

import "testing"

func TestParseEmptyArgsNilRequest(t *testing.T) {
	if req, _ := Parse([]string{}); req != nil {
		t.Error("expected nil Request")
	}
	if req, _ := Parse(nil); req != nil {
		t.Error("expected nil Request")
	}
}
