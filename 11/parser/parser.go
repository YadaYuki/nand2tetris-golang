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

// ParseProgram is Parser for all program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.ParseStatement()
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

func (p *Parser) ParseStatement() ast.Statement {
	if p.curToken.Type != token.KEYWORD {
		panic(fmt.Sprintf("Initial Token Type should be KEYWORD. got %s ", p.curToken.Type))
	}
	return p.ParseKeyWord()
}

func (p *Parser) ParseKeyWord() ast.Statement {
	keyWord, _ := tokenizer.KeyWord(p.curToken)
	switch keyWord {
	case token.LET:
		return p.ParseLetStatement()
	case token.RETURN:
		return p.ParseReturnStatement()
	case token.DO:
		return p.ParseDoStatement()
	case token.VAR:
		return p.ParseVarDecStatement()
	case token.STATIC:
		return p.ParseClassVarDecStatement()
	case token.FIELD:
		return p.ParseClassVarDecStatement()
	case token.IF:
		return p.ParseIfStatement()
	case token.WHILE:
		return p.ParseWhileStatement()
	case token.CLASS:
		return p.ParseClassStatement()
	case token.METHOD:
		return p.ParseSubroutineDecStatement()
	case token.CONSTRUCTOR:
		return p.ParseSubroutineDecStatement()
	case token.FUNCTION:
		return p.ParseSubroutineDecStatement()
	default:
		return nil
	}
}

func (p *Parser) ParseClassStatement() *ast.ClassStatement {
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

		classVarDec := p.ParseClassVarDecStatement()
		stmt.ClassVarDecList = append(stmt.ClassVarDecList, *classVarDec)
		p.advanceToken()
	}
	stmt.SubroutineDecList = []ast.SubroutineDecStatement{}
	for token.KeyWord(p.curToken.Literal) == token.CONSTRUCTOR || token.KeyWord(p.curToken.Literal) == token.FUNCTION || token.KeyWord(p.curToken.Literal) == token.METHOD {
		subroutineDec := p.ParseSubroutineDecStatement()
		stmt.SubroutineDecList = append(stmt.SubroutineDecList, *subroutineDec)
		p.advanceToken()
	}
	if token.Symbol(p.curToken.Literal) != token.RBRACE {
		return nil
	}
	return stmt
}

func (p *Parser) ParseSubroutineDecStatement() *ast.SubroutineDecStatement {
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
	stmt.ParameterList = p.ParseParameterListStatement()
	if token.Symbol(p.curToken.Literal) != token.RPAREN {
		return nil
	}
	p.advanceToken()
	stmt.SubroutineBody = p.ParseSubroutineBodyStatement()
	return stmt
}

func (p *Parser) ParseSubroutineBodyStatement() *ast.SubroutineBodyStatement {
	stmt := &ast.SubroutineBodyStatement{Token: p.curToken}
	if token.Symbol(p.curToken.Literal) != token.LBRACE {
		return nil
	}
	p.advanceToken()
	stmt.VarDecList = []ast.VarDecStatement{}
	for token.KeyWord(p.curToken.Literal) == token.VAR {
		varDec := p.ParseVarDecStatement()
		stmt.VarDecList = append(stmt.VarDecList, *varDec)
		p.advanceToken()
	}
	// NOTE: originally, should call ParseBlockStatement.
	// But, the block statement in subroutine body does not start "{".
	// so, implement original Parser here.
	stmt.Statements = &ast.BlockStatement{}
	stmt.Statements.Statements = []ast.Statement{}
	for token.Symbol(p.curToken.Literal) != token.RBRACE && !p.curTokenIs(token.EOF) {
		statement := p.ParseStatement()
		if statement != nil {
			stmt.Statements.Statements = append(stmt.Statements.Statements, statement)
		}
		p.advanceToken()
	}
	return stmt
}

// TODO:Add Error Handling
func (p *Parser) ParseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectNext(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = p.curToken
	p.advanceToken()

	if token.Symbol(p.curToken.Literal) == token.LBRACKET {
		p.advanceToken()

		stmt.Idx = p.ParseExpression()
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
	stmt.Value = p.ParseExpression()
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) != token.SEMICOLON {
		return nil
	}
	return stmt
}

func (p *Parser) ParseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) == token.SEMICOLON {
		return stmt
	}
	stmt.Value = p.ParseExpression()
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) != token.SEMICOLON {
		return nil
	}
	return stmt
}

func (p *Parser) ParseDoStatement() *ast.DoStatement {
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

	stmt.ExpressionListStmt = p.ParseExpressionListStatement()
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) != token.SEMICOLON {
		return nil
	}

	return stmt
}

func (p *Parser) ParseVarDecStatement() *ast.VarDecStatement {
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

func (p *Parser) ParseClassVarDecStatement() *ast.ClassVarDecStatement {
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

func (p *Parser) ParseIfStatement() *ast.IfStatement {
	stmt := &ast.IfStatement{Token: p.curToken}
	if p.expectNext(token.SYMBOL) {
		if token.Symbol(p.curToken.Literal) != token.LPAREN {
			return nil
		}
	}
	stmt.Condition = p.ParseExpression()
	p.advanceToken()

	if token.Symbol(p.curToken.Literal) != token.RPAREN {
		return nil
	}
	p.advanceToken()

	stmt.Consequence = p.ParseBlockStatement()
	p.advanceToken()
	if token.KeyWord(p.curToken.Literal) == token.ELSE {
		p.advanceToken()
		if p.expectNext(token.SYMBOL) {
			return nil
		}
		stmt.Alternative = p.ParseBlockStatement()
	}
	return stmt
}

func (p *Parser) ParseWhileStatement() *ast.WhileStatement {
	stmt := &ast.WhileStatement{Token: p.curToken}
	if p.expectNext(token.SYMBOL) {
		if token.Symbol(p.curToken.Literal) != token.LPAREN {
			return nil
		}
	}
	p.advanceToken()
	stmt.Condition = p.ParseExpression()
	p.advanceToken()

	if token.Symbol(p.curToken.Literal) != token.RPAREN {
		return nil
	}
	p.advanceToken()
	stmt.Statements = p.ParseBlockStatement()
	return stmt
}

func (p *Parser) ParseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	p.advanceToken()
	block.Statements = []ast.Statement{}
	for token.Symbol(p.curToken.Literal) != token.RBRACE && !p.curTokenIs(token.EOF) {
		stmt := p.ParseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.advanceToken()
	}
	return block
}

func (p *Parser) ParseExpressionListStatement() *ast.ExpressionListStatement {
	expressionListStmt := &ast.ExpressionListStatement{Token: p.curToken}
	p.advanceToken()
	expressionListStmt.ExpressionList = []ast.Expression{}
	for token.Symbol(p.curToken.Literal) != token.RPAREN {
		expression := p.ParseExpression()
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

func (p *Parser) ParseParameterListStatement() *ast.ParameterListStatement {
	parameterListStmt := &ast.ParameterListStatement{Token: p.curToken}
	p.advanceToken()
	parameterListStmt.ParameterList = []ast.ParameterStatement{}
	for token.Symbol(p.curToken.Literal) != token.RPAREN {
		parameterStmt := p.ParseParameterStatement()
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

func (p *Parser) ParseParameterStatement() *ast.ParameterStatement {
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

func (p *Parser) ParseExpression() ast.Expression {
	expressionToken := p.curToken
	prefixTerm := p.ParseTerm()
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
		suffixTerm := p.ParseTerm()
		return &ast.InfixExpression{Left: prefixTerm, Operator: operator, Right: suffixTerm}
	}
}

func (p *Parser) ParseTerm() ast.Term {
	switch p.curToken.Type {
	case token.INTCONST:
		return p.ParseIntegerConstTerm()
	case token.IDENTIFIER:
		if token.Symbol(p.nextToken.Literal) == token.LPAREN {
			return p.ParseSubroutineCallTerm()
		}
		if token.Symbol(p.nextToken.Literal) == token.DOT {
			return p.ParseSubroutineCallTerm()
		}
		if token.Symbol(p.nextToken.Literal) == token.LBRACKET {
			return p.ParseArrayElementTerm()
		}
		return p.ParseIdentifierTerm()
	case token.STARTINGCONST:
		return p.ParseStringConstTerm()
	case token.SYMBOL:
		if token.Symbol(p.curToken.Literal) == token.LPAREN {
			return p.ParseBracketTerm()
		}
		if token.Symbol(p.curToken.Literal) == token.MINUS || token.Symbol(p.curToken.Literal) == token.BANG {
			return p.ParsePrefixTerm()
		}
	case token.KEYWORD:
		return p.ParseKeyWordConstTerm()
	}

	return nil
}

func (p *Parser) ParseIntegerConstTerm() ast.Term {
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		panic(fmt.Sprintf("could not Parse %q as integer", p.curToken.Literal))
	}
	return &ast.IntergerConstTerm{Token: p.curToken, Value: value}
}

func (p *Parser) ParseIdentifierTerm() ast.Term {
	return &ast.IdentifierTerm{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) ParseStringConstTerm() ast.Term {
	return &ast.StringConstTerm{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) ParseKeyWordConstTerm() ast.Term {
	if token.KeyWord(p.curToken.Literal) != token.NULL && token.KeyWord(p.curToken.Literal) != (token.TRUE) && token.KeyWord(p.curToken.Literal) != token.FALSE && token.KeyWord(p.curToken.Literal) != token.THIS {
		panic(fmt.Sprintf("could not Parse %s as keywordConst", p.curToken.Literal))
	}
	return &ast.KeywordConstTerm{Token: p.curToken, KeyWord: token.KeyWord(p.curToken.Literal)}
}

func (p *Parser) ParseSubroutineCallTerm() ast.Term {
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

	subroutineCallTerm.ExpressionListStmt = p.ParseExpressionListStatement()
	return subroutineCallTerm
}

func (p *Parser) ParseArrayElementTerm() ast.Term {
	arrayElementTerm := &ast.ArrayElementTerm{Token: p.curToken, ArrayName: p.curToken}
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) != token.LBRACKET {
		return nil
	}
	p.advanceToken()
	idx := p.ParseExpression()
	arrayElementTerm.Idx = idx
	p.advanceToken()
	if token.Symbol(p.curToken.Literal) != token.RBRACKET {
		return nil
	}
	return arrayElementTerm
}

func (p *Parser) ParsePrefixTerm() ast.Term {
	prefixTerm := &ast.PrefixTerm{Token: p.curToken, Prefix: token.Symbol(p.curToken.Literal)}
	p.advanceToken()
	prefixTerm.Value = p.ParseTerm()
	return prefixTerm
}

func (p *Parser) ParseBracketTerm() ast.Term {
	bracketTerm := &ast.BracketTerm{Token: p.curToken}
	p.advanceToken()
	expression := p.ParseExpression()
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
