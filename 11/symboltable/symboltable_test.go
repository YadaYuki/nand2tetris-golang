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
