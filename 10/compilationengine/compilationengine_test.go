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
	program := ce.parseIfStatement()
	t.Log(program)
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
