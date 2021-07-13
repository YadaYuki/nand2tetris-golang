package code

import (
	"assembly/ast"
	// "assembly/value"
	"testing"
)

func TestBinary(t *testing.T) {
	testCases := []struct {
		command   ast.Command
		binaryStr string
	}{
		{&ast.ACommand{Value: 100}, "0000000001100100"},
	}
	for _, tt := range testCases {
		binaryStr := Binary(tt.command)
		if tt.binaryStr != binaryStr {
			t.Fatalf("binaryStr should be %s, got %s ", tt.binaryStr, binaryStr)
		}
	}
}
