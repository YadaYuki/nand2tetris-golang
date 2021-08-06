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
func newCompilationEngine() *CompilationEngine {
	vmWriter := vmwriter.New("test.vm", 0644)
	symbolTable := symboltable.New()
	ce := New(vmWriter, symbolTable)
	return ce
}

func TestVarDecStatements(t *testing.T) {
	input := "var int temp;"
	p := newParser(input)
	ast := p.ParseProgram()
	ce := newCompilationEngine()
	ce.CompileProgram(ast)
	if !bytes.Equal([]byte("if-goto hoge"+value.NEW_LINE), ce.VMCode) {
		t.Fatalf("VarDecStatement VMCode should be %s, got %s", "hoge", ce.VMCode)
	}
}

func TestExpression(t *testing.T) {
	testCases := []struct {
		expressionInput string
		vmCode          string
	}{
		{"7", "push constant 7" + value.NEW_LINE},
		{"7 + 8", "push constant 7" + value.NEW_LINE + "push constant 8" + value.NEW_LINE + "add" + value.NEW_LINE},
		{"2 * 2", "push constant 2" + value.NEW_LINE + "push constant 2" + value.NEW_LINE + "call mul 2" + value.NEW_LINE},
		{"4 * 3", "push constant 4" + value.NEW_LINE + "push constant 3" + value.NEW_LINE + "call mul 2" + value.NEW_LINE},
	}

	for _, tt := range testCases {
		p := newParser(tt.expressionInput)
		ast := p.ParseExpression()
		ce := newCompilationEngine()
		ce.CompileExpression(ast)
		if !bytes.Equal([]byte(tt.vmCode), ce.VMCode) {
			t.Fatalf("Expression VMCode should be %s, got %s", tt.vmCode, ce.VMCode)
		}
	}
}
