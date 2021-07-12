package ast

import (
	"assembly/value"
	"testing"
)

func TestACommand(t *testing.T) {
	val := 10000
	aCommand := ACommand{value: val}
	if aCommand.String() != "@10000"+value.NEW_LINE {
		t.Fatalf("aCommand.String() should be %s, got %s ", "@10000"+value.NEW_LINE, aCommand.String())
	}
}

func TestCCommand(t *testing.T) {
	testCases := []struct {
		command    CCommand
		commandStr string
	}{
		{CCommand{comp: "-1", dest: "M"}, "M=-1" + value.NEW_LINE},
		{CCommand{comp: "D", jump: "JMP"}, "D;JMP" + value.NEW_LINE},
		{CCommand{comp: "D|A", dest: "AM", jump: "JMP"}, "AM=D|A;JMP" + value.NEW_LINE},
	}
	for _, tt := range testCases {
		if tt.commandStr != tt.command.String() {
			t.Fatalf("cCommand.String() should be %s, got %s ", tt.commandStr, tt.command.String())
		}
	}
}
