package ast

import "testing"

func TestArithmeticCommandString(t *testing.T) {
	testCases := []struct {
		command    *ArithmeticCommand
		commandStr string
	}{
		{&ArithmeticCommand{C_ARITHMETIC, ADD}, "add"},
		{&ArithmeticCommand{C_ARITHMETIC, NEG}, "neg"},
		{&ArithmeticCommand{C_ARITHMETIC, AND}, "and"},
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
		{&PushCommand{C_PUSH, PUSH, ARGUMENT, 4}, "push argument 4"},
		{&PushCommand{C_PUSH, PUSH, LOCAL, 111}, "push local 111"},
		{&PushCommand{C_PUSH, PUSH, THIS, 12}, "push this 12"},
	}
	for _, tt := range testCases {
		if tt.commandStr != tt.command.String() {
			t.Fatalf("command.String() should be %s , but got %s", tt.commandStr, tt.command.String())
		}
	}
}

func TestPopCommandString(t *testing.T) {
	testCases := []struct {
		command    *PopCommand
		commandStr string
	}{
		{&PopCommand{C_POP, POP, ARGUMENT, 4}, "pop argument 4"},
		{&PopCommand{C_POP, POP, LOCAL, 111}, "pop local 111"},
		{&PopCommand{C_POP, POP, THIS, 12}, "pop this 12"},
	}
	for _, tt := range testCases {
		if tt.commandStr != tt.command.String() {
			t.Fatalf("command.String() should be %s , but got %s", tt.commandStr, tt.command.String())
		}
	}
}

func TestLabelCommandString(t *testing.T) {
	testCases := []struct {
		command    *LabelCommand
		commandStr string
	}{
		{&LabelCommand{C_LABEL, LABEL, "IF_ELSE"}, "label IF_ELSE"},
	}
	for _, tt := range testCases {
		if tt.commandStr != tt.command.String() {
			t.Fatalf("command.String() should be %s , but got %s", tt.commandStr, tt.command.String())
		}
	}
}

func TestGotoCommandString(t *testing.T) {
	testCases := []struct {
		command    *GotoCommand
		commandStr string
	}{
		{&GotoCommand{C_LABEL, GOTO, "IF_ELSE"}, "goto IF_ELSE"},
	}
	for _, tt := range testCases {
		if tt.commandStr != tt.command.String() {
			t.Fatalf("command.String() should be %s , but got %s", tt.commandStr, tt.command.String())
		}
	}
}

func TestIfCommandString(t *testing.T) {
	testCases := []struct {
		command    *IfCommand
		commandStr string
	}{
		{&IfCommand{C_IF, IF_GOTO, "IF_ELSE"}, "if-goto IF_ELSE"},
	}
	for _, tt := range testCases {
		if tt.commandStr != tt.command.String() {
			t.Fatalf("command.String() should be %s , but got %s", tt.commandStr, tt.command.String())
		}
	}
}

func TestCallCommandString(t *testing.T) {
	testCases := []struct {
		command    *CallCommand
		commandStr string
	}{
		{&CallCommand{C_CALL, CALL, "func", 10}, "call func 10"},
	}
	for _, tt := range testCases {
		if tt.commandStr != tt.command.String() {
			t.Fatalf("command.String() should be %s , but got %s", tt.commandStr, tt.command.String())
		}
	}
}

func TestFunctionCommandString(t *testing.T) {
	testCases := []struct {
		command    *FunctionCommand
		commandStr string
	}{
		{&FunctionCommand{C_FUNCTION, FUNCTION, "func", 10}, "function func 10"},
	}
	for _, tt := range testCases {
		if tt.commandStr != tt.command.String() {
			t.Fatalf("command.String() should be %s , but got %s", tt.commandStr, tt.command.String())
		}
	}
}

func TestReturnCommandString(t *testing.T) {
	testCases := []struct {
		command    *ReturnCommand
		commandStr string
	}{
		{&ReturnCommand{C_RETURN, RETURN}, "return"},
	}
	for _, tt := range testCases {
		if tt.commandStr != tt.command.String() {
			t.Fatalf("command.String() should be %s , but got %s", tt.commandStr, tt.command.String())
		}
	}
}
