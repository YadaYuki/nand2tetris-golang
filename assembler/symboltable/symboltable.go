package symboltable

import (
	"fmt"
)

type SymbolTable struct {
	SymbolTableDict map[string]int
}

func getInitialSymbolTable() map[string]int {
	initialSymbolTable := map[string]int{"SP": 0, "LCL": 1, "ARG": 2, "THIS": 3, "THAT": 4, "SCREEN": 16384, "KBD": 24576}
	// initialize Register Address
	for i := 0; i < 16; i++ {
		initialSymbolTable[fmt.Sprintf("R%d", i)] = i
	}
	return initialSymbolTable
}

func New() *SymbolTable {
	initialSymbolTable := getInitialSymbolTable()
	return &SymbolTable{SymbolTableDict: initialSymbolTable}
}

func (st *SymbolTable) Contains(symbol string) bool {
	_, ok := st.SymbolTableDict[symbol]
	return ok
}

func (st *SymbolTable) GetAddress(symbol string) (int, error) {
	contains := st.Contains(symbol)
	if !contains {
		return -1, fmt.Errorf("%s is not Contained in symbolTable", symbol)
	}
	return st.SymbolTableDict[symbol], nil
}

func (st *SymbolTable) AddEntry(symbol string, address int) error {
	st.SymbolTableDict[symbol] = address
	return nil
}
