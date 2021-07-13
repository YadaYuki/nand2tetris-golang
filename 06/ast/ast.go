package ast

import (
	"assembly/value"
	"fmt"
	"strings"
)

type Command interface {
	String() string
}

type ACommand struct {
	Value int
}

func (aCommand *ACommand) String() string {
	return fmt.Sprintf("@%d", aCommand.Value) + value.NEW_LINE
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
