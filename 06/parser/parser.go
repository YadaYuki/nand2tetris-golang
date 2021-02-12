package parser

import (
	"errors"
	"fmt"
	"strings"
)

// CommandType is type of command
type CommandType int

const (
	aCommand CommandType = iota
	cCommand
	lCommand
)

func (command CommandType) String() string {
	switch command {
	case aCommand:
		return "A_COMMAND"
	case cCommand:
		return "C_COMMAND"
	case lCommand:
		return "L_COMMAND"
	default:
		return "Unknown"
	}
}

// GetCommandType is function get command type from command
func GetCommandType(commandStr string) (c CommandType, err error) {
	s := strings.TrimSpace(commandStr)

	if s[0:1] == "@" {
		return aCommand, nil
	}
	if strings.LastIndexAny(s, "(") == 0 && strings.LastIndexAny(s, ")") == len(s)-1 {
		return lCommand, nil
	}

	// TODO: clarify syntax of C_COMMAND
	if strings.Contains(s, ";") || strings.Contains(s, "=") {
		return cCommand, nil
	}

	return lCommand, errors.New("Invalid CommandType")
}

// GetSymbol returns Symbol name
func GetSymbol(commandStr string) (symbol string, err error) {
	s := strings.TrimSpace(commandStr)
	commandType, err := GetCommandType(s)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if commandType == aCommand {
		return s[1:], nil
	}
	if commandType == lCommand {
		return s[1 : len(s)-1], nil
	}
	return "", err
}
