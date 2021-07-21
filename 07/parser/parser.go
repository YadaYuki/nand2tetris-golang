package parser

import (
	"VMtranslator/ast"
	"VMtranslator/value"
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
