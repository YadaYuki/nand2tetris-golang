package parser

type CommandType int

const (
	CArithmetic CommandType = iota
	CPush
	CPop
	CLabel
	CGoto
	CIf
	CFunction
	CReturn
	CCall
)
