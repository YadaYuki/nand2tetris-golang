package main

import (
	"VMtranslator/ast"
	"VMtranslator/codewriter"
	"VMtranslator/parser"
	"io/ioutil"
)

func main() {
	vm, _ := ioutil.ReadFile("MemoryAccess/BasicTest/BasicTest.vm")
	parser := parser.New(string(vm))
	codeWriter := codewriter.New("MemoryAccess/BasicTest/BasicTest.asm", "BasicTest")
	for parser.HasMoreCommand() {
		switch parser.CommandType() {
		case ast.C_PUSH:
			command, _ := parser.ParsePush()
			codeWriter.WritePushPop(command)
		case ast.C_POP:
			command, _ := parser.ParsePop()
			codeWriter.WritePushPop(command)
		case ast.C_ARITHMETIC:
			command, _ := parser.ParseArithmetic()
			codeWriter.WriteArithmetic(command)
		case ast.C_LABEL:
			command, _ := parser.ParseLabel()
			codeWriter.WriteLabel(command)
		case ast.C_GOTO:
			command, _ := parser.ParseGoto()
			codeWriter.WriteGoto(command)
		case ast.C_FUNCTION:
			command, _ := parser.ParseFunction()
			codeWriter.WriteFunction(command)
		case ast.C_CALL:
			command, _ := parser.ParseCall()
			codeWriter.WriteCall(command)
		case ast.C_RETURN:
			command, _ := parser.ParseReturn()
			codeWriter.WriteReturn(command)
		}
		parser.Advance()
	}
	codeWriter.Close()
}
