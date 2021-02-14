package symbol_table

import (
	"errors"
)

var symbolTable = map[string]int{}

// CurrentAddress is current ROM Address. It stores custom variable
// var CurrentAddress = 1024

// AddEntry add symbol to symbol table
func AddEntry(symbol string, address int) {
	symbolTable[symbol] = address
}

// Contains returns whether symbol table has address correspond to symbol
func Contains(symbol string) bool {
	if address := symbolTable[symbol]; address == 0 {
		return false
	}
	return true
}

// GetAddress returns address(int) of symbol
func GetAddress(symbol string) (address int, err error) {
	if symbolTableHasAddress := Contains(symbol); symbolTableHasAddress == false {
		return 0, errors.New("symbol table does not have the address")
	}
	return symbolTable[symbol], nil
}
