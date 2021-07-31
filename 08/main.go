package main

import (
	"VMtranslator/ast"
	"VMtranslator/codewriter"
	"VMtranslator/parser"
	"io/ioutil"
)

func main() {
	vm, err := ioutil.ReadFile("FunctionCalls/NestedCall/Sys.vm")
	if err != nil {
		panic(err)
	}
	parser := parser.New(string(vm))
	codeWriter := codewriter.New("FunctionCalls/NestedCall/Sys.asm", "Main")
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
