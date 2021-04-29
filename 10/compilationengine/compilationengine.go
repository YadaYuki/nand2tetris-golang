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
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[token.Symbol]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}

// func (ce *CompilationEngine) nextPrecedence() int {
// 	if p,ok := precedences[ce.nextToken.Type];ok{
// 		return p
// 	}
// 	return LOWEST
// }

// func (ce *CompilationEngine) curPrecedence() int {
// 	if p,ok := precedences[ce.curToken.Type];ok{
// 		return p
// 	}
// 	return LOWEST
// }

func (ce *CompilationEngine) registerSingle(tokenType token.TokenType, fn singleParseFn) {
	ce.singleParseFns[tokenType] = fn
}

func (ce *CompilationEngine) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	ce.infixParseFns[tokenType] = fn
}

// New is initializer of compilation engine
func New(jt *tokenizer.JackTokenizer) *CompilationEngine {
	ce := &CompilationEngine{jt: jt}
	ce.singleParseFns = make(map[token.TokenType]singleParseFn)
	ce.registerSingle(token.STARTINGCONST, ce.parseStringConstExpression)
	ce.registerSingle(token.INTCONST, ce.parseIntegerConstExpression)
	ce.registerSingle(token.IDENTIFIER, ce.parseIdentifierExpression)
	// ce.registerSingle(token.BANG,ce.parseSingleExpression)
	// ce.registerSingle(token.LPAREN,ce.parseSingleGroupedExpression)
	// ce.registerSingle(token.IF,ce.parseIfExpression)
	// ce.registerSingle(token.FUNCTION,ce.parseFunctionLiteral)
	// ce.infixParseFns = make(map[token.TokenType]infixParseFn)
	// ce.registerInfix(token.PLUS,ce.parseInfixExpression)
	// ce.registerInfix(token.MINUS,ce.parseInfixExpression)
	// ce.registerInfix(token.ASTERISK,ce.parseInfixExpression)
	// ce.registerInfix(token.SLASH,ce.parseInfixExpression)
	// ce.registerInfix(token.LT,ce.parseInfixExpression)
	// ce.registerInfix(token.GT,ce.parseInfixExpression)
	// ce.registerInfix(token.EQ,ce.parseInfixExpression)
	// ce.registerInfix(token.NOT_EQ,ce.parseInfixExpression)
	// ce.registerInfix(token.LPAREN,ce.parseCallFunctionExpression)
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
	default:
		return nil
	}
}

// TODO:Add Error Handling
func (ce *CompilationEngine) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: ce.curToken}
	if !ce.expectNext(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: ce.curToken, Value: ce.curToken.Literal}
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.ASSIGN {
		return nil
	}
	stmt.Symbol = ce.curToken
	ce.advanceToken()
	// TODO: add parse expression
	// stmt.LetValue = ce.parseExpression(LOWEST)
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
	// stmt.ReturnValue = ce.parseExpression(LOWEST)
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.SEMICOLON {
		return nil
	}
	return stmt
}

func (ce *CompilationEngine) parseDoStatement() *ast.DoStatement {
	stmt := &ast.DoStatement{Token: ce.curToken}
	ce.advanceToken()
	// stmt.SubroutineCall = ce.parseExpression(LOWEST)
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.SEMICOLON {
		return nil
	}
	return stmt
}

func (ce *CompilationEngine) parseVarDecStatement() *ast.VarDecStatement {
	stmt := &ast.VarDecStatement{Token: ce.curToken, Identifiers: []*ast.Identifier{}}
	if ce.expectNext(token.KEYWORD) {
		if token.KeyWord(ce.curToken.Literal) != token.INT && token.KeyWord(ce.curToken.Literal) != token.BOOLEAN && token.KeyWord(ce.curToken.Literal) != token.CHAR {
			return nil
		}
	}
	stmt.ValueType = ce.curToken
	for token.Symbol(ce.curToken.Literal) != token.SEMICOLON {
		ce.advanceToken()
		identifier := &ast.Identifier{Token: ce.curToken, Value: ce.curToken.Literal}
		stmt.Identifiers = append(stmt.Identifiers, identifier)
		ce.advanceToken() //
	}
	return stmt
}

func (ce *CompilationEngine) parseClassVarDecStatement() *ast.ClassVarDecStatement {
	stmt := &ast.ClassVarDecStatement{Token: ce.curToken, Identifiers: []*ast.Identifier{}}
	if ce.expectNext(token.KEYWORD) {
		if token.KeyWord(ce.curToken.Literal) != token.INT && token.KeyWord(ce.curToken.Literal) != token.BOOLEAN && token.KeyWord(ce.curToken.Literal) != token.CHAR {
			return nil
		}
	}
	stmt.ValueType = ce.curToken
	for token.Symbol(ce.curToken.Literal) != token.SEMICOLON {
		ce.advanceToken()
		identifier := &ast.Identifier{Token: ce.curToken, Value: ce.curToken.Literal}
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
		expression := ce.parseExpression(LOWEST)
		if expression != nil {
			expressionListStmt.ExpressionList = append(expressionListStmt.ExpressionList, expression)
		}
		ce.advanceToken()
		ce.advanceToken()
	}
	return expressionListStmt
}

func (ce *CompilationEngine) parseExpression(precedence int) ast.Expression {
	prefix := ce.singleParseFns[ce.curToken.Type]
	if prefix == nil {
		return nil
	}
	// leftExp := single()
	// // TODO:Fix to SEMICOLON
	// for !p.nextTokenIs(token.SYMBOL) && precedence < ce.nextPrecedence() {
	// 	infix := ce.infixParseFns(ce.nextToken.Type)
	// 	if infix == nil {
	// 		return leftExp
	// 	}
	// 	ce.nextToken()
	// 	leftExp = infix(leftExp)
	// }
	// return leftExp
	return prefix()
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
	idx := ce.parseExpression(LOWEST)
	expression.Value = &ast.ArrayElementTerm{Token: ce.curToken, ArrayName: ce.curToken.Literal, Idx: idx}
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.RBRACKET {
		return nil
	}
	return expression
}

func (ce *CompilationEngine) parsePrefixExpression() ast.Expression {
	expression := &ast.SingleExpression{Token: ce.curToken}
	prefixTerm := &ast.PrefixTerm{Token: ce.curToken, Prefix: token.Symbol(ce.curToken.Literal)}
	ce.advanceToken()
	prefixTerm.Value = ce.parseTerm()
	expression.Value = prefixTerm
	return expression
}

func (ce *CompilationEngine) parseBracketExpression() ast.Expression {
	expression := &ast.SingleExpression{Token: ce.curToken}
	bracketTerm := &ast.BracketTerm{Token: ce.curToken}
	ce.advanceToken()
	exp := ce.parseExpression(LOWEST)
	bracketTerm.Value = exp
	expression.Value = bracketTerm
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.RPAREN {
		return nil
	}
	return expression
}

// TODO:Implement Term Dict
func (ce *CompilationEngine) parseTerm() ast.Term {
	return &ast.IntergerConstTerm{Token: token.Token{Type: token.INTCONST, Literal: "4"}, Value: 4}
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
