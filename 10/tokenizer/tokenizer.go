package tokenizer

import (
	"errors"
	"jack_compiler/token"
)

// JackTokenizer has member necessary for parsing
type JackTokenizer struct {
	input        string // code input
	position     int
	readPosition int
	ch           byte
}

var symbolMap = map[byte]bool{'{': true, '}': true, '(': true, ')': true, '[': true, ']': true, '.': true, ':': true, ',': true, ';': true, '+': true, '-': true, '*': true, '/': true, '&': true, '|': true, '<': true, '>': true, '=': true, '~': true}
var keywordMap = map[string]bool{"class": true, "constructor": true, "function": true, "method": true, "field": true, "static": true, "var": true, "int": true, "char": true, "boolean": true, "void": true, "true": true, "false": true, "null": true, "this": true, "let": true, "do": true, "if": true, "else": true, "while": true, "return": true}

// New is initializer of jack tokenizer
func New(input string) *JackTokenizer {
	jt := &JackTokenizer{input: input, ch: input[0], readPosition: 0, position: 0}
	return jt
}

// HasMoreTokens returns whether hasMoreToken
func (jackTokenizer *JackTokenizer) HasMoreTokens() bool {
	return len(jackTokenizer.input) != jackTokenizer.readPosition
}

// Advance returns next token
func (jackTokenizer *JackTokenizer) Advance() (advanceToken token.Token, err error) {
	var tok token.Token
	jackTokenizer.skipWhitespace()
	if jackTokenizer.HasMoreTokens() == false {
		return token.Token{}, errors.New("jackTokenizer has no more token")
	}
	if _, ok := symbolMap[jackTokenizer.ch]; ok {
		tok = token.Token{Type: token.SYMBOL, Literal: string(jackTokenizer.ch)}
	} else if isLetter(jackTokenizer.ch) { // KEYWORD or IDENTI
		word := jackTokenizer.readWord()
		if _, ok := keywordMap[word]; ok {
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

func (jackTokenizer *JackTokenizer) readChar() {
	jackTokenizer.ch = jackTokenizer.input[jackTokenizer.readPosition]
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
