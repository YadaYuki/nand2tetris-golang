package code

import (
	"assembler/ast"
	// "assembler/value"
	"testing"
)

func TestBinary(t *testing.T) {
	testCases := []struct {
		command   ast.Command
		binaryStr string
	}{
		{&ast.ACommand{Value: 100}, "0000000001100100"},
		{&ast.CCommand{Comp: "A", Dest: "D"}, "1110110000010000"},
		{&ast.CCommand{Comp: "M", Dest: "D"}, "1110001100001000"},
		{&ast.CCommand{Comp: "D|A", Dest: "AM", Jump: "JMP"}, "1110010101101111"},
	}
	for _, tt := range testCases {
		binaryStr := Binary(tt.command)
		if tt.binaryStr != binaryStr {
			t.Fatalf("binaryStr should be %s, got %s ", tt.binaryStr, binaryStr)
		}
	}
}
