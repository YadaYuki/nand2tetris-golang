package main

import (
	"fmt"
	"jack_compiler/compilationengine"
	"jack_compiler/tokenizer"
)

func main() {
	jt := tokenizer.New(`
    class Main {

   constructor Square new(int Ax, int Ay, int Asize) {
    let x = Ax;
    let y = Ay;
    let size = Asize;
    do draw;
    return x;
 }
    function void main() {
        var int length;
        var int i, sum;
        return a;
      }
    }`)
	ce := compilationengine.New(jt)
	fmt.Println(ce.ParseProgram().Xml())

}
