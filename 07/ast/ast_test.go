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

func TestPushCommandString(t *testing.T) {
	testCases := []struct {
		command    *PushCommand
		commandStr string
	}{
		{&PushCommand{PUSH, ARGUMENT, 4}, "push argument 4"},
		{&PushCommand{PUSH, LOCAL, 111}, "push local 111"},
		{&PushCommand{PUSH, THIS, 12}, "push this 12"},
	}
	for _, tt := range testCases {
		if tt.commandStr != tt.command.String() {
			t.Fatalf("command.String() should be %s , but got %s", tt.commandStr, tt.command.String())
		}
	}
}
