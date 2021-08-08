package ast

import (
	"assembler/value"
	"fmt"
	"strings"
)

type CommandType string

const (
	A_COMMAND CommandType = "A_COMMAND"
	C_COMMAND CommandType = "C_COMMAND"
	L_COMMAND CommandType = "L_COMMAND"
)

type Command interface {
	String() string
}

type ACommand struct {
	Value    int
	ValueStr string
}

func (aCommand *ACommand) String() string {
	return fmt.Sprintf("@%s", aCommand.ValueStr) + value.NEW_LINE
}

type CCommand struct {
	Comp string
	Dest string
	Jump string
}

func (cCommand *CCommand) String() string {
	commandStr := fmt.Sprintf("%s=%s;%s", cCommand.Dest, cCommand.Comp, cCommand.Jump) + value.NEW_LINE
	if cCommand.Jump == "" {
		commandStr = strings.Replace(commandStr, ";", "", 1)
	} else if cCommand.Dest == "" {
		commandStr = strings.Replace(commandStr, "=", "", 1)
	}
	return commandStr
}

type LCommand struct {
	Symbol string
}

func (lCommand *LCommand) String() string {
	return fmt.Sprintf("(%s)", lCommand.Symbol)
}
