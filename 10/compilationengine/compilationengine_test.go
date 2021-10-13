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
	testCases := []struct {
		input                  string
		expectedClassName      string
		expectedSubroutineName string
	}{
		{"do HogeClass.HogeFunc();", "HogeClass", "HogeFunc"},
		{"do HogeFunc();", "", "HogeFunc"},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		doStmt := ce.parseDoStatement()
		if token.KeyWord(doStmt.TokenLiteral()) != token.DO {
			t.Errorf("s.TokenLiteral not %s. got %q", token.DO, doStmt.TokenLiteral())
		}
		if doStmt.ClassName.Literal != tt.expectedClassName {
			t.Errorf("doStmt.ClassName.Literal not %s. got %s", tt.expectedClassName, doStmt.ClassName.Literal)
		}
		if doStmt.SubroutineName.Literal != tt.expectedSubroutineName {
			t.Errorf("doStmt.SubroutineName.Literal not %s. got %s", tt.expectedSubroutineName, doStmt.SubroutineName.Literal)
		}
	}
}

func TestVarDecStatements(t *testing.T) {
	testCases := []struct {
		input               string
		expectedValueType   string
		expectedIdentifiers []string
	}{
		{"var int a;", "int", []string{"a"}},
		{"var char a;", "char", []string{"a"}},
		{"var boolean a;", "boolean", []string{"a"}},
		{"var ClassName a;", "ClassName", []string{"a"}},
		{"var int a,b,c;", "int", []string{"a", "b", "c"}},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		varDecStmt := ce.parseVarDecStatement()
		if varDecStmt.TokenLiteral() != string(token.VAR) {
			t.Errorf("varDecStmt.TokenLiteral not %s. got %q", token.VAR, varDecStmt.TokenLiteral())
		}
		if varDecStmt.ValueType.Literal != tt.expectedValueType {
			t.Errorf("varDecStmt.ValueType.Literal no %s . got %s", tt.expectedValueType, varDecStmt.ValueType.Literal)
		}
		for i, ident := range varDecStmt.Identifiers {
			if tt.expectedIdentifiers[i] != ident.Literal {
				t.Errorf("identifiers no %s . got %s", tt.expectedIdentifiers[i], ident.Literal)
			}
		}
	}
}

func TestClassVarDecStatements(t *testing.T) {
	testCases := []struct {
		input                   string
		expectedClassVarDecType token.KeyWord
		expectedValueType       string
		expectedIdentifiers     []string
	}{
		{"static int a;", token.STATIC, string(token.INT), []string{"a"}},
		{"field int a;", token.FIELD, string(token.INT), []string{"a"}},
		{"static boolean a;", token.STATIC, string(token.BOOLEAN), []string{"a"}},
		{"static char a;", token.STATIC, string(token.CHAR), []string{"a"}},
		{"static ClassName a;", token.STATIC, "ClassName", []string{"a"}},
		{"static int a,b,c;", token.STATIC, string(token.INT), []string{"a", "b", "c"}},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		classVarDecStmt := ce.parseClassVarDecStatement()
		if classVarDecStmt == nil {
			t.Fatalf("parseClassVarDecStatement() returned nil")
		}
		if tt.expectedClassVarDecType != token.KeyWord(classVarDecStmt.Token.Literal) {
			t.Fatalf("ClassVarDecType should be %s. got %s", tt.expectedClassVarDecType, classVarDecStmt.Token.Literal)
		}
		if tt.expectedValueType != classVarDecStmt.ValueType.Literal {
			t.Fatalf("valueType should be %s . got %s", tt.expectedValueType, classVarDecStmt.ValueType.Literal)
		}
		for i, ident := range classVarDecStmt.Identifiers {
			if tt.expectedIdentifiers[i] != ident.Literal {
				t.Fatalf("identifiers should be %s . got %s", tt.expectedIdentifiers[i], ident.Literal)
			}
		}
	}
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

func TestParseIdentifierTerm(t *testing.T) {
	input := `hoge`
	jt := tokenizer.New(input)
	ce := New(jt)
	identifierTerm := ce.parseIdentifierTerm()
	if identifierTerm.String() != "hoge" {
		t.Fatalf("identifierTerm.Value is not hoge,got = %s", identifierTerm.Value)
	}
}

func TestParseStringConstTerm(t *testing.T) {
	input := `"hoge hoge hoge"`
	jt := tokenizer.New(input)
	ce := New(jt)
	stringConstTerm := ce.parseStringConstTerm()
	if stringConstTerm.Value != "hoge hoge hoge" {
		t.Fatalf("stringConstTerm.Value is not %s,got = %s", "hoge hoge hoge", stringConstTerm.Value)
	}
}

func TestParseIntegerConstTerm(t *testing.T) {
	input := `1234`
	jt := tokenizer.New(input)
	ce := New(jt)
	intergerConstTerm := ce.parseIntegerConstTerm()
	if intergerConstTerm.Value != 1234 {
		t.Fatalf("intergerConstTerm.Value is not %d,got = %d", 1234, intergerConstTerm.Value)
	}
}

func TestParseKeyWordConstTerm(t *testing.T) {
	testCases := []struct {
		input           string
		expectedKeyWord token.KeyWord
	}{
		{"true", token.TRUE},
		{"false", token.FALSE},
		{"null", token.NULL},
		{"this", token.THIS},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		keyWordConstTerm := ce.parseKeyWordConstTerm()
		if keyWordConstTerm.KeyWord != tt.expectedKeyWord {
			t.Fatalf("keyWordConstTerm.KeyWord is not %s,got = %s", tt.expectedKeyWord, keyWordConstTerm.KeyWord)
		}
	}
}

func TestParseIfStatement(t *testing.T) {
	testCases := []struct {
		input                   string
		expectedConditionString string
		expectedHasAlternative  bool
	}{
		{`if(a){}`, "a", false},
		{`if(a){return;}else{}`, "a", true},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		ifStmt := ce.parseIfStatement()
		if ifStmt == nil {
			t.Fatalf("parseIfStatement() returned nil")
		}
		if ifStmt.Condition.String() != tt.expectedConditionString {
			t.Fatalf("ifStmt.Condition.String() should be %s,got = %s", tt.expectedConditionString, ifStmt.Condition.String())
		}
		hasAlternative := ifStmt.Alternative != nil
		if hasAlternative != tt.expectedHasAlternative {
			t.Fatalf("hasAlternative should be %t,got = %t", tt.expectedHasAlternative, hasAlternative)
		}
	}
}

func TestParseWhileStatement(t *testing.T) {
	testCases := []struct {
		input                   string
		expectedConditionString string
		expectedStatementCount  int
	}{
		{`while(x=1){}`, "x=1", 0},
		{`while(a){}`, "a", 0},
		{`while(a){do Hoge();return 1;}`, "a", 2},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		whileStmt := ce.parseWhileStatement()
		if whileStmt.Condition.String() != tt.expectedConditionString {
			t.Fatalf("whileStmt.Condition.String() should be %s,got = %s", tt.expectedConditionString, whileStmt.Condition.String())
		}
		if len(whileStmt.Statements.Statements) != tt.expectedStatementCount {
			t.Fatalf("len(whileStmt.Statements.Statements)  is not %d,got = %d", tt.expectedStatementCount, len(whileStmt.Statements.Statements))
		}
	}
}

func TestParseExpressionListStatement(t *testing.T) {
	testCases := []struct {
		input                   string
		expectedExpressionCount int
	}{
		{"()", 0},
		{"(a,b,c,d,e,f)", 6},
		{"(a,1+1,c,2*2)", 4},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		expressionListStmt := ce.parseExpressionListStatement()
		if len(expressionListStmt.ExpressionList) != tt.expectedExpressionCount {
			t.Fatalf("len(program.Statements) should be %d,got = %d", tt.expectedExpressionCount, len(expressionListStmt.ExpressionList))
		}
	}
}

func TestParseSubroutineCallTerm(t *testing.T) {
	testCases := []struct {
		input                  string
		expectedClassName      string
		expectedSubroutineName string
	}{
		{`HogeClass.fuga()`, "HogeClass", "fuga"},
		{`fuga()`, "", "fuga"},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		subroutineCallTerm := ce.parseSubroutineCallTerm()
		if subroutineCallTerm.ClassName.Literal != tt.expectedClassName {
			t.Fatalf("subroutineCallTerm.ClassName.Literal is not %s ,got = %s", tt.expectedClassName, subroutineCallTerm.ClassName.Literal)
		}
		if subroutineCallTerm.SubroutineName.Literal != tt.expectedSubroutineName {
			t.Fatalf("subroutineCallTerm.VarName.Literal is not %s ,got = %s", tt.expectedSubroutineName, subroutineCallTerm.SubroutineName.Literal)
		}
	}
}

func TestParseArrayElementTerm(t *testing.T) {
	testCases := []struct {
		input             string
		expectedArrayName string
		expectedIdxString string
	}{
		{"y[1]", "y", "1"},
		{"hoge[a]", "hoge", "a"},
		{"x[a+1]", "x", "a+1"},
		{"x[(5+5)+4]", "x", "(5+5)+4"},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		arrayElemTerm := ce.parseArrayElementTerm()
		if arrayElemTerm.ArrayName.Literal != tt.expectedArrayName {
			t.Fatalf("arrayElemTerm.ArrayName.Literal is not %s,got = %s", tt.expectedArrayName, arrayElemTerm.ArrayName.Literal)
		}
		if arrayElemTerm.Idx.String() != tt.expectedIdxString {
			t.Fatalf("arrayElemTerm.Idx.String() is not %s,got = %s", tt.expectedIdxString, arrayElemTerm.Idx.String())
		}
	}
}
func TestParsePrefixTerm(t *testing.T) {
	testCases := []struct {
		input               string
		expectedPrefix      token.Symbol
		expectedValueString string
	}{
		{"-124", token.MINUS, "124"},
		{"~124", token.BANG, "124"},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		prefixTerm := ce.parsePrefixTerm()
		if prefixTerm.Value.String() != tt.expectedValueString {
			t.Fatalf("prefixTerm.Value.String() is not %s ,got = %s", tt.expectedValueString, prefixTerm.Value.String())
		}
		if prefixTerm.Prefix != tt.expectedPrefix {
			t.Fatalf("prefixTerm.Prefix is not %s ,got = %s", tt.expectedPrefix, prefixTerm.Prefix)
		}
	}
}

func TestParseBracketTerm(t *testing.T) {
	testCases := []struct {
		input               string
		expectedValueString string
	}{
		{"(4)", "4"},
		{"(-1)", "-1"},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		bracketTerm := ce.parseBracketTerm()
		if bracketTerm.Value.String() != tt.expectedValueString {
			t.Fatalf("value.Value is not %s,got %s", tt.expectedValueString, bracketTerm.Value.String())
		}
	}
}

func TestParseParameterStatement(t *testing.T) {

	testCases := []struct {
		input              string
		expectedTypeString string
		expectedName       string
	}{
		{"int hoge", string(token.INT), "hoge"},
		{"char fuga", string(token.CHAR), "fuga"},
		{"boolean pepe", string(token.BOOLEAN), "pepe"},
		{"HogeClass papa", "HogeClass", "papa"},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		parameterStmt := ce.parseParameterStatement()
		if parameterStmt == nil {
			t.Fatalf("parse '%s' returned nil ", tt.input)
		}
		if parameterStmt.ValueType.Literal != tt.expectedTypeString {
			t.Fatalf("stmt.Type is not %s ,got = %s", tt.expectedTypeString, parameterStmt.ValueType.Literal)
		}
		if parameterStmt.Name != tt.expectedName {
			t.Fatalf("stmt.Name is not %s,got = %s", tt.expectedName, parameterStmt.Name)
		}
	}
}

func TestParseParameterListStatement(t *testing.T) {
	type TestCase struct {
		expectedTypeString string
		expectedIdentifier string
	}
	testCases := []struct {
		input              string
		expectedParameters []TestCase
	}{
		{`()`, []TestCase{}},
		{`(int hoge,char fuga,boolean pepe,HogeClass papa)`, []TestCase{
			{string(token.INT), "hoge"},
			{string(token.CHAR), "fuga"},
			{string(token.BOOLEAN), "pepe"},
			{"HogeClass", "papa"},
		},
		},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		parameterListStmt := ce.parseParameterListStatement()
		for idx, testCase := range tt.expectedParameters {
			if parameterListStmt.ParameterList[idx].Name != testCase.expectedIdentifier {
				t.Fatalf("stmt.ParameterList[%d].Name is not %s,got %s", idx, testCase.expectedIdentifier, parameterListStmt.ParameterList[idx].Name)
			}
			if parameterListStmt.ParameterList[idx].ValueType.Literal != testCase.expectedTypeString {
				t.Fatalf("stmt.ParameterList[%d].Type is not %s,got %s", idx, testCase.expectedTypeString, parameterListStmt.ParameterList[idx].ValueType.Literal)
			}
		}
	}
}

func TestParseBlockStatement(t *testing.T) {
	testCases := []struct {
		input                  string
		expectedStatementCount int
	}{
		{`{}`, 0},
		{`{return 1;do Hoge();}`, 2},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		blockStmt := ce.parseBlockStatement()
		if blockStmt == nil {
			t.Fatalf("parseBlockStatement() returned nil")
		}
		if tt.expectedStatementCount != len(blockStmt.Statements) {
			t.Fatalf("StatementCount should be %d. got %d", tt.expectedStatementCount, len(blockStmt.Statements))
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
	testCases := []struct {
		input               string
		expectedVarDecCount int
		expectedStmtCount   int
	}{
		{"{}", 0, 0},
		{`{var int a,b,c;var int length;var char casdfasdf;var boolean a1,b2,cx;}`, 4, 0},
		{`{var int a,b,c;var int length;var char casdfasdf;var boolean a1,b2,cx;do draw();return x; }`, 4, 2},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		subroutineBodyStmt := ce.parseSubroutineBodyStatement()
		if len(subroutineBodyStmt.VarDecList) != tt.expectedVarDecCount {
			t.Fatalf("len(subroutineBodyStmt.VarDecList) should be %d ,got = %d", len(subroutineBodyStmt.VarDecList), tt.expectedVarDecCount)
		}
		if len(subroutineBodyStmt.Statements) != tt.expectedStmtCount {
			t.Fatalf("len(subroutineBodyStmt.Statements) should be %d ,got = %d", len(subroutineBodyStmt.Statements), tt.expectedStmtCount)
		}
	}
}

func TestParseSubroutineDecStatement(t *testing.T) {
	testCases := []struct {
		input                  string
		expectedSubroutineType token.KeyWord
		expectedReturnType     string
		expectedName           string
	}{
		{"method void hoge (){}", token.METHOD, "void", "hoge"},
		{"constructor void hoge (){}", token.CONSTRUCTOR, "void", "hoge"},
		{"function void hoge (){}", token.FUNCTION, "void", "hoge"},
		{"method int hoge (){}", token.METHOD, string(token.INT), "hoge"},
		{"method char hoge (){}", token.METHOD, string(token.CHAR), "hoge"},
		{"method boolean hoge (){}", token.METHOD, string(token.BOOLEAN), "hoge"},
		{"method ClassName hoge (){}", token.METHOD, "ClassName", "hoge"},
	}
	for _, tt := range testCases {
		jt := tokenizer.New(tt.input)
		ce := New(jt)
		subroutineDecStmt := ce.parseSubroutineDecStatement()
		if token.KeyWord(subroutineDecStmt.Token.Literal) != tt.expectedSubroutineType {
			t.Fatalf("subroutineDecStmt.Token.Literal is not %s ,got = %s", subroutineDecStmt.Token.Literal, tt.expectedSubroutineType)
		}
		if subroutineDecStmt.ReturnType.Literal != tt.expectedReturnType {
			t.Fatalf("subroutineDecStmt.ReturnType.Literal  is not %s ,got = %s", subroutineDecStmt.ReturnType.Literal, tt.expectedReturnType)
		}
		if subroutineDecStmt.Name.Literal != tt.expectedName {
			t.Fatalf("subroutineDecStmt.Name.Literal should be %s ,got = %s", tt.expectedName, subroutineDecStmt.Name.Literal)
		}
	}

}
