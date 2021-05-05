package main

import (
	"fmt"
	"jack_compiler/compilationengine"
	"jack_compiler/tokenizer"
)

func main() {
	jt := tokenizer.New(`
	class Main {
    static boolean test;   
                           
    function void main() {
      var SquareGame game;
      let game = SquareGame.new();
      do game.run();
      do game.dispose();
      return;
    }

    function void more() {  
        var int i, j;       
        var String s;
        var Array a;
        if (false) {
            let s = "string constant";
            let s = null;
            let a[1] = a[2];
        }
        else {              
            let i = i * (-j);
            let j = j / (-2);  
            let i = i | j;
        }
        return;
    }
}
	
	`)
	ce := compilationengine.New(jt)
	fmt.Println(ce.ParseProgram())
}
