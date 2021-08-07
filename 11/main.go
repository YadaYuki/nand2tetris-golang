package main

import (
	"jack_compiler/compilationengine"
	"jack_compiler/parser"
	"jack_compiler/symboltable"
	"jack_compiler/tokenizer"
	"jack_compiler/vmwriter"
)

func main() {
	jt := tokenizer.New(`do Output.printInt(1 + (2 * 3));`)
	parser := parser.New(jt)
	ast := parser.ParseDoStatement()
	vm := vmwriter.New("sample.vm", 0644)
	st := symboltable.New()
	ce := compilationengine.New(vm, st)
	ce.CompileDoStatement(ast)
	ce.Close()
}
