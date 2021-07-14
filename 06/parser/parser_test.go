package parser

import "testing"

func TestAdvance(t *testing.T) {
	p := New("sample")
	p.Advance()
	if p.currentCommandIdx != 1 {
		t.Fatalf("p.currentCommandIdx should be 1 , but got %d", p.currentCommandIdx)
	}
}

func TestHasMoreToken(t *testing.T) {
	p := New(`sample
	hoge`)
	p.Advance()
	if !p.HasMoreToken() {
		t.Fatal("p.HasMoreToken should be true , but got false")
	}
	p.Advance()
	if p.HasMoreToken() {
		t.Fatalf("p.HasMoreToken should be false , but got true")
	}
}
