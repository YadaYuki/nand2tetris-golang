package parser

import "strings"

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

	return cCommand, nil
}
