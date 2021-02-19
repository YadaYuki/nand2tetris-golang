package code_writer

import (
	"VMtranslator/parser"
	"errors"
	"strconv"
)

// GetPushPop get
func GetPushPop(commandType parser.CommandType, segment string, index int) (assembly string, err error) {

	if commandType == parser.CPush {
		if segment == "constant" {
			return "@" + strconv.Itoa(index) + "\n" + "D=A\n" + "@SP\n" + "A=M\n" + "M=D", nil
		}
	}
	if commandType == parser.CPop {
		return "", nil
	}
	return "", errors.New("Invalid Command Type")
}
