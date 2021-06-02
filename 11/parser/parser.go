package parser

import (
	"fmt"
	"jack_compiler/ast"
	"jack_compiler/token"
	"jack_compiler/tokenizer"
	"strconv"
)

// Parser is struct
type Parser struct {
	jt        *tokenizer.JackTokenizer
	curToken  token.Token
	nextToken token.Token
}

// New is initializer of compilation engine
func New(jt *tokenizer.JackTokenizer) *Parser {
	p := &Parser{jt: jt}
	p.advanceToken()
	p.advanceToken()
	return p
}

// ParseProgram is parser for all program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.advanceToken()
	}
	return program
}

func (p *Parser) advanceToken() {
	p.curToken = p.nextToken
	p.nextToken, _ = p.jt.Advance()
}

func (p *Parser) parseStatement() ast.Statement {
	if p.curToken.Type != token.KEYWORD {
		panic(fmt.Sprintf("Initial Token Type should be KEYWORD. got %s ", p.curToken.Type))
	}
	return p.parseKeyWord()
}

func (p *Parser) parseKeyWord() ast.Statement {
	keyWord, _ := tokenizer.KeyWord(p.curToken)
	switch keyWord {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.DO:
		return p.parseDoStatement()
	case token.VAR:
		return p.parseVarDecStatement()
	case token.STATIC:
		return p.parseClassVarDecStatement()
	case token.FIELD:
		return p.parseClassVarDecStatement()
	case token.IF:
		return p.parseIfStatement()
	case token.WHILE:
		return p.parseWhileStatement()
	case token.CLASS:
		return p.parseClassStatement()
	case token.METHOD:
		return p.parseSubroutineDecStatement()
	case token.CONSTRUCTOR:
		return p.parseSubroutineDecStatement()
	case token.FUNCTION:
		return p.parseSubroutineDecStatement()
	default:
		return nil
	}
}

func (p *Parser) parseClassStatement() *ast.ClassStatement {
	stmt := &ast.ClassStatement{Token: p.curToken}
	if !p.expectNext(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = p.curToken
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) != token.LBRACE {
		return nil
	}
	p.advanceToken()
	stmt.ClassVarDecList = []ast.ClassVarDecStatement{}
	for token.KeyWord(p.curToken.Literal) == token.STATIC || token.KeyWord(p.curToken.Literal) == token.FIELD {

		classVarDec := p.parseClassVarDecStatement()
		stmt.ClassVarDecList = append(stmt.ClassVarDecList, *classVarDec)
		p.advanceToken()
	}
	stmt.SubroutineDecList = []ast.SubroutineDecStatement{}
	for token.KeyWord(p.curToken.Literal) == token.CONSTRUCTOR || token.KeyWord(p.curToken.Literal) == token.FUNCTION || token.KeyWord(p.curToken.Literal) == token.METHOD {
		subroutineDec := p.parseSubroutineDecStatement()
		stmt.SubroutineDecList = append(stmt.SubroutineDecList, *subroutineDec)
		p.advanceToken()
	}
	if token.Symbol(p.curToken.Literal) != token.RBRACE {
		return nil
	}
	return stmt
}

func (p *Parser) parseSubroutineDecStatement() *ast.SubroutineDecStatement {
	stmt := &ast.SubroutineDecStatement{Token: p.curToken}
	if !p.nextTokenIs(token.IDENTIFIER) && !p.nextTokenIs(token.KEYWORD) {
		return nil
	}
	p.advanceToken()
	stmt.ReturnType = p.curToken
	if !p.expectNext(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = p.curToken
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) != token.LPAREN {
		return nil
	}
	stmt.ParameterList = p.parseParameterListStatement()
	if token.Symbol(p.curToken.Literal) != token.RPAREN {
		return nil
	}
	p.advanceToken()
	stmt.SubroutineBody = p.parseSubroutineBodyStatement()
	return stmt
}

func (p *Parser) parseSubroutineBodyStatement() *ast.SubroutineBodyStatement {
	stmt := &ast.SubroutineBodyStatement{Token: p.curToken}
	if token.Symbol(p.curToken.Literal) != token.LBRACE {
		return nil
	}
	p.advanceToken()
	stmt.VarDecList = []ast.VarDecStatement{}
	for token.KeyWord(p.curToken.Literal) == token.VAR {
		varDec := p.parseVarDecStatement()
		stmt.VarDecList = append(stmt.VarDecList, *varDec)
		p.advanceToken()
	}
	// NOTE: originally, should call parseBlockStatement.
	// But, the block statement in subroutine body does not start "{".
	// so, implement original parser here.
	stmt.Statements = &ast.BlockStatement{}
	stmt.Statements.Statements = []ast.Statement{}
	for token.Symbol(p.curToken.Literal) != token.RBRACE && !p.curTokenIs(token.EOF) {
		statement := p.parseStatement()
		if statement != nil {
			stmt.Statements.Statements = append(stmt.Statements.Statements, statement)
		}
		p.advanceToken()
	}
	return stmt
}

// TODO:Add Error Handling
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectNext(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = p.curToken
	p.advanceToken()

	if token.Symbol(p.curToken.Literal) == token.LBRACKET {
		p.advanceToken()

		stmt.Idx = p.parseExpression()
		p.advanceToken()

		if token.Symbol(p.curToken.Literal) != token.RBRACKET {
			return nil
		}
		p.advanceToken()
	}

	if token.Symbol(p.curToken.Literal) != token.ASSIGN {
		return nil
	}

	stmt.Symbol = p.curToken
	p.advanceToken()
	stmt.Value = p.parseExpression()
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) != token.SEMICOLON {
		return nil
	}
	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) == token.SEMICOLON {
		return stmt
	}
	stmt.Value = p.parseExpression()
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) != token.SEMICOLON {
		return nil
	}
	return stmt
}

func (p *Parser) parseDoStatement() *ast.DoStatement {
	stmt := &ast.DoStatement{Token: p.curToken}
	p.advanceToken()

	if token.Symbol(p.nextToken.Literal) == token.DOT {
		stmt.ClassName = p.curToken
		p.advanceToken() // className
		p.advanceToken() // token.DOT
	}
	stmt.VarName = p.curToken
	p.advanceToken()

	if token.Symbol(p.curToken.Literal) != token.LPAREN {
		return nil
	}

	stmt.ExpressionListStmt = p.parseExpressionListStatement()
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) != token.SEMICOLON {
		return nil
	}

	return stmt
}

func (p *Parser) parseVarDecStatement() *ast.VarDecStatement {
	stmt := &ast.VarDecStatement{Token: p.curToken, Identifiers: []token.Token{}}
	if token.KeyWord(p.nextToken.Literal) != token.INT && token.KeyWord(p.nextToken.Literal) != token.BOOLEAN && token.KeyWord(p.nextToken.Literal) != token.CHAR && !p.nextTokenIs(token.IDENTIFIER) {
		return nil
	}
	p.advanceToken()
	stmt.ValueType = p.curToken
	for token.Symbol(p.curToken.Literal) != token.SEMICOLON {
		p.advanceToken()
		identifier := p.curToken
		stmt.Identifiers = append(stmt.Identifiers, identifier)
		p.advanceToken() //
	}
	return stmt
}

func (p *Parser) parseClassVarDecStatement() *ast.ClassVarDecStatement {
	stmt := &ast.ClassVarDecStatement{Token: p.curToken, Identifiers: []token.Token{}}
	p.advanceToken()

	if token.KeyWord(p.curToken.Literal) != token.INT && token.KeyWord(p.curToken.Literal) != token.BOOLEAN && token.KeyWord(p.curToken.Literal) != token.CHAR && !p.curTokenIs(token.IDENTIFIER) {
		return nil
	}
	stmt.ValueType = p.curToken
	p.advanceToken()

	for {
		identifier := p.curToken
		stmt.Identifiers = append(stmt.Identifiers, identifier)
		p.advanceToken()
		// TODO:refactoring
		if token.Symbol(p.curToken.Literal) == token.COMMA {
			p.advanceToken()
			continue
		} else if token.Symbol(p.curToken.Literal) == token.SEMICOLON {
			break
		} else {
			return nil
		}
	}
	if token.Symbol(p.curToken.Literal) != token.SEMICOLON {
		return nil
	}
	return stmt
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	stmt := &ast.IfStatement{Token: p.curToken}
	if p.expectNext(token.SYMBOL) {
		if token.Symbol(p.curToken.Literal) != token.LPAREN {
			return nil
		}
	}
	stmt.Condition = p.parseExpression()
	p.advanceToken()

	if token.Symbol(p.curToken.Literal) != token.RPAREN {
		return nil
	}
	p.advanceToken()

	stmt.Consequence = p.parseBlockStatement()
	p.advanceToken()
	if token.KeyWord(p.curToken.Literal) == token.ELSE {
		p.advanceToken()
		if p.expectNext(token.SYMBOL) {
			return nil
		}
		stmt.Alternative = p.parseBlockStatement()
	}
	return stmt
}

func (p *Parser) parseWhileStatement() *ast.WhileStatement {
	stmt := &ast.WhileStatement{Token: p.curToken}
	if p.expectNext(token.SYMBOL) {
		if token.Symbol(p.curToken.Literal) != token.LPAREN {
			return nil
		}
	}
	p.advanceToken()
	stmt.Condition = p.parseExpression()
	p.advanceToken()

	if token.Symbol(p.curToken.Literal) != token.RPAREN {
		return nil
	}
	p.advanceToken()
	stmt.Statements = p.parseBlockStatement()
	return stmt
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	p.advanceToken()
	block.Statements = []ast.Statement{}
	for token.Symbol(p.curToken.Literal) != token.RBRACE && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.advanceToken()
	}
	return block
}

func (p *Parser) parseExpressionListStatement() *ast.ExpressionListStatement {
	expressionListStmt := &ast.ExpressionListStatement{Token: p.curToken}
	p.advanceToken()
	expressionListStmt.ExpressionList = []ast.Expression{}
	for token.Symbol(p.curToken.Literal) != token.RPAREN {
		expression := p.parseExpression()
		if expression != nil {
			expressionListStmt.ExpressionList = append(expressionListStmt.ExpressionList, expression)
		}
		p.advanceToken()
		if token.Symbol(p.curToken.Literal) == token.RPAREN {
			break
		} else if token.Symbol(p.curToken.Literal) == token.COMMA {
			p.advanceToken()
		} else {
			return nil
		}
	}
	return expressionListStmt
}

func (p *Parser) parseParameterListStatement() *ast.ParameterListStatement {
	parameterListStmt := &ast.ParameterListStatement{Token: p.curToken}
	p.advanceToken()
	parameterListStmt.ParameterList = []ast.ParameterStatement{}
	for token.Symbol(p.curToken.Literal) != token.RPAREN {
		parameterStmt := p.parseParameterStatement()
		if parameterStmt == nil {
			return nil
		}
		parameterListStmt.ParameterList = append(parameterListStmt.ParameterList, *parameterStmt)
		p.advanceToken()
		if token.Symbol(p.curToken.Literal) == token.RPAREN {
			break
		}
		p.advanceToken()
	}
	return parameterListStmt
}

func (p *Parser) parseParameterStatement() *ast.ParameterStatement {
	parameterStmt := &ast.ParameterStatement{Token: p.curToken}
	if p.curToken.Type != token.KEYWORD {
		return nil
	}
	parameterStmt.Type = token.KeyWord(p.curToken.Literal)
	p.advanceToken()
	if p.curToken.Type != token.IDENTIFIER {
		return nil
	}
	parameterStmt.Name = p.curToken.Literal
	return parameterStmt
}

func (p *Parser) parseExpression() ast.Expression {
	expressionToken := p.curToken
	prefixTerm := p.parseTerm()
	if token.Symbol(p.nextToken.Literal) != token.ASSIGN &&
		token.Symbol(p.nextToken.Literal) != token.PLUS &&
		token.Symbol(p.nextToken.Literal) != token.MINUS &&
		token.Symbol(p.nextToken.Literal) != token.ASTERISK &&
		token.Symbol(p.nextToken.Literal) != token.SLASH &&
		token.Symbol(p.nextToken.Literal) != token.LT &&
		token.Symbol(p.nextToken.Literal) != token.GT &&
		token.Symbol(p.nextToken.Literal) != token.EQ &&
		token.Symbol(p.nextToken.Literal) != token.NOT_EQ {
		return &ast.SingleExpression{Token: expressionToken, Value: prefixTerm}
	} else {
		p.advanceToken()
		operator := p.curToken
		p.advanceToken()
		suffixTerm := p.parseTerm()
		return &ast.InfixExpression{Left: prefixTerm, Operator: operator, Right: suffixTerm}
	}
}

func (p *Parser) parseTerm() ast.Term {
	switch p.curToken.Type {
	case token.INTCONST:
		return p.parseIntegerConstTerm()
	case token.IDENTIFIER:
		if token.Symbol(p.nextToken.Literal) == token.LPAREN {
			return p.parseSubroutineCallTerm()
		}
		if token.Symbol(p.nextToken.Literal) == token.DOT {
			return p.parseSubroutineCallTerm()
		}
		if token.Symbol(p.nextToken.Literal) == token.LBRACKET {
			return p.parseArrayElementTerm()
		}
		return p.parseIdentifierTerm()
	case token.STARTINGCONST:
		return p.parseStringConstTerm()
	case token.SYMBOL:
		if token.Symbol(p.curToken.Literal) == token.LPAREN {
			return p.parseBracketTerm()
		}
		if token.Symbol(p.curToken.Literal) == token.MINUS || token.Symbol(p.curToken.Literal) == token.BANG {
			return p.parsePrefixTerm()
		}
	case token.KEYWORD:
		return p.parseKeyWordConstTerm()
	}

	return nil
}

func (p *Parser) parseIntegerConstTerm() ast.Term {
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		panic(fmt.Sprintf("could not parse %q as integer", p.curToken.Literal))
	}
	return &ast.IntergerConstTerm{Token: p.curToken, Value: value}
}

func (p *Parser) parseIdentifierTerm() ast.Term {
	return &ast.IdentifierTerm{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseStringConstTerm() ast.Term {
	return &ast.StringConstTerm{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseKeyWordConstTerm() ast.Term {
	if token.KeyWord(p.curToken.Literal) != token.NULL && token.KeyWord(p.curToken.Literal) != (token.TRUE) && token.KeyWord(p.curToken.Literal) != token.FALSE && token.KeyWord(p.curToken.Literal) != token.THIS {
		panic(fmt.Sprintf("could not parse %s as keywordConst", p.curToken.Literal))
	}
	return &ast.KeywordConstTerm{Token: p.curToken, KeyWord: token.KeyWord(p.curToken.Literal)}
}

func (p *Parser) parseSubroutineCallTerm() ast.Term {
	subroutineCallTerm := &ast.SubroutineCallTerm{Token: p.curToken}
	if token.Symbol(p.nextToken.Literal) == token.DOT {
		subroutineCallTerm.ClassName = p.curToken
		p.advanceToken() // className
		p.advanceToken() // token.DOT
	}
	subroutineCallTerm.VarName = p.curToken
	p.advanceToken()

	if token.Symbol(p.curToken.Literal) != token.LPAREN {
		return nil
	}

	subroutineCallTerm.ExpressionListStmt = p.parseExpressionListStatement()
	return subroutineCallTerm
}

func (p *Parser) parseArrayElementTerm() ast.Term {
	arrayElementTerm := &ast.ArrayElementTerm{Token: p.curToken, ArrayName: p.curToken}
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) != token.LBRACKET {
		return nil
	}
	p.advanceToken()
	idx := p.parseExpression()
	arrayElementTerm.Idx = idx
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) != token.RBRACKET {
		return nil
	}
	return arrayElementTerm
}

func (p *Parser) parsePrefixTerm() ast.Term {
	prefixTerm := &ast.PrefixTerm{Token: p.curToken, Prefix: token.Symbol(p.curToken.Literal)}
	p.advanceToken()
	prefixTerm.Value = p.parseTerm()
	return prefixTerm
}

func (p *Parser) parseBracketTerm() ast.Term {
	bracketTerm := &ast.BracketTerm{Token: p.curToken}
	p.advanceToken()
	expression := p.parseExpression()
	bracketTerm.Value = expression
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) != token.RPAREN {
		return nil
	}
	return bracketTerm
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) nextTokenIs(t token.TokenType) bool {
	return p.nextToken.Type == t
}

func (p *Parser) expectNext(t token.TokenType) bool {
	if p.nextTokenIs(t) {
		p.advanceToken()
		return true
	}
	return false
}
