package ast

import (
	"bytes"
	"jack_compiler/token"
	"strings"
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

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}
func (p *Program) Xml() string {
	var out bytes.Buffer
	out.WriteString("<expression>")
	// TODO: Implement
	// for _, stmt := range p.Statements {
	// 	out.WriteString(stmt.Xml())
	// }
	out.WriteString("</expression>")
	return out.String()
}

// LetStatement is Ast of "let"
type LetStatement struct {
	// TODO:Add array element []
	Token  token.Token // KEYWORD:"let"
	Name   *Identifier
	Symbol token.Token // Symbol:"="
	Value  Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	if ls.Value != nil {
		out.WriteString(ls.Symbol.Literal)
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (ls *LetStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString("<letStatement>")
	out.WriteString(keywordXml(ls.TokenLiteral()))
	out.WriteString(ls.Name.Xml())
	if ls.Value != nil {
		out.WriteString(symbolXml(ls.Symbol.Literal))
		// TODO:implement expression Xml
		// out.WriteString(ls.Value.Xml())
	}
	out.WriteString("</letStatement>")
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

func (i *Identifier) Xml() string { return "<identifier> " + i.Value + " </identifier>" }

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
type ReturnStatement struct {
	Token token.Token // KEYWORD
	Value Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString("<returnStatement>")
	out.WriteString(keywordXml(rs.TokenLiteral()))
	if rs.Value != nil {
		// TODO:implement expression Xml
		// out.WriteString(rs.Value.String())
	}
	out.WriteString("</returnStatement>")
	return out.String()
}

type DoStatement struct {
	Token          token.Token // Keyword:"do"
	SubroutineCall Expression
}

func (ds *DoStatement) statementNode() {}

func (ds *DoStatement) TokenLiteral() string { return ds.Token.Literal }

func (ds *DoStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ds.TokenLiteral() + " ")
	out.WriteString(ds.SubroutineCall.String())
	out.WriteString(";")
	return out.String()
}

func (ds *DoStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString("<doStatement>")
	out.WriteString(keywordXml(ds.TokenLiteral()))
	// out.WriteString(keywordXml(ds.SubroutineCall.Xml))
	out.WriteString("</doStatement>")
	return out.String()
}

type VarDecStatement struct {
	Token       token.Token // Keyword:"var"
	ValueType   token.Token // "int","char","boolean",{class name}
	Identifiers []*Identifier
}

func (vds *VarDecStatement) statementNode() {}

func (vds *VarDecStatement) TokenLiteral() string { return vds.Token.Literal }

func (vds *VarDecStatement) String() string {
	var out bytes.Buffer
	out.WriteString(vds.TokenLiteral() + " ")
	out.WriteString(vds.ValueType.Literal + " ")
	out.WriteString(vds.Identifiers[0].String())
	for _, identifier := range vds.Identifiers[1:] {
		out.WriteString("," + identifier.String())
	}
	out.WriteString(";")
	return out.String()
}
func (vds *VarDecStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString("<varDec>")
	out.WriteString(keywordXml(vds.TokenLiteral()))
	out.WriteString(keywordXml(vds.ValueType.Literal))
	out.WriteString(vds.Identifiers[0].Xml())
	for _, identifier := range vds.Identifiers[1:] {
		out.WriteString(keywordXml(","))
		out.WriteString(identifier.Xml())
	}
	out.WriteString(symbolXml(";"))
	out.WriteString("</varDec>")
	return out.String()
}

type ClassVarDecStatement struct {
	Token       token.Token // Keyword:"static","field"
	ValueType   token.Token // "int","char","boolean",{class name}
	Identifiers []*Identifier
}

func (cvds *ClassVarDecStatement) statementNode() {}

func (cvds *ClassVarDecStatement) TokenLiteral() string { return cvds.Token.Literal }

func (cvds *ClassVarDecStatement) String() string {
	var out bytes.Buffer
	out.WriteString(cvds.TokenLiteral() + " ")
	out.WriteString(cvds.ValueType.Literal + " ")
	out.WriteString(cvds.Identifiers[0].String())
	for _, identifier := range cvds.Identifiers[1:] {
		out.WriteString("," + identifier.String())
	}
	out.WriteString(";")
	return out.String()
}
func (cvds *ClassVarDecStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString("<classVarDec>")
	out.WriteString(keywordXml(cvds.TokenLiteral()))
	out.WriteString(keywordXml(cvds.ValueType.Literal))
	out.WriteString(cvds.Identifiers[0].Xml())
	for _, identifier := range cvds.Identifiers[1:] {
		out.WriteString(keywordXml(","))
		out.WriteString(identifier.Xml())
	}
	out.WriteString(symbolXml(";"))
	out.WriteString("</classVarDec>")
	return out.String()
}

type BaseExpression struct {
	Token token.Token // 式の最初のトークン
	Terms []TermStatement
}

func (bs *BaseExpression) statementNode() {}

func (bs *BaseExpression) TokenLiteral() string { return es.Token.Literal }

func (bs *BaseExpression) String() string {
	// if bs.Expression != nil {
	// 	return es.Expression.String()
	// }
	return ""
}

type TermStatement struct {
	Token token.Token // 式の最初のトークン
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(")")
	out.WriteString(fl.Body.String())
	return out.String()
}

type SubroutineCallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (cfe *SubroutineCallExpression) expressionNode()      {}
func (cfe *SubroutineCallExpression) TokenLiteral() string { return cfe.Token.Literal }

func (cfe *SubroutineCallExpression) String() string {

	var out bytes.Buffer
	args := []string{}

	for _, a := range cfe.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(cfe.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ","))
	out.WriteString(")")

	return out.String()
}

func keywordXml(keyword string) string       { return "<keyword> " + keyword + " </keyword>" }
func symbolXml(symbol string) string         { return "<symbol> " + symbol + " </symbol>" }
func identifierXml(identifier string) string { return "<identifier> " + identifier + " </identifier>" }
