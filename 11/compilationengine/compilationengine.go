package compilationengine

import (
	"errors"
	"jack_compiler/ast"
	"jack_compiler/symboltable"
	"jack_compiler/vmwriter"
)

// CompilationEngine is struct
type CompilationEngine struct {
	*vmwriter.VMWriter
	*symboltable.SymbolTable
}

// New is initializer of compilation engine
func New(vm *vmwriter.VMWriter, st *symboltable.SymbolTable) *CompilationEngine {
	ce := &CompilationEngine{VMWriter: vm, SymbolTable: st}
	return ce
}

func (ce *CompilationEngine) CompileProgram(programAst *ast.Program) {
	for _, stmtAst := range programAst.Statements {
		ce.compileStatement(stmtAst)
	}
}

func (ce *CompilationEngine) compileStatement(statementAst ast.Statement) error {
	switch statementAst := statementAst.(type) {
	case *ast.VarDecStatement:
		return ce.compileVarDec(statementAst)
	default:
		return errors.New("statementAst type: %T is not valid")
	}
}

func (ce *CompilationEngine) compileVarDec(varDecAst *ast.VarDecStatement) error {
	ce.WriteIf("hoge")
	return nil
}
