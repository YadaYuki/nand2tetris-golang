package tokenizer

import (
	"errors"
	"jack/compiler/token"
)

// JackTokenizer has member necessary for parsing
type JackTokenizer struct {
	input        string // code input
	position     int
	readPosition int
	ch           byte
}

// New is initializer of jack tokenizer
func New(input string) *JackTokenizer {
	jt := &JackTokenizer{input: input, ch: input[0], readPosition: 0, position: 0}
	return jt
}

// HasMoreTokens returns whether hasMoreToken
func (jackTokenizer *JackTokenizer) HasMoreTokens() bool {
	return len(jackTokenizer.input) > jackTokenizer.readPosition
}

// Advance returns next token
func (jackTokenizer *JackTokenizer) Advance() (advanceToken token.Token, err error) {
	var tok token.Token
	jackTokenizer.skipWhitespace()
	if jackTokenizer.HasMoreTokens() == false {
		return token.Token{Type: token.EOF, Literal: ""}, nil
	}
	if _, ok := token.SymbolMap[jackTokenizer.ch]; ok {
		tok = token.Token{Type: token.SYMBOL, Literal: string(jackTokenizer.ch)}
	} else if isLetter(jackTokenizer.ch) { // KEYWORD or IDENTIFIER
		word := jackTokenizer.readWord()
		if _, ok := token.KeyWordMap[word]; ok {
			tok = token.Token{Type: token.KEYWORD, Literal: word}
		} else {
			tok = token.Token{Type: token.IDENTIFIER, Literal: word}
		}
	} else if isNumber(jackTokenizer.ch) {
		word := jackTokenizer.readNumber()
		tok = token.Token{Type: token.INTCONST, Literal: word}
	} else if isSingleQuote(jackTokenizer.ch) {
		word := jackTokenizer.readString()
		tok = token.Token{Type: token.STARTINGCONST, Literal: word[1:]}
	} else {
		return tok, errors.New("invalide ch")
	}
	jackTokenizer.readChar()
	return tok, nil
}

// KeyWord returns keyword type
func KeyWord(tok token.Token) (keyword token.KeyWord, err error) {
	if tok.Type != token.KEYWORD {
		return token.NULL, errors.New("KeyWord Function can call only token type is KEYWORD")
	}
	return token.KeyWordMap[tok.Literal], nil
}

func (jackTokenizer *JackTokenizer) readChar() {
	if jackTokenizer.readPosition >= len(jackTokenizer.input) {
		jackTokenizer.ch = 0
	} else {
		jackTokenizer.ch = jackTokenizer.input[jackTokenizer.readPosition]
	}
	jackTokenizer.position = jackTokenizer.readPosition
	jackTokenizer.readPosition++
}

func (jackTokenizer *JackTokenizer) readWord() string {
	position := jackTokenizer.position
	for isLetter(jackTokenizer.ch) || isNumber(jackTokenizer.ch) || isUnderline(jackTokenizer.ch) {
		jackTokenizer.readChar()
	}
	return jackTokenizer.input[position:jackTokenizer.position]
}

func (jackTokenizer *JackTokenizer) readNumber() string {
	position := jackTokenizer.position
	for isNumber(jackTokenizer.ch) {
		jackTokenizer.readChar()
	}
	return jackTokenizer.input[position:jackTokenizer.position]
}

func (jackTokenizer *JackTokenizer) readString() string {
	position := jackTokenizer.position
	for {
		jackTokenizer.readChar()
		if isSingleQuote(jackTokenizer.ch) {
			break
		}
	}
	return jackTokenizer.input[position:jackTokenizer.position]
}

func (jackTokenizer *JackTokenizer) skipWhitespace() {
	for jackTokenizer.ch == ' ' || jackTokenizer.ch == '\t' || jackTokenizer.ch == '\n' || jackTokenizer.ch == '\r' {
		jackTokenizer.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isUnderline(ch byte) bool {
	return ch == '_'
}

func isSingleQuote(ch byte) bool {
	return ch == '\''
}
