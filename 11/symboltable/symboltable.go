package symboltable

type SymbolTable struct {
}

type Symbol struct {
	Name    string
	VarKind VarKind
	VarType string
	Idx     int
}

type VarKind string

const (
	STATIC   VarKind = "static"
	FIELD    VarKind = "field"
	ARGUMENT VarKind = "argument"
	VAR      VarKind = "var"
)

func New() *SymbolTable {
	return &SymbolTable{}
}
