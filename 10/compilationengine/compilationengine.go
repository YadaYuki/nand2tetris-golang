package compilationengine

import (
	"fmt"
	"jack_compiler/ast"
	"jack_compiler/token"
	"jack_compiler/tokenizer"
	"strconv"
)

type (
	singleParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// CompilationEngine is struct
type CompilationEngine struct {
	jt             *tokenizer.JackTokenizer
	errors         []string
	curToken       token.Token
	nextToken      token.Token
	singleParseFns map[token.TokenType]singleParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

const (
	_ int = iota

	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

// var precedences = map[token.Symbol]int{
// 	token.EQ:       EQUALS,
// 	token.NOT_EQ:   EQUALS,
// 	token.LT:       LESSGREATER,
// 	token.GT:       LESSGREATER,
// 	token.PLUS:     SUM,
// 	token.MINUS:    SUM,
// 	token.SLASH:    PRODUCT,
// 	token.ASTERISK: PRODUCT,
// 	token.LPAREN:   CALL,
// }

// func (ce *CompilationEngine) nextPrecedence() int {
// 	if p,ok := precedences[ce.nextToken.Type];ok{
// 		return p
// 	}
// 	return
// }

// func (ce *CompilationEngine) curPrecedence() int {
// 	if p,ok := precedences[ce.curToken.Type];ok{
// 		return p
// 	}
// 	return
// }

// func (ce *CompilationEngine) registerSingle(tokenType token.TokenType, fn singleParseFn) {
// 	ce.singleParseFns[tokenType] = fn
// }

// func (ce *CompilationEngine) registerInfix(tokenType token.TokenType, fn infixParseFn) {
// 	ce.infixParseFns[tokenType] = fn
// }

// New is initializer of compilation engine
func New(jt *tokenizer.JackTokenizer) *CompilationEngine {
	ce := &CompilationEngine{jt: jt}
	ce.advanceToken()
	ce.advanceToken()
	return ce
}

// ParseProgram is parser for all program
func (ce *CompilationEngine) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for ce.curToken.Type != token.EOF {
		stmt := ce.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		ce.advanceToken()
	}
	return program
}

func (ce *CompilationEngine) advanceToken() {
	ce.curToken = ce.nextToken
	ce.nextToken, _ = ce.jt.Advance()
}

func (ce *CompilationEngine) parseStatement() ast.Statement {
	if ce.curToken.Type != token.KEYWORD {
		panic(fmt.Sprintf("Initial Token Type should be KEYWORD. got %s ", ce.curToken.Type))
	}
	return ce.parseKeyWord()
}

func (ce *CompilationEngine) parseKeyWord() ast.Statement {
	keyWord, _ := tokenizer.KeyWord(ce.curToken)
	switch keyWord {
	case token.LET:
		return ce.parseLetStatement()
	case token.RETURN:
		return ce.parseReturnStatement()
	case token.DO:
		return ce.parseDoStatement()
	case token.VAR:
		return ce.parseVarDecStatement()
	case token.STATIC:
		return ce.parseClassVarDecStatement()
	case token.FIELD:
		return ce.parseClassVarDecStatement()
	case token.IF:
		return ce.parseIfStatement()
	case token.WHILE:
		return ce.parseWhileStatement()
	case token.CLASS:
		return ce.parseClassStatement()
	case token.METHOD:
		return ce.parseSubroutineDecStatement()
	case token.CONSTRUCTOR:
		return ce.parseSubroutineDecStatement()
	case token.FUNCTION:
		return ce.parseSubroutineDecStatement()
	default:
		return nil
	}
}

func (ce *CompilationEngine) parseClassStatement() *ast.ClassStatement {
	stmt := &ast.ClassStatement{Token: ce.curToken}
	if !ce.expectNext(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = ce.curToken
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.LBRACE {
		return nil
	}
	stmt.Statements = ce.parseBlockStatement()
	if token.Symbol(ce.curToken.Literal) != token.RBRACE {
		return nil
	}
	ce.advanceToken()
	return stmt
}

func (ce *CompilationEngine) parseSubroutineDecStatement() *ast.SubroutineDecStatement {
	stmt := &ast.SubroutineDecStatement{Token: ce.curToken}
	if !ce.nextTokenIs(token.IDENTIFIER) && !ce.nextTokenIs(token.KEYWORD) {
		return nil
	}
	ce.advanceToken()
	stmt.ReturnType = ce.curToken
	if !ce.expectNext(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = ce.curToken
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.LPAREN {
		return nil
	}
	stmt.ParameterList = ce.parseParameterListStatement()
	if token.Symbol(ce.curToken.Literal) != token.RPAREN {
		return nil
	}
	ce.advanceToken()
	stmt.Statements = ce.parseBlockStatement()
	return stmt
}

// TODO:Add Error Handling
func (ce *CompilationEngine) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: ce.curToken}
	if !ce.expectNext(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = ce.curToken
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.ASSIGN {
		return nil
	}
	stmt.Symbol = ce.curToken
	ce.advanceToken()
	// TODO: add parse expression
	stmt.Value = ce.parseExpression()
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.SEMICOLON {
		return nil
	}
	return stmt
}

func (ce *CompilationEngine) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: ce.curToken}
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) == token.SEMICOLON {
		return stmt
	}
	// stmt.ReturnValue = ce.parseExpression()
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.SEMICOLON {
		return nil
	}
	return stmt
}

func (ce *CompilationEngine) parseDoStatement() *ast.DoStatement {
	stmt := &ast.DoStatement{Token: ce.curToken}
	ce.advanceToken()
	stmt.SubroutineCall = ce.curToken
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.SEMICOLON {
		return nil
	}
	return stmt
}

func (ce *CompilationEngine) parseVarDecStatement() *ast.VarDecStatement {
	stmt := &ast.VarDecStatement{Token: ce.curToken, Identifiers: []token.Token{}}
	if ce.expectNext(token.KEYWORD) {
		if token.KeyWord(ce.curToken.Literal) != token.INT && token.KeyWord(ce.curToken.Literal) != token.BOOLEAN && token.KeyWord(ce.curToken.Literal) != token.CHAR {
			return nil
		}
	}
	stmt.ValueType = ce.curToken
	for token.Symbol(ce.curToken.Literal) != token.SEMICOLON {
		ce.advanceToken()
		identifier := ce.curToken
		stmt.Identifiers = append(stmt.Identifiers, identifier)
		ce.advanceToken() //
	}
	return stmt
}

func (ce *CompilationEngine) parseClassVarDecStatement() *ast.ClassVarDecStatement {
	stmt := &ast.ClassVarDecStatement{Token: ce.curToken, Identifiers: []token.Token{}}
	if ce.expectNext(token.KEYWORD) {
		if token.KeyWord(ce.curToken.Literal) != token.INT && token.KeyWord(ce.curToken.Literal) != token.BOOLEAN && token.KeyWord(ce.curToken.Literal) != token.CHAR {
			return nil
		}
	}
	stmt.ValueType = ce.curToken
	for token.Symbol(ce.curToken.Literal) != token.SEMICOLON {
		ce.advanceToken()
		identifier := ce.curToken
		stmt.Identifiers = append(stmt.Identifiers, identifier)
		ce.advanceToken() //
	}
	return stmt
}

func (ce *CompilationEngine) parseIfStatement() *ast.IfStatement {
	stmt := &ast.IfStatement{Token: ce.curToken}
	if ce.expectNext(token.SYMBOL) {
		if token.Symbol(ce.curToken.Literal) != token.LPAREN {
			return nil
		}
	}
	// TODO:Add parseExpression
	for token.Symbol(ce.curToken.Literal) != token.RPAREN {
		ce.advanceToken()
	}

	ce.advanceToken()
	stmt.Consequence = ce.parseBlockStatement()
	ce.advanceToken()
	if token.KeyWord(ce.curToken.Literal) == token.ELSE {
		ce.advanceToken()
		if ce.expectNext(token.SYMBOL) {
			return nil
		}
		stmt.Alternative = ce.parseBlockStatement()
	}
	return stmt
}

func (ce *CompilationEngine) parseWhileStatement() *ast.WhileStatement {
	stmt := &ast.WhileStatement{Token: ce.curToken}
	if ce.expectNext(token.SYMBOL) {
		if token.Symbol(ce.curToken.Literal) != token.LPAREN {
			return nil
		}
	}
	// TODO:Add parseExpression
	for token.Symbol(ce.curToken.Literal) != token.RPAREN {
		ce.advanceToken()
	}

	ce.advanceToken()
	stmt.Statements = ce.parseBlockStatement()
	ce.advanceToken()
	return stmt
}

func (ce *CompilationEngine) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: ce.curToken}
	ce.advanceToken()
	block.Statements = []ast.Statement{}
	for token.Symbol(ce.curToken.Literal) != token.RBRACE && !ce.curTokenIs(token.EOF) {
		stmt := ce.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		ce.advanceToken()
	}
	return block
}

func (ce *CompilationEngine) parseExpressionListStatement() *ast.ExpressionListStatement {
	expressionListStmt := &ast.ExpressionListStatement{Token: ce.curToken}
	ce.advanceToken()
	expressionListStmt.ExpressionList = []ast.Expression{}
	for token.Symbol(ce.curToken.Literal) != token.RPAREN && !ce.curTokenIs(token.EOF) {
		expression := ce.parseExpression()
		if expression != nil {
			expressionListStmt.ExpressionList = append(expressionListStmt.ExpressionList, expression)
		}
		ce.advanceToken()
		ce.advanceToken()
	}
	return expressionListStmt
}

func (ce *CompilationEngine) parseParameterListStatement() *ast.ParameterListStatement {
	parameterListStmt := &ast.ParameterListStatement{Token: ce.curToken}
	ce.advanceToken()
	parameterListStmt.ParameterList = []ast.ParameterStatement{}
	for token.Symbol(ce.curToken.Literal) != token.RPAREN {
		parameterStmt := ce.parseParameterStatement()
		if parameterStmt == nil {
			return nil
		}
		parameterListStmt.ParameterList = append(parameterListStmt.ParameterList, *parameterStmt)
		ce.advanceToken()
		if token.Symbol(ce.curToken.Literal) == token.RPAREN {
			break
		}
		ce.advanceToken()
	}
	return parameterListStmt
}

func (ce *CompilationEngine) parseParameterStatement() *ast.ParameterStatement {
	parameterStmt := &ast.ParameterStatement{Token: ce.curToken}
	if ce.curToken.Type != token.KEYWORD {
		return nil
	}
	parameterStmt.Type = token.KeyWord(ce.curToken.Literal)
	ce.advanceToken()
	if ce.curToken.Type != token.IDENTIFIER {
		return nil
	}
	parameterStmt.Name = ce.curToken.Literal
	return parameterStmt
}

func (ce *CompilationEngine) parseExpression() ast.Expression {
	// leftExp := single()
	// TODO:Fix to SEMICOLON
	// for !p.nextTokenIs(token.SYMBOL) && precedence < ce.nextPrecedence() {
	// 	infix := ce.infixParseFns(ce.nextToken.Type)
	// 	if infix == nil {
	// 		return leftExp
	// 	}
	// 	ce.nextToken()
	// 	leftExp = infix(leftExp)
	// }
	return &ast.SingleExpression{}
}

func (ce *CompilationEngine) parseTerm() ast.Term {
	// TODO:Fix to SEMICOLON
	// for !p.nextTokenIs(token.SYMBOL) && precedence < ce.nextPrecedence() {
	// 	infix := ce.infixParseFns(ce.nextToken.Type)
	// 	if infix == nil {
	// 		return leftExp
	// 	}
	// 	ce.advanceToken()
	// 	leftExp = infix(leftExp)
	// }
	return &ast.BracketTerm{}
}

func (ce *CompilationEngine) parseIntegerConstExpression() ast.Expression {
	expression := &ast.SingleExpression{Token: ce.curToken}
	value, err := strconv.ParseInt(ce.curToken.Literal, 0, 64)
	if err != nil {
		panic(fmt.Sprintf("could not parse %q as integer", ce.curToken.Literal))
	}
	expression.Value = &ast.IntergerConstTerm{Token: ce.curToken, Value: value}
	return expression
}

func (ce *CompilationEngine) parseIdentifierExpression() ast.Expression {
	expression := &ast.SingleExpression{Token: ce.curToken, Value: &ast.IdentifierTerm{Token: ce.curToken, Value: ce.curToken.Literal}}
	return expression
}

func (ce *CompilationEngine) parseStringConstExpression() ast.Expression {
	expression := &ast.SingleExpression{Token: ce.curToken, Value: &ast.StringConstTerm{Token: ce.curToken, Value: ce.curToken.Literal}}
	return expression
}

func (ce *CompilationEngine) parseSubroutineCallExpression() ast.Expression {
	expression := &ast.SingleExpression{Token: ce.curToken}
	ce.advanceToken()
	expressionListStmt := ce.parseExpressionListStatement()
	expression.Value = &ast.SubroutineCallTerm{Token: ce.curToken, FunctionName: ce.curToken.Literal, ExpressionListStmt: *expressionListStmt}
	return expression
}

func (ce *CompilationEngine) parseArrayElementExpression() ast.Expression {
	expression := &ast.SingleExpression{Token: ce.curToken}
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.LBRACKET {
		return nil
	}
	ce.advanceToken()
	idx := ce.parseExpression()
	expression.Value = &ast.ArrayElementTerm{Token: ce.curToken, ArrayName: ce.curToken.Literal, Idx: idx}
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.RBRACKET {
		return nil
	}
	return expression
}

func (ce *CompilationEngine) parsePrefixExpression() ast.Term {
	prefixTerm := &ast.PrefixTerm{Token: ce.curToken, Prefix: token.Symbol(ce.curToken.Literal)}
	ce.advanceToken()
	prefixTerm.Value = ce.parseTerm()
	return prefixTerm
}

func (ce *CompilationEngine) parseBracketExpression() ast.Expression {
	expression := &ast.SingleExpression{Token: ce.curToken}
	bracketTerm := &ast.BracketTerm{Token: ce.curToken}
	ce.advanceToken()
	exp := ce.parseExpression()
	bracketTerm.Value = exp
	expression.Value = bracketTerm
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.RPAREN {
		return nil
	}
	return expression
}

func (ce *CompilationEngine) curTokenIs(t token.TokenType) bool {
	return ce.curToken.Type == t
}

func (ce *CompilationEngine) nextTokenIs(t token.TokenType) bool {
	return ce.nextToken.Type == t
}

func (ce *CompilationEngine) expectNext(t token.TokenType) bool {
	if ce.nextTokenIs(t) {
		ce.advanceToken()
		return true
	}
	return false
}
