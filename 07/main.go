package main

import (
	"VMtranslator/code_writer"
	"VMtranslator/parser"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
		if commandType == parser.CPush {
			segment, _ := parser.GetArg1(s)
			index, _ := parser.GetArg2(s)
			indexInt, _ := strconv.Atoi(index)
			assembly, _ := code_writer.GetPushPop(commandType, segment, indexInt)
			fmt.Println(assembly)
		}
		if commandType == parser.CArithmetic {
			assembly, _ := code_writer.GetArithmetic(s)
			fmt.Println(assembly)
		}
	}
}
