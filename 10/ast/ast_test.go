package ast

import (
	"jack_compiler/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.KEYWORD, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "myVar"},
					Value: "myVar",
				},
				Symbol: token.Token{Type: token.SYMBOL, Literal: "="},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "let myVar=anotherVar;" {
		t.Errorf("program.String() wrong. got = %q", program.String())
	}
}

func TestLetXml(t *testing.T) {
	letStatement := &LetStatement{
		Token: token.Token{Type: token.KEYWORD, Literal: "let"},
		Name: &Identifier{
			Token: token.Token{Type: token.IDENTIFIER, Literal: "myVar"},
			Value: "myVar",
		},
		Symbol: token.Token{Type: token.SYMBOL, Literal: "="},
		Value: &Identifier{
			Token: token.Token{Type: token.IDENTIFIER, Literal: "anotherVar"},
			Value: "anotherVar",
		},
	}
	t.Log(letStatement.String())
	t.Log(letStatement.Xml())
}

func TestDoXml(t *testing.T) {
	doStatement := &DoStatement{
		Token: token.Token{Type: token.KEYWORD, Literal: "do"},
		SubroutineCall: &Identifier{
			Token: token.Token{Type: token.IDENTIFIER, Literal: "anotherVar"},
			Value: "anotherVar",
		},
		// SubroutineCall: &SubroutineCallExpression{
		// 	Token: token.Token{Type: token.IDENTIFIER, Literal: "anotherVar"},
		// 	Value: "anotherVar",
		// 	Arguments:
		// },
	}
	t.Log(doStatement.String())
	t.Log(doStatement.Xml())
}

func TestVarDecString(t *testing.T) {
	varDecStatement := &VarDecStatement{
		Token:     token.Token{Type: token.KEYWORD, Literal: "var"},
		ValueType: token.Token{Type: token.KEYWORD, Literal: "int"},
		Identifiers: []*Identifier{
			&Identifier{
				Token: token.Token{Type: token.IDENTIFIER, Literal: "hogehoge"},
				Value: "hogehgoe",
			},
			&Identifier{
				Token: token.Token{Type: token.IDENTIFIER, Literal: "hogehoge"},
				Value: "hogehgoe",
			},
			&Identifier{
				Token: token.Token{Type: token.IDENTIFIER, Literal: "hogehoge"},
				Value: "hogehgoe",
			},
		},
	}
	t.Log(varDecStatement.String())
	t.Log(varDecStatement.Xml())
}

func TestClassVarDecString(t *testing.T) {
	classVarDecStatement := &ClassVarDecStatement{
		Token:     token.Token{Type: token.KEYWORD, Literal: "static"},
		ValueType: token.Token{Type: token.KEYWORD, Literal: "int"},
		Identifiers: []*Identifier{
			&Identifier{
				Token: token.Token{Type: token.IDENTIFIER, Literal: "hogehoge"},
				Value: "hogehgoe",
			},
			&Identifier{
				Token: token.Token{Type: token.IDENTIFIER, Literal: "hogehoge"},
				Value: "hogehgoe",
			},
			&Identifier{
				Token: token.Token{Type: token.IDENTIFIER, Literal: "hogehoge"},
				Value: "hogehgoe",
			},
		},
	}
	t.Log(classVarDecStatement.String())
	t.Log(classVarDecStatement.Xml())
}

func TestSingleExpressionString(t *testing.T) {
	SingleExpression := &SingleExpression{
		Token: token.Token{Type: token.INTCONST, Literal: "4"},
		Value: &IntergerConstTerm{
			Token: token.Token{Type: token.INTCONST, Literal: "4"},
			Value: 4,
		},
	}
	t.Log(SingleExpression.String())
	t.Log(SingleExpression.Xml())
}

func TestKeywordConstTermString(t *testing.T) {
	SingleExpression := &SingleExpression{
		Token: token.Token{Type: token.KEYWORD, Literal: "true"},
		Value: &KeywordConstTerm{
			Token:   token.Token{Type: token.KEYWORD, Literal: "true"},
			KeyWord: token.TRUE,
		},
	}
	t.Log(SingleExpression.String())
	t.Log(SingleExpression.Xml())
}

func TestIfStatementString(t *testing.T) {
	varDecStatement := &VarDecStatement{
		Token:     token.Token{Type: token.KEYWORD, Literal: "var"},
		ValueType: token.Token{Type: token.KEYWORD, Literal: "int"},
		Identifiers: []*Identifier{
			&Identifier{
				Token: token.Token{Type: token.IDENTIFIER, Literal: "hogehoge"},
				Value: "hogehgoe",
			},
			&Identifier{
				Token: token.Token{Type: token.IDENTIFIER, Literal: "hogehoge"},
				Value: "hogehgoe",
			},
			&Identifier{
				Token: token.Token{Type: token.IDENTIFIER, Literal: "hogehoge"},
				Value: "hogehgoe",
			},
		},
	}

	ifStatement := &IfStatement{
		Token: token.Token{Type: token.KEYWORD, Literal: "if"},
		Condition: &SingleExpression{
			Token: token.Token{Type: token.INTCONST, Literal: "4"},
			Value: &IntergerConstTerm{
				Token: token.Token{Type: token.INTCONST, Literal: "4"},
				Value: 4,
			},
		},
		Consequence: &BlockStatement{
			Token: token.Token{Type: token.SYMBOL, Literal: "{"},
			Statements: []Statement{
				varDecStatement,
			},
		},
	}
	t.Log(ifStatement.String())
	t.Log(ifStatement.Xml())
}

func TestExpressionListStatementString(t *testing.T) {
	expressionListStatement := &ExpressionListStatement{
		Token: token.Token{Type: token.SYMBOL, Literal: "("},
		ExpressionList: []Expression{
			&SingleExpression{
				Token: token.Token{Type: token.INTCONST, Literal: "4"},
				Value: &IntergerConstTerm{
					Token: token.Token{Type: token.INTCONST, Literal: "4"},
					Value: 4,
				},
			},
			&SingleExpression{
				Token: token.Token{Type: token.INTCONST, Literal: "4"},
				Value: &IntergerConstTerm{
					Token: token.Token{Type: token.INTCONST, Literal: "4"},
					Value: 4,
				},
			},
		},
	}
	t.Log(expressionListStatement.String())
	t.Log(expressionListStatement.Xml())
}

func TestSubroutineCallTermString(t *testing.T) {
	expressionListStatement := &ExpressionListStatement{
		Token: token.Token{Type: token.SYMBOL, Literal: "("},
		ExpressionList: []Expression{
			&SingleExpression{
				Token: token.Token{Type: token.INTCONST, Literal: "4"},
				Value: &IntergerConstTerm{
					Token: token.Token{Type: token.INTCONST, Literal: "4"},
					Value: 4,
				},
			},
			&SingleExpression{
				Token: token.Token{Type: token.INTCONST, Literal: "4"},
				Value: &IntergerConstTerm{
					Token: token.Token{Type: token.INTCONST, Literal: "4"},
					Value: 4,
				},
			},
		},
	}
	subroutineCallTerm := &SubroutineCallTerm{
		Token:              token.Token{Type: token.IDENTIFIER, Literal: "hoge"},
		FunctionName:       "hoge",
		ExpressionListStmt: *expressionListStatement,
	}
	t.Log(subroutineCallTerm.String())
	t.Log(subroutineCallTerm.Xml())
}

func TestArrayElementTermString(t *testing.T) {
	arrayElementTerm := &ArrayElementTerm{
		Token:     token.Token{Type: token.IDENTIFIER, Literal: "hoge"},
		ArrayName: "hoge",
		Idx: &SingleExpression{
			Token: token.Token{Type: token.INTCONST, Literal: "4"},
			Value: &IntergerConstTerm{
				Token: token.Token{Type: token.INTCONST, Literal: "4"},
				Value: 4,
			},
		},
	}
	t.Log(arrayElementTerm.String())
	t.Log(arrayElementTerm.Xml())
}

func TestPrefixTermString(t *testing.T) {
	prefixTerm := &PrefixTerm{
		Token:  token.Token{Type: token.SYMBOL, Literal: "-"},
		Prefix: token.MINUS,
		Value: &IntergerConstTerm{
			Token: token.Token{Type: token.INTCONST, Literal: "4"},
			Value: 4,
		},
	}
	t.Log(prefixTerm.String())
	t.Log(prefixTerm.Xml())
}
