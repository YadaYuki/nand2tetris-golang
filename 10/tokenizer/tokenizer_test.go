package tokenizer

import (
	"jack/compiler/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	class Main {
    static boolean test;   
                           

    function void main() {
        var SquareGame game;
        let game = game;
        do game.run();
        do game.dispose();
        return;
    }
}`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.KEYWORD, "class"},
		{token.IDENTIFIER, "Main"},
		{token.SYMBOL, "{"},
		{token.KEYWORD, "static"},
		{token.KEYWORD, "boolean"},
		{token.IDENTIFIER, "test"},
		{token.SYMBOL, ";"},
		{token.KEYWORD, "function"},
		{token.KEYWORD, "void"},
		{token.IDENTIFIER, "main"},
		{token.SYMBOL, "("},
		{token.SYMBOL, ")"},
		{token.SYMBOL, "{"},
		{token.KEYWORD, "var"},
		{token.IDENTIFIER, "SquareGame"},
		{token.IDENTIFIER, "game"},
		{token.SYMBOL, ";"},
		{token.KEYWORD, "let"},
		{token.IDENTIFIER, "game"},
		{token.SYMBOL, "="},
		{token.IDENTIFIER, "game"},
		{token.SYMBOL, ";"},
		{token.KEYWORD, "do"},
		{token.IDENTIFIER, "game"},
		{token.SYMBOL, "."},
		{token.IDENTIFIER, "run"},
		{token.SYMBOL, "("},
		{token.SYMBOL, ")"},
		{token.SYMBOL, ";"},
		{token.KEYWORD, "do"},
		{token.IDENTIFIER, "game"},
		{token.SYMBOL, "."},
		{token.IDENTIFIER, "dispose"},
		{token.SYMBOL, "("},
		{token.SYMBOL, ")"},
		{token.SYMBOL, ";"},
		{token.KEYWORD, "return"},
		{token.SYMBOL, ";"},
		{token.SYMBOL, "}"},
		{token.SYMBOL, "}"},
	}
	jt :=  New(input)
	for i,tt := range tests {
		tok,_ := jt.Advance()
		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q,got %q. -tokenliteral : expected=%q,got %q",i,tt.expectedType,tok.Type,)
			
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - tokenliteral wrong. expected=%q,got %q",i,tt.expectedLiteral,tok.Literal)
		}
	}
}
