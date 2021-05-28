package symboltable

type SymbolTable struct {
	ClassScopeSymbolTable  map[string]Symbol
	MethodScopeSymbolTable map[string]Symbol
	Scope                  Scope
	CurrentIdx             int
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
	return &SymbolTable{}
}

// TODO: TDD!
func (st *SymbolTable) StartSubroutine() {
	st.MethodScopeSymbolTable = map[string]Symbol{}
	st.Scope = SubroutineScope
}

func (st *SymbolTable) Define(name string, varType string, varKind string) {

}

func (st *SymbolTable) VarCount(varKind string) int {
	return -1
}

func (st *SymbolTable) KindOf(name string) VarKind {
	return STATIC
}

func (st *SymbolTable) TypeOf(name string) string {
	return ""
}

func (st *SymbolTable) IndexOf(name string) int {
	return -1
}
