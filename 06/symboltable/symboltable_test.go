package symboltable

import "testing"

func TestContains(t *testing.T) {
	symbolTable := New()
	testCases := []struct {
		symbol   string
		contains bool
	}{
		{"R1", true}, {"SP", true}, {"HOGE", false},
	}
	for _, tt := range testCases {
		contains := symbolTable.Contains(tt.symbol)
		if contains != tt.contains {
			t.Fatalf("Contains(%s) should return %t ,got %t", tt.symbol, tt.contains, contains)
		}
	}
}
