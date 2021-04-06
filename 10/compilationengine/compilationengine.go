package compilationengine

import (
	"jack/compiler/ast"
	"jack/compiler/token"
	"jack/compiler/tokenizer"
	"strconv"
)

type (
	prefixParseFn func() ast.Expression
  infixParseFn  func(ast.Expression) ast.Expression
)

// CompilationEngine is struct
type CompilationEngine struct {
	jt *tokenizer.JackTokenizer
	errors  []string
	curToken  token.Token
	nextToken token.Token
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns map[token.TokenType]infixParseFn
}

const (
	_ int =iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)


var precedences = map[token.TokenType]int{
	token.EQ: EQUALS,
	token.NOT_EQ: EQUALS,
	token.LT: LESSGREATER,
	token.GT: LESSGREATER,
	token.PLUS:  SUM,
	token.MINUS:  SUM,
	token.SLASH:PRODUCT,
	token.ASTERISK:PRODUCT,
	token.LPAREN : CALL
}

func (ce *CompilationEngine) nextPrecedence() int {
	if p,ok := precedences[ce.nextToken.Type];ok{
		return p
	}
	return LOWEST
}

func (ce *CompilationEngine) curPrecedence() int {
	if p,ok := precedences[ce.curToken.Type];ok{
		return p
	}
	return LOWEST
}

func (ce *CompilationEngine) registerPrefix(tokenType token.TokenType, fn prefixParseFn){
	ce.prefixParseFns[tokenType] = fn
}

func (ce *CompilationEngine) registerInfix(tokenType token.TokenType, fn infixParseFn){
	ce.infixParseFns[tokenType] = fn
}

// New is initializer of compilation engine
func New(jt *tokenizer.JackTokenizer) *CompilationEngine {
	ce := &CompilationEngine{jt: jt}
	ce.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	ce.registerPrefix(token.IDENTIFIER,ce.parseIdentifier)
	ce.registerPrefix(token.INTCONST,ce.parseIntConst)
	ce.registerPrefix(token.MINUS,ce.parsePrefixExpression)
	ce.registerPrefix(token.BANG,ce.parsePrefixExpression)
	ce.registerPrefix(token.LPAREN,ce.parsePrefixGroupedExpression)
	ce.registerPrefix(token.IF,ce.parseIfExpression)
	ce.registerPrefix(token.FUNCTION,ce.parseFunctionLiteral)
	ce.infixParseFns = make(map[token.TokenType]infixParseFn)
	ce.registerInfix(token.PLUS,ce.parseInfixExpression)
	ce.registerInfix(token.MINUS,ce.parseInfixExpression)
	ce.registerInfix(token.ASTERISK,ce.parseInfixExpression)
	ce.registerInfix(token.SLASH,ce.parseInfixExpression)
	ce.registerInfix(token.LT,ce.parseInfixExpression)
	ce.registerInfix(token.GT,ce.parseInfixExpression)
	ce.registerInfix(token.EQ,ce.parseInfixExpression)
	ce.registerInfix(token.NOT_EQ,ce.parseInfixExpression)
	ce.registerInfix(token.LPAREN,ce.parseCallFunctionExpression)
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

	switch ce.curToken.Type {
	// case token.SYMBOL:
	// return nil
	case token.KEYWORD:
		return ce.parseKeyWord()
	// case token.IDENTIFIER:
	// return nil
	// case token.INTCONST:
	// return nil
	// case token.STARTINGCONST:
	// return nil
	default:
		return ce.parseExpressionStatement()
	}
}

func (ce *CompilationEngine) parseKeyWord() ast.Statement {
	keyWord, _ := tokenizer.KeyWord(ce.curToken)
	switch keyWord {
	case token.LET:
		return ce.parseLetStatement()
	case token.RETURN:
		return ce.parseReturnStatement()
	default:
		return nil
	}
}

func (ce *CompilationEngine) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: ce.curToken}
	if !ce.expectNext(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: ce.curToken, Value: ce.curToken.Literal}
	if !ce.expectNext(token.SYMBOL) {
		// TODO:Add ASSIGN
		return nil
	}

	// TODO: Add Expression
	for {
		if ce.curTokenIs(token.SYMBOL) {
			// TODO:Add SEMICOLON
			break
		}
		ce.advanceToken()
	}
	return stmt
}

func (ce *CompilationEngine) parseReturnStatement() *ast.ReturnStatement{
	stmt := &ast.ReturnStatement{Token:ce.curToken}
		// TODO: Add Expression
		for {
			if ce.curTokenIs(token.SYMBOL) {
				// TODO:Add SEMICOLON
				break
			}
			ce.advanceToken()
		}
		return stmt
}

func (ce *CompilationEngine) parseExpressionStatement() *ast.ExpressionStatement{
	stmt := &ast.ExpressionStatement{Token:ce.curToken}
	stmt.Expression = ce.parseExpression(LOWEST)
	if ce.nextTokenIs(token.SYMBOL){
		// TODO:Add SEMICOLON
		ce.advanceToken()
	}
	return stmt
}

func (ce *CompilationEngine) parseExpression(precedence int) ast.Expression{
	prefix := ce.prefixParseFns[ce.curToken.Type]
	if prefix == nil{
		return nil
	}
	leftExp := prefix() 
	// TODO:Fix to SEMICOLON
	for !p.nextTokenIs(token.SYMBOL) && precedence < ce.nextPrecedence() {
		infix := ce.infixParseFns(ce.nextToken.Type)
		if infix == nil{
			return leftExp
		}
		ce.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (ce *CompilationEngine) parsePrefixGroupedExpression() ast.Expression{
	ce.advanceToken()
	exp := ce.parseExpression(LOWEST)
	if ce.expectNext(token.RPAREN){
		return nil
	}
	return exp
}

func (ce *CompilationEngine) parseIdentifier() ast.Expression{
	return &ast.Identifier{Token:ce.curToken,Value:ce.curToken.Literal}
}

func (ce *CompilationEngine) parseIntConst() ast.Expression{
	il := &ast.IntConst{Token:ce.curToken}
	value,err := strconv.ParseInt(ce.curToken.Literal,0,64)
	if err != nil{
		return nil
	}
	il.Value = value
	return il
}

func (ce *CompilationEngine) parsePrefixExpression() ast.Expression{
	expression := &ast.PrefixExpression{
		Token:ce.curToken,
		Operator: ce.curToken.Literal,
	}
	ce.advanceToken()
	expression.Right = ce.parseExpression(PREFIX)
	return expression
}

func (ce *CompilationEngine) parseInfixExpression() ast.Expression{
	expression := &ast.InfixExpression{
		Token: ce.curToken,
		Operator: ce.curToken.Literal,
		Left: left,
	}
	precedence := ce.curPrecedence()
	ce.advanceToken()
	expression.Right = ce.parseExpression()
	return expression
}
 

func (ce *CompilationEngine) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Toke:ce.curToken}
	if !ce.expectNext(token.LPAREN) {
		return nil
	}
	ce.advanceToken()
	expression.Condition = ce.parseExpression(LOWEST)
	if !ce.expectNext(token.RPAREN) {
		return nil
	}

	if !ce.expectNext(token.LBRACE) {
		return nil
	}
	expression.Consequence = ce.parseBlockStatement()
	if ce.curTokenIs(toke.ELSE) {
		ce.advanceToken()
		if !ce.expectNext(token.LBRACE) {
			return nil
		}
		expression.Alternative = ce.parseBlockStatement()
	}
	return expression
}

func (ce *CompilationEngine) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token:ce.curToken}
	block.Statements = []ast.Statement{}
	ce.advanceToken()
	for !ce.curTokenIs(token.RBRACE) && !ce.curTokenIs(token.EOF) {
		stmt := ce.parseStatement()
		if stmt != nil{
			block.Statements = append(block.Statements,stmt)
		}
		ce.advanceToken()
	}
	return block
}

func (ce *CompilationEngine) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: ce.curToken}
	
	if !ce.expectNext(token.LPAREN) {
		return nil
	}
	
	lit.Parameters = ce.parseFunctionParameters()
	
	if !ce.expectNext(token.LBRACE) {
		return nil
	}
	
	lit.Body = ce.parseBlockStatement()

	return lit
}

func (ce *CompilationEngine) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}
	if ce.nextTokenIs(token.RPAREN) {
		ce.advanceToken()
		return identifiers
	}
	ce.advanceToken()
	ident := &ast.Identifier{Token:ce.curToken,Value: ce.curToken.Literal}
	identifiers = append(identifiers,ident)
	for ce.nextTokenIs(token.COMMA) {
		ce.advanceToken()
		ce.advanceToken()
		ident := &ast.Identifier{Token:ce.curToken,Value :ce.curToken.Literal}
		identifiers = append(identifiers,ident)
	}
	if !ce.expectNext(token.RPAREN){
		return nil
	}
	return identifiers
}

func (ce *CompilationEngine) parseCallFunctionExpression() []ast.Expression {
	args := []ast.Expression{}
	if ce.nextTokenIs(token.RPAREN) {
		ce.advanceToken()
		return args
	}
	ce.advanceToken()
	args = append(args,ce.parseExpression(LOWEST))
	for ce.nextTokenIs(token.COMMA) {
		ce.advanceToken()
		ce.advanceToken()
		args = append(args,ce.parseExpression(LOWEST))
	}
	if !expectNext(token.RPAREN) {
		return nil
	}
	return args
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



