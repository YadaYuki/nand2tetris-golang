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
		return nil
	}
	codeWriter.writeAssembly(assembly)
	return nil
}

func getPushAssembly(pushCommand *ast.PushCommand) (string, error) {
	switch pushCommand.Segment {
	case ast.CONSTANT:
		return getPushConstantAssembly(pushCommand), nil
	}
	return "", fmt.Errorf("%T couldn't convert to assembly", pushCommand)
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
