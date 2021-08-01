package parser

import (
	"VMtranslator/ast"
	"VMtranslator/value"
	"fmt"
	"strconv"
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
	case ast.LABEL:
		return ast.C_LABEL
	case ast.GOTO:
		return ast.C_GOTO
	case ast.IF_GOTO:
		return ast.C_IF
	case ast.CALL:
		return ast.C_CALL
	case ast.FUNCTION:
		return ast.C_FUNCTION
	case ast.RETURN:
		return ast.C_RETURN
	default:
		return ast.C_EMPTY
	}
}

func (p *Parser) Advance() {
	for {
		p.CurrentCommandIdx++
		if !p.HasMoreCommand() {
			break
		}
		p.CurrentTokenIdx = 0
		p.CurrentCommandTokenArr = strings.Split(p.CommandStrArr[p.CurrentCommandIdx], value.SPACE)
		if p.CommandType() != ast.C_EMPTY {
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

func (p *Parser) Arg2() (int, error) {
	switch p.CommandType() {
	case ast.C_PUSH, ast.C_POP, ast.C_FUNCTION, ast.C_CALL:
		arg2, err := strconv.Atoi(p.CurrentCommandTokenArr[2])
		if err != nil {
			return -1, err
		}
		return arg2, nil
	default:
		return -1, fmt.Errorf("%s cannnot call Arg2()", p.CommandType())
	}
}

func (p *Parser) ParsePush() (*ast.PushCommand, error) {
	arg1, err := p.Arg1()
	if err != nil {
		return nil, err
	}
	arg2, err := p.Arg2()
	if err != nil {
		return nil, err
	}
	command := &ast.PushCommand{Comamnd: ast.C_PUSH, Symbol: ast.PUSH, Segment: ast.SegmentType(arg1), Index: arg2}
	return command, nil
}

func (p *Parser) ParsePop() (*ast.PopCommand, error) {
	arg1, err := p.Arg1()
	if err != nil {
		return nil, err
	}
	arg2, err := p.Arg2()
	if err != nil {
		return nil, err
	}
	command := &ast.PopCommand{Comamnd: ast.C_POP, Symbol: ast.PUSH, Segment: ast.SegmentType(arg1), Index: arg2}
	return command, nil
}

func (p *Parser) ParseLabel() (*ast.LabelCommand, error) {
	arg1, err := p.Arg1()
	if err != nil {
		return nil, err
	}
	command := &ast.LabelCommand{Command: ast.C_LABEL, Symbol: ast.LABEL, LabelName: arg1}
	return command, nil
}

func (p *Parser) ParseGoto() (*ast.GotoCommand, error) {
	arg1, err := p.Arg1()
	if err != nil {
		return nil, err
	}
	command := &ast.GotoCommand{Command: ast.C_GOTO, Symbol: ast.GOTO, LabelName: arg1}
	return command, nil
}

func (p *Parser) ParseIf() (*ast.IfCommand, error) {
	arg1, err := p.Arg1()
	if err != nil {
		return nil, err
	}
	command := &ast.IfCommand{Command: ast.C_IF, Symbol: ast.IF_GOTO, LabelName: arg1}
	return command, nil
}
func (p *Parser) ParseCall() (*ast.CallCommand, error) {
	arg1, err := p.Arg1()
	if err != nil {
		return nil, err
	}
	arg2, err := p.Arg2()
	if err != nil {
		return nil, err
	}
	command := &ast.CallCommand{Command: ast.C_CALL, Symbol: ast.CALL, FunctionName: arg1, NumArgs: arg2}
	return command, nil
}

func (p *Parser) ParseFunction() (*ast.FunctionCommand, error) {
	arg1, err := p.Arg1()
	if err != nil {
		return nil, err
	}
	arg2, err := p.Arg2()
	if err != nil {
		return nil, err
	}
	command := &ast.FunctionCommand{Command: ast.C_CALL, Symbol: ast.CALL, FunctionName: arg1, NumLocals: arg2}
	return command, nil
}

func (p *Parser) ParseReturn() (*ast.ReturnCommand, error) {
	return &ast.ReturnCommand{Command: ast.C_RETURN, Symbol: ast.RETURN}, nil
}

func (p *Parser) ParseArithmetic() (*ast.ArithmeticCommand, error) {
	arg1, err := p.Arg1()
	if err != nil {
		return nil, err
	}
	command := &ast.ArithmeticCommand{Command: ast.C_ARITHMETIC, Symbol: ast.CommandSymbol(arg1)}
	return command, nil
}
