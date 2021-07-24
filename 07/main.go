package main

import (
	"VMtranslator/ast"
	"VMtranslator/codewriter"
	"VMtranslator/parser"
	"io/ioutil"
)

func main() {
	vm, _ := ioutil.ReadFile("StackArithmetic/StackTest/StackTest.vm")
	parser := parser.New(string(vm))
	codeWriter := codewriter.New("StackTest.asm")
	for parser.HasMoreCommand() {
		switch parser.CommandType() {
		case ast.C_PUSH:
			command, _ := parser.ParsePush()
			codeWriter.WritePushPop(command)
		case ast.C_ARITHMETIC:
			command, _ := parser.ParseArithmetic()
			codeWriter.WriteArithmetic(command)
		}
		parser.Advance()
	}
	codeWriter.Close()
}
