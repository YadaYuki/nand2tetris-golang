package compilationengine

import (
	"errors"
	"fmt"
	"jack_compiler/ast"
	"jack_compiler/symboltable"
	"jack_compiler/token"
	"jack_compiler/vmwriter"
)

// CompilationEngine is struct
type CompilationEngine struct {
	*vmwriter.VMWriter
	*symboltable.SymbolTable
	ClassName string
}

// New is initializer of compilation engine
func New(className string, vm *vmwriter.VMWriter, st *symboltable.SymbolTable) *CompilationEngine {
	ce := &CompilationEngine{VMWriter: vm, SymbolTable: st, ClassName: className}
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

func (ce *CompilationEngine) CompileInit() {
	jackBasicLibraries := []string{
		"Math", "Output", "Keyboard", "Memory", "Screen", "Sys",
	}
	for i := range jackBasicLibraries {
		ce.WriteCall(fmt.Sprintf("%s.init", jackBasicLibraries[i]), 0)
	}
}

func (ce *CompilationEngine) CompileVarDec(varDecAst *ast.VarDecStatement) error {
	return nil
}

func (ce *CompilationEngine) CompileReturnStatement(statementAst *ast.ReturnStatement) error {
	ce.WriteReturn()
	return nil
}

func (ce *CompilationEngine) CompileExpression(expressionAst ast.Expression) error {
	switch c := expressionAst.(type) {
	case *ast.SingleExpression:
		return ce.CompileSingleExpression(c)
	case *ast.InfixExpression:
		return ce.CompileInfixExpression(c)
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
	switch token.Symbol(infixExpressionAst.Operator.Literal) {
	case token.PLUS:
		{
			ce.WriteArithmetic(vmwriter.ADD)
			return nil
		}
	case token.ASTERISK:
		{
			ce.WriteCall("Math.multiply", 2) // TODO: add mul logic in VM Layer.
			return nil
		}
	}
	return nil
}

func (ce *CompilationEngine) CompileTerm(termAst ast.Term) error {
	switch c := termAst.(type) {
	case *ast.IntergerConstTerm:
		return ce.CompileIntergerConstTerm(c)
	case *ast.BracketTerm:
		return ce.CompileBracketTerm(c)
	}
	return nil
}

func (ce *CompilationEngine) CompileIntergerConstTerm(intergerConstTerm *ast.IntergerConstTerm) error {
	ce.WritePush(vmwriter.CONST, int(intergerConstTerm.Value))
	return nil
}

func (ce *CompilationEngine) CompileBracketTerm(bracketTerm *ast.BracketTerm) error {
	return ce.CompileExpression(bracketTerm.Value)
}

func (ce *CompilationEngine) CompileDoStatement(doStatement *ast.DoStatement) error {
	expressionListStmt := doStatement.ExpressionListStmt
	for i := range expressionListStmt.ExpressionList {
		ce.CompileExpression(expressionListStmt.ExpressionList[i])
	}
	ce.WriteCall(fmt.Sprintf("%s.%s", doStatement.ClassName.String(), doStatement.VarName.String()), len(expressionListStmt.ExpressionList))
	ce.WritePop(vmwriter.TEMP, 0)
	return nil
}
