package symboltable

import "fmt"

type SymbolTable struct {
	SymbolTableDict      map[string]int
	CurrentVariableCount int
}

func New() *SymbolTable {
	initialSymbolTable := map[string]int{"SP": 0, "LCL": 1, "ARG": 2, "THIS": 3, "THAT": 4, "SCREEN": 16384, "KBD": 24576}
	// initialize Register Address
	for i := 0; i < 16; i++ {
		initialSymbolTable[fmt.Sprintf("R%d", i)] = i
	}
	return &SymbolTable{SymbolTableDict: initialSymbolTable, CurrentVariableCount: 0}
}
