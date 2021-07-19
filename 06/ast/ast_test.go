package ast

import (
	"assembly/value"
	"fmt"
	"testing"
)

func TestACommand(t *testing.T) {
	val := 10000
	aCommand := ACommand{Value: val, ValueStr: fmt.Sprintf("%d", val)}
	if aCommand.String() != "@10000"+value.NEW_LINE {
		t.Fatalf("aCommand.String() should be %s, got %s ", "@10000"+value.NEW_LINE, aCommand.String())
	}
}

func TestCCommand(t *testing.T) {
	testCases := []struct {
		command    CCommand
		commandStr string
	}{
		{CCommand{Comp: "-1", Dest: "M"}, "M=-1" + value.NEW_LINE},
		{CCommand{Comp: "D", Jump: "JMP"}, "D;JMP" + value.NEW_LINE},
		{CCommand{Comp: "D|A", Dest: "AM", Jump: "JMP"}, "AM=D|A;JMP" + value.NEW_LINE},
	}
	for _, tt := range testCases {
		if tt.commandStr != tt.command.String() {
			t.Fatalf("cCommand.String() should be %s, got %s ", tt.commandStr, tt.command.String())
		}
	}
}
