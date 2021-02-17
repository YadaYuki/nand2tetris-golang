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
		commandType, _ := parser.GetCommandType(s)
		fmt.Println(commandType)
	}
}
