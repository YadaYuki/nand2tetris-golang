package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"vm_translator/ast"
	"vm_translator/codewriter"
	"vm_translator/parser"
)

func getVmFileListInDir(dirPath string) ([]string, error) {
	vmPathPattern := filepath.Join(dirPath, "*.vm")
	vmFileListInDir, err := filepath.Glob(vmPathPattern)
	if err != nil {
		return []string{}, err
	}
	return vmFileListInDir, nil
}

func removeExt(filename string) string {
	return strings.Trim(filename, filepath.Ext(filename))
}

func translateVm(className string, vmDirname string, asmDirname string) {
	vmFileList, err := getVmFileListInDir(vmDirname)
	if err != nil {
		panic(err)
	}
	// join all vm code in dir.
	vmCodeList, vmClassNameList := []string{}, []string{}
	for _, vmFile := range vmFileList {
		vmCode, err := ioutil.ReadFile(vmFile)
		if err != nil {
			panic(err)
		}
		filename := filepath.Base(vmFile)
		vmClassNameList = append(vmClassNameList, removeExt(filename))
		vmCodeList = append(vmCodeList, string(vmCode))
	}
	codeWriter := codewriter.New(fmt.Sprintf("%s/%s.asm", asmDirname, className))
	// writeInit
	codeWriter.WriteInit()
	for i := range vmCodeList {
		parser := parser.New(vmCodeList[i])
		codeWriter.SetVmClassName(vmClassNameList[i])
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
			case ast.C_IF:
				command, _ := parser.ParseIf()
				codeWriter.WriteIf(command)
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
	}
	codeWriter.Close()

}

func main() {
	className, vmDirname, asmDirname := flag.String("class", "SimpleFunction", "name of vm class"), flag.String("vm-dir", "FunctionCalls/SimpleFunction", "dirname of vm"), flag.String("asm-dir", "FunctionCalls/SimpleFunction", "dirname of asm")
	flag.Parse()
	translateVm(*className, *vmDirname, *asmDirname)
}
