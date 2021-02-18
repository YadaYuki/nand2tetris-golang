package main

import (
	"VMtranslator/parser"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	filename := flag.Args()[0]
	// TODO: if filename is directory/ parse all .vm file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		s := fileScanner.Text()
		// commandType, err := parser.GetCommandType(s)
		// if commandType{

		// }
		arg1, _ := parser.GetArg1(s)
		fmt.Println(arg1)
	}
}
