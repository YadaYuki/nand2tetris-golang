package codewriter

import (
	"bytes"
	"io/ioutil"
	"testing"
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
