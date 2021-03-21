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
	return leftExp
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
