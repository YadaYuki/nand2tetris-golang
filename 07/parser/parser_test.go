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

func TestAdvance(t *testing.T) {
	testCases := []struct {
		p                       *Parser
		commandTypeAfterAdvance ast.CommandType
	}{
		{&Parser{CurrentCommandIdx: 0, CurrentCommandTokenArr: []string{"push", "local", "1"}, CommandStrArr: []string{"push local 1", "", "add"}}, ast.C_ARITHMETIC},
		{&Parser{CurrentCommandIdx: 2, CurrentCommandTokenArr: []string{"push", "local", "1"}, CommandStrArr: []string{"push local 1", "", "add", "pop local 2"}}, ast.C_POP},
	}
	for _, tt := range testCases {
		tt.p.Advance()
		if tt.p.CommandType() != tt.commandTypeAfterAdvance {
			t.Fatalf("p.CommandType after Advance should be %s , but got %s", tt.commandTypeAfterAdvance, tt.p.CommandType())
		}
	}
}

func TestArg1(t *testing.T) {
	testCases := []struct {
		p    *Parser
		arg1 string
	}{
		{&Parser{CurrentCommandIdx: 0, CommandStrArr: []string{"push local 1"}, CurrentCommandTokenArr: []string{"push", "local", "1"}}, "local"},
		{&Parser{CurrentCommandIdx: 0, CommandStrArr: []string{""}, CurrentCommandTokenArr: []string{""}}, ""},
		{&Parser{CurrentCommandIdx: 0, CommandStrArr: []string{"add"}, CurrentCommandTokenArr: []string{"add"}}, "add"},
	}

	for _, tt := range testCases {
		if arg1, _ := tt.p.Arg1(); arg1 != tt.arg1 {
			t.Fatalf("p.Arg1 should be %s , but got %s", tt.arg1, arg1)
		}
	}
}
