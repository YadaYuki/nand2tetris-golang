package parser

import (
	"VMtranslator/ast"
	"VMtranslator/value"
	"fmt"
	"strings"
)

type Parser struct {
	CurrentCommandIdx      int
	CurrentTokenIdx        int
	CommandStrArr          []string // ["push local 1",....]
	CurrentCommandTokenArr []string // ["push","local","1"]
	input                  string
}

func New(input string) *Parser {
	CommandStrArr := strings.Split(input, value.NEW_LINE)
	InitialCurrentCommandTokenArr := strings.Split(CommandStrArr[0], value.SPACE)
	return &Parser{input: input, CurrentCommandIdx: 0, CurrentTokenIdx: 0, CommandStrArr: CommandStrArr, CurrentCommandTokenArr: InitialCurrentCommandTokenArr}
}

func (p *Parser) HasMoreCommand() bool {
	return len(p.CommandStrArr) > p.CurrentCommandIdx
}

func (p *Parser) CommandType() ast.CommandType {
	if p.CommandStrArr[p.CurrentCommandIdx] == "" {
		return ast.C_EMPTY
	}
	curretnCommandPrefix := ast.CommandSymbol(p.CurrentCommandTokenArr[0])
	switch curretnCommandPrefix {
	case ast.PUSH:
		return ast.C_PUSH
	case ast.POP:
		return ast.C_POP
	case ast.ADD, ast.SUB, ast.NEG, ast.EQ, ast.GT, ast.LT, ast.AND, ast.OR, ast.NOT:
		return ast.C_ARITHMETIC
	default:
		return ast.C_EMPTY
	}
}

func (p *Parser) Advance() {
	for {
		p.CurrentCommandIdx++
		p.CurrentTokenIdx = 0
		p.CurrentCommandTokenArr = strings.Split(p.CommandStrArr[p.CurrentCommandIdx], value.SPACE)
		if p.CommandType() != ast.C_EMPTY || !p.HasMoreCommand() {
			break
		}
	}
}

func (p *Parser) Arg1() (string, error) {
	switch p.CommandType() {
	case ast.C_ARITHMETIC:
		return p.CurrentCommandTokenArr[0], nil // return arithmetic symbol. "add","eq"...
	case ast.C_PUSH, ast.C_POP, ast.C_LABEL, ast.C_GOTO, ast.C_IF, ast.C_FUNCTION, ast.C_CALL:
		return p.CurrentCommandTokenArr[1], nil
	default:
		return "", fmt.Errorf("%s cannnot call Arg1()", p.CommandType())
	}
}
