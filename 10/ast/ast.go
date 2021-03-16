package ast

import (
	"jack_compiler/token"
)

// Node is Node of AST
type Node interface {
	TokenLiteral() string
}

// Statement is Statement Node of AST
type Statement interface {
	Node
	statementNode()
}

// Expression is Expressiont Node of AST
type Expression interface {
	Node
	expressionNode()
}

// LetStatement is Ast of "let"
type LetStatement struct {
	Statement
	Name  *Identifier
	Value *Expression
}

// Identifier is let variable type
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns literal of token
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
