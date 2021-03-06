package main

import (
	"fmt"
	"jack_compiler/tokenizer"
)

func main() {
	// flag.Parse()
	// filename := flag.Args()[0]
	// fmt.Println(filename)
	// file, err := os.Open(filename)
	// if err != nil {
	// 	log.Fatal("%s", err)
	// }
	// fileScanner := bufio.NewScanner(file)
	// for fileScanner.Scan() {
	// 	s := fileScanner.Text()
	// 	fmt.Println(s)
	// }
	jt := tokenizer.New("")
	fmt.Println(jt.Advance())
}
