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

func TestGetAddress(t *testing.T) {
	symbolTable := New()
	testCases := []struct {
		symbol  string
		address int
	}{
		{"R1", 1}, {"SP", 0}, {"HOGE", -1},
	}
	for _, tt := range testCases {
		address, _ := symbolTable.GetAddress(tt.symbol)
		if address != tt.address {
			t.Fatalf("%s's address should be %d, got %d", tt.symbol, tt.address, address)
		}
	}
}
