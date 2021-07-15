package parser

import (
	"assembly/value"
	"strings"
)

type Parser struct {
	input             string
	commandStrList    []string
	currentCommandIdx int
	readPosition      int
}

func New(input string) *Parser {
	parser := &Parser{input: input, commandStrList: strings.Split(input, value.NEW_LINE), currentCommandIdx: 0, readPosition: 0}
	return parser
}

func (p *Parser) Advance() {
	p.currentCommandIdx++
}

func (p *Parser) HasMoreCommand() bool {
	return len(p.commandStrList) > p.currentCommandIdx
}
func (p *Parser) skipWhiteSpace() {
	for p.commandStrList[p.currentCommandIdx][p.readPosition] == value.SPACE || p.commandStrList[p.currentCommandIdx][p.readPosition] == value.TAB {
		p.readPosition++
	}
}
