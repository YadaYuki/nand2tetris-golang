package main

import (
	"VMtranslator/ast"
	"VMtranslator/codewriter"
	"VMtranslator/parser"
	"io/ioutil"
)

func main() {
	vm, _ := ioutil.ReadFile("StackArithmetic/SimpleAdd/SimpleAdd.vm")
	parser := parser.New(string(vm))
	codeWriter := codewriter.New("SimpleAdd.asm")

	for parser.HasMoreCommand() {
		switch parser.CommandType() {

		case ast.C_PUSH:
			arg1, _ := parser.Arg1()
			arg2, _ := parser.Arg2()
			command := &ast.PushCommand{Comamnd: ast.C_PUSH, Symbol: ast.PUSH, Segment: ast.SegmentType(arg1), Index: arg2}
			codeWriter.WritePushPop(command)
		case ast.C_ARITHMETIC:
			arg1, _ := parser.Arg1()
			command := &ast.ArithmeticCommand{Command: ast.C_ARITHMETIC, Symbol: ast.CommandSymbol(arg1)}
			codeWriter.WriteArithmetic(command)
		}
		parser.Advance()
	}
	codeWriter.Close()
}
