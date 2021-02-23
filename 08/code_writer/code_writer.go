package code_writer

import (
	"VMtranslator/parser"
	"errors"
	"strconv"
	"strings"
)

// Flag is for distinguish lt,eq... instructions
var flag = 0

// GetPushPop get
func GetPushPop(commandType parser.CommandType, segment string, index int) (assembly string, err error) {

	if commandType == parser.CPush {
		switch segment {
		case "constant":
			return "@" + strconv.Itoa(index) + "\n" + "D=A\n" + "@SP\n" + "A=M\n" + "M=D" + "@SP\n" + "M=M+1\n", nil
		case "local":
			return getPushSegmentAssembly("LCL", index), nil
		case "argument":
			return getPushSegmentAssembly("ARG", index), nil
		case "this":
			return getPushSegmentAssembly("THIS", index), nil
		case "that":
			return getPushSegmentAssembly("THAT", index), nil
		case "pointer":
			return getPushPointerAssembly(index), nil
		case "temp":
			return getPushTempAssembly(index), nil
		case "static":
			return getPushStaticAssembly(index), nil
		}

	}
	// TODO: add pop command
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

// GetWriteInit return vm initialize assembly
func GetWriteInit() (assembly string) {
	// set SP 256
	return "@256\n" + "D=A\n" + "@SP" + "M=D\n"
}

// GetWriteLabel convert vm "label" to assembly "label"
func GetWriteLabel(label string) (assembly string) {
	return "(" + label + ")\n"
}

// GetWriteGoto convert vm "goto" to assembly
func GetWriteGoto(label string) (assembly string) {
	return "@" + label + "\n" + "0;JMP\n"
}

// GetWriteIf convert vm "if-goto" to assembly
func GetWriteIf(label string) (assembly string) {
	return "@SP\n" + "A=M\n" + "D=M\n" + "@SP\n" + "M=M-1\n" + "@" + label + "\n" + "D;JNE\n"
}

// sub module

var setDtoStackAssembly = "@SP\n" + "A=M\n" + "M=D" + "@SP\n" + "M=M+1\n"

func getCompareAssembly(assemblyCommand string) string {
	flag++
	return "@SP\n" + "A=M-1\n" + "D=M\n" + "D=D-M\n" + "@SP\n" + "M=M-1\n" + "@TRUE" + strconv.Itoa(flag) + "\n" + "D;" + assemblyCommand + "\n" + "@SP\n" + "A=M\n" + "M=0\n" + "@NEXT" + strconv.Itoa(flag) + "\n" + "0;JMP\n" + "(TRUE" + strconv.Itoa(flag) + ")\n" + "@SP\n" + "A=M\n" + "M=-1\n" + "(NEXT" + strconv.Itoa(flag) + ")"
}

func getPushSegmentAssembly(assemblyCommand string, index int) string {
	return "@" + strconv.Itoa(index) + "\n" + "D=A\n" + "@" + assemblyCommand + "\n" + "A=M+D\n" + "D=M\n" +
		setDtoStackAssembly
}

func getPushTempAssembly(index int) string {
	return "@" + strconv.Itoa(index) + "\n" + "D=A\n" + "@THAT\n" + "A=A+D\n" + "A=A+1\n" + "D=M\n" +
		setDtoStackAssembly
}

func getPushPointerAssembly(index int) string {
	return "@" + strconv.Itoa(index) + "\n" + "D=A\n" + "@THIS\n" + "A=A+D\n" + "D=M\n" + // set pointer value to D
		setDtoStackAssembly
}

func getPushStaticAssembly(index int) string {
	StaticHeadAddress := "16"
	return "@" + strconv.Itoa(index) + "\n" + "D=A\n" + "@" + StaticHeadAddress + "\n" + "A=D+A\n" + "D=M\n" + setDtoStackAssembly
}
