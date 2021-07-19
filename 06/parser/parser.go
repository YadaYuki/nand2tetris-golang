package parser

import (
	"assembly/ast"
	"assembly/symboltable"
	"assembly/value"
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	*symboltable.SymbolTable
	input             string
	commandStrList    []string
	currentCommandIdx int
	readPosition      int
}

func New(input string, symbolTable *symboltable.SymbolTable) *Parser {
	parser := &Parser{input: input, commandStrList: strings.Split(input, value.NEW_LINE), currentCommandIdx: 0, readPosition: 0, SymbolTable: symbolTable}
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
	if commansStr == "" {
		return ""
	}
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
		commands = append(commands, command)
		p.Advance()
	}
	return commands, nil
}

func (p *Parser) ParseCommand() (ast.Command, error) {
	p.removeWhiteSpace()
	switch p.CommandType() {
	case ast.A_COMMAND:
		aCommand, err := p.parseACommand()
		p.resetReadPosition()
		if err != nil {
			return nil, err
		}
		return aCommand, nil
	case ast.C_COMMAND:
		cCommand, err := p.parseCCommand()
		p.resetReadPosition()
		if err != nil {
			return nil, err
		}
		return cCommand, nil
	case ast.L_COMMAND:
		lCommand, err := p.parseLCommand()
		p.resetReadPosition()
		if err != nil {
			return nil, err
		}
		return lCommand, nil
	default:
		return nil, fmt.Errorf("%s is invalid Command Type ", p.commandStrList[p.currentCommandIdx])
	}
}

func (p *Parser) parseACommand() (*ast.ACommand, error) {
	p.readChar() // read "@"
	valueStr := ""
	for p.hasMoreChar() {
		c := p.commandStrList[p.currentCommandIdx][p.readPosition]
		valueStr += string(c)
		p.readChar()
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		if p.Contains(valueStr) {
			value, _ = p.GetAddress(valueStr)
		}
	}
	return &ast.ACommand{ValueStr: valueStr, Value: value}, nil
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

	for c := p.commandStrList[p.currentCommandIdx][p.readPosition]; c != '='; c = p.commandStrList[p.currentCommandIdx][p.readPosition] {
		dest += string(c)
		p.readChar()
	}
	return dest
}

func (p *Parser) parseComp() string {
	comp := ""
	for p.hasMoreChar() {
		c := p.commandStrList[p.currentCommandIdx][p.readPosition]
		if c == ';' {
			break
		}
		if c == '/' {
			break
		}
		comp += string(c)
		p.readChar()
	}
	return comp
}

func (p *Parser) parseJump() string {
	jump := ""
	for p.hasMoreChar() {
		c := p.commandStrList[p.currentCommandIdx][p.readPosition]
		if !isLetter(c) && !isNumber(c) && !isUnderline(c) {
			break
		}
		jump += string(c)
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

func (p *Parser) Symbol() (string, error) {
	switch p.CommandType() {
	case ast.A_COMMAND:
		aCommand, _ := p.parseACommand()
		p.resetReadPosition()
		return aCommand.ValueStr, nil
	case ast.L_COMMAND:
		lCommand, _ := p.parseLCommand()
		p.resetReadPosition()
		return lCommand.Symbol, nil
	default:
		return "", fmt.Errorf("%s does not have Symbol ", p.CommandType())
	}
}

func (p *Parser) removeWhiteSpace() {
	p.commandStrList[p.currentCommandIdx] = strings.Replace(p.commandStrList[p.currentCommandIdx], string(value.SPACE), "", -1)
	p.commandStrList[p.currentCommandIdx] = strings.Replace(p.commandStrList[p.currentCommandIdx], string(value.TAB), "", -1)
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

func (p *Parser) ResetParseIdx() {
	p.resetReadPosition()
	p.currentCommandIdx = 0
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
