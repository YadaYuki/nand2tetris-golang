package parser

import (
	"errors"
	"strings"
)

type CommandType int

const (
	CArithmetic CommandType = iota
	CPush
	CPop
	CLabel
	CGoto
	CIf
	CFunction
	CReturn
	CCall
)

func (command CommandType) String() string {
	switch command {
	case CArithmetic:
		return "C_ARITHMETIC"
	case CPush:
		return "C_PUSH"
	case CPop:
		return "C_POP"
	case CLabel:
		return "C_LABEL"
	case CGoto:
		return "C_GOTO"
	case CIf:
		return "C_IF"
	case CFunction:
		return "C_FUNCTION"
	case CReturn:
		return "C_RETURN"
	case CCall:
		return "C_CALL"
	default:
		return "Unknown"
	}
}

func GetCommandType(commandStr string) (c CommandType, err error) {
	s := strings.TrimSpace(commandStr)
	arithmeticCommand := map[string]int{"add": 1, "sub": 1, "neg": 1, "eq": 1, "gt": 1, "lt": 1, "and": 1, "or": 1, "not": 1}
	if _, ok := arithmeticCommand[s]; ok {
		return CArithmetic, nil
	}
	return CCall, errors.New("Invalid CommandType")
}
