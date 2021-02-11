package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"assembly/parser"
)

func main() {
	file, err := os.Open("rect/Rect.asm")

	if err != nil {
		log.Fatalf("%s", err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for {
		s, err := parser.Advance(fileScanner)
		if err != nil {
			break
		}
		fmt.Println(s)
	}
}
