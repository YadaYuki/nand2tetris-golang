package main

import (
	"fmt"
	"jack_compiler/compilationengine"
	"jack_compiler/tokenizer"
)

func main() {
	jt := tokenizer.New(`
    function void main() {
        var Array a;
        var int length;
        var int i, sum;
      	let i = 0;
				while (i < length) {
           let a = 0;
					 let sum = sum + a[i];
					 let i = i + 1;
      	}
	}`)
	ce := compilationengine.New(jt)
	fmt.Println(ce.ParseProgram().Xml())
}
