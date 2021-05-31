package main

import (
	"fmt"
	"jack_compiler/compilationengine"
	"jack_compiler/symboltable"
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
	st := symboltable.New()
	ce := compilationengine.New(jt, st)
	fmt.Println(ce.ParseProgram().Xml())
}
