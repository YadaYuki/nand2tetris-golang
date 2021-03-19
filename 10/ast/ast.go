package ast

import (
	"jack/compiler/token"
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

// Expression is Expression Node of AST
type Expression interface {
	Node
	expressionNode()
}

// Program is Ast of all program
type Program struct {
	Statements []Statement
}

// LetStatement is Ast of "let"
type LetStatement struct {
	Statement
	Token token.Token
	Name  *Identifier
	Value Expression
}

// Identifier is variable identifier type
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns literal of token
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
