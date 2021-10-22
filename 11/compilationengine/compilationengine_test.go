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
		{"4 - 2", "push constant 4" + value.NEW_LINE + "push constant 2" + value.NEW_LINE + "sub" + value.NEW_LINE},
		{"4 = 2", "push constant 4" + value.NEW_LINE + "push constant 2" + value.NEW_LINE + "eq" + value.NEW_LINE},
		{"4 & 2", "push constant 4" + value.NEW_LINE + "push constant 2" + value.NEW_LINE + "and" + value.NEW_LINE},
		{"4 | 2", "push constant 4" + value.NEW_LINE + "push constant 2" + value.NEW_LINE + "or" + value.NEW_LINE},
		{"4 > 2", "push constant 4" + value.NEW_LINE + "push constant 2" + value.NEW_LINE + "gt" + value.NEW_LINE},
		{"4 < 2", "push constant 4" + value.NEW_LINE + "push constant 2" + value.NEW_LINE + "lt" + value.NEW_LINE},
		{"4 / 2", "push constant 4" + value.NEW_LINE + "push constant 2" + value.NEW_LINE + "call Math.divide 2" + value.NEW_LINE},
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

func TestCompilePrefixTerm(t *testing.T) {
	testCases := []struct {
		expressionInput string
		vmCode          string
	}{
		{"-7", "push constant 7" + value.NEW_LINE + "neg" + value.NEW_LINE},
		{"~4", "push constant 4" + value.NEW_LINE + "not" + value.NEW_LINE},
	}
	for _, tt := range testCases {
		p := newParser(tt.expressionInput)
		prefixTermAst := p.ParsePrefixTerm()
		ce := newCompilationEngine("Main")
		ce.CompilePrefixTerm(prefixTermAst)
		if !bytes.Equal([]byte(tt.vmCode), ce.VMCode) {
			t.Fatalf("prefixTermAst VMCode should be %s, got %s", tt.vmCode, ce.VMCode)
		}
	}
}

func TestCompileIdentifierTerm(t *testing.T) {
	testCases := []struct {
		identifierTermInput string
		varKindInput        symboltable.VarKind
		vmCode              string
	}{
		{"a", symboltable.ARGUMENT, "push argument 0" + value.NEW_LINE},
		{"b", symboltable.VAR, "push local 0" + value.NEW_LINE},
	}
	for _, tt := range testCases {
		p := newParser(tt.identifierTermInput)
		identifierTermAst := p.ParseIdentifierTerm()
		ce := newCompilationEngine("Main")
		// 関数スコープで変数をシンボルテーブルに登録する。
		ce.StartSubroutine()
		ce.Define(tt.identifierTermInput, "int", tt.varKindInput)

		ce.CompileIdentifierTerm(identifierTermAst)
		if !bytes.Equal([]byte(tt.vmCode), ce.VMCode) {
			t.Fatalf("identifierTermAst VMCode should be %s, got %s", tt.vmCode, ce.VMCode)
		}
	}
}

func TestCompileSubroutineCallTerm(t *testing.T) {
	testCases := []struct {
		subroutineCallTermInput string
		vmCode                  string
	}{
		{"Main.add()", "call Main.add 0" + value.NEW_LINE},
		{"Main.add(1,2)", "push constant 1" + value.NEW_LINE + "push constant 2" + value.NEW_LINE + "call Main.add 2" + value.NEW_LINE},
	}
	for _, tt := range testCases {
		p := newParser(tt.subroutineCallTermInput)
		subroutineCallTermAst := p.ParseSubroutineCallTerm()
		ce := newCompilationEngine("Main")
		ce.CompileSubroutineCallTerm(subroutineCallTermAst)
		if !bytes.Equal([]byte(tt.vmCode), ce.VMCode) {
			t.Fatalf("identifierTermAst VMCode should be %s, got %s", tt.vmCode, ce.VMCode)
		}
	}
}

func TestStringConstTerm(t *testing.T) {
	testCases := []struct {
		expressionInput string
		vmCode          string
	}{
		{`"a"`, "push constant 1" + value.NEW_LINE + "call String.new 1" + value.NEW_LINE + "push constant 97" + value.NEW_LINE + "call String.appendChar 2" + value.NEW_LINE},
		{`"abcde"`, "push constant 5" + value.NEW_LINE + "call String.new 1" + value.NEW_LINE + "push constant 97" + value.NEW_LINE + "call String.appendChar 2" + value.NEW_LINE + "push constant 98" + value.NEW_LINE + "call String.appendChar 2" + value.NEW_LINE + "push constant 99" + value.NEW_LINE + "call String.appendChar 2" + value.NEW_LINE + "push constant 100" + value.NEW_LINE + "call String.appendChar 2" + value.NEW_LINE + "push constant 101" + value.NEW_LINE + "call String.appendChar 2" + value.NEW_LINE},
	}
	for _, tt := range testCases {
		p := newParser(tt.expressionInput)
		ast := p.ParseExpression()
		ce := newCompilationEngine("Main")
		ce.CompileExpression(ast)
		if !bytes.Equal([]byte(tt.vmCode), ce.VMCode) {
			t.Fatalf("StringConstTerm VMCode should be %s, got %s", tt.vmCode, ce.VMCode)
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

func TestClassStatement(t *testing.T) {
	testCases := []struct {
		expressionInput string
		vmCode          string
	}{
		{"Class Main {}", ""},
		{"Class Main {function void main(){}}", "function Main.main 0" + value.NEW_LINE},
	}
	for _, tt := range testCases {
		p := newParser(tt.expressionInput)
		ast := p.ParseClassStatement()
		ce := newCompilationEngine("Main")
		ce.CompileClassStatement(ast)
		if !bytes.Equal([]byte(tt.vmCode), ce.VMCode) {
			t.Fatalf("classStatement VMCode should be %s, got %s", tt.vmCode, ce.VMCode)
		}
	}
}
