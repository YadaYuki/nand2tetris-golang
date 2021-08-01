package main

import (
	"VMtranslator/ast"
	"VMtranslator/codewriter"
	"VMtranslator/parser"
	"VMtranslator/value"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func getVmFileListInDir(dirPath string) ([]string, error) {
	vmPathPattern := filepath.Join(dirPath, "*.vm")
	vmFileListInDir, err := filepath.Glob(vmPathPattern)
	if err != nil {
		return []string{}, err
	}
	return vmFileListInDir, nil
}

func main() {
	//
	vmFileList, err := getVmFileListInDir("FunctionCalls/FibonacciElement")
	if err != nil {
		panic(err)
	}
	// join all vm code in dir.
	vmCodeList := []string{}
	for _, vmFile := range vmFileList {
		vmCode, err := ioutil.ReadFile(vmFile)
		if err != nil {
			panic(err)
		}
		vmCodeList = append(vmCodeList, string(vmCode))
	}
	vm := strings.Join(vmCodeList, value.NEW_LINE)

	parser := parser.New(string(vm))
	codeWriter := codewriter.New("FunctionCalls/NestedCall/Sys.asm", "Main")

	// writeInit
	codeWriter.WriteInit()

	for parser.HasMoreCommand() {
		switch parser.CommandType() {
		case ast.C_PUSH:
			command, _ := parser.ParsePush()
			codeWriter.WritePushPop(command)
		case ast.C_POP:
			command, _ := parser.ParsePop()
			codeWriter.WritePushPop(command)
		case ast.C_ARITHMETIC:
			command, _ := parser.ParseArithmetic()
			codeWriter.WriteArithmetic(command)
		case ast.C_LABEL:
			command, _ := parser.ParseLabel()
			codeWriter.WriteLabel(command)
		case ast.C_GOTO:
			command, _ := parser.ParseGoto()
			codeWriter.WriteGoto(command)
		case ast.C_FUNCTION:
			command, _ := parser.ParseFunction()
			codeWriter.WriteFunction(command)
		case ast.C_CALL:
			command, _ := parser.ParseCall()
			codeWriter.WriteCall(command)
		case ast.C_RETURN:
			command, _ := parser.ParseReturn()
			codeWriter.WriteReturn(command)
		}
		parser.Advance()
	}
	codeWriter.Close()
}
