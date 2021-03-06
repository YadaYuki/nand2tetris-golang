package jack_tokenizer

import "fmt"

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

func (jackTokenizer *JackTokenizer) HasMoreTokens() {
	fmt.Print(jackTokenizer.input)
}
