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
	singleExpression := &SingleExpression{
		Token: token.Token{Type: token.INTCONST, Literal: "4"},
		Value: &IntergerConstTerm{
			Token: token.Token{Type: token.INTCONST, Literal: "4"},
			Value: 4,
		},
	}
	t.Log(singleExpression.String())
	t.Log(singleExpression.Xml())
}
