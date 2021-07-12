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
	value int
}

func (aCommand *ACommand) String() string {
	return fmt.Sprintf("@%d", aCommand.value) + value.NEW_LINE
}

type CCommand struct {
	comp string
	dest string
	jump string
}

func (cCommand *CCommand) String() string {
	commandStr := fmt.Sprintf("%s=%s;%s", cCommand.dest, cCommand.comp, cCommand.jump) + value.NEW_LINE
	if cCommand.jump == "" {
		commandStr = strings.Replace(commandStr, ";", "", 1)
	} else if cCommand.dest == "" {
		commandStr = strings.Replace(commandStr, "=", "", 1)
	}
	return commandStr
}
