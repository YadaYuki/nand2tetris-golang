package compilationengine

import (
	"jack/compiler/ast"
	"jack/compiler/token"
	"jack/compiler/tokenizer"
)

// CompilationEngine is struct
type CompilationEngine struct {
	jt *tokenizer.JackTokenizer

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
		ce.nextToken()
	}
	return nil
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
		return nil
	}
}

func (ce *CompilationEngine) parseKeyWord() ast.Statement {
	switch tokenizer.KeyWord(ce.curToken) {
	case token.LET:
		return ce.parseLetStatement()
	default:
		return nil
	}
}

func (ce *CompilationEngine) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: ce.curToken}
	if !ce.expectNext(token.IDENTIFIER) {
		return nil
	}
	return nil
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
	} else {
		return false
	}
}
