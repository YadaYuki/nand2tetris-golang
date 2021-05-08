package main

import (
	"fmt"
	"jack_compiler/compilationengine"
	"jack_compiler/tokenizer"
)

func main() {
	jt := tokenizer.New(`
	class Main {
  function void main() {
  var Array a;
  var int length;
  var int i, sum;
	let length = Keyboard.readInt("HOW MANY NUMBERS? ");
	let a = Array.new(length);
	let i = 0;
	while (i < length) {
	    let a[i] = Keyboard.readInt("ENTER THE NEXT NUMBER: ");
	    let i = i + 1;
	}
	let i = 0;
	let sum = 0;
	while (i < length) {
	    let sum = sum + a[i];
	    let i = i + 1;
	}
	do Output.printString("THE AVERAGE IS: ");
	do Output.printInt(sum / length);
	do Output.println();
	return;
	}
}
`)
	ce := compilationengine.New(jt)
	fmt.Println(ce.ParseProgram().Xml())
}
