package ast

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
