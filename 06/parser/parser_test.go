package parser

import (
	"assembler/ast"
	"assembler/symboltable"
	"testing"
)

var st = symboltable.New()

func TestAdvance(t *testing.T) {
	p := New("sample", st)
	p.Advance()
	if p.currentCommandIdx != 1 {
		t.Fatalf("p.currentCommandIdx should be 1 , but got %d", p.currentCommandIdx)
	}
}

func TestHasMoreCommand(t *testing.T) {
	p := New(`sample
	hoge`, st)
	p.Advance()
	if !p.HasMoreCommand() {
		t.Fatal("p.HasMoreCommand should be true , but got false")
	}
	p.Advance()
	if p.HasMoreCommand() {
		t.Fatalf("p.HasMoreCommand should be false , but got true")
	}
}

func TestCommandType(t *testing.T) {
	testCases := []struct {
		parser      *Parser
		commandType ast.CommandType
	}{
		{&Parser{commandStrList: []string{"@10"}, currentCommandIdx: 0, SymbolTable: st}, ast.A_COMMAND},
		{&Parser{commandStrList: []string{"D=M"}, currentCommandIdx: 0, SymbolTable: st}, ast.C_COMMAND},
		{&Parser{commandStrList: []string{"(SAMPLE)"}, currentCommandIdx: 0, SymbolTable: st}, ast.L_COMMAND},
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
		parser   *Parser
		valueStr string
	}{
		{&Parser{commandStrList: []string{"@10"}, currentCommandIdx: 0}, "10"},
		{&Parser{commandStrList: []string{"@100"}, currentCommandIdx: 0}, "100"},
	}
	for _, tt := range testCases {
		command, _ := tt.parser.parseACommand()
		if command.ValueStr != tt.valueStr {
			t.Fatalf("command.Value Should be %s, got %s", command.ValueStr, tt.valueStr)
		}
	}
}
func TestParseCCommand(t *testing.T) {
	testCases := []struct {
		parser  *Parser
		command *ast.CCommand
	}{
		{&Parser{commandStrList: []string{"D=M//HOGE"}, currentCommandIdx: 0, SymbolTable: st}, &ast.CCommand{
			Comp: "M",
			Dest: "D",
			Jump: "",
		}},
		{&Parser{commandStrList: []string{"D=D-M"}, currentCommandIdx: 0, SymbolTable: st}, &ast.CCommand{
			Comp: "D-M",
			Dest: "D",
			Jump: "",
		}},
		{&Parser{commandStrList: []string{"0;JMP"}, currentCommandIdx: 0, SymbolTable: st}, &ast.CCommand{
			Comp: "0",
			Dest: "",
			Jump: "JMP",
		}},
		{&Parser{commandStrList: []string{"AM=D|A;JMP"}, currentCommandIdx: 0, SymbolTable: st}, &ast.CCommand{Comp: "D|A", Dest: "AM", Jump: "JMP"}},
	}
	for _, tt := range testCases {
		command, _ := tt.parser.parseCCommand()
		if command.Comp != tt.command.Comp {
			t.Fatalf("command.Comp Should be %s, got %s", tt.command.Comp, command.Comp)
		}
		if command.Dest != tt.command.Dest {
			t.Fatalf("command.Dest Should be %s, got %s", tt.command.Dest, command.Dest)
		}
		if command.Jump != tt.command.Jump {
			t.Fatalf("command.Jump Should be %s, got %s", tt.command.Jump, command.Jump)
		}
	}
}

func TestParseLCommand(t *testing.T) {
	testCases := []struct {
		parser *Parser
		symbol string
	}{
		{&Parser{commandStrList: []string{"(HOGE)"}, currentCommandIdx: 0, SymbolTable: st}, "HOGE"},
		{&Parser{commandStrList: []string{"(FUGA)"}, currentCommandIdx: 0, SymbolTable: st}, "FUGA"},
	}
	for _, tt := range testCases {
		command, _ := tt.parser.parseLCommand()
		if command.Symbol != tt.symbol {
			t.Fatalf("command.Symbol Should be %s, got %s", command.Symbol, tt.symbol)
		}
	}
}

func TestParseSymbol(t *testing.T) {
	testCases := []struct {
		parser *Parser
		symbol string
	}{
		{&Parser{commandStrList: []string{"(HOGE)"}, currentCommandIdx: 0, SymbolTable: st}, "HOGE"},
		{&Parser{commandStrList: []string{"@FUGA"}, currentCommandIdx: 0, SymbolTable: st}, "FUGA"},
	}
	for _, tt := range testCases {
		symbol, _ := tt.parser.Symbol()
		if symbol != tt.symbol {
			t.Fatalf("symbol Should be %s, got %s", symbol, tt.symbol)
		}
	}
}
