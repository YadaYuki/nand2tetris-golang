package main

import (
	"Assembly/parser"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("rect/Rect.asm")

	if err != nil {
		log.Fatalf("%s", err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		fmt.Println(parser.GetCommandType("hoge"))
	}
}
