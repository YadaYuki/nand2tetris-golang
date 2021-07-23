package codewriter

import (
	"io/ioutil"
)

type CodeWriter struct {
	Filename string
	Assembly []byte
}

func New(filename string) *CodeWriter {
	return &CodeWriter{Filename: filename, Assembly: []byte{}}
}

func (codeWriter *CodeWriter) Close() {
	err := ioutil.WriteFile(codeWriter.Filename, codeWriter.Assembly, 0644)
	if err != nil {
		panic(err)
	}
}

func (codeWriter *CodeWriter) writeAssembly(assembly string) {
	codeWriter.Assembly = append(codeWriter.Assembly, []byte(assembly)...)
}
