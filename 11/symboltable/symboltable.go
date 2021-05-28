package symboltable

import (
	"errors"
)

type SymbolTable struct {
	ClassScopeSymbolTable  map[string]Symbol
	MethodScopeSymbolTable map[string]Symbol
	Scope                  Scope
	CurrentStaticIdx       int
	CurrentFieldIdx        int
	CurrentArgumentIdx     int
	CurrentVarIdx          int
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
	st.Scope = SubroutineScope
}

func (st *SymbolTable) Define(name string, varType string, varKind VarKind) error {

	symbol := Symbol{}
	symbol.Name = name
	symbol.VarKind = varKind
	symbol.VarType = varType

	switch varKind {
	case STATIC:
		symbol.Idx = st.CurrentStaticIdx
		st.CurrentStaticIdx++
	case FIELD:
		symbol.Idx = st.CurrentFieldIdx
		st.CurrentFieldIdx++
	case ARGUMENT:
		symbol.Idx = st.CurrentArgumentIdx
		st.CurrentArgumentIdx++
	case VAR:
		symbol.Idx = st.CurrentVarIdx
		st.CurrentVarIdx++
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
	table := map[string]Symbol{}
	switch st.Scope {
	case SubroutineScope:
		table = st.MethodScopeSymbolTable
	case ClassScope:
		table = st.ClassScopeSymbolTable
	}
	symbol, ok := table[name]
	if !ok {
		return ""
	}
	return symbol.VarKind
}

func (st *SymbolTable) TypeOf(name string) string {
	table := map[string]Symbol{}
	switch st.Scope {
	case SubroutineScope:
		table = st.MethodScopeSymbolTable
	case ClassScope:
		table = st.ClassScopeSymbolTable
	}
	symbol, ok := table[name]
	if !ok {
		return ""
	}
	return symbol.VarType
}

func (st *SymbolTable) IndexOf(name string) int {
	return -1
}
