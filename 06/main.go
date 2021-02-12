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
		s := fileScanner.Text()
		if len(s) > 0 {
			commandType, err := parser.GetCommandType(s)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(commandType)
		}
	}
}
