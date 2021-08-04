package compilationengine

import (
	"errors"
	"jack_compiler/ast"
	"jack_compiler/symboltable"
	"jack_compiler/token"
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
		ce.CompileStatement(stmtAst)
	}
}

func (ce *CompilationEngine) CompileStatement(statementAst ast.Statement) error {
	switch statementAst := statementAst.(type) {
	case *ast.VarDecStatement:
		return ce.CompileVarDec(statementAst)
	default:
		return errors.New("statementAst type: %T is not valid")
	}
}

func (ce *CompilationEngine) CompileVarDec(varDecAst *ast.VarDecStatement) error {
	ce.WriteIf("hoge")
	return nil
}

func (ce *CompilationEngine) CompileExpression(expressionAst ast.Expression) error {
	switch c := expressionAst.(type) {
	case *ast.SingleExpression:
		ce.CompileSingleExpression(c)
	case *ast.InfixExpression:
		ce.CompileInfixExpression(c)
	}
	return nil
}

func (ce *CompilationEngine) CompileSingleExpression(singleExpressionAst *ast.SingleExpression) error {
	ce.CompileTerm(singleExpressionAst.Value)
	return nil
}

func (ce *CompilationEngine) CompileInfixExpression(infixExpressionAst *ast.InfixExpression) error {
	ce.CompileTerm(infixExpressionAst.Left)
	ce.CompileTerm(infixExpressionAst.Right)
	ce.WriteArithmetic(ce.getArithmeticCommand(token.Symbol(infixExpressionAst.Operator.Literal)))
	return nil
}

func (ce *CompilationEngine) getArithmeticCommand(symbol token.Symbol) vmwriter.Command {
	switch symbol {
	case token.PLUS:
		return vmwriter.ADD
		// case token.ASTERISK:
		// 	return vmwriter.
	}
	return ""
}

func (ce *CompilationEngine) CompileTerm(termAst ast.Term) error {
	switch c := termAst.(type) {
	case *ast.IntergerConstTerm:
		return ce.CompileIntergerConstTerm(c)
	}
	return nil
}

func (ce *CompilationEngine) CompileIntergerConstTerm(intergerConstTerm *ast.IntergerConstTerm) error {
	ce.WritePush(vmwriter.CONST, int(intergerConstTerm.Value))
	return nil
}
