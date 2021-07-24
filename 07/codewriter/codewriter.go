package codewriter

import (
	"VMtranslator/ast"
	"VMtranslator/value"
	"fmt"
	"io/ioutil"
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
		return nil
	}
	codeWriter.writeAssembly(arithmeticAssembly)
	return nil
}

func getArithmeticAssembly(arithmeticCommand *ast.ArithmeticCommand) (string, error) {
	switch arithmeticCommand.Symbol {
	case ast.ADD:
		return getAddCommandAssembly(), nil
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
