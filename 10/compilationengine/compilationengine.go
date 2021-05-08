package compilationengine

import (
	"fmt"
	"jack_compiler/ast"
	"jack_compiler/token"
	"jack_compiler/tokenizer"
	"strconv"
)

// CompilationEngine is struct
type CompilationEngine struct {
	jt        *tokenizer.JackTokenizer
	curToken  token.Token
	nextToken token.Token
}

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

	if token.Symbol(ce.curToken.Literal) == token.LBRACKET {
		ce.advanceToken()

		stmt.Idx = ce.parseExpression()
		ce.advanceToken()

		if token.Symbol(ce.curToken.Literal) != token.RBRACKET {
			return nil
		}
		ce.advanceToken()
	}

	if token.Symbol(ce.curToken.Literal) != token.ASSIGN {
		return nil
	}

	stmt.Symbol = ce.curToken
	ce.advanceToken()
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
	stmt.Value = ce.parseExpression()
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
	if token.KeyWord(ce.nextToken.Literal) != token.INT && token.KeyWord(ce.nextToken.Literal) != token.BOOLEAN && token.KeyWord(ce.nextToken.Literal) != token.CHAR && !ce.nextTokenIs(token.IDENTIFIER) {
		return nil
	}
	ce.advanceToken()
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
	if token.KeyWord(ce.nextToken.Literal) != token.INT && token.KeyWord(ce.nextToken.Literal) != token.BOOLEAN && token.KeyWord(ce.nextToken.Literal) != token.CHAR && !ce.nextTokenIs(token.IDENTIFIER) {
		return nil
	}
	stmt.ValueType = ce.curToken
	for token.Symbol(ce.curToken.Literal) != token.SEMICOLON {
		ce.advanceToken()
		identifier := ce.curToken
		stmt.Identifiers = append(stmt.Identifiers, identifier)
		ce.advanceToken()
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
	stmt.Condition = ce.parseExpression()
	ce.advanceToken()

	if token.Symbol(ce.curToken.Literal) != token.RPAREN {
		return nil
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
	ce.advanceToken()
	stmt.Condition = ce.parseExpression()
	ce.advanceToken()

	if token.Symbol(ce.curToken.Literal) != token.RPAREN {
		return nil
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
	expressionToken := ce.curToken
	prefixTerm := ce.parseTerm()
	if token.Symbol(ce.nextToken.Literal) != token.ASSIGN &&
		token.Symbol(ce.nextToken.Literal) != token.PLUS &&
		token.Symbol(ce.nextToken.Literal) != token.MINUS &&
		token.Symbol(ce.nextToken.Literal) != token.ASTERISK &&
		token.Symbol(ce.nextToken.Literal) != token.SLASH &&
		token.Symbol(ce.nextToken.Literal) != token.LT &&
		token.Symbol(ce.nextToken.Literal) != token.GT &&
		token.Symbol(ce.nextToken.Literal) != token.EQ &&
		token.Symbol(ce.nextToken.Literal) != token.NOT_EQ {
		return &ast.SingleExpression{Token: expressionToken, Value: prefixTerm}
	} else {
		ce.advanceToken()
		operator := ce.curToken
		ce.advanceToken()
		suffixTerm := ce.parseTerm()
		return &ast.InfixExpression{Left: prefixTerm, Operator: operator, Right: suffixTerm}
	}
}

func (ce *CompilationEngine) parseTerm() ast.Term {
	switch ce.curToken.Type {
	case token.INTCONST:
		return ce.parseIntegerConstTerm()
	case token.IDENTIFIER:
		if token.Symbol(ce.nextToken.Literal) == token.LPAREN {
			return ce.parseSubroutineCallTerm()
		}
		if token.Symbol(ce.nextToken.Literal) == token.LBRACKET {
			return ce.parseArrayElementTerm()
		}
		return ce.parseIdentifierTerm()
	case token.STARTINGCONST:
		return ce.parseStringConstTerm()
	case token.SYMBOL:
		if token.Symbol(ce.curToken.Literal) == token.LPAREN {
			return ce.parseBracketTerm()
		}
		if token.Symbol(ce.curToken.Literal) == token.MINUS || token.Symbol(ce.curToken.Literal) == token.BANG {
			return ce.parsePrefixTerm()
		}
	case token.KEYWORD:
		return ce.parseKeyWordConstTerm()
	}

	return nil
}

func (ce *CompilationEngine) parseIntegerConstTerm() ast.Term {
	value, err := strconv.ParseInt(ce.curToken.Literal, 0, 64)
	if err != nil {
		panic(fmt.Sprintf("could not parse %q as integer", ce.curToken.Literal))
	}
	return &ast.IntergerConstTerm{Token: ce.curToken, Value: value}
}

func (ce *CompilationEngine) parseIdentifierTerm() ast.Term {
	return &ast.IdentifierTerm{Token: ce.curToken, Value: ce.curToken.Literal}
}

func (ce *CompilationEngine) parseStringConstTerm() ast.Term {
	return &ast.StringConstTerm{Token: ce.curToken, Value: ce.curToken.Literal}
}

func (ce *CompilationEngine) parseKeyWordConstTerm() ast.Term {
	if token.KeyWord(ce.curToken.Literal) != token.NULL && token.KeyWord(ce.curToken.Literal) != (token.TRUE) && token.KeyWord(ce.curToken.Literal) != token.FALSE && token.KeyWord(ce.curToken.Literal) != token.THIS {
		panic(fmt.Sprintf("could not parse %s as keywordConst", ce.curToken.Literal))
	}
	return &ast.KeywordConstTerm{Token: ce.curToken, KeyWord: token.KeyWord(ce.curToken.Literal)}
}

func (ce *CompilationEngine) parseSubroutineCallTerm() ast.Term {
	ce.advanceToken()
	expressionListStmt := ce.parseExpressionListStatement()
	return &ast.SubroutineCallTerm{Token: ce.curToken, FunctionName: ce.curToken.Literal, ExpressionListStmt: *expressionListStmt}
}

func (ce *CompilationEngine) parseArrayElementTerm() ast.Term {
	arrayElementTerm := &ast.ArrayElementTerm{Token: ce.curToken, ArrayName: ce.curToken.Literal}
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.LBRACKET {
		return nil
	}
	ce.advanceToken()
	idx := ce.parseExpression()
	arrayElementTerm.Idx = idx
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.RBRACKET {
		return nil
	}
	return arrayElementTerm
}

func (ce *CompilationEngine) parsePrefixTerm() ast.Term {
	prefixTerm := &ast.PrefixTerm{Token: ce.curToken, Prefix: token.Symbol(ce.curToken.Literal)}
	ce.advanceToken()
	prefixTerm.Value = ce.parseTerm()
	return prefixTerm
}

func (ce *CompilationEngine) parseBracketTerm() ast.Term {
	bracketTerm := &ast.BracketTerm{Token: ce.curToken}
	ce.advanceToken()
	expression := ce.parseExpression()
	bracketTerm.Value = expression
	ce.advanceToken()
	if token.Symbol(ce.curToken.Literal) != token.RPAREN {
		return nil
	}
	return bracketTerm
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
