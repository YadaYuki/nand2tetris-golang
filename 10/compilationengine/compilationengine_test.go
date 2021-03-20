package compilationengine

import (
	"jack/compiler/ast"
	"jack/compiler/tokenizer"
	"testing"
	
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let hoge = 111;
		let foobar = 838383;
		`

	jt := tokenizer.New(input)
	ce := New(jt)
	program := ce.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 4 {
		t.Fatalf("program.Statements does not contain 4 statements. got=%d", len(program.Statements))
	}
	testCases := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"hoge"},
		{"foobar"},
	}
	for i, tt := range testCases {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	// TODO:Fix panic

	// if s.TokenLiteral() != "let" {
	// 	t.Errorf("s.TokenLiteral not 'let'. got %q", s.TokenLiteral())
	// }
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got %T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'.got '%s'", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'.got '%s'", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T){
	input := `
	return x;
	return 1;
	return ;
`
	jt := tokenizer.New(input)
	ce := New(jt)
	program := ce.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}
	testCases := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"1"},
		{},
	}
	for i,tt := range testCases{
		stmt := program.Statements[i]
		if !testReturnStatement(t,stmt,tt.expectedIdentifier){
			return 
		}
	}
}

func testReturnStatement(t *testing.T, s ast.Statement, name string) bool {	
	if s.TokenLiteral() != "return" {
		t.Errorf("s.TokenLiteral not 'return'. got %q", s.TokenLiteral())
		return false
	}
	_, ok := s.(*ast.ReturnStatement)
	if !ok {
		t.Errorf("s not *ast.ReturnStatement. got %T", s)
		return false
	}
	return true
}

func TestIdentifierExpression(t *testing.T){
	input := "foobar;"
	jt := tokenizer.New(input)
	ce := New(jt)
	program := ce.ParseProgram()
	if len(program.Statements) != 1{
		t.Fatalf("program has not enough statements. got=%d",len(program.Statements))
	}
	stmt,ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok{
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",program.Statements[0])
	}
	ident,ok := stmt.Expression.(*ast.Identifier)
	if !ok{
		t.Fatalf("exp not *ast.Identifier. got=%T",stmt.Expression)
	}
	if ident.Value != "foobar"{
		t.Errorf("ident.Value not %s. got %s","foobar",ident.Value)
	}
	if ident.TokenLiteral() != "foobar"{
		t.Errorf("ident.TokenLiteral() not %s. got %s","foobar",ident.TokenLiteral())
	}
}

