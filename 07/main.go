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
	codeWriter := codewriter.New("BasicTest.asm", "BasicTest")
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
		}
		parser.Advance()
	}
	codeWriter.Close()
}
