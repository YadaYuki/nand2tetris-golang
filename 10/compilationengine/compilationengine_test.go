package compilationengine

import (
	"jack_compiler/ast"
	"jack_compiler/tokenizer"
	"testing"
	// "fmt"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x=5;
		let y=10;
		let hoge=111;
		let foo=838383;
		let bar="hogehoge";
		`

	jt := tokenizer.New(input)
	ce := New(jt)
	program := ce.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 5 {
		t.Fatalf("program.Statements does not contain 5 statements. got=%d", len(program.Statements))
	}
	testCases := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"hoge"},
		{"foo"},
		{"bar"},
	}
	for i, tt := range testCases {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got %q", s.TokenLiteral())
	}
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
	// if s.String() == fmt.Sprintf("let %s = %s;", s.TokenLiteral(), letStmt.Value.String()) {
	// 	return true
	// }
	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
	return x;
	return 1;
	return;
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
	for i, tt := range testCases {
		stmt := program.Statements[i]
		if !testReturnStatement(t, stmt, tt.expectedIdentifier) {
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

func TestDoStatements(t *testing.T) {
	input := `
	do x;
	do 1;
	do a;
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
	for i, tt := range testCases {
		stmt := program.Statements[i]
		if !testDoStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testDoStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "do" {
		t.Errorf("s.TokenLiteral not 'do'. got %q", s.TokenLiteral())
		return false
	}
	_, ok := s.(*ast.DoStatement)
	if !ok {
		t.Errorf("s not *ast.DoStatement. got %T", s)
		return false
	}
	return true
}

func TestVarDecStatements(t *testing.T) {
	input := `
	var int a,b,c;
	var char casdfasdf;
	var boolean a1,b2,cx;
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
		expectedValueType   string
		expectedIdentifiers []string
	}{
		{"int", []string{"a", "b", "c"}},
		{"char", []string{"casdfasdf"}},
		{"boolean", []string{"a1", "b2", "cx"}},
	}
	for i, tt := range testCases {
		stmt := program.Statements[i]
		if !testVarDecStatement(t, stmt, tt.expectedValueType, tt.expectedIdentifiers) {
			return
		}
	}
}

func testVarDecStatement(t *testing.T, s ast.Statement, expectedValueType string, identifiers []string) bool {
	if s.TokenLiteral() != "var" {
		t.Errorf("s.TokenLiteral not 'var'. got %q", s.TokenLiteral())
		return false
	}
	vds, ok := s.(*ast.VarDecStatement)
	if !ok {
		t.Errorf("s not *ast.VarDecStatement. got %T", s)
		return false
	}
	if vds.ValueType.Literal != expectedValueType {
		t.Errorf("valueType no %s . got %s", expectedValueType, vds.ValueType.Literal)
		return false
	}
	for i, ident := range vds.Identifiers {
		if identifiers[i] != ident.String() {
			t.Errorf("identifiers no %s . got %s", identifiers[i], ident)
			return false
		}
	}
	return true
}

func TestClassVarDecStatements(t *testing.T) {
	input := `
	static int a,b,c;
	field char casdfasdf;
	static boolean a1,b2,cx;
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
		expectedValueType   string
		expectedIdentifiers []string
	}{
		{"int", []string{"a", "b", "c"}},
		{"char", []string{"casdfasdf"}},
		{"boolean", []string{"a1", "b2", "cx"}},
	}
	for i, tt := range testCases {
		stmt := program.Statements[i]
		if !testClassVarDecStatement(t, stmt, tt.expectedValueType, tt.expectedIdentifiers) {
			return
		}
	}
}

func testClassVarDecStatement(t *testing.T, s ast.Statement, expectedValueType string, identifiers []string) bool {
	if s.TokenLiteral() != "static" && s.TokenLiteral() != "field" {
		t.Errorf("s.TokenLiteral not 'static' and 'field'. got %q", s.TokenLiteral())
		return false
	}
	vds, ok := s.(*ast.ClassVarDecStatement)
	if !ok {
		t.Errorf("s not *ast.ClassVarDecStatement. got %T", s)
		return false
	}
	if vds.ValueType.Literal != expectedValueType {
		t.Errorf("valueType no %s . got %s", expectedValueType, vds.ValueType.Literal)
		return false
	}
	for i, ident := range vds.Identifiers {
		if identifiers[i] != ident.String() {
			t.Errorf("identifiers no %s . got %s", identifiers[i], ident)
			return false
		}
	}
	return true
}

func TestParseIntConstTermExpression(t *testing.T) {
	input := `33`
	jt := tokenizer.New(input)
	ce := New(jt)
	expression := ce.parseExpression(LOWEST)
	singleExpression, ok := expression.(*ast.SingleExpression)
	if !ok {
		t.Fatalf("expression is not ast.SingleExpression,got = %T", expression)
	}
	intergerConstTerm, ok := singleExpression.Value.(*ast.IntergerConstTerm)
	if !ok {
		t.Fatalf("intergerConstTerm is not ast.IntergerConstTerm,got = %T", intergerConstTerm)
	}
	if intergerConstTerm.Value != 33 {
		t.Fatalf("intergerConstTerm.Value is not 1,got = %d", intergerConstTerm.Value)
	}
}

func TestParseIdentifierTermExpression(t *testing.T) {
	input := `hoge`
	jt := tokenizer.New(input)
	ce := New(jt)
	expression := ce.parseExpression(LOWEST)
	singleExpression, ok := expression.(*ast.SingleExpression)
	if !ok {
		t.Fatalf("expression is not ast.SingleExpression,got = %T", expression)
	}
	identifierTerm, ok := singleExpression.Value.(*ast.IdentifierTerm)
	if !ok {
		t.Fatalf("identifierTerm is not ast.IdentifierTerm,got = %T", identifierTerm)
	}
	if identifierTerm.Value != "hoge" {
		t.Fatalf("identifierTerm.Value is not hoge,got = %s", identifierTerm.Value)
	}
}

func TestParseStringConstTermExpression(t *testing.T) {
	input := `"hoge"`
	jt := tokenizer.New(input)
	ce := New(jt)
	expression := ce.parseExpression(LOWEST)
	singleExpression, ok := expression.(*ast.SingleExpression)
	if !ok {
		t.Fatalf("expression is not ast.SingleExpression,got = %T", expression)
	}
	stringConstTerm, ok := singleExpression.Value.(*ast.StringConstTerm)
	if !ok {
		t.Fatalf("stringConstTerm is not ast.StringConstTerm,got = %T", stringConstTerm)
	}
	if stringConstTerm.Value != "hoge" {
		t.Fatalf("stringConstTerm.Value is not hoge,got = %s", stringConstTerm.Value)
	}
}
