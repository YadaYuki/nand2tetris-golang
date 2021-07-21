package parser

import (
	"VMtranslator/ast"
	"testing"
)

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

func TestCommandType(t *testing.T) {
	testCases := []struct {
		p           *Parser
		commandType ast.CommandType
	}{
		{&Parser{CurrentCommandIdx: 0, CurrentCommandTokenArr: []string{"push", "local", "1"}, CommandStrArr: []string{"push local 1", "", "add"}}, ast.C_PUSH},
		{&Parser{CurrentCommandIdx: 1, CurrentCommandTokenArr: []string{}, CommandStrArr: []string{"push local 1", ""}}, ast.C_EMPTY},
		{&Parser{CurrentCommandIdx: 2, CurrentCommandTokenArr: []string{"add"}, CommandStrArr: []string{"push local 1", "", "add"}}, ast.C_ARITHMETIC},
	}
	for _, tt := range testCases {
		if tt.p.CommandType() != tt.commandType {
			t.Fatalf("p.CommandType should be %s , but got %s", tt.commandType, tt.p.CommandType())
		}
	}
}
