package main

import (
	"assembler/ast"
	"assembler/code"
	"assembler/parser"
	"assembler/symboltable"
	"assembler/value"
	"flag"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Assemble(input string) (binaryArr []string, err error) {
	st := symboltable.New()
	p := parser.New(input, st)
	// first path
	currentBinaryCount := 0
	for p.HasMoreCommand() {
		switch p.CommandType() {
		case ast.A_COMMAND, ast.C_COMMAND:
			currentBinaryCount++
		case ast.L_COMMAND:
			symbol, _ := p.Symbol()
			p.AddEntry(symbol, currentBinaryCount)
		}
		p.Advance()
	}
	p.ResetParseIdx()

	customVariableCount := 0
	INTIAL_VARIABLE_COUNT := 16
	for p.HasMoreCommand() {
		if p.CommandType() == ast.A_COMMAND {
			symbol, _ := p.Symbol()
			_, err := strconv.Atoi(symbol)
			if !p.Contains(symbol) && err != nil {
				p.AddEntry(symbol, INTIAL_VARIABLE_COUNT+customVariableCount)
				customVariableCount++
			}
		}
		p.Advance()
	}
	p.ResetParseIdx()
	// second path
	for p.HasMoreCommand() {
		command, _ := p.ParseCommand()
		if p.CommandType() == ast.A_COMMAND || p.CommandType() == ast.C_COMMAND {
			binaryArr = append(binaryArr, code.Binary(command))
		}
		p.Advance()
	}
	return binaryArr, nil
}

func AssembleAsmFile(asmFilename string, hackFilename string) error {
	asm, _ := ioutil.ReadFile(asmFilename)
	input := string(asm)
	binaryArr, _ := Assemble(input)
	ioutil.WriteFile(hackFilename, []byte(strings.Join(binaryArr, value.NEW_LINE)), os.ModePerm)
	return nil
}

func main() {
	asmFilename, hackFilename := flag.String("asm", "add/Add.asm", "filename of asm"), flag.String("hack", "Main.hack", "filename of hack")
	flag.Parse()
	AssembleAsmFile(*asmFilename, *hackFilename)
}
