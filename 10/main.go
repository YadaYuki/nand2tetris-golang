package main

import (
	"fmt"
	"jack_compiler/compilationengine"
	"jack_compiler/tokenizer"
)

func main() {
	jt := tokenizer.New(`
	let i = 0;
	let sum = 0;
	
	while (i < length) {
	    let sum = sum + a[i];
	    let i = i + 1;
	}}
	
	do Output.printString("THE AVERAGE IS: ");
	do Output.printInt(sum / length);
	do Output.println();
	
	`)
	ce := compilationengine.New(jt)
	fmt.Println(ce.ParseProgram().Xml())
}
