package ast

import "fmt"

type CommandType string

const (
	PUSH CommandType = "push"
	POP  CommandType = "pop"
)

type AritimeticCommandType CommandType

const (
	ADD AritimeticCommandType = "add"
	SUB AritimeticCommandType = "sub"
	NEG AritimeticCommandType = "neg"
	EQ  AritimeticCommandType = "eq"
	GT  AritimeticCommandType = "gt"
	LT  AritimeticCommandType = "lt"
	AND AritimeticCommandType = "and"
	OR  AritimeticCommandType = "or"
	NOT AritimeticCommandType = "not"
)

type VMCommand interface {
	String()
}

type ArithmeticCommand struct {
	Command AritimeticCommandType
}

func (arithmeticCommand *ArithmeticCommand) String() string {
	return string(arithmeticCommand.Command)
}

type SegmentType string

const (
	ARGUMENT SegmentType = "argument"
	LOCAL    SegmentType = "local"
	STATIC   SegmentType = "static"
	CONSTANT SegmentType = "constant"
	THIS     SegmentType = "this"
	THAT     SegmentType = "that"
	POINTER  SegmentType = "pointer"
	TEMP     SegmentType = "temp"
)

type MemoryAccessCommand interface {
	VMCommand
}

type PushCommand struct {
	Comamnd CommandType // push
	Segment SegmentType
	Index   int
}

func (pushCommand *PushCommand) String() string {
	return fmt.Sprintf("%s %s %d", pushCommand.Comamnd, pushCommand.Segment, pushCommand.Index)
}
