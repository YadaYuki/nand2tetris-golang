package main

import (
	"fmt"
	"jack/compiler/tokenizer"
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
}
