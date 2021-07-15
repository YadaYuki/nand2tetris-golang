package parser

import "testing"

func TestAdvance(t *testing.T) {
	p := New("sample")
	p.Advance()
	if p.currentCommandIdx != 1 {
		t.Fatalf("p.currentCommandIdx should be 1 , but got %d", p.currentCommandIdx)
	}
}

func TestHasMoreCommand(t *testing.T) {
	p := New(`sample
	hoge`)
	p.Advance()
	if !p.HasMoreCommand() {
		t.Fatal("p.HasMoreCommand should be true , but got false")
	}
	p.Advance()
	if p.HasMoreCommand() {
		t.Fatalf("p.HasMoreCommand should be false , but got true")
	}
}

func TestSkipWhiteSpace(t *testing.T) {
	testCases := []struct {
		parser                          *Parser
		readPositionAfterSkipWhiteSpace int
	}{
		{&Parser{commandStrList: []string{"input"}, currentCommandIdx: 0}, 0},
		{&Parser{commandStrList: []string{" input"}, currentCommandIdx: 0}, 1},
		{&Parser{commandStrList: []string{"   input"}, currentCommandIdx: 0}, 3},
		{&Parser{commandStrList: []string{"   \tinput"}, currentCommandIdx: 0}, 4},
	}
	for _, tt := range testCases {
		tt.parser.skipWhiteSpace()
		if tt.parser.readPosition != tt.readPositionAfterSkipWhiteSpace {
			t.Fatalf("parser.readPosition should be %d,got %d", tt.readPositionAfterSkipWhiteSpace, tt.parser.readPosition)
		}
		if tt.parser.commandStrList[tt.parser.currentCommandIdx][tt.parser.readPosition] != byte('i') {
			t.Fatalf("readChar should be `i`,got %c", tt.parser.commandStrList[tt.parser.currentCommandIdx][tt.parser.readPosition])
		}
	}
}
