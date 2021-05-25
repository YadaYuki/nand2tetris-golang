package symboltable

type SymbolTable struct {
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
