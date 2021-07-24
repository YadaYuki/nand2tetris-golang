package codewriter

import (
	"VMtranslator/ast"
	"VMtranslator/value"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
)

type CodeWriter struct {
	Filename string
	Assembly []byte
}

func New(filename string) *CodeWriter {
	return &CodeWriter{Filename: filename, Assembly: []byte{}}
}

func (codeWriter *CodeWriter) Close() {
	err := ioutil.WriteFile(codeWriter.Filename, codeWriter.Assembly, 0644)
	if err != nil {
		panic(err)
	}
}

func (codeWriter *CodeWriter) WritePushPop(command ast.MemoryAccessCommand) error {
	var assembly string
	switch c := command.(type) {
	case *ast.PushCommand:
		pushAssembly, err := getPushAssembly(c)
		if err != nil {
			return err
		}
		assembly = pushAssembly
	}
	codeWriter.writeAssembly(assembly)
	return nil
}

func (codeWriter *CodeWriter) WriteArithmetic(command *ast.ArithmeticCommand) error {
	arithmeticAssembly, err := getArithmeticAssembly(command)
	if err != nil {
		return err
	}
	codeWriter.writeAssembly(arithmeticAssembly)
	return nil
}

func getArithmeticAssembly(arithmeticCommand *ast.ArithmeticCommand) (string, error) {
	switch arithmeticCommand.Symbol {
	case ast.ADD:
		return getAddCommandAssembly(), nil
	case ast.SUB:
		return getSubCommandAssembly(), nil
	case ast.NEG:
		return getNegCommandAssembly(), nil
	case ast.NOT:
		return getNotCommandAssembly(), nil
	case ast.AND:
		return getAndCommandAssembly(), nil
	case ast.OR:
		return getOrCommandAssembly(), nil
	case ast.GT, ast.LT, ast.EQ:
		return getCompareAssembly(arithmeticCommand.Symbol), nil
	}
	return "", fmt.Errorf("%T couldn't convert to arithmeticAssembly", arithmeticCommand)
}

func getAddCommandAssembly() string {
	assembly := ""
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE   // read value which SP points to into M
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-1 points to into M
	assembly += "D=M" + value.NEW_LINE                            // set M to D
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-2 points to into M
	assembly += "M=M+D" + value.NEW_LINE                          // set RAM[SP-2] = RAM[SP-2] + RAM[SP-1]
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE // decrement SP
	return assembly
}

func getSubCommandAssembly() string {
	assembly := ""
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE   // read value which SP points to into M
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-1 points to into M
	assembly += "D=M" + value.NEW_LINE                            // set M to D
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-2 points to into M
	assembly += "M=M-D" + value.NEW_LINE                          // set RAM[SP-2] = RAM[SP-2] - RAM[SP-1]
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE // decrement SP
	return assembly
}

func getAndCommandAssembly() string {
	assembly := ""
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE   // read value which SP points to into M
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-1 points to into M
	assembly += "D=M" + value.NEW_LINE                            // set M to D
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-2 points to into M
	assembly += "M=M&D" + value.NEW_LINE                          // set RAM[SP-2] = RAM[SP-2] and RAM[SP-1]
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE // decrement SP
	return assembly
}

func getOrCommandAssembly() string {
	assembly := ""
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE   // read value which SP points to into M
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-1 points to into M
	assembly += "D=M" + value.NEW_LINE                            // set M to D
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-2 points to into M
	assembly += "M=M|D" + value.NEW_LINE                          // set RAM[SP-2] = RAM[SP-2] or RAM[SP-1]
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE // decrement SP
	return assembly
}

func getNegCommandAssembly() string {
	assembly := ""
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE // read value which SP points to into M
	assembly += "A=A-1" + value.NEW_LINE                        // set value which SP-1 points to into M
	assembly += "M=-M" + value.NEW_LINE                         // set RAM[SP-1] = -RAM[SP-1]
	return assembly
}

func getNotCommandAssembly() string {
	assembly := ""
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE // read value which SP points to into M
	assembly += "A=A-1" + value.NEW_LINE                        // set value which SP-1 points to into M
	assembly += "M=!M" + value.NEW_LINE                         // set RAM[SP-1] = !RAM[SP-1]
	return assembly
}

func getCompareAssembly(compareCommandSymbol ast.CommandSymbol) string {
	assembly := ""
	// set x(RAM[SP-2]) - y(RAM[SP-1]) to D(==x-y)
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE + "A=M" + value.NEW_LINE + "D=M" + value.NEW_LINE // set RAM[SP-1]=y to D
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE + "A=M" + value.NEW_LINE                          // set RAM[SP-2]=x to M
	assembly += "D=M-D" + value.NEW_LINE                                                                            // set x - y to D
	compareAddressFlag := strconv.Itoa(rand.Intn(1000000))
	// jump based on D
	switch compareCommandSymbol {
	case ast.EQ:
		assembly += "@TRUE" + compareAddressFlag + value.NEW_LINE + "D;JEQ" + value.NEW_LINE
	case ast.GT:
		assembly += "@TRUE" + compareAddressFlag + value.NEW_LINE + "D;JGT" + value.NEW_LINE
	case ast.LT:
		assembly += "@TRUE" + compareAddressFlag + value.NEW_LINE + "D;JLT" + value.NEW_LINE
	}
	assembly += "M=0" + value.NEW_LINE + "@NEXT" + compareAddressFlag + value.NEW_LINE + "0;JMP" + value.NEW_LINE // if false set 0 to RAM[SP-2] & jump to NEXT(to prevent TRUE process)
	assembly += "(TRUE" + compareAddressFlag + ")" + value.NEW_LINE + "M=-1" + value.NEW_LINE                     // if true set -1 to RAM[SP-2]
	assembly += "(NEXT" + compareAddressFlag + ")" + value.NEW_LINE                                               // NEXT Addr
	assembly += "@SP" + value.NEW_LINE + "M=M+1" + value.NEW_LINE                                                 // increment SP
	return assembly
}

func getPushAssembly(pushCommand *ast.PushCommand) (string, error) {
	switch pushCommand.Segment {
	case ast.CONSTANT:
		return getPushConstantAssembly(pushCommand), nil
	}
	return "", fmt.Errorf("%T couldn't convert to pushAssembly", pushCommand)
}

func getPushConstantAssembly(pushCommand *ast.PushCommand) string {
	assembly := ""
	assembly += "@" + strconv.Itoa(pushCommand.Index) + value.NEW_LINE + "D=A" + value.NEW_LINE // set constant value to D
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE                                 // read value which SP points to into M
	assembly += "M=D" + value.NEW_LINE                                                          // set D to M
	assembly += "@SP" + value.NEW_LINE + "M=M+1" + value.NEW_LINE                               // increment SP
	return assembly
}

func (codeWriter *CodeWriter) writeAssembly(assembly string) {
	codeWriter.Assembly = append(codeWriter.Assembly, []byte(assembly)...)
}
