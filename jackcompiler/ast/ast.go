package ast

// TODO: Xml() function to encoding/xml package. https://pkg.go.dev/encoding/xml

import (
	"bytes"
	"jackcompiler/token"
	"strconv"
	"strings"
)

// Node is Node of AST
type Node interface {
	TokenLiteral() string
	String() string
	Xml() string
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

type Term interface {
	Node
	termNode()
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
	for _, stmt := range p.Statements {
		out.WriteString(stmt.Xml())
	}
	return out.String()
}

type ClassStatement struct {
	Token             token.Token // KEYWORD:"class"
	Name              token.Token
	ClassVarDecList   []ClassVarDecStatement
	SubroutineDecList []SubroutineDecStatement
}

func (cs *ClassStatement) statementNode() {}

func (cs *ClassStatement) TokenLiteral() string { return cs.Token.Literal }

func (cs *ClassStatement) String() string {
	var out bytes.Buffer
	out.WriteString(cs.TokenLiteral() + " ")
	out.WriteString(cs.Name.Literal)
	out.WriteString("{")
	for _, classVarDec := range cs.ClassVarDecList {
		out.WriteString(classVarDec.String())
	}
	for _, subroutineDec := range cs.SubroutineDecList {
		out.WriteString(subroutineDec.String())
	}
	out.WriteString("}")
	return out.String()
}

func (cs *ClassStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString("<class>")
	out.WriteString(cs.Token.Xml())
	out.WriteString(cs.Name.Xml())
	out.WriteString(symbolXml("{"))
	for _, classVarDec := range cs.ClassVarDecList {
		out.WriteString(classVarDec.Xml())
	}
	for _, subroutineDec := range cs.SubroutineDecList {
		out.WriteString(subroutineDec.Xml())
	}
	out.WriteString(symbolXml("}"))
	out.WriteString("</class>")
	return out.String()
}

type SubroutineDecStatement struct {
	Token          token.Token // KEYWORD:"constructor" or "method" or "function"
	ReturnType     token.Token // KEYWORD:"void" or IDENTIFIER
	Name           token.Token // IDENTIFIER
	ParameterList  *ParameterListStatement
	SubroutineBody *SubroutineBodyStatement
}

func (sds *SubroutineDecStatement) statementNode() {}

func (sds *SubroutineDecStatement) TokenLiteral() string { return sds.Token.Literal }

func (sds *SubroutineDecStatement) String() string {
	var out bytes.Buffer
	out.WriteString(sds.TokenLiteral() + " ")
	out.WriteString(sds.ReturnType.Literal + " ")
	out.WriteString(sds.Name.Literal)
	out.WriteString(sds.ParameterList.String())
	out.WriteString(sds.SubroutineBody.String())
	return out.String()
}

func (sds *SubroutineDecStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString("<subroutineDec>")
	out.WriteString(sds.Token.Xml())
	out.WriteString(sds.ReturnType.Xml())
	out.WriteString(identifierXml(sds.Name.Literal))
	out.WriteString(sds.ParameterList.Xml())
	out.WriteString(sds.SubroutineBody.Xml())
	out.WriteString("</subroutineDec>")
	return out.String()
}

type SubroutineBodyStatement struct {
	Token      token.Token // "{"
	VarDecList []VarDecStatement
	Statements []Statement
}

func (sbs *SubroutineBodyStatement) statementNode() {}

func (sbs *SubroutineBodyStatement) TokenLiteral() string { return sbs.Token.Literal }

func (sbs *SubroutineBodyStatement) String() string {
	var out bytes.Buffer
	out.WriteString("{")
	for _, varDec := range sbs.VarDecList {
		out.WriteString(varDec.String())
	}
	for _, stmt := range sbs.Statements {
		out.WriteString(stmt.String())
	}
	out.WriteString("}")
	return out.String()
}

func (sbs *SubroutineBodyStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString("<subroutineBody>")
	out.WriteString(symbolXml("{"))
	for _, varDec := range sbs.VarDecList {
		out.WriteString(varDec.Xml())
	}
	out.WriteString("<statements>")
	for _, stmt := range sbs.Statements {
		out.WriteString(stmt.Xml())
	}
	out.WriteString("</statements>")
	out.WriteString(symbolXml("}"))
	out.WriteString("</subroutineBody>")
	return out.String()
}

// LetStatement is Ast of "let"
type LetStatement struct {
	Token  token.Token // KEYWORD:"let"
	Name   token.Token
	Idx    Expression
	Symbol token.Token // Symbol:"="
	Value  Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.Literal)
	if ls.Idx != nil {
		out.WriteString("[" + ls.Idx.String() + "]")
	}
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
	out.WriteString(ls.Token.Xml())
	out.WriteString(identifierXml(ls.Name.Literal))
	if ls.Idx != nil {
		out.WriteString(symbolXml("[") + ls.Idx.Xml() + symbolXml("]"))
	}
	if ls.Value != nil {
		out.WriteString(symbolXml(ls.Symbol.Literal))
		out.WriteString(ls.Value.Xml())
	}
	out.WriteString(symbolXml(";"))
	out.WriteString("</letStatement>")
	return out.String()
}

// ReturnStatement is Ast of "return"
type ReturnStatement struct {
	Token token.Token // KEYWORD
	Value Expression
}

func (ds *ReturnStatement) statementNode() {}

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
	out.WriteString(rs.Token.Xml())
	if rs.Value != nil {
		out.WriteString(rs.Value.Xml())
	}
	out.WriteString(symbolXml(";"))
	out.WriteString("</returnStatement>")
	return out.String()
}

type DoStatement struct {
	Token              token.Token // Keyword:"do"
	ClassName          token.Token
	SubroutineName     token.Token
	ExpressionListStmt *ExpressionListStatement
}

func (ds *DoStatement) statementNode() {}

func (ds *DoStatement) TokenLiteral() string { return ds.Token.Literal }

func (ds *DoStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ds.TokenLiteral() + " ")
	if ds.ClassName.Literal != "" {
		out.WriteString(ds.ClassName.Literal + ".")
	}
	out.WriteString(ds.SubroutineName.Literal)
	out.WriteString(ds.ExpressionListStmt.String())
	out.WriteString(";")
	return out.String()
}

func (ds *DoStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString("<doStatement>")
	out.WriteString(ds.Token.Xml())
	if ds.ClassName.Literal != "" {
		out.WriteString(identifierXml(ds.ClassName.Literal) + symbolXml("."))
	}
	out.WriteString(identifierXml(ds.SubroutineName.Literal))
	out.WriteString(ds.ExpressionListStmt.Xml())
	out.WriteString(symbolXml(";"))
	out.WriteString("</doStatement>")
	return out.String()
}

type VarDecStatement struct {
	Token       token.Token // Keyword:"var"
	ValueType   token.Token // "int","char","boolean",{class name}
	Identifiers []token.Token
}

func (vds *VarDecStatement) statementNode() {}

func (vds *VarDecStatement) TokenLiteral() string { return vds.Token.Literal }

func (vds *VarDecStatement) String() string {
	var out bytes.Buffer
	out.WriteString(vds.TokenLiteral() + " ")
	out.WriteString(vds.ValueType.Literal + " ")
	identifiersStringLs := []string{}
	for _, identifier := range vds.Identifiers {
		identifiersStringLs = append(identifiersStringLs, identifier.Literal)
	}
	out.WriteString(strings.Join(identifiersStringLs, ","))
	out.WriteString(";")
	return out.String()
}
func (vds *VarDecStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString("<varDec>")
	out.WriteString(vds.Token.Xml())
	out.WriteString(vds.ValueType.Xml())
	identifiersXmlLs := []string{}
	for _, identifier := range vds.Identifiers {
		identifiersXmlLs = append(identifiersXmlLs, identifierXml(identifier.Literal))
	}
	out.WriteString(strings.Join(identifiersXmlLs, symbolXml(",")))
	out.WriteString(symbolXml(";"))
	out.WriteString("</varDec>")
	return out.String()
}

type ClassVarDecStatement struct {
	Token       token.Token // Keyword:"static","field"
	ValueType   token.Token // "int","char","boolean",{class name}
	Identifiers []token.Token
}

func (cvds *ClassVarDecStatement) statementNode() {}

func (cvds *ClassVarDecStatement) TokenLiteral() string { return cvds.Token.Literal }

func (cvds *ClassVarDecStatement) String() string {
	var out bytes.Buffer
	out.WriteString(cvds.TokenLiteral() + " ")
	out.WriteString(cvds.ValueType.Literal + " ")
	identifiersStringLs := []string{}
	for _, identifier := range cvds.Identifiers {
		identifiersStringLs = append(identifiersStringLs, identifier.Literal)
	}
	out.WriteString(strings.Join(identifiersStringLs, ","))
	out.WriteString(";")
	return out.String()
}
func (cvds *ClassVarDecStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString("<classVarDec>")
	out.WriteString(cvds.Token.Xml())
	out.WriteString(cvds.ValueType.Xml())
	identifiersStringXml := []string{}
	for _, identifier := range cvds.Identifiers {
		identifiersStringXml = append(identifiersStringXml, identifierXml(identifier.Literal))
	}
	out.WriteString(strings.Join(identifiersStringXml, symbolXml(",")))
	out.WriteString(symbolXml(";"))
	out.WriteString("</classVarDec>")
	return out.String()
}

type IfStatement struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ifs *IfStatement) statementNode() {}

func (ifs *IfStatement) TokenLiteral() string { return ifs.Token.Literal }

func (ifs *IfStatement) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ifs.Condition.String())
	out.WriteString(" ")
	out.WriteString(ifs.Consequence.String())
	if ifs.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ifs.Alternative.String())
	}
	return out.String()
}

func (ifs *IfStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString("<ifStatement>")
	out.WriteString(keywordXml("if"))
	out.WriteString(symbolXml("("))
	out.WriteString(ifs.Condition.Xml())
	out.WriteString(symbolXml(")"))
	out.WriteString(symbolXml("{"))
	out.WriteString(ifs.Consequence.Xml())
	out.WriteString(symbolXml("}"))
	if ifs.Alternative != nil {
		out.WriteString(keywordXml("else"))
		out.WriteString(symbolXml("{"))
		out.WriteString(ifs.Alternative.Xml())
		out.WriteString(symbolXml("}"))
	}
	out.WriteString("</ifStatement>")
	return out.String()
}

type WhileStatement struct {
	Token      token.Token
	Condition  Expression
	Statements *BlockStatement
}

func (ws *WhileStatement) statementNode() {}

func (ws *WhileStatement) TokenLiteral() string { return ws.Token.Literal }

func (ws *WhileStatement) String() string {
	var out bytes.Buffer
	out.WriteString("while")
	out.WriteString("(")
	out.WriteString(ws.Condition.String())
	out.WriteString(")")
	out.WriteString(" ")
	out.WriteString(ws.Statements.String())
	return out.String()
}

func (ws *WhileStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString("<whileStatement>")
	out.WriteString(keywordXml("while"))
	out.WriteString(symbolXml("("))
	out.WriteString(ws.Condition.Xml())
	out.WriteString(symbolXml(")"))
	out.WriteString(symbolXml("{"))
	out.WriteString(ws.Statements.Xml())
	out.WriteString(symbolXml("}"))
	out.WriteString("</whileStatement>")
	return out.String()
}

type BlockStatement struct {
	Token      token.Token // symbol,{
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (bs *BlockStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString("<statements>")
	for _, s := range bs.Statements {
		out.WriteString(s.Xml())
	}
	out.WriteString("</statements>")
	return out.String()
}

type ExpressionListStatement struct {
	Token          token.Token // symbol,(
	ExpressionList []Expression
}

func (els *ExpressionListStatement) statementNode() {}

func (els *ExpressionListStatement) TokenLiteral() string { return els.Token.Literal }

func (els *ExpressionListStatement) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	expressionListStrs := []string{}
	for _, s := range els.ExpressionList {
		expressionListStrs = append(expressionListStrs, s.String())
	}
	out.WriteString(strings.Join(expressionListStrs, ","))
	out.WriteString(")")
	return out.String()
}

func (els *ExpressionListStatement) Xml() string {
	var out bytes.Buffer
	// TODO: Fix from "(" to token.LPAWN. others as well
	out.WriteString(symbolXml("("))
	out.WriteString("<expressionList>")
	expressionListXmls := []string{}
	for _, s := range els.ExpressionList {
		expressionListXmls = append(expressionListXmls, s.Xml())
	}
	out.WriteString(strings.Join(expressionListXmls, symbolXml((","))))
	out.WriteString("</expressionList>")
	out.WriteString(symbolXml(")"))
	return out.String()
}

//関数定義時の引数リスト
type ParameterListStatement struct {
	Token         token.Token // symbol,(
	ParameterList []ParameterStatement
}

func (pls *ParameterListStatement) statementNode() {}

func (pls *ParameterListStatement) TokenLiteral() string { return pls.Token.Literal }

func (pls *ParameterListStatement) String() string {
	if len(pls.ParameterList) == 0 {
		return "()"
	}
	parameterList := []string{}
	for _, s := range pls.ParameterList {
		parameterList = append(parameterList, s.String())
	}
	return "(" + strings.Join(parameterList, ",") + ")"
}

func (pls *ParameterListStatement) Xml() string {
	var out bytes.Buffer
	out.WriteString(symbolXml("("))
	out.WriteString("<parameterList>")
	parameterListXml := []string{}
	for _, s := range pls.ParameterList {
		parameterListXml = append(parameterListXml, s.Xml())
	}
	out.WriteString(strings.Join(parameterListXml, symbolXml(",")))
	out.WriteString("</parameterList>")
	out.WriteString(symbolXml(")"))
	return out.String()
}

type ParameterStatement struct {
	Token     token.Token // 式の最初のトークン
	ValueType token.Token // "int","char","boolean",{class name}
	Name      string
}

func (ps *ParameterStatement) statementNode() {}

func (ps *ParameterStatement) TokenLiteral() string { return ps.Token.Literal }

func (ps *ParameterStatement) String() string {
	return ps.ValueType.String() + " " + ps.Name
}

func (ps *ParameterStatement) Xml() string {
	return ps.ValueType.Xml() + identifierXml(ps.Name)
}

type SingleExpression struct {
	Token token.Token // 式の最初のトークン
	Value Term
}

func (pe *SingleExpression) expressionNode() {}

func (pe *SingleExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *SingleExpression) String() string {
	return pe.Value.String()
}

func (pe *SingleExpression) Xml() string {
	var out bytes.Buffer
	out.WriteString("<expression>")
	out.WriteString(pe.Value.Xml())
	out.WriteString("</expression>")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token // 式の最初のトークン
	Left     Term
	Operator token.Token
	Right    Term
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *InfixExpression) String() string {
	return ie.Left.String() + ie.Operator.Literal + ie.Right.String()
}

func (ie *InfixExpression) Xml() string {
	var out bytes.Buffer
	out.WriteString("<expression>")
	out.WriteString(ie.Left.Xml())
	switch token.Symbol(ie.Operator.String()) {
	case token.LT:
		out.WriteString(symbolXml("&lt;"))
	case token.GT:
		out.WriteString(symbolXml("&gt;"))
	case token.AMP:
		out.WriteString(symbolXml("&amp;"))
	default:
		out.WriteString(symbolXml(ie.Operator.Literal))
	}
	out.WriteString(ie.Right.Xml())
	out.WriteString("</expression>")
	return out.String()
}

type IntergerConstTerm struct {
	Token token.Token
	Value int64
}

func (ict *IntergerConstTerm) termNode() {}

func (ict *IntergerConstTerm) TokenLiteral() string { return ict.Token.Literal }

func (ict *IntergerConstTerm) String() string {
	return strconv.FormatInt(ict.Value, 10)
}
func (ict *IntergerConstTerm) Xml() string {
	return termXml(ict.Token.Xml())
}

type StringConstTerm struct {
	Token token.Token
	Value string
}

func (sct *StringConstTerm) termNode() {}

func (sct *StringConstTerm) TokenLiteral() string { return sct.Token.Literal }

func (sct *StringConstTerm) String() string {
	return `"` + sct.Value + `"`
}
func (sct *StringConstTerm) Xml() string {
	return termXml(sct.Token.Xml())
}

type IdentifierTerm struct {
	Token token.Token
	Value string
}

func (ict *IdentifierTerm) termNode() {}

func (ict *IdentifierTerm) TokenLiteral() string { return ict.Token.Literal }

func (ict *IdentifierTerm) String() string {
	return ict.Value
}
func (ict *IdentifierTerm) Xml() string {
	return termXml(identifierXml(ict.String()))
}

type KeywordConstTerm struct {
	Token   token.Token
	KeyWord token.KeyWord
}

func (kct *KeywordConstTerm) termNode() {}

func (kct *KeywordConstTerm) TokenLiteral() string { return kct.Token.Literal }

func (kct *KeywordConstTerm) String() string {
	return string(kct.KeyWord)
}

func (kct *KeywordConstTerm) Xml() string {
	return termXml(kct.Token.Xml())
}

type SubroutineCallTerm struct {
	Token              token.Token // FunctionName
	ClassName          token.Token
	SubroutineName     token.Token
	ExpressionListStmt *ExpressionListStatement
}

func (sct *SubroutineCallTerm) termNode() {}

func (sct *SubroutineCallTerm) TokenLiteral() string { return sct.Token.Literal }

func (sct *SubroutineCallTerm) String() string {
	var out bytes.Buffer
	if sct.ClassName.Literal != "" {
		out.WriteString(sct.ClassName.Literal + ".")
	}
	out.WriteString(sct.SubroutineName.Literal)
	out.WriteString(sct.ExpressionListStmt.String())
	return out.String()
}

func (sct *SubroutineCallTerm) Xml() string {
	var out bytes.Buffer
	if sct.ClassName.Literal != "" {
		out.WriteString(sct.ClassName.Xml() + symbolXml("."))
	}
	out.WriteString(sct.SubroutineName.Xml())
	out.WriteString(sct.ExpressionListStmt.Xml())
	return termXml(out.String())
}

type ArrayElementTerm struct {
	Token     token.Token // Identifier
	ArrayName token.Token
	Idx       Expression
}

func (aet *ArrayElementTerm) termNode() {}

func (aet *ArrayElementTerm) TokenLiteral() string { return aet.Token.Literal }

func (aet *ArrayElementTerm) String() string {
	var out bytes.Buffer
	out.WriteString(aet.ArrayName.String())
	out.WriteString("[" + aet.Idx.String() + "]")
	return out.String()
}

func (aet *ArrayElementTerm) Xml() string {
	var out bytes.Buffer
	out.WriteString(aet.ArrayName.Xml())
	out.WriteString(symbolXml("["))
	out.WriteString(aet.Idx.Xml())
	out.WriteString(symbolXml("]"))
	return termXml(out.String())
}

type PrefixTerm struct {
	Token  token.Token  // "-","~"
	Prefix token.Symbol // "-","~"
	Value  Term
}

func (pt *PrefixTerm) termNode() {}

func (pt *PrefixTerm) TokenLiteral() string { return pt.Token.Literal }

func (pt *PrefixTerm) String() string {
	return string(pt.Prefix) + pt.Value.String()
}

func (pt *PrefixTerm) Xml() string {
	return termXml(symbolXml(string(pt.Prefix)) + pt.Value.Xml())
}

type BracketTerm struct {
	Token token.Token // "("
	Value Expression
}

func (bt *BracketTerm) termNode() {}

func (bt *BracketTerm) TokenLiteral() string { return bt.Token.Literal }

func (bt *BracketTerm) String() string {
	return "(" + bt.Value.String() + ")"
}

func (bt *BracketTerm) Xml() string {
	return termXml(symbolXml("(") + bt.Value.Xml() + symbolXml(")"))
}

func keywordXml(keyword string) string       { return "<keyword> " + keyword + " </keyword>" }
func symbolXml(symbol string) string         { return "<symbol> " + symbol + " </symbol>" }
func identifierXml(identifier string) string { return "<identifier> " + identifier + " </identifier>" }
func termXml(term string) string             { return "<term>" + term + "</term>" }
