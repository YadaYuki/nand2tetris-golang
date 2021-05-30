package vmwriter

import (
	"io/fs"
	"io/ioutil"
)

type VMWriter struct {
	VMCode   []byte
	Filename string
	perm     fs.FileMode
}

type Command string

const (
	ADD Command = "add"
	SUB Command = "sub"
	NEG Command = "neg"
	EQ  Command = "eq"
	GT  Command = "gt"
	LT  Command = "lt"
	AND Command = "and"
	OR  Command = "or"
	NOT Command = "not"
)

type Segment string

const (
	CONST   Segment = "const"
	ARG     Segment = "arg"
	LOCAL   Segment = "local"
	STATIC  Segment = "static"
	THIS    Segment = "this"
	THAT    Segment = "that"
	POINTER Segment = "pointer"
	TEMP    Segment = "temp"
)

func New(filename string, permission fs.FileMode) *VMWriter {
	return &VMWriter{Filename: filename, VMCode: []byte{}, perm: permission}
}

func (vm *VMWriter) WritePush(segment Segment, idx int) {
}

func (vm *VMWriter) WritePop(segment Segment, idx int) {
}

func (vm *VMWriter) WriteArithmetic(command Command) {
}

func (vm *VMWriter) WriteLabel(label string) {
}

func (vm *VMWriter) WriteGoto(label string) {
}

func (vm *VMWriter) WriteIf(label string) {
}

func (vm *VMWriter) WriteCall(name string, nArgs int) {
}

func (vm *VMWriter) WriteFunction(name string, nArgs int) {
}

func (vm *VMWriter) WriteReturn() {
}

func (vm *VMWriter) Close() {
	err := ioutil.WriteFile(vm.Filename, vm.VMCode, vm.perm)
	if err != nil {
		panic(err)
	}
}
