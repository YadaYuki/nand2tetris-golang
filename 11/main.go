package main

import (
	"flag"
	"io/ioutil"
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
	ast := parser.ParseProgram()
	vm := vmwriter.New("vm/Main.vm", 0644)
	st := symboltable.New()
	ce := compilationengine.New("Main", vm, st)
	ce.CompileInit()
	ce.CompileProgram(ast)
	ce.Close()
}
