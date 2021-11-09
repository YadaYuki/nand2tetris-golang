package codewriter

import (
	"bytes"
	"io/ioutil"
	"testing"
	"vmtranslator/ast"
)

func TestClose(t *testing.T) {
	assembly := []byte("Hello,World")
	filename := "test.asm"
	codeWriter := &CodeWriter{
		Assembly: assembly, Filename: filename,
	}
	codeWriter.Close()
	content, _ := ioutil.ReadFile(filename)
	if !bytes.Equal(content, assembly) {
		t.Fatalf("assembly should be %s. got %s", assembly, content)
	}
}

func TestWriteAssembly(t *testing.T) {
	assembly := "Hello,World"
	codeWriter := New("test.asm")
	codeWriter.writeAssembly(string(assembly))
	if !bytes.Equal(codeWriter.Assembly, []byte(assembly)) {
		t.Fatalf("assembly should be %s. got %s", assembly, codeWriter.Assembly)
	}
}

func TestGetPushAssembly(t *testing.T) {
	testCases := []struct {
		pushCommand *ast.PushCommand
		assembly    string
	}{
		{&ast.PushCommand{Comamnd: ast.C_PUSH, Symbol: ast.PUSH, Segment: ast.CONSTANT, Index: 1111}, "@1111\r\nD=A\r\n@SP\r\nA=M\r\nM=D\r\n@SP\r\nM=M+1\r\n"},
	}
	codeWriter := New("test.asm")
	for _, tt := range testCases {
		assembly, _ := codeWriter.getPushAssembly(tt.pushCommand)
		if !bytes.Equal([]byte(assembly), []byte(tt.assembly)) {
			t.Fatalf("assembly should be %s. got %s", tt.assembly, assembly)
		}
	}
}

func TestGetArithmeticAssembly(t *testing.T) {
	testCases := []struct {
		arithmeticCommand *ast.ArithmeticCommand
		assembly          string
	}{
		{arithmeticCommand: &ast.ArithmeticCommand{Command: ast.C_ARITHMETIC, Symbol: ast.ADD}, assembly: ""},
	}
	codeWriter := New("test.asm")
	for _, tt := range testCases {
		assembly, _ := codeWriter.getArithmeticAssembly(tt.arithmeticCommand)
		if !bytes.Equal([]byte(assembly), []byte(tt.assembly)) {
			t.Fatalf("assembly should be %s. got %s", tt.assembly, assembly)
		}
	}
}
