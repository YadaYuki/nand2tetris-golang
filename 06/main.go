package main

import (
	"assembly/code"
	"assembly/parser"
	"fmt"
)

//TODO: test

func main() {
	input := `@2
D=A
@3
D=D+A
@0
M=D
AM=D|A;JMP`
	p := parser.New(input)
	commands, _ := p.ParseAssembly()
	for _, command := range commands {
		fmt.Println(code.Binary(command))
	}
}
