package code

import (
	"assembly/ast"
	"fmt"
)

func Binary(command ast.Command) string {
	switch c := command.(type) {
	case *ast.ACommand:
		return fmt.Sprintf("%016b", c.Value)
	case *ast.CCommand:
		cCommandBinaryPrefix := "111"
		return cCommandBinaryPrefix + Comp(c.Comp) + Dest(c.Dest) + Jump(c.Jump)
	}
	return ""
}

// GetDestBinary return Binary Code Correspond to dest label
func Dest(dest string) string {
	if dest == "" {
		return "000"
	}
	destBinaryMap := map[string]string{"M": "001", "D": "010", "MD": "011", "A": "100", "AM": "101", "AD": "110", "AMD": "111"}
	destBinary := destBinaryMap[dest]
	return destBinary
}

// GetJumpBinary return Binary Code Correspond to dest label
func Jump(jump string) string {
	if jump == "" {
		return "000"
	}
	jumpBinaryMap := map[string]string{"JGT": "001", "JEQ": "010", "JGE": "011", "JLT": "100", "JNE": "101", "JLE": "110", "JMP": "111"}
	jumpBinary := jumpBinaryMap[jump]
	return jumpBinary
}

func Comp(comp string) string {
	compBinaryMap := map[string]string{
		// a = 0
		"0":   "0101010",
		"1":   "0111111",
		"-1":  "0111010",
		"D":   "0001100",
		"A":   "0110000",
		"!D":  "0001101",
		"!A":  "0110001",
		"-D":  "0001111",
		"-A":  "0110011",
		"D+1": "0011111",
		"A+1": "0110111",
		"D-1": "0001110",
		"A-1": "0110010",
		"D+A": "0000010",
		"D-A": "0010011",
		"A-D": "0000111",
		"D&A": "0000000",
		"D|A": "0010101",
		// a = 1
		"M":   "1110000",
		"!M":  "1110001",
		"-M":  "1110011",
		"M+1": "1110111",
		"M-1": "1110010",
		"D+M": "1000010",
		"D-M": "1010011",
		"M-D": "1000111",
		"D&M": "1000000",
		"D|M": "1010101",
	}
	return compBinaryMap[comp]
}
