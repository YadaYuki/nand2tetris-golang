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

// GetDest returns machine language　Correspond to dest label
func GetDest(commandStr string) (symbol string, err error) {
	s := strings.TrimSpace(commandStr)
	commandType, err := GetCommandType(s)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if commandType != cCommand {
		return "", errors.New("only C_COMMAND has dest label")
	}
	if strings.Contains(s, "=") == false {
		return "null", nil
	}
	dest := strings.Split(s, "=")[0]
	return dest, nil
}

// GetJump returns machine language　Correspond to dest label
func GetJump(commandStr string) (symbol string, err error) {
	s := strings.TrimSpace(commandStr)
	commandType, err := GetCommandType(s)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if commandType != cCommand {
		return "", errors.New("only C_COMMAND has jump label")
	}
	if strings.Contains(s, ";") == false {
		return "null", nil
	}
	jump := strings.Split(s, ";")[1]
	return jump, nil
}

// GetComp returns machine language　Correspond to dest label
func GetComp(commandStr string) (symbol string, err error) {
	s := strings.TrimSpace(commandStr)
	commandType, err := GetCommandType(s)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if commandType != cCommand {
		return "", errors.New("only C_COMMAND has jump label")
	}
	jump, err := GetJump(s)
	if jump == "null" {
		comp := strings.Split(s, "=")[1]
		return comp, nil
	}
	dest, err := GetDest(s)
	if dest == "null" {
		comp := strings.Split(s, ";")[0]
		return comp, nil
	}
	comp := strings.Split((strings.Split(s, "=")[1]), ";")[0]
	return comp, nil
}
