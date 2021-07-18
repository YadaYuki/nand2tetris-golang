package main

import (
	"assembly/ast"
	"assembly/parser"
	"assembly/symboltable"
	"fmt"
	"strconv"
)

//TODO: test

func main() {
	input := `(LOOP)
@FUGA
D=A
@HOGE
D=D+A
@0
M=D
AM=D|A;JMP`
	st := symboltable.New()
	p := parser.New(input, st)
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
	fmt.Println(p.SymbolTableDict)
}
