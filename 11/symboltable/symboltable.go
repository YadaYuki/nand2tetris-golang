package symboltable

import (
	"errors"
)

type SymbolTable struct {
	ClassScopeSymbolTable  map[string]Symbol
	MethodScopeSymbolTable map[string]Symbol
	Scope                  Scope
	currentStaticIdx       int
	currentFieldIdx        int
	currentArgumentIdx     int
	currentVarIdx          int
}

type Symbol struct {
	Name    string
	VarKind VarKind
	VarType string // not definite enum for this member. because varType is not only KeyWord.(className,FUnctionName...)
	Idx     int
}

type VarKind string

const (
	STATIC   VarKind = "static"
	FIELD    VarKind = "field"
	ARGUMENT VarKind = "argument"
	VAR      VarKind = "var"
	NONE     VarKind = "none"
)

type Scope int

const (
	SubroutineScope Scope = iota
	ClassScope      Scope = iota
)

func New() *SymbolTable {
	return &SymbolTable{Scope: ClassScope, MethodScopeSymbolTable: map[string]Symbol{}, ClassScopeSymbolTable: map[string]Symbol{}}
}

// TODO: TDD!
func (st *SymbolTable) StartSubroutine() {
	st.MethodScopeSymbolTable = map[string]Symbol{}
	st.currentArgumentIdx = 0
	st.currentVarIdx = 0
	st.Scope = SubroutineScope
}

func (st *SymbolTable) Define(name string, varType string, varKind VarKind) error {

	symbol := Symbol{}
	symbol.Name = name
	symbol.VarKind = varKind
	symbol.VarType = varType

	switch varKind {
	case STATIC:
		symbol.Idx = st.currentStaticIdx
		st.currentStaticIdx++
	case FIELD:
		symbol.Idx = st.currentFieldIdx
		st.currentFieldIdx++
	case ARGUMENT:
		symbol.Idx = st.currentArgumentIdx
		st.currentArgumentIdx++
	case VAR:
		symbol.Idx = st.currentVarIdx
		st.currentVarIdx++
	default:
		return errors.New("") // TODO: Add Error
	}

	switch st.Scope {
	case SubroutineScope:
		st.MethodScopeSymbolTable[name] = symbol
	case ClassScope:
		st.ClassScopeSymbolTable[name] = symbol
	default:
		return errors.New("") // TODO: Add Error
	}
	return nil
}

func (st *SymbolTable) VarCount(varKind VarKind) int {
	table := map[string]Symbol{}
	switch st.Scope {
	case SubroutineScope:
		table = st.MethodScopeSymbolTable
	case ClassScope:
		table = st.ClassScopeSymbolTable
	}
	varCount := 0
	for _, symbol := range table {
		if symbol.VarKind == varKind {
			varCount++
		}
	}
	return varCount
}

func (st *SymbolTable) KindOf(name string) VarKind {
	symbol, ok := st.MethodScopeSymbolTable[name]
	if ok {
		return symbol.VarKind
	}
	symbol, ok = st.ClassScopeSymbolTable[name]
	if ok {
		return symbol.VarKind
	}
	return ""
}

func (st *SymbolTable) TypeOf(name string) string {
	symbol, ok := st.MethodScopeSymbolTable[name]
	if ok {
		return symbol.VarType
	}
	symbol, ok = st.ClassScopeSymbolTable[name]
	if ok {
		return symbol.VarType
	}
	return ""
}

func (st *SymbolTable) IndexOf(name string) int {
	symbol, ok := st.MethodScopeSymbolTable[name]
	if ok {
		return symbol.Idx
	}
	symbol, ok = st.ClassScopeSymbolTable[name]
	if ok {
		return symbol.Idx
	}
	return -1
}
