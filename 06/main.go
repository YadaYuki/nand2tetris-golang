package main

import (
	"assembly/ast"
	"assembly/code"
	"assembly/parser"
	"assembly/symboltable"
	"fmt"
	"io/ioutil"
	"strconv"
)

//TODO: test

func main() {
	asm, _ := ioutil.ReadFile("add/Add.asm")
	input := string(asm)
	st := symboltable.New()
	p := parser.New(input, st)
	// first path
	customVariableCount := 0
	INTIAL_VARIABLE_COUNT := 16
	for i := 0; p.HasMoreCommand(); i++ {
		switch p.CommandType() {
		case ast.A_COMMAND:
			symbol, _ := p.Symbol()
			_, err := strconv.Atoi(symbol)
			if err == nil { // not symbol
				break
			}
			err = p.AddEntry(symbol, INTIAL_VARIABLE_COUNT+customVariableCount)
			if err != nil { // already registered.
				break
			}
			customVariableCount++
		case ast.L_COMMAND:
			symbol, _ := p.Symbol()
			p.AddEntry(symbol, i)
		}
		p.Advance()
	}
	p.ResetParseIdx()
	// second path
	// binary := ""
	for p.HasMoreCommand() {
		command, _ := p.ParseCommand()
		if p.CommandType() == ast.A_COMMAND || p.CommandType() == ast.C_COMMAND {
			// c, ok := command.(*ast.CCommand)
			fmt.Println(code.Binary(command))
		}
		p.Advance()
	}
}
