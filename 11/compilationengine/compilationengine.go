package compilationengine

import (
	"jack_compiler/parser"
	"jack_compiler/symboltable"
	"jack_compiler/vmwriter"
)

// CompilationEngine is struct
type CompilationEngine struct {
	parser *parser.Parser
	vm     *vmwriter.VMWriter
	st     *symboltable.SymbolTable
}

// New is initializer of compilation engine
func New(parser *parser.Parser, vm *vmwriter.VMWriter, st *symboltable.SymbolTable) *CompilationEngine {
	ce := &CompilationEngine{parser: parser, vm: vm, st: st}
	return ce
}
