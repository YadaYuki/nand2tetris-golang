package ast

import "fmt"

type CommandType string

const (
	C_PUSH       CommandType = "C_PUSH"
	C_POP        CommandType = "C_POP"
	C_ARITHMETIC CommandType = "C_ARITHMETIC"
	C_LABEL      CommandType = "C_LABEL"
	C_GOTO       CommandType = "C_GOTO"
	C_IF         CommandType = "C_IF"
	C_FUNCTION   CommandType = "C_FUNCTION"
	C_RETURN     CommandType = "C_RETURN"
	C_CALL       CommandType = "C_CALL"
	C_EMPTY      CommandType = "C_EMPTY"
)

type CommandSymbol string

const (
	PUSH    CommandSymbol = "push"
	POP     CommandSymbol = "pop"
	LABEL   CommandSymbol = "label"
	GOTO    CommandSymbol = "goto"
	IF_GOTO CommandSymbol = "if-goto"
	CALL    CommandSymbol = "call"
	ADD     CommandSymbol = "add"
	SUB     CommandSymbol = "sub"
	NEG     CommandSymbol = "neg"
	EQ      CommandSymbol = "eq"
	GT      CommandSymbol = "gt"
	LT      CommandSymbol = "lt"
	AND     CommandSymbol = "and"
	OR      CommandSymbol = "or"
	NOT     CommandSymbol = "not"
)

type VMCommand interface {
	String() string
}

type ArithmeticCommand struct {
	Command CommandType   // C_ARITHMETIC
	Symbol  CommandSymbol // "add","lt"...
}

func (arithmeticCommand *ArithmeticCommand) String() string {
	return string(arithmeticCommand.Symbol)
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
	Comamnd CommandType // C_PUSH
	Symbol  CommandSymbol
	Segment SegmentType
	Index   int
}

func (pushCommand *PushCommand) String() string {
	return fmt.Sprintf("%s %s %d", pushCommand.Symbol, pushCommand.Segment, pushCommand.Index)
}

type PopCommand struct {
	Comamnd CommandType // C_POP
	Symbol  CommandSymbol
	Segment SegmentType
	Index   int
}

func (popCommand *PopCommand) String() string {
	return fmt.Sprintf("%s %s %d", popCommand.Symbol, popCommand.Segment, popCommand.Index)
}

type LabelCommand struct {
	Command   CommandType   // C_LABEL
	Symbol    CommandSymbol // label
	LabelName string
}

func (labelCommand *LabelCommand) String() string {
	return fmt.Sprintf("%s %s", labelCommand.Symbol, labelCommand.LabelName)
}

type GotoCommand struct {
	Command   CommandType   // C_GOTO
	Symbol    CommandSymbol // goto
	LabelName string
}

func (gotoCommand *GotoCommand) String() string {
	return fmt.Sprintf("%s %s", gotoCommand.Symbol, gotoCommand.LabelName)
}

type IfCommand struct {
	Command   CommandType   // C_IF
	Symbol    CommandSymbol // if-goto
	LabelName string
}

func (ifCommand *IfCommand) String() string {
	return fmt.Sprintf("%s %s", ifCommand.Symbol, ifCommand.LabelName)
}

type CallCommand struct {
	Command      CommandType   // C_CALL
	Symbol       CommandSymbol // call
	FunctionName string
	numArgs      int
}

func (callCommand *CallCommand) String() string {
	return fmt.Sprintf("%s %s %d", callCommand.Symbol, callCommand.FunctionName, callCommand.numArgs)
}

// type CallCommand struct {
// 	Command      CommandType   // C_CALL
// 	Symbol       CommandSymbol // if-goto
// 	FunctionName string
// 	numArgs      int
// }

// func (callCommand *CallCommand) String() string {
// 	return fmt.Sprintf("%s %s %d", callCommand.Symbol, callCommand.FunctionName, callCommand.numArgs)
// }
