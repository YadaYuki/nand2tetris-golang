package symboltable

type SymbolTable struct {
	ClassScopeSymbolTable  map[string]Symbol
	MethodScopeSymbolTable map[string]Symbol
	Scope                  int
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

func (SymbolTable *st) StartSubroutine() {
}

func (SymbolTable *st) Define(name string, varType string, varKind string) {
}

func (SymbolTable *st) VarCount(varKind string) int {
}

func (SymbolTable *st) KindOf(name string) VarKind {
}

func (SymbolTable *st) TypeOf(name string) string {
}

func (SymbolTable *st) IndexOf(name string) int {
}
