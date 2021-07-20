package ast

type AritimeticCommandType string

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
	ArthmeticCommand AritimeticCommandType
}

func (arithmeticCommand *ArithmeticCommand) String() string {
	return string(arithmeticCommand.ArthmeticCommand)
}