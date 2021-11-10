package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"jackcompiler/ast"
	"jackcompiler/compilationengine"
	"jackcompiler/parser"
	"jackcompiler/symboltable"
	"jackcompiler/tokenizer"
	"jackcompiler/vmwriter"
	"os"
	"path/filepath"
)

func getJackFileListInDir(dirPath string) ([]string, error) {
	vmPathPattern := filepath.Join(dirPath, "*.jack")
	vmFileListInDir, err := filepath.Glob(vmPathPattern)
	if err != nil {
		return []string{}, err
	}
	return vmFileListInDir, nil
}

func main() {
	flag.Parse()
	pathToJack := flag.Arg(0)

	jackFileList := []string{}

	fileInfo, _ := os.Stat(pathToJack)
	if fileInfo.IsDir() {
		jackFileListInDir, err := getJackFileListInDir(pathToJack)
		if err != nil {
			panic(err)
		}
		jackFileList = jackFileListInDir
	} else {
		jackFileList = []string{pathToJack}
	}

	for _, jackFilename := range jackFileList {
		jackCode, err := ioutil.ReadFile(jackFilename)
		if err != nil {
			panic(err)
		}

		jt := tokenizer.New(string(jackCode))
		parser := parser.New(jt)
		programAst := parser.ParseProgram()
		classStmt, ok := programAst.Statements[0].(*ast.ClassStatement)
		if !ok {
			panic(fmt.Sprintf("Statement[0] should be ClassStatement, but got %T", classStmt))
		}
		className := classStmt.Name.Literal
		vm := vmwriter.New(fmt.Sprintf("vm/program/%s.vm", className), 0644)
		st := symboltable.New()
		ce := compilationengine.New(className, vm, st)
		ce.CompileProgram(programAst)
		ce.Close()
	}
}
