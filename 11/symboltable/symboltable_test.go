package symboltable

import (
	"reflect"
	"testing"
)

func TestStartSubroutine(t *testing.T) {
	st := New()
	st.StartSubroutine()
	if st.Scope != SubroutineScope {
		t.Errorf("st.Scope not SubroutineScope. got %d", st.Scope)
	}
	if !reflect.DeepEqual(st.MethodScopeSymbolTable, map[string]Symbol{}) {
		t.Errorf("st.MethodScopeSymbolTable not map[string]Symbol{}. got %v", st.MethodScopeSymbolTable)
	}
}

func TestDefine(t *testing.T) {
	st := New()
	testCases := []struct {
		name    string
		varKind VarKind
		varType string
		Idx     int
	}{
		{"a", STATIC, "int", 0},
		{"b", FIELD, "int", 0},
		{"c", ARGUMENT, "int", 0},
		{"d", VAR, "int", 0},
		{"e", VAR, "int", 1},
	}
	for _, tt := range testCases {
		st.Define(tt.name, tt.varType, tt.varKind)
	}
	for _, tt := range testCases {
		symbol := Symbol{
			tt.name,
			tt.varKind,
			tt.varType,
			tt.Idx,
		}
		if !reflect.DeepEqual(st.ClassScopeSymbolTable[tt.name], symbol) {
			t.Errorf("st.ClassScopeSymbolTable not %v. got %v", symbol, st.ClassScopeSymbolTable)
		}
	}
}

// symbolTable := map[string]Symbol{"a": {"a", "static", "int", 0}}
