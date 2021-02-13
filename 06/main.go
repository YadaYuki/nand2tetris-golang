package main

import (
	"Assembly/code"
	"Assembly/parser"
	"Assembly/util"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("rect/RectL.asm")

	if err != nil {
		log.Fatalf("%s", err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		s := fileScanner.Text()
		if len(s) > 0 {
			commandType, _ := parser.GetCommandType(s)

			binaryCode := ""
			if commandType == parser.CCommand {
				dest, err := parser.GetDest(s)
				if err != nil {
					log.Fatal(err)
				}
				jump, err := parser.GetJump(s)
				if err != nil {
					log.Fatal(err)
				}
				comp, err := parser.GetComp(s)
				if err != nil {
					log.Fatal(err)
				}
				destBinary, err := code.GetDestBinary(dest)
				if err != nil {
					log.Fatal(err)
				}
				jumpBinary, err := code.GetJumpBinary(jump)
				if err != nil {
					log.Fatal(err)
				}
				compBinary, err := code.GetCompBinary(comp)
				if err != nil {
					log.Fatal(err)
				}
				binaryCode = "111" + compBinary + destBinary + jumpBinary
			}
			if commandType == parser.LCommand {
			}
			if commandType == parser.ACommand {
				symbolStr, _ := parser.GetSymbol(s)
				if symbol, err := strconv.Atoi(symbolStr); err == nil {
					binaryCode = "0" + util.Fill(strconv.FormatInt(int64(symbol), 2), "0", 15)
				}
			}
			// TODO:implement Write file

			fmt.Println(binaryCode)
		}
	}
}
