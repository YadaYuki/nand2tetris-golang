package main

import (
	"Assembly/code"
	"Assembly/parser"
	"Assembly/symbol_table"
	"Assembly/util"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//TODO: test

func main() {
	file, err := os.Open("rect/Rect.asm")
	if err != nil {
		log.Fatalf("%s", err)
	}

	fileScanner := bufio.NewScanner(file)
	currentRomAddress := 0
	for fileScanner.Scan() {
		s := fileScanner.Text()
		if isComment := strings.Index(s, "//"); len(s) > 0 && isComment == -1 {
			commandType, err := parser.GetCommandType(s)
			if err != nil {
				log.Fatal(err)
			}
			if commandType == parser.CCommand || commandType == parser.ACommand {
				currentRomAddress++
			}
			if commandType == parser.LCommand {
				symbol, _ := parser.GetSymbol(s)
				symbol_table.AddEntry(symbol, currentRomAddress+1)
			}
		}
	}

	// TODO:Fix File reader
	file2, err := os.Open("rect/Rect.asm")
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer file2.Close()
	fileScanner = bufio.NewScanner(file2)
	currentCustomVariableAddress := 16
	for fileScanner.Scan() {
		s := fileScanner.Text()

		if isComment := strings.Index(s, "//"); len(s) > 0 && isComment == -1 {
			commandType, err := parser.GetCommandType(s)
			if err != nil {
				log.Fatal(err)
			}
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
			if commandType == parser.ACommand {
				symbol, err := parser.GetSymbol(s)
				symbolInt, err := strconv.Atoi(symbol)
				if err == nil {
					binaryCode = "0" + util.Fill(strconv.FormatInt(int64(symbolInt), 2), "0", 15)
				}
				if err != nil { // custom variable
					if contains := symbol_table.Contains(symbol); contains == true {
						address, _ := symbol_table.GetAddress(symbol)
						binaryCode = "0" + util.Fill(strconv.FormatInt(int64(address), 2), "0", 15)
					} else {
						binaryCode = "0" + util.Fill(strconv.FormatInt(int64(currentCustomVariableAddress), 2), "0", 15)
						symbol_table.AddEntry(symbol, currentCustomVariableAddress)
						currentCustomVariableAddress++
					}
				}
			}
			if commandType != parser.LCommand {
				fmt.Println(binaryCode)
			}
		}
	}
	defer file2.Close()
}
