package ast

import (
	"jack/compiler/token"
)

// Node is Node of AST
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) TokenLiteral() string{
	if len(p.Statements) > 0{
		return p.Statements[0].TokenLiteral()
	}else{
		return ""
	}
}

// LetStatement is Ast of "let"
type LetStatement struct {
	Token token.Token // KEYWORD
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}


// Identifier is variable identifier type
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns literal of token
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }


// ReturnStatement is Ast of "return"
type ReturnStatement struct{
	Token token.Token // KEYWORD
	Value Expression
}
func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {return rs.Token.Literal}

type ExpressionStatement struct {
	Token token.Token // 式の最初のトークン
	Expression Expression  
}

func (es *ExpressionStatement) statementNode(){}

func (es *ExpressionStatement) TokenLiteral() string {return es.Token.Literal}

