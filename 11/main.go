package main

import (
	"fmt"
	"jack_compiler/parser"
	"jack_compiler/tokenizer"
)

func main() {
	jt := tokenizer.New(`
	class Main {
		function void main() {
			 do Output.printInt(1 + (2 * 3));
			 return;
		}
 }
`)
	parser := parser.New(jt)
	fmt.Println(parser.ParseProgram().Xml())
}
