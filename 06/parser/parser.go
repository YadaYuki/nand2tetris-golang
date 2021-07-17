package parser

import (
	"assembly/ast"
	"assembly/value"
	"fmt"
	"strconv"
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

func (p *Parser) CommandType() ast.CommandType {
	commansStr := p.commandStrList[p.currentCommandIdx]
	// switch by prefix char
	commandStrPrefix := commansStr[0]
	switch commandStrPrefix {
	case '@':
		return ast.A_COMMAND
	case '0', '1', 'D', 'A', '!', '-', 'M':
		return ast.C_COMMAND
	case '(':
		return ast.L_COMMAND
	default:
		return ""
	}
}

func (p *Parser) ParseAssembly() ([]ast.Command, error) {
	commands := []ast.Command{}
	for p.HasMoreCommand() {
		command, _ := p.ParseCommand()
		p.resetReadPosition()
		commands = append(commands, command)
		p.Advance()
	}
	return commands, nil
}

func (p *Parser) ParseCommand() (ast.Command, error) {
	switch p.CommandType() {
	case ast.A_COMMAND:
		aCommand, err := p.parseACommand()
		if err != nil {
			return nil, err
		}
		return aCommand, nil
	case ast.C_COMMAND:
		cCommand, err := p.parseCCommand()
		if err != nil {
			return nil, err
		}
		return cCommand, nil
	default:
		return nil, fmt.Errorf("%s is invalid Command Type ", p.commandStrList[p.currentCommandIdx])
	}
}

func (p *Parser) parseACommand() (*ast.ACommand, error) {
	p.readChar() // read "@"
	valueStr := ""
	for p.hasMoreChar() {
		valueStr += string(p.commandStrList[p.currentCommandIdx][p.readPosition])
		p.readChar()
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return nil, err
	}
	return &ast.ACommand{Value: value}, nil
}

func (p *Parser) parseCCommand() (*ast.CCommand, error) {
	dest, comp, jump := "", "", ""
	hasEqual := strings.Contains(p.commandStrList[p.currentCommandIdx], "=")
	hasSemicolon := strings.Contains(p.commandStrList[p.currentCommandIdx], ";")
	if hasEqual && hasSemicolon { // parseDest
		dest = p.parseDest()
		p.readChar() // read "="
		comp = p.parseComp()
		p.readChar() // read ";"
		jump = p.parseJump()
	} else if hasEqual {
		dest = p.parseDest() // read "="
		p.readChar()
		comp = p.parseComp()
	} else if hasSemicolon {
		comp = p.parseComp()
		p.readChar() // read ";"
		jump = p.parseJump()
	}
	return &ast.CCommand{Dest: dest, Comp: comp, Jump: jump}, nil
}

func (p *Parser) parseDest() string {
	dest := ""
	for p.commandStrList[p.currentCommandIdx][p.readPosition] != '=' {
		dest += string(p.commandStrList[p.currentCommandIdx][p.readPosition])
		p.readChar()
	}

	return dest
}

func (p *Parser) parseComp() string {
	comp := ""
	for p.hasMoreChar() {
		if p.commandStrList[p.currentCommandIdx][p.readPosition] == ';' {
			break
		}
		comp += string(p.commandStrList[p.currentCommandIdx][p.readPosition])
		p.readChar()
	}
	return comp
}

func (p *Parser) parseJump() string {
	jump := ""
	for p.hasMoreChar() {
		jump += string(p.commandStrList[p.currentCommandIdx][p.readPosition])
		p.readChar()
	}
	return jump
}

func (p *Parser) parseLCommand() (*ast.LCommand, error) {
	p.readChar() // read '('
	valueStr := ""
	for p.commandStrList[p.currentCommandIdx][p.readPosition] != ')' {
		valueStr += string(p.commandStrList[p.currentCommandIdx][p.readPosition])
		p.readChar()
	}
	return &ast.LCommand{Symbol: valueStr}, nil
}

func (p *Parser) skipWhiteSpace() {
	for p.hasMoreChar() && (p.commandStrList[p.currentCommandIdx][p.readPosition] == value.SPACE || p.commandStrList[p.currentCommandIdx][p.readPosition] == value.TAB) {
		p.readChar()
	}
}

func (p *Parser) readChar() {
	p.readPosition++
}

func (p *Parser) hasMoreChar() bool {
	return len(p.commandStrList[p.currentCommandIdx]) > p.readPosition
}
func (p *Parser) resetReadPosition() {
	p.readPosition = 0
}
