package compilationengine

import (
	"bytes"
	"jack_compiler/parser"
	"jack_compiler/symboltable"
	"jack_compiler/tokenizer"
	"jack_compiler/value"
	"jack_compiler/vmwriter"
	"testing"
	// "fmt"
)

func newParser(input string) *parser.Parser {
	jt := tokenizer.New(input)
	p := parser.New(jt)
	return p
}
func newCompilationEngine(className string) *CompilationEngine {
	vmWriter := vmwriter.New("test.vm", 0644)
	symbolTable := symboltable.New()
	ce := New(className, vmWriter, symbolTable)
	return ce
}

func TestExpression(t *testing.T) {
	testCases := []struct {
		expressionInput string
		vmCode          string
	}{
		{"7", "push constant 7" + value.NEW_LINE},
		{"7 + 8", "push constant 7" + value.NEW_LINE + "push constant 8" + value.NEW_LINE + "add" + value.NEW_LINE},
		{"2 * 2", "push constant 2" + value.NEW_LINE + "push constant 2" + value.NEW_LINE + "call Math.multiply 2" + value.NEW_LINE},
		{"4 * 3", "push constant 4" + value.NEW_LINE + "push constant 3" + value.NEW_LINE + "call Math.multiply 2" + value.NEW_LINE},
		{"(2+3)*(5+4)", "push constant 2" + value.NEW_LINE + "push constant 3" + value.NEW_LINE + "add" + value.NEW_LINE + "push constant 5" + value.NEW_LINE + "push constant 4" + value.NEW_LINE + "add" + value.NEW_LINE + "call Math.multiply 2" + value.NEW_LINE},
	}

	for _, tt := range testCases {
		p := newParser(tt.expressionInput)
		ast := p.ParseExpression()
		ce := newCompilationEngine("Main")
		ce.CompileExpression(ast)
		if !bytes.Equal([]byte(tt.vmCode), ce.VMCode) {
			t.Fatalf("Expression VMCode should be %s, got %s", tt.vmCode, ce.VMCode)
		}
	}
}

func TestDoStatement(t *testing.T) {
	testCases := []struct {
		expressionInput string
		vmCode          string
	}{
		{"do Output.printInt(1);", "push constant 1" + value.NEW_LINE + "call Output.printInt 1" + value.NEW_LINE + "pop temp 0" + value.NEW_LINE},
		{"do Output.printInt(1,3,4);", "push constant 1" + value.NEW_LINE + "push constant 3" + value.NEW_LINE + "push constant 4" + value.NEW_LINE + "call Output.printInt 3" + value.NEW_LINE + "pop temp 0" + value.NEW_LINE},
		{"do Output.printInt(1 + (2*3));", "push constant 1" + value.NEW_LINE + "push constant 2" + value.NEW_LINE + "push constant 3" + value.NEW_LINE + "call Math.multiply 2" + value.NEW_LINE + "add" + value.NEW_LINE + "call Output.printInt 1" + value.NEW_LINE + "pop temp 0" + value.NEW_LINE},
	}
	for _, tt := range testCases {
		p := newParser(tt.expressionInput)
		ast := p.ParseDoStatement()
		ce := newCompilationEngine("Main")
		ce.CompileDoStatement(ast)
		if !bytes.Equal([]byte(tt.vmCode), ce.VMCode) {
			t.Fatalf("doStatement VMCode should be %s, got %s", tt.vmCode, ce.VMCode)
		}
	}
}

func TestReturnStatement(t *testing.T) {
	testCases := []struct {
		expressionInput string
		vmCode          string
	}{
		{"return;", "push constant 0" + value.NEW_LINE + "return" + value.NEW_LINE},
		{"return 1;", "push constant 1" + value.NEW_LINE + "return" + value.NEW_LINE},
		{"return 1+2;", "push constant 1" + value.NEW_LINE + "push constant 2" + value.NEW_LINE + "add" + value.NEW_LINE + "return" + value.NEW_LINE},
	}
	for _, tt := range testCases {
		p := newParser(tt.expressionInput)
		ast := p.ParseReturnStatement()
		ce := newCompilationEngine("Main")
		ce.CompileReturnStatement(ast)
		if !bytes.Equal([]byte(tt.vmCode), ce.VMCode) {
			t.Fatalf("returnStatement VMCode should be %s, got %s", tt.vmCode, ce.VMCode)
		}
	}
}

func TestSubroutineDecStatement(t *testing.T) {
	testCases := []struct {
		expressionInput string
		vmCode          string
	}{
		{"function void main (){}", "function Main.main 0" + value.NEW_LINE},
		{"function void main (){do Output.printInt();}", "function Main.main 0" + value.NEW_LINE + "call Output.printInt 0" + value.NEW_LINE + "pop temp 0" + value.NEW_LINE},
		{"function void main (){return;}", "function Main.main 0" + value.NEW_LINE + "push constant 0" + value.NEW_LINE + "return" + value.NEW_LINE},
	}
	for _, tt := range testCases {
		p := newParser(tt.expressionInput)
		ast := p.ParseSubroutineDecStatement()
		ce := newCompilationEngine("Main")
		ce.CompileSubroutineDecStatement(ast)
		if !bytes.Equal([]byte(tt.vmCode), ce.VMCode) {
			t.Fatalf("subroutineDecStatement VMCode should be %s, got %s", tt.vmCode, ce.VMCode)
		}
	}
}
