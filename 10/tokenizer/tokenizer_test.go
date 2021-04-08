package tokenizer

import (
	"jack/compiler/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.SYMBOL, "="},
		{token.SYMBOL, "+"},
		{token.SYMBOL, "("},
		{token.SYMBOL, ")"},
		{token.SYMBOL, "{"},
		{token.SYMBOL, "}"},
		{token.SYMBOL, ","},
		{token.SYMBOL, ";"},
		{token.EOF, ""},
	}
	jt :=  New(input)
	for i,tt := range tests {
		tok,_ := jt.Advance()
		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q,got %q",i,tt.expectedType,tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q,got %q",i,tt.expectedLiteral,tok.Literal)
		}
	}
}
