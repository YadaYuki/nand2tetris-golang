package parser

import (
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
