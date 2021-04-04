package main

import (
	"fmt"
	"jack/compiler/tokenizer"
	// "jack/compiler/compilationengine"
)

func main() {
	jt := tokenizer.New(`!1`)
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
