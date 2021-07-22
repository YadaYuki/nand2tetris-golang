package codewriter

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestClose(t *testing.T) {

	assembly := []byte("Hello,World")
	filename := "test.asm"
	codeWriter := &CodeWirter{
		Assembly: assembly, Filename: filename,
	}
	codeWriter.Close()
	content, _ := ioutil.ReadFile(filename)
	if !bytes.Equal(content, assembly) {
		t.Fatalf("assembly should be %s. got %s", assembly, content)
	}
}
