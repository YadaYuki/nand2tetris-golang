package vmwriter

import (
	"jack_compiler/value"
)

type VMWriter struct {
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

func New() *VMWriter {
	return &VMWriter{}
}

func (vm *VMWriter) WritePush() string {
	return "" + value.NEW_LINE
}
