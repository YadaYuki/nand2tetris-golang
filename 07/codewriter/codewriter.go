package codewriter

import (
	"io/ioutil"
)

type CodeWirter struct {
	Filename string
	Assembly []byte
}

func New(filename string) *CodeWirter {
	return &CodeWirter{Filename: filename, Assembly: []byte{}}
}

func (codeWriter *CodeWirter) Close() {
	err := ioutil.WriteFile(codeWriter.Filename, codeWriter.Assembly, 0644)
	if err != nil {
		panic(err)
	}
}
