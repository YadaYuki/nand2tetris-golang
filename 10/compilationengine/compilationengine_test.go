package compilationengine

import (
	"jack_compiler/ast"
	"jack_compiler/token"
	"jack_compiler/tokenizer"
	"testing"
	// "fmt"
)

func TestLetStatements(t *testing.T) {
	input := `let x=5;
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
	if letStmt.Name.Literal != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'.got '%s'", name, letStmt.Name.Literal)
		return false
	}
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
		if identifiers[i] != ident.Literal {
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
		if identifiers[i] != ident.Literal {
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
	expression := ce.parseExpression()
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
	expression := ce.parseExpression()
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
	expression := ce.parseExpression()
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

func TestParseIfStatement(t *testing.T) {
	input := `
	 if(x=1){
		do x;
		do 1;
		do a;
	 }else{
		do x;
		do 1;
		do a;
	 }`
	jt := tokenizer.New(input)
	ce := New(jt)
	program := ce.ParseProgram()
	if len(program.Statements) != 1 {
		t.Fatalf("len(program.Statements) is not 1,got = %d", len(program.Statements))
	}
	ifStmt, ok := program.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("ifStmt is not ast.IfStatement,got = %T", ifStmt)
	}
	if len(ifStmt.Consequence.Statements) != 3 {
		t.Fatalf("len(ifStmt.Consequence)  is not 3,got = %d", len(ifStmt.Consequence.Statements))
	}
}

func TestParseWhileStatement(t *testing.T) {
	input := `
	 while(x=1){
		do x;
		do 1;
		do a;
	 }`
	jt := tokenizer.New(input)
	ce := New(jt)
	program := ce.ParseProgram()
	if len(program.Statements) != 1 {
		t.Fatalf("len(program.Statements) is not 1,got = %d", len(program.Statements))
	}
	whileStmt, ok := program.Statements[0].(*ast.WhileStatement)
	if !ok {
		t.Fatalf("whileStmt is not ast.WhileStatement,got = %T", whileStmt)
	}
	if len(whileStmt.Statements.Statements) != 3 {
		t.Fatalf("len(whileStmt.Statements.Statements)  is not 3,got = %d", len(whileStmt.Statements.Statements))
	}
	// t.Log(whileStmt.Xml())
}

func TestParseExpressionListStatement(t *testing.T) {
	input := `(a,b,c,d,e,f)`
	jt := tokenizer.New(input)
	ce := New(jt)
	expressionListStmt := ce.parseExpressionListStatement()
	if len(expressionListStmt.ExpressionList) != 6 {
		t.Fatalf("len(program.Statements) is not 1,got = %d", len(expressionListStmt.ExpressionList))
	}
}

func TestParseSubroutineCallExpression(t *testing.T) {
	input := `hoge(a,b,c,d,e,f)`
	jt := tokenizer.New(input)
	ce := New(jt)
	expression := ce.parseSubroutineCallExpression()
	singleExpression, ok := expression.(*ast.SingleExpression)
	if !ok {
		t.Fatalf("expression is not ast.SingleExpression,got = %T", expression)
	}
	subroutineCallTerm, ok := singleExpression.Value.(*ast.SubroutineCallTerm)
	if !ok {
		t.Fatalf("subroutineCallTerm is not ast.SubroutineCallTerm,got = %T", subroutineCallTerm)
	}
	if len(subroutineCallTerm.ExpressionListStmt.ExpressionList) != 6 {
		t.Fatalf("len(subroutineCallTerm.ExpressionListStmt.ExpressionList) is not 6,got = %d", len(subroutineCallTerm.ExpressionListStmt.ExpressionList))
	}
}

func TestParseArrayElementExpression(t *testing.T) {
	input := `hoge[a]`
	jt := tokenizer.New(input)
	ce := New(jt)
	expression := ce.parseArrayElementExpression()
	singleExpression, ok := expression.(*ast.SingleExpression)
	if !ok {
		t.Fatalf("expression is not ast.SingleExpression,got = %T", expression)
	}
	arrayElementTerm, ok := singleExpression.Value.(*ast.ArrayElementTerm)
	if !ok {
		t.Fatalf("arrayElementTerm is not ast.ArrayElementTerm,got = %T", arrayElementTerm)
	}
	if arrayElementTerm.Idx.TokenLiteral() == "4" {
		t.Fatalf("arrayElementTerm.Idx.TokenLiteral() is not `4`,got = %s", arrayElementTerm.Idx.TokenLiteral())
	}
}
func TestParsePrefixExpression(t *testing.T) {
	input := `-4`
	jt := tokenizer.New(input)
	ce := New(jt)
	term := ce.parsePrefixExpression()
	prefixTerm, ok := term.(*ast.PrefixTerm)
	if !ok {
		t.Fatalf("prefixTerm is not ast.PrefixTerm,got = %T", prefixTerm)
	}
	if prefixTerm.Prefix != token.MINUS {
		t.Fatalf("prefixTerm.Prefix is not token.MINUS,got = %s", prefixTerm.Prefix)
	}
}

func TestParseBracketExpression(t *testing.T) {
	input := `(4)`
	jt := tokenizer.New(input)
	ce := New(jt)
	expression := ce.parseBracketExpression()
	singleExpression, ok := expression.(*ast.SingleExpression)
	if !ok {
		t.Fatalf("expression is not ast.SingleExpression,got = %T", expression)
	}
	bracketTerm, ok := singleExpression.Value.(*ast.BracketTerm)
	if !ok {
		t.Fatalf("bracketTerm is not ast.BracketTerm,got = %T", bracketTerm)
	}
	// t.Log(bracketTerm.Value.TokenLiteral())
	value, ok := bracketTerm.Value.(*ast.SingleExpression)
	if !ok {
		t.Fatalf("bracketTerm.Value is not ast.*ast.SingleExpression,got = %T", bracketTerm.Value)
	}
	intergerConstTerm, ok := value.Value.(*ast.IntergerConstTerm)
	if !ok {
		t.Fatalf("intergerConstTerm.Value is not ast.*ast.IntergerConstTerm,got = %T", bracketTerm.Value)
	}
	if intergerConstTerm.Value != 4 {
		t.Fatalf("value.Value is not 4,got = %d", value.Value)
	}
	t.Log(expression.Xml())
}

func TestParseParameterStatement(t *testing.T) {
	input := `int hoge`
	jt := tokenizer.New(input)
	ce := New(jt)
	stmt := ce.parseParameterStatement()
	if stmt.Name != "hoge" {
		t.Fatalf("stmt.Name is not hoge,got = %s", stmt.Name)
	}
	if stmt.Type != token.INT {
		t.Fatalf("stmt.Type is not token.INT,got = %s", stmt.Type)
	}
}

func TestParseParameterListStatement(t *testing.T) {
	input := `(int hoge,char fuga,boolean pepe)`
	jt := tokenizer.New(input)
	ce := New(jt)
	stmt := ce.parseParameterListStatement()
	if len(stmt.ParameterList) != 3 {
		t.Fatalf("len(stmt.ParameterList) is not 3 ,got = %d", len(stmt.ParameterList))
	}
	testCases := []struct {
		expectedType       token.KeyWord
		expectedIdentifier string
	}{
		{token.INT, "hoge"},
		{token.CHAR, "fuga"},
		{token.BOOLEAN, "pepe"},
	}
	for idx, testCase := range testCases {
		if stmt.ParameterList[idx].Name != testCase.expectedIdentifier {
			t.Fatalf("stmt.ParameterList[%d].Name is not %s,got %s", idx, testCase.expectedIdentifier, stmt.ParameterList[idx].Name)
		}
		if stmt.ParameterList[idx].Type != testCase.expectedType {
			t.Fatalf("stmt.ParameterList[%d].Type is not %s,got %s", idx, testCase.expectedType, stmt.ParameterList[idx].Type)
		}
	}
}

func TestParseClassStatement(t *testing.T) {
	input := `class hoge {
		let x=5;
		let y=10;
		let hoge=111;
		let foo=838383;
		let bar="hogehoge";
	}`
	jt := tokenizer.New(input)
	ce := New(jt)
	stmt := ce.parseClassStatement()
	if stmt.Name.Literal != "hoge" {
		t.Fatalf("stmt.Name  is not hoge ,got = %s", stmt.Name)
	}
	if token.KeyWord(stmt.Token.Literal) != token.CLASS {
		t.Fatalf("token.KeyWord(stmt.Token.Literal)  is not token.CLASS ,got = %s", token.KeyWord(stmt.Token.Literal))
	}
	if len(stmt.Statements.Statements) != 5 {
		t.Fatalf("len(stmt.Statements.Statements) is not 5 ,got = %d", len(stmt.Statements.Statements))
	}
}

func TestParseSubroutineDecStatement(t *testing.T) {
	input := `method void fuga (int hoge,boolean fuga){
		let x=5;
		let y=10;
		let hoge=111;
		let foo=838383;
		let bar="hogehoge";
		return hoge;
	}`
	jt := tokenizer.New(input)
	ce := New(jt)
	stmt := ce.ParseProgram()
	if len(stmt.Statements) != 1 {
		t.Fatalf("len(stmt.Statements)  is not hoge ,got = %d", len(stmt.Statements))
	}
	subroutineDecStmt, ok := stmt.Statements[0].(*ast.SubroutineDecStatement)
	if !ok {
		t.Fatalf("stmt.Statements[0]  is not SubroutineDecStatement ,got = %T", stmt.Statements[0])
	}
	if token.KeyWord(subroutineDecStmt.ReturnType.Literal) != token.VOID {
		t.Fatalf("subroutineDecStmt.ReturnType.Literal  is not void ,got = %s", subroutineDecStmt.ReturnType.Literal)
	}
	if len(subroutineDecStmt.ParameterList.ParameterList) != 2 {
		t.Fatalf("len(subroutineDecStmt.ParameterList.ParameterList)  is not 2 ,got = %d", len(subroutineDecStmt.ParameterList.ParameterList))
	}
	if len(subroutineDecStmt.Statements.Statements) != 5 {
		t.Fatalf("len(subroutineDecStmt.Statements.Statements)  is not 5 ,got = %d", len(subroutineDecStmt.Statements.Statements))
	}
}
