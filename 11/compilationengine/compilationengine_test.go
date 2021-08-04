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

var commonVmWriter *vmwriter.VMWriter = vmwriter.New("test.vm", 0644)
var commonSymbolTable *symboltable.SymbolTable = symboltable.New()

func newParser(input string) *parser.Parser {
	jt := tokenizer.New(input)
	p := parser.New(jt)
	return p
}
func newCompilationEngine() *CompilationEngine {
	ce := New(commonVmWriter, commonSymbolTable)
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
