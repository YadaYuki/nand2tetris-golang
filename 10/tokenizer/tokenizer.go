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

// New is initializer of jack tokenizer
func New(input string) *JackTokenizer {
	jt := &JackTokenizer{input: input}
	return jt
}

// HasMoreTokens returns whether hasMoreToken
func (jackTokenizer *JackTokenizer) HasMoreTokens() bool {
	return len(jackTokenizer.input) != jackTokenizer.readPosition
}

// Advance returns next token
func (jackTokenizer *JackTokenizer) Advance() (advanceToken token.Token, err error) {
	symbolMap := map[byte]bool{'{': true, '}': true, '(': true, ')': true, '[': true, ']': true, '.': true, ':': true, ',': true, ';': true, '+': true, '-': true, '*': true, '/': true, '&': true, '|': true, '<': true, '>': true, '=': true, '~': true}
	if jackTokenizer.HasMoreTokens() == false {
		return token.Token{}, errors.New("jackTokenizer has no more token")
	}
	if _, ok := symbolMap[jackTokenizer.ch]; ok {
		return token.Token{Type: token.SYMBOL, Literal: string(jackTokenizer.ch)}, nil
	}
	jackTokenizer.readChar()
	return token.Token{Type: token.KEYWORD, Literal: "HGOE"}, errors.New("hgoehgoe")
}

func (jackTokenizer *JackTokenizer) readChar() {
	jackTokenizer.ch = jackTokenizer.input[jackTokenizer.readPosition]
	jackTokenizer.position = jackTokenizer.readPosition
	jackTokenizer.readPosition++
}
