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
	case *ast.DoStatement:
		return ce.CompileDoStatement(statementAst)
	case *ast.ReturnStatement:
		return ce.CompileReturnStatement(statementAst)
	case *ast.ClassStatement:
		return ce.CompileClassStatement(statementAst)
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
	if statementAst.Value != nil {
		ce.CompileExpression(statementAst.Value)
	} else {
		ce.WritePush(vmwriter.CONST, 0)
	}
	ce.WriteReturn()
	return nil
}

func (ce *CompilationEngine) CompileClassStatement(statementAst *ast.ClassStatement) error {
	_, subroutineDecList := statementAst.ClassVarDecList, statementAst.SubroutineDecList
	// for range varDecList {
	// }
	for _, subroutineDec := range subroutineDecList {
		ce.CompileSubroutineDecStatement(&subroutineDec)
	}
	return nil
}

func (ce *CompilationEngine) CompileSubroutineDecStatement(subroutineDecStmtAst *ast.SubroutineDecStatement) error {
	ce.WriteFunction(fmt.Sprintf("%s.%s", ce.ClassName, subroutineDecStmtAst.Name.Literal), 0)
	ce.StartSubroutine()
	ce.CompileParameterListStatement(subroutineDecStmtAst.ParameterList)
	ce.CompileSubroutineBodyStatement(subroutineDecStmtAst.SubroutineBody)
	return nil
}

func (ce *CompilationEngine) CompileSubroutineBodyStatement(subroutineBodyStmt *ast.SubroutineBodyStatement) {
	for _, varDecStmt := range subroutineBodyStmt.VarDecList {
		for _, identifier := range varDecStmt.Identifiers {
			ce.Define(identifier.Literal, varDecStmt.ValueType.Literal, symboltable.VAR)
		}
	}
	for _, stmt := range subroutineBodyStmt.Statements {
		ce.CompileStatement(stmt)
	}
}

func (ce *CompilationEngine) CompileParameterListStatement(parameterListStmtAst *ast.ParameterListStatement) error {
	for _, stmt := range parameterListStmtAst.ParameterList {
		// シンボルテーブルにArgumentを登録。
		ce.Define(stmt.Name, stmt.ValueType.Literal, symboltable.ARGUMENT)
	}
	return nil
}

func (ce *CompilationEngine) CompileLetStatement(letStatement *ast.LetStatement) error {
	ce.CompileExpression(letStatement.Value)

	varKind := ce.KindOf(letStatement.Name.Literal)
	indexOf := ce.IndexOf(letStatement.Name.Literal)
	switch varKind {
	case symboltable.ARGUMENT:
		ce.WritePop(vmwriter.ARG, indexOf)
		return nil
	case symboltable.STATIC:
		ce.WritePop(vmwriter.STATIC, indexOf)
		return nil
	case symboltable.FIELD:
		ce.WritePop(vmwriter.THIS, indexOf)
		return nil
	case symboltable.VAR:
		ce.WritePop(vmwriter.LOCAL, indexOf)
		return nil
	}
	return nil // TODO:Error,fmt.Errorf("Identifier ...")
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
	return ce.CompileTerm(singleExpressionAst.Value)
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
	case token.MINUS:
		{
			ce.WriteArithmetic(vmwriter.SUB)
			return nil
		}
	case token.EQ:
		{
			ce.WriteArithmetic(vmwriter.EQ)
			return nil
		}
	case token.AMP:
		{
			ce.WriteArithmetic(vmwriter.AND)
			return nil
		}
	case token.OR:
		{
			ce.WriteArithmetic(vmwriter.OR)
			return nil
		}
	case token.GT:
		{
			ce.WriteArithmetic(vmwriter.GT)
			return nil
		}
	case token.LT:
		{
			ce.WriteArithmetic(vmwriter.LT)
			return nil
		}
	case token.ASTERISK:
		{
			ce.WriteCall("Math.multiply", 2)
			return nil
		}
	case token.SLASH:
		{
			ce.WriteCall("Math.divide", 2)
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
	case *ast.StringConstTerm:
		return ce.CompileStringConstTerm(c)
	case *ast.IdentifierTerm:
		return ce.CompileIdentifierTerm(c)
	case *ast.SubroutineCallTerm:
		return ce.CompileSubroutineCallTerm(c)
	}
	return nil
}

func (ce *CompilationEngine) CompileIntergerConstTerm(intergerConstTerm *ast.IntergerConstTerm) error {
	ce.WritePush(vmwriter.CONST, int(intergerConstTerm.Value))
	return nil

}
func (ce *CompilationEngine) CompileSubroutineCallTerm(subroutineCallTerm *ast.SubroutineCallTerm) error {
	ce.CompileExpressionListStatement(subroutineCallTerm.ExpressionListStmt)
	ce.WriteCall(fmt.Sprintf("%s.%s", subroutineCallTerm.ClassName.String(), subroutineCallTerm.SubroutineName.String()), len(subroutineCallTerm.ExpressionListStmt.ExpressionList))
	return nil
}

func (ce *CompilationEngine) CompileBracketTerm(bracketTerm *ast.BracketTerm) error {
	return ce.CompileExpression(bracketTerm.Value)
}

func (ce *CompilationEngine) CompileIdentifierTerm(identifierTerm *ast.IdentifierTerm) error {
	varKind := ce.KindOf(identifierTerm.TokenLiteral())
	indexOf := ce.IndexOf(identifierTerm.TokenLiteral())
	switch varKind {
	case symboltable.ARGUMENT:
		ce.WritePush(vmwriter.ARG, indexOf)
		return nil
	case symboltable.STATIC:
		ce.WritePush(vmwriter.STATIC, indexOf)
		return nil
	case symboltable.FIELD:
		ce.WritePush(vmwriter.THIS, indexOf)
		return nil
	case symboltable.VAR:
		ce.WritePush(vmwriter.LOCAL, indexOf)
		return nil
	}
	return nil // TODO:Error,fmt.Errorf("Identifier ...")
}

func (ce *CompilationEngine) CompileStringConstTerm(stringConstTerm *ast.StringConstTerm) error {
	ce.WritePush(vmwriter.CONST, len(stringConstTerm.Value))
	ce.WriteCall("String.new", 1)
	for _, c := range stringConstTerm.Value {
		ce.WritePush(vmwriter.CONST, int(c))
		ce.WriteCall("String.appendChar", 2)
	}
	return nil
}

func (ce *CompilationEngine) CompilePrefixTerm(prefixTerm *ast.PrefixTerm) error {
	ce.CompileTerm(prefixTerm.Value)
	switch prefixTerm.Prefix {
	case token.MINUS:
		ce.WriteArithmetic(vmwriter.NEG)
		return nil
	case token.BANG:
		ce.WriteArithmetic(vmwriter.NOT)
		return nil
	}
	return fmt.Errorf("prefixTerm.Prefix should be '-' or '~'. But got %s", prefixTerm.Prefix)
}

func (ce *CompilationEngine) CompileDoStatement(doStatement *ast.DoStatement) error {
	ce.CompileExpressionListStatement(doStatement.ExpressionListStmt)
	ce.WriteCall(fmt.Sprintf("%s.%s", doStatement.ClassName.String(), doStatement.SubroutineName.String()), len(doStatement.ExpressionListStmt.ExpressionList))
	ce.WritePop(vmwriter.TEMP, 0)
	return nil
}

func (ce *CompilationEngine) CompileExpressionListStatement(expressionListStmt *ast.ExpressionListStatement) error {
	for i := range expressionListStmt.ExpressionList {
		ce.CompileExpression(expressionListStmt.ExpressionList[i])
	}
	return nil
}
