package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"jack_compiler/ast"
	"jack_compiler/compilationengine"
	"jack_compiler/parser"
	"jack_compiler/symboltable"
	"jack_compiler/tokenizer"
	"jack_compiler/vmwriter"
)

func main() {
	flag.Parse()
	jackFilename := flag.Arg(0)
	jackCode, err := ioutil.ReadFile(jackFilename)
	if err != nil {
		panic(err)
	}
	jt := tokenizer.New(string(jackCode))
	parser := parser.New(jt)
	programAst := parser.ParseProgram()
	classStmt, ok := programAst.Statements[0].(*ast.ClassStatement)
	if !ok {
		panic(fmt.Sprintf("Statement[0] should be ClassStatement, but got %T", classStmt))
	}
	className := classStmt.Name.Literal

	vm := vmwriter.New(fmt.Sprintf("vm/%s.vm", className), 0644)
	st := symboltable.New()
	ce := compilationengine.New(className, vm, st)
	ce.CompileProgram(programAst)
	ce.Close()
}
