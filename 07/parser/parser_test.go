package parser

import "testing"

func TestHasMoreCommand(t *testing.T) {
	testCases := []struct {
		p              *Parser
		hasMoreCommand bool
	}{
		{&Parser{CurrentCommandIdx: 2, CommandStrArr: []string{"", ""}}, false},
		{&Parser{CurrentCommandIdx: 1, CommandStrArr: []string{"", ""}}, true},
	}
	for _, tt := range testCases {
		if tt.p.HasMoreCommand() != tt.hasMoreCommand {
			t.Fatalf("p.HasMoreCommand should be %T , but got %T", tt.p.HasMoreCommand(), tt.hasMoreCommand)
		}
	}
}
