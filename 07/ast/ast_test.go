package ast

import "testing"

func TestArithmeticCommandString(t *testing.T) {
	testCases := []struct {
		command    *ArithmeticCommand
		commandStr string
	}{
		{&ArithmeticCommand{ADD}, "add"},
		{&ArithmeticCommand{NEG}, "neg"},
		{&ArithmeticCommand{AND}, "and"},
	}
	for _, tt := range testCases {
		if tt.commandStr != tt.command.String() {
			t.Fatalf("command.String() should be %s , but got %s", tt.commandStr, tt.command.String())
		}
	}
}
