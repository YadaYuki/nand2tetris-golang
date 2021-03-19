package main

import (
	"fmt"
	"jack/compiler/tokenizer"
	"jack/compiler/compilationengine"
)

func main() {
	TODO:Fix bug related near EOF
	jt := tokenizer.New(`var 
	aa = 
	'123'; 
	cc = 'asdfasdfasdfhogehoge'
	bb = 133; 
	
	class {
		hoge = 123
	}        `)
	for jt.HasMoreTokens() {
		token, err := jt.Advance()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(token)
	}
// 	input := `
// 	return x ;
// 	return 1 ;
// 	return ;
// `
// jt := tokenizer.New(input)
// ce := compilationengine.New(jt)
// program := ce.ParseProgram()
}
