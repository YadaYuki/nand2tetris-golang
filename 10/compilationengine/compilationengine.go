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

func (ce *CompilationEngine) ParseProgram() *ast.Program {
	return nil
}

func (ce *CompilationEngine) advanceToken() {
	ce.curToken = ce.nextToken
	ce.nextToken, _ = ce.jt.Advance()
}

func (ce *CompilationEngine) parseStatement() {}
