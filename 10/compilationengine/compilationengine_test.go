package compilationengine

import (
	"jack_compiler/ast"
	"jack_compiler/token"
	"jack_compiler/tokenizer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	testCases := []struct {
		input               string
		expectedName        string
		expectedValueString string
	}{
		{`let y=10;`, "y", "10"},
		{`let hoge =111;`, "hoge", "111"},
		{`let x = 5+5;`, "x", "5+5"},
		{`let x = (5+5)+4;`, "x", "(5+5)+4"},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		letStmt := ce.parseLetStatement()
		if letStmt == nil {
			t.Fatalf("parseLetStatement() returned nil")
		}
		if letStmt.TokenLiteral() != "let" {
			t.Errorf("s.TokenLiteral not 'let'. got %q", letStmt.TokenLiteral())
		}
		if letStmt.Name.Literal != tt.expectedName {
			t.Errorf("letStmt.Name.TokenLiteral() not '%s'.got '%s'", tt.expectedName, letStmt.Name.Literal)
		}
		if letStmt.Value.String() != tt.expectedValueString {
			t.Errorf("letStmt.Value.String() not '%s'.got '%s'", tt.expectedValueString, letStmt.Value.String())
		}
	}
}

func TestLetArrayElemetStatements(t *testing.T) {
	testCases := []struct {
		input               string
		expectedName        string
		expectedIdxString   string
		expectedValueString string
	}{
		{`let y[1]=10;`, "y", "1", "10"},
		{`let hoge[a] =111;`, "hoge", "a", "111"},
		{`let x[a+1] = 5+5;`, "x", "a+1", "5+5"},
		{`let x[(5+5)+4] = (5+5)+4;`, "x", "(5+5)+4", "(5+5)+4"},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		letStmt := ce.parseLetStatement()
		if letStmt == nil {
			t.Fatalf("parseLetStatement() returned nil")
		}
		if letStmt.TokenLiteral() != "let" {
			t.Errorf("s.TokenLiteral not 'let'. got %q", letStmt.TokenLiteral())
		}
		if letStmt.Name.Literal != tt.expectedName {
			t.Errorf("letStmt.Name.TokenLiteral() not '%s'.got '%s'", tt.expectedName, letStmt.Name.Literal)
		}
		if letStmt.Value.String() != tt.expectedValueString {
			t.Errorf("letStmt.Value.String() not '%s'.got '%s'", tt.expectedValueString, letStmt.Value.String())
		}
		if letStmt.Idx.String() != tt.expectedIdxString {
			t.Errorf("letStmt.Idx.String() not '%s'.got '%s'", tt.expectedIdxString, letStmt.Idx.String())
		}
	}
}

func TestReturnStatements(t *testing.T) {
	testCases := []struct {
		input               string
		expectedValueString string
	}{
		{`return x;`, "x"},
		{`return 1;`, "1"},
		{`return 1+1;`, "1+1"},
		{`return;`, ""},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		returnStmt := ce.parseReturnStatement()
		if returnStmt == nil {
			t.Fatalf("parseReturnStatement() returned nil")
		}
		if returnStmt.TokenLiteral() != string(token.RETURN) {
			t.Errorf("s.TokenLiteral not %s. got %q", token.RETURN, returnStmt.TokenLiteral())
		}
		if tt.expectedValueString != "" { // 戻り値が存在する場合
			if returnStmt.Value.String() != tt.expectedValueString {
				t.Errorf("returnStmt.Value.String() not '%s'.got '%s'", tt.expectedValueString, returnStmt.Value.String())
			}
		} else {
			if returnStmt.Value != nil {
				t.Errorf("returnStmt.Value should be nil. got %s ", returnStmt.Value.String())
			}
		}

	}
}

func TestParseDoStatements(t *testing.T) {
	input := `
	do ClassName.VarName(a,b,c,d,e);
`
	jt := tokenizer.New(input)
	ce := New(jt)
	program := ce.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}
	testCases := []struct {
		expectedClassName string
		expectedVarName   string
	}{
		{"ClassName", "VarName"},
	}
	for i, tt := range testCases {
		stmt := program.Statements[i]
		if !testDoStatement(t, stmt, tt.expectedClassName, tt.expectedVarName) {
			return
		}
	}
}

func testDoStatement(t *testing.T, s ast.Statement, className string, varName string) bool {
	if s.TokenLiteral() != "do" {
		t.Errorf("s.TokenLiteral not 'do'. got %q", s.TokenLiteral())
		return false
	}
	doStmt, ok := s.(*ast.DoStatement)
	if !ok {
		t.Errorf("s not *ast.DoStatement. got %T", s)
		return false
	}
	if doStmt.ClassName.Literal != className {
		t.Errorf("doStmt.VarName.Literal not %s. got %s", className, doStmt.ClassName.Literal)
		return false
	}
	if doStmt.VarName.Literal != varName {
		t.Errorf("doStmt.VarName.Literal not %s. got %s", varName, doStmt.VarName.Literal)
		return false
	}
	return true
}

func TestVarDecStatements(t *testing.T) {
	input := `
	var int a,b,c;
	var int length;
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
	input := `33+33`
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
	
	 }else{
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

func TestParseSubroutineCallTermExpression(t *testing.T) {
	input := `hoge.fuga(a,b,c,d,e,f)`
	jt := tokenizer.New(input)
	ce := New(jt)
	expression := ce.parseExpression()
	singleExpression, ok := expression.(*ast.SingleExpression)
	if !ok {
		t.Fatalf("expression is not ast.SingleExpression,got = %T", expression)
	}
	subroutineCallTerm, ok := singleExpression.Value.(*ast.SubroutineCallTerm)
	if !ok {
		t.Fatalf("subroutineCallTerm is not ast.SubroutineCallTerm,got = %T", subroutineCallTerm)
	}
	if subroutineCallTerm.ClassName.Literal != "hoge" {
		t.Fatalf("subroutineCallTerm.ClassName.Literal is not hoge,got = %s", subroutineCallTerm.ClassName.Literal)
	}
	if subroutineCallTerm.VarName.Literal != "fuga" {
		t.Fatalf("subroutineCallTerm.VarName.Literal is not hoge,got = %s", subroutineCallTerm.VarName.Literal)
	}
	if len(subroutineCallTerm.ExpressionListStmt.ExpressionList) != 6 {
		t.Fatalf("len(subroutineCallTerm.ExpressionListStmt.ExpressionList) is not 6,got = %d", len(subroutineCallTerm.ExpressionListStmt.ExpressionList))
	}
}

func TestParseArrayElementExpression(t *testing.T) {
	input := `hoge[a]`
	jt := tokenizer.New(input)
	ce := New(jt)
	expression := ce.parseExpression()
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
	input := `-124`
	jt := tokenizer.New(input)
	ce := New(jt)
	expression := ce.parseExpression()
	singleExpression, ok := expression.(*ast.SingleExpression)
	if !ok {
		t.Fatalf("expression is not ast.SingleExpression,got = %T", expression)
	}
	prefixTerm, ok := singleExpression.Value.(*ast.PrefixTerm)
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
	expression := ce.parseExpression()
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
	testCases := []struct {
		input                      string
		expectedClassName          string
		expectedClassVarDecCount   int
		expectedSubroutineDecCount int
	}{
		{`class Hoge{}`, "Hoge", 0, 0},
		{`class Hoge{field int x, y;field int size;}`, "Hoge", 2, 0},
		{`class Hoge{constructor Square new(int Ax, int Ay, int Asize) {} method void dispose() {}}`, "Hoge", 0, 2},
		{`class Hoge{field int x, y;field int size; constructor Square new(int Ax, int Ay, int Asize) {} method void dispose() {}}`, "Hoge", 2, 2},
	}

	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		stmt := ce.parseClassStatement()
		if stmt.Name.Literal != tt.expectedClassName {
			t.Fatalf("ClassName is not  %s,got = %s", tt.expectedClassName, stmt.Name.Literal)
		}
		if token.KeyWord(stmt.Token.Literal) != token.CLASS {
			t.Fatalf("token.KeyWord(stmt.Token.Literal)  is not token.CLASS ,got = %s", token.KeyWord(stmt.Token.Literal))
		}
		if len(stmt.ClassVarDecList) != tt.expectedClassVarDecCount {
			t.Fatalf("len(stmt.ClassVarDecList) is not %d ,got = %d", tt.expectedClassVarDecCount, len(stmt.ClassVarDecList))
		}
		if len(stmt.SubroutineDecList) != tt.expectedSubroutineDecCount {
			t.Fatalf("len(stmt.SubroutineDecList) is not %d ,got = %d", tt.expectedSubroutineDecCount, len(stmt.SubroutineDecList))
		}
	}
}

func TestParseSubroutineBodyStatement(t *testing.T) {
	input := `{
		var int a,b,c;
		var int length;
		var char casdfasdf;
		var boolean a1,b2,cx;
		let x=5;
		let y=10;
		let hoge=111;
		let foo=838383;
		let bar="hogehoge";
	}`
	jt := tokenizer.New(input)
	ce := New(jt)
	stmt := ce.parseSubroutineBodyStatement()
	if len(stmt.VarDecList) != 4 {
		t.Fatalf("len(stmt.VarDecList)  is not 4 ,got = %d", len(stmt.VarDecList))
	}
	if len(stmt.Statements.Statements) != 5 {
		t.Fatalf("len(stmt.Statements.Statements)  is not 5 ,got = %d", len(stmt.Statements.Statements))
	}
}

func TestParseSubroutineDecStatement(t *testing.T) {
	input := `method void fuga (int hoge,boolean fuga){
		var int a,b,c;
		var int length;
		var char casdfasdf;
		var boolean a1,b2,cx;
		let x=5;
		let y=10;
		let hoge=111;
		let foo=838383;
		let bar="hogehoge";
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
	if len(subroutineDecStmt.SubroutineBody.VarDecList) != 4 {
		t.Fatalf("len(subroutineDecStmt.Statements.Statements)  is not 4 ,got = %d", len(subroutineDecStmt.SubroutineBody.VarDecList))
	}
	if len(subroutineDecStmt.SubroutineBody.Statements.Statements) != 5 {
		t.Fatalf("len(subroutineDecStmt.Statements.Statements)  is not 5 ,got = %d", len(subroutineDecStmt.SubroutineBody.Statements.Statements))
	}
}
