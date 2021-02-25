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
	return "@256\n" + "D=A\n" + "@SP\n" + "M=D\n"
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

var setDtoStackAssembly = "@SP\n" + "A=M\n" + "M=D\n" + "@SP\n" + "M=M+1\n"
var popStackToDAssembly = "@SP\n" + "A=M\n" + "D=M\n" + "M=M-1\n"

// GetWriteCall convert vm "call (functionName) (numArgs)" to assembly
func GetWriteCall(functionName string, numArgs int) (assembly string) {
	flag++
	returnAddress := "RETURN" + strconv.Itoa(flag)
	return "@" + returnAddress + "\n" + "D=A\n" + setDtoStackAssembly + // push return-address
		"@LCL\n" + "A=M\n" + "D=M\n" + setDtoStackAssembly + //  push LCL
		"@ARG\n" + "A=M\n" + "D=M\n" + setDtoStackAssembly + // push ARG
		"@THIS\n" + "A=M\n" + "D=M\n" + setDtoStackAssembly + // push THIS
		"@THAT\n" + "A=M\n" + "D=M\n" + setDtoStackAssembly + // push THAT
		"@" + strconv.Itoa(numArgs) + "\n" + "D=A\n" + "@SP\n" + "A=M\n" + "D=M-D\n" + "@5\n" + "D=D-A\n" + "@ARG\n" + "M=D\n" + // ARG=SP-n-5
		"@SP\n" + "A=M\n" + "D=M\n" + "@LCL\n" + "A=M\n" + "M=D\n" + // LCL = SP
		"@" + functionName + "\n" + "0;JMP\n" + // goto function
		"(" + returnAddress + ")\n"
}

// GetWriteFunction convert vm "function (functionName) (numLocals)" to assembly
func GetWriteFunction(functionName string, numLocals int) (assembly string) {
	// initialize local variable by 0
	var assemblyByte []byte
	assemblyByte = append(assemblyByte, "("+functionName+")\n"...)
	pushZeroToStackAssembly := "@0\n" + "D=A\n" + setDtoStackAssembly
	for i := 0; i < numLocals; i++ {
		assemblyByte = append(assemblyByte, pushZeroToStackAssembly...)
	}
	return string(assemblyByte)
}

// GetWriteReturn convert vm "return" to assembly
func GetWriteReturn() (assembly string) {
	return "@LCL\n" + "A=M\n" + "D=M\n" + "@FRAME\n" + "M=D\n" + // FRAME=LCL
		"@5\n" + "D=A\n" + "@FRAME\n" + "D=M-D\n" + "A=D\n" + "D=M\n" + "@RET\n" + "M=D\n" + // RET=*(FRAME-5)
		popStackToDAssembly + "@ARG\n" + "A=M\n" + "A=M\n" + "M=D\n" + // *ARG=pop()
		"@ARG\n" + "A=M\n" + "D=M+1\n" + "@SP\n" + "A=M\n" + "M=D\n" + // SP=ARG+1
		"@1\n" + "D=A\n" + "@FRAME\n" + "A=M\n" + "A=M-D\n" + "A=M\n" + "D=M\n" + "@THAT\n" + "A=M\n" + "M=D\n" + // THAT=*(FRAME-1)
		"@2\n" + "D=A\n" + "@FRAME\n" + "A=M\n" + "A=M-D\n" + "A=M\n" + "D=M\n" + "@THIS\n" + "A=M\n" + "M=D\n" + // THIS=*(FRAME-2)
		"@3\n" + "D=A\n" + "@FRAME\n" + "A=M\n" + "A=M-D\n" + "A=M\n" + "D=M\n" + "@ARG\n" + "A=M\n" + "M=D\n" + // ARG=*(FRAME-3)
		"@4\n" + "D=A\n" + "@FRAME\n" + "A=M\n" + "A=M-D\n" + "A=M\n" + "D=M\n" + "@LCL\n" + "A=M\n" + "M=D\n" + // LCL=*(FRAME-4)
		"@RET\n" + "0;JMP\n" + // goto RET
}

// sub module

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
