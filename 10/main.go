package main

import (
	"fmt"
	"jack_compiler/compilationengine"
	"jack_compiler/tokenizer"
)

func main() {
	jt := tokenizer.New(`
	let a[i] = Keyboard.readInt("ENTER THE NEXT NUMBER: ");
`)
	ce := compilationengine.New(jt)
	fmt.Println(ce.ParseProgram().Xml())
}
