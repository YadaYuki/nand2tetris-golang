package vmwriter

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"jack_compiler/value"
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
	pushVmCode := fmt.Sprintf("push %s %d", segment, idx) + value.NEW_LINE
	vm.writeData(pushVmCode)
}

func (vm *VMWriter) WritePop(segment Segment, idx int) {
	popVmCode := fmt.Sprintf("pop %s %d", segment, idx) + value.NEW_LINE
	vm.writeData(popVmCode)
}

func (vm *VMWriter) WriteArithmetic(command Command) {
	vm.writeData(string(command) + value.NEW_LINE)
}

func (vm *VMWriter) WriteLabel(label string) {
	labelVmCode := fmt.Sprintf("label %s", label) + value.NEW_LINE
	vm.writeData(labelVmCode)
}

func (vm *VMWriter) WriteGoto(label string) {
	gotoVmCode := fmt.Sprintf("goto %s", label) + value.NEW_LINE
	vm.writeData(gotoVmCode)
}

func (vm *VMWriter) WriteIf(label string) {
	ifVmCode := fmt.Sprintf("if-goto %s", label) + value.NEW_LINE
	vm.writeData(ifVmCode)
}

func (vm *VMWriter) WriteCall(name string, nArgs int) {
	callVmCode := fmt.Sprintf("call %s %d", name, nArgs) + value.NEW_LINE
	vm.writeData(callVmCode)
}

func (vm *VMWriter) WriteFunction(name string, nLocals int) {
	functionVmCode := fmt.Sprintf("function %s %d", name, nLocals) + value.NEW_LINE
	vm.writeData(functionVmCode)
}

func (vm *VMWriter) WriteReturn() {
}

func (vm *VMWriter) writeData(vmCode string) {
	vm.VMCode = append(vm.VMCode, []byte(vmCode)...)
}

func (vm *VMWriter) Close() {
	err := ioutil.WriteFile(vm.Filename, vm.VMCode, vm.perm)
	if err != nil {
		panic(err)
	}
}
