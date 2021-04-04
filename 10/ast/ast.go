package ast

import (
	"jack/compiler/token"
	"bytes"
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

func (p *Program) String() string{
	var out bytes.Buffer
	for _,stmt := range p.Statements{
		out.WriteString(stmt.String())
	}
	return out.String()
}

// LetStatement is Ast of "let"
type LetStatement struct {
	Token token.Token // KEYWORD
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}

func (ls *LetStatement) String() string{
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString("=")
	if ls.Value != nil{
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Identifier is variable identifier type
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns literal of token
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string {
	return i.Value
}

type IntConst struct {
	Token token.Token
	Value int64
}

func (i *IntConst) expressionNode() {}

// TokenLiteral returns literal of token
func (i *IntConst) TokenLiteral() string { return i.Token.Literal }

func (i *IntConst) String() string {
	return i.Token.Literal
}

// ReturnStatement is Ast of "return"
type ReturnStatement struct{
	Token token.Token // KEYWORD
	Value Expression
}
func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {return rs.Token.Literal}

func (rs *ReturnStatement) String() string{
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.Value != nil{
		out.WriteString(rs.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token token.Token // 式の最初のトークン
	Expression Expression  
}

func (es *ExpressionStatement) statementNode(){}

func (es *ExpressionStatement) TokenLiteral() string {return es.Token.Literal}

func (es *ExpressionStatement) String() string{
	if es.Expression != nil{
		return es.Expression.String()
	}
	return ""
}

type PrefixExpression struct {
	Token token.Token
	Operator string
	Right Expression
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string {return pe.Token.Literal}
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token token.Token
	Left  Expression
	Operator string
	Right Expression
}

func (ie *InfixExpression) expressionNode(){}
func (ie *InfixExpression) TokenLiteral() string{ return ie.Token.Literal }

func (ie *InfixExpression) String() string{
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " +ie.Operator+ " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}


