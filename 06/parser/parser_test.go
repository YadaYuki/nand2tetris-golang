package parser

import (
	"assembly/ast"
	"testing"
)

func TestAdvance(t *testing.T) {
	p := New("sample")
	p.Advance()
	if p.currentCommandIdx != 1 {
		t.Fatalf("p.currentCommandIdx should be 1 , but got %d", p.currentCommandIdx)
	}
}

func TestHasMoreCommand(t *testing.T) {
	p := New(`sample
	hoge`)
	p.Advance()
	if !p.HasMoreCommand() {
		t.Fatal("p.HasMoreCommand should be true , but got false")
	}
	p.Advance()
	if p.HasMoreCommand() {
		t.Fatalf("p.HasMoreCommand should be false , but got true")
	}
}

func TestSkipWhiteSpace(t *testing.T) {
	testCases := []struct {
		parser                          *Parser
		readPositionAfterSkipWhiteSpace int
	}{
		{&Parser{commandStrList: []string{"input"}, currentCommandIdx: 0}, 0},
		{&Parser{commandStrList: []string{" input"}, currentCommandIdx: 0}, 1},
		{&Parser{commandStrList: []string{"   input"}, currentCommandIdx: 0}, 3},
		{&Parser{commandStrList: []string{"   \tinput"}, currentCommandIdx: 0}, 4},
	}
	for _, tt := range testCases {
		tt.parser.skipWhiteSpace()
		if tt.parser.readPosition != tt.readPositionAfterSkipWhiteSpace {
			t.Fatalf("parser.readPosition should be %d,got %d", tt.readPositionAfterSkipWhiteSpace, tt.parser.readPosition)
		}
		if tt.parser.commandStrList[tt.parser.currentCommandIdx][tt.parser.readPosition] != byte('i') {
			t.Fatalf("readChar should be `i`,got %c", tt.parser.commandStrList[tt.parser.currentCommandIdx][tt.parser.readPosition])
		}
	}
}

func TestCommandType(t *testing.T) {
	testCases := []struct {
		parser      *Parser
		commandType ast.CommandType
	}{
		{&Parser{commandStrList: []string{"@10"}, currentCommandIdx: 0}, ast.A_COMMAND},
		{&Parser{commandStrList: []string{"D=M"}, currentCommandIdx: 0}, ast.C_COMMAND},
	}
	for _, tt := range testCases {
		commandType := tt.parser.CommandType()
		if commandType != tt.commandType {
			t.Fatalf("commandType() Should be %s, got %s", tt.commandType, commandType)
		}
	}
}

func TestParseACommand(t *testing.T) {
	testCases := []struct {
		parser *Parser
		value  int
	}{
		{&Parser{commandStrList: []string{"@10"}, currentCommandIdx: 0}, 10},
		{&Parser{commandStrList: []string{"@100"}, currentCommandIdx: 0}, 100},
	}
	for _, tt := range testCases {
		command, _ := tt.parser.parseACommand()
		if command.Value != tt.value {
			t.Fatalf("command.Value Should be %d, got %d", command.Value, tt.value)
		}
	}
}
