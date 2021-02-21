package code_writer

import (
	"VMtranslator/parser"
	"errors"
	"strconv"
	"strings"
)

// Flag is for distinguish lt,eq... instructions
var Flag = 0

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

// GetArithmetic returns assembly for arithmetic command
func GetArithmetic(commandStr string) (assembly string, err error) {
	s := strings.TrimSpace(commandStr)
	switch s {
	case "add":
		return "@SP\n" + "A=M\n" + "D=M\n" + "A=A-1\n" + "D=D+M\n" + "M=D\n" + "@SP\n" + "M=M-1\n", nil
	case "sub":
		return "@SP\n" + "A=M\n" + "D=M\n" + "A=A-1\n" + "M=M-D\n" + "@SP\n" + "M=M-1\n", nil
	case "neg":
		return "@SP\n" + "A=M\n" + "D=M\n" + "M=-D\n", nil
	case "and":
		return "@SP\n" + "A=M\n" + "D=M\n" + "A=A-1\n" + "D=D&M\n" + "M=D\n" + "@SP\n" + "M=M-1\n", nil
	case "or":
		return "@SP\n" + "A=M\n" + "D=M\n" + "A=A-1\n" + "D=D|M\n" + "M=D\n" + "@SP\n" + "M=M-1\n", nil
	case "not":
		return "@SP\n" + "A=M\n" + "D=M\n" + "M=!D\n", nil
	case "eq":
		return getCompareAssembly("JEQ"), nil
	case "gt":
		return getCompareAssembly("JGT"), nil
	case "lt":
		return getCompareAssembly("JLT"), nil
	default:
		return "", errors.New("invalid arithmetic command")
	}
}

func getCompareAssembly(assemblyCommand string) string {
	Flag++
	return "@SP\n" + "A=M-1\n" + "D=M\n" + "D=D-M\n" + "@SP\n" + "M=M-1\n" + "@TRUE" + strconv.Itoa(Flag) + "\n" + "D;" + assemblyCommand + "\n" + "@SP\n" + "A=M\n" + "M=0\n" + "@NEXT" + strconv.Itoa(Flag) + "\n" + "0;JMP\n" + "(TRUE" + strconv.Itoa(Flag) + ")\n" + "@SP\n" + "A=M\n" + "M=-1\n" + "(NEXT" + strconv.Itoa(Flag) + ")"
}
