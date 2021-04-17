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

func TestXml(t *testing.T) {
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
	t.Log(letStatement.Xml())
}
