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
	labelFlag int
}

// New is initializer of compilation engine
func New(className string, vm *vmwriter.VMWriter, st *symboltable.SymbolTable) *CompilationEngine {
	ce := &CompilationEngine{VMWriter: vm, SymbolTable: st, ClassName: className, labelFlag: 0}
	return ce
}

func (ce *CompilationEngine) CompileProgram(programAst *ast.Program) {
	for _, stmtAst := range programAst.Statements {
		ce.CompileStatement(stmtAst)
	}
}

func (ce *CompilationEngine) CompileStatement(statementAst ast.Statement) error {
	switch statementAst := statementAst.(type) {
	case *ast.DoStatement:
		return ce.CompileDoStatement(statementAst)
	case *ast.ReturnStatement:
		return ce.CompileReturnStatement(statementAst)
	case *ast.ClassStatement:
		return ce.CompileClassStatement(statementAst)
	case *ast.LetStatement:
		return ce.CompileLetStatement(statementAst)
	case *ast.IfStatement:
		return ce.CompileIfStatement(statementAst)
	case *ast.WhileStatement:
		return ce.CompileWhileStatement(statementAst)
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
	varDecList, subroutineDecList := statementAst.ClassVarDecList, statementAst.SubroutineDecList
	for _, varDec := range varDecList {
		varKind := symboltable.VarKind(varDec.Token.Literal)
		for _, ident := range varDec.Identifiers {
			ce.Define(ident.Literal, varDec.ValueType.Literal, varKind)
		}
	}

	for _, subroutineDec := range subroutineDecList {
		ce.CompileSubroutineDecStatement(&subroutineDec)
	}
	return nil
}

func (ce *CompilationEngine) CompileSubroutineDecStatement(subroutineDecStmtAst *ast.SubroutineDecStatement) error {
	switch subroutineDecStmtAst.Token.Literal {
	case string(token.CONSTRUCTOR):
		return ce.CompileSubroutineDecConstructorStatement(subroutineDecStmtAst)
	case string(token.METHOD):
		return ce.CompileSubroutineDecMethodStatement(subroutineDecStmtAst)
	case string(token.FUNCTION):
		return ce.CompileSubroutineDecFunctionStatement(subroutineDecStmtAst)
	}
	return nil // TODO:error
}

func (ce *CompilationEngine) CompileSubroutineDecConstructorStatement(subroutineDecStmtAst *ast.SubroutineDecStatement) error {
	// ローカル変数の数を計算する
	localVarCount := 0
	for _, varDec := range subroutineDecStmtAst.SubroutineBody.VarDecList {
		localVarCount += len(varDec.Identifiers)
	}
	// thisポインタをローカル変数として扱うため、その分加算する。
	localVarCount += 1
	ce.WriteFunction(fmt.Sprintf("%s.%s", ce.ClassName, subroutineDecStmtAst.Name.Literal), localVarCount)

	// フィールド変数分だけ、メモリを確保し → thisポインタにオブジェクトの先頭アドレスを格納する。
	fieldVarCount := ce.VarCount(symboltable.FIELD)
	ce.WritePush(vmwriter.CONST, fieldVarCount)
	ce.WriteCall("Memory.alloc", 1)
	ce.WritePop(vmwriter.LOCAL, 0)

	ce.StartSubroutine()
	ce.Define(string(token.THIS), ce.ClassName, symboltable.VAR)
	ce.CompileParameterListStatement(subroutineDecStmtAst.ParameterList)
	ce.CompileSubroutineBodyStatement(subroutineDecStmtAst.SubroutineBody)
	return nil
}

func (ce *CompilationEngine) CompileSubroutineDecMethodStatement(subroutineDecStmtAst *ast.SubroutineDecStatement) error {
	// ローカル変数の数を計算する
	localVarCount := 0
	for _, varDec := range subroutineDecStmtAst.SubroutineBody.VarDecList {
		localVarCount += len(varDec.Identifiers)
	}
	ce.WriteFunction(fmt.Sprintf("%s.%s", ce.ClassName, subroutineDecStmtAst.Name.Literal), localVarCount)

	ce.StartSubroutine()
	// thisを第一引数として定義
	ce.Define(string(token.THIS), ce.ClassName, symboltable.ARGUMENT)

	ce.CompileParameterListStatement(subroutineDecStmtAst.ParameterList)
	ce.CompileSubroutineBodyStatement(subroutineDecStmtAst.SubroutineBody)

	return nil
}

func (ce *CompilationEngine) CompileSubroutineDecFunctionStatement(subroutineDecStmtAst *ast.SubroutineDecStatement) error {
	// ローカル変数の数を計算する
	localVarCount := 0
	for _, varDec := range subroutineDecStmtAst.SubroutineBody.VarDecList {
		localVarCount += len(varDec.Identifiers)
	}
	ce.WriteFunction(fmt.Sprintf("%s.%s", ce.ClassName, subroutineDecStmtAst.Name.Literal), localVarCount)

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
	if letStatement.Idx != nil { // 配列の要素に対する代入であった場合。
		return ce.CompileLetArrayElementStatement(letStatement)
	}
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
		thisVarKind := ce.KindOf(string(token.THIS))
		thisIndexOf := ce.IndexOf(string(token.THIS))
		switch thisVarKind {
		case symboltable.ARGUMENT:
			ce.WritePush(vmwriter.ARG, thisIndexOf)
		case symboltable.VAR:
			ce.WritePush(vmwriter.LOCAL, thisIndexOf)
		default:
			return nil // TODO:Error,fmt.Errorf("Identifier ...")
		}
		ce.WritePop(vmwriter.POINTER, 0)
		ce.WritePop(vmwriter.THIS, indexOf)
		return nil
	case symboltable.VAR:
		ce.WritePop(vmwriter.LOCAL, indexOf)
		return nil
	}
	return nil // TODO:Error,fmt.Errorf("Identifier ...")
}

func (ce *CompilationEngine) CompileLetArrayElementStatement(letStatement *ast.LetStatement) error {
	// 配列の先頭アドレス(addr)をstackにpushする
	varKind := ce.KindOf(letStatement.Name.Literal)
	indexOf := ce.IndexOf(letStatement.Name.Literal)
	switch varKind {
	case symboltable.ARGUMENT:
		ce.WritePush(vmwriter.ARG, indexOf)
	// case symboltable.STATIC:
	// 	ce.WritePush(vmwriter.STATIC, indexOf)
	//
	// case symboltable.FIELD:
	// 	ce.WritePush(vmwriter.THIS, indexOf)
	//
	case symboltable.VAR:
		ce.WritePush(vmwriter.LOCAL, indexOf)
	default:
		return nil // TODO:error
	}
	// (addr + idx)
	ce.CompileExpression(letStatement.Idx)
	ce.WriteArithmetic(vmwriter.ADD)
	// NOTE: CompileExpressionで配列の値参照が生じた際、pointer 0の値が書き換えられてしまうので、代入先であるアドレスを一時的にtempに保存する。
	ce.WritePop(vmwriter.TEMP, 0)
	ce.CompileExpression(letStatement.Value)
	ce.WritePush(vmwriter.TEMP, 0)

	ce.WritePop(vmwriter.POINTER, 1)
	ce.WritePop(vmwriter.THAT, 0)
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
	case *ast.ArrayElementTerm:
		return ce.CompileArrayElementTerm(c)
	case *ast.PrefixTerm:
		return ce.CompilePrefixTerm(c)
	case *ast.KeywordConstTerm:
		return ce.CompileKeywordConstTerm(c)
	}
	return nil
}

func (ce *CompilationEngine) CompileIntergerConstTerm(intergerConstTerm *ast.IntergerConstTerm) error {
	ce.WritePush(vmwriter.CONST, int(intergerConstTerm.Value))
	return nil
}

func (ce *CompilationEngine) CompileSubroutineCallTerm(subroutineCallTerm *ast.SubroutineCallTerm) error {

	argumentNum := len(subroutineCallTerm.ExpressionListStmt.ExpressionList)
	// method呼び出しの場合は、インスタンスのアドレスを第一引数としてpushする
	if ce.IndexOf(subroutineCallTerm.ClassName.Literal) != -1 {
		typeOf := ce.TypeOf(subroutineCallTerm.ClassName.Literal)
		ce.CompileIdentifierTerm(&ast.IdentifierTerm{Token: subroutineCallTerm.ClassName, Value: subroutineCallTerm.ClassName.Literal})
		argumentNum += 1
		ce.CompileExpressionListStatement(subroutineCallTerm.ExpressionListStmt)
		ce.WriteCall(fmt.Sprintf("%s.%s", typeOf, subroutineCallTerm.SubroutineName.String()), argumentNum)
		return nil
	}
	ce.CompileExpressionListStatement(subroutineCallTerm.ExpressionListStmt)
	ce.WriteCall(fmt.Sprintf("%s.%s", subroutineCallTerm.ClassName.String(), subroutineCallTerm.SubroutineName.String()), argumentNum)
	return nil
}

func (ce *CompilationEngine) CompileBracketTerm(bracketTerm *ast.BracketTerm) error {
	return ce.CompileExpression(bracketTerm.Value)
}

func (ce *CompilationEngine) CompileArrayElementTerm(arrayElementTerm *ast.ArrayElementTerm) error {
	// 配列の先頭アドレスをpushする
	varKind := ce.KindOf(arrayElementTerm.TokenLiteral())
	indexOf := ce.IndexOf(arrayElementTerm.TokenLiteral())
	switch varKind {
	case symboltable.ARGUMENT:
		ce.WritePush(vmwriter.ARG, indexOf)
	// case symboltable.STATIC:
	// 	ce.WritePush(vmwriter.STATIC, indexOf)
	// case symboltable.FIELD:
	// 	ce.WritePush(vmwriter.THIS, indexOf)
	case symboltable.VAR:
		ce.WritePush(vmwriter.LOCAL, indexOf)
	default:
		return nil // TODO:Error
	}

	// [先頭アドレス + idx]にアクセス.
	ce.CompileExpression(arrayElementTerm.Idx)
	ce.WriteArithmetic(vmwriter.ADD)
	ce.WritePop(vmwriter.POINTER, 1)
	ce.WritePush(vmwriter.THAT, 0)

	return nil
}

func (ce *CompilationEngine) CompileIdentifierTerm(identifierTerm *ast.IdentifierTerm) error {
	varKind := ce.KindOf(identifierTerm.TokenLiteral())
	indexOf := ce.IndexOf(identifierTerm.TokenLiteral())
	switch varKind {
	case symboltable.ARGUMENT:
		ce.WritePush(vmwriter.ARG, indexOf)
		return nil
	// case symboltable.STATIC:
	// 	ce.WritePush(vmwriter.STATIC, indexOf)
	// 	return nil
	case symboltable.FIELD:
		thisVarKind := ce.KindOf(string(token.THIS))
		thisIndexOf := ce.IndexOf(string(token.THIS))
		switch thisVarKind {
		case symboltable.ARGUMENT:
			ce.WritePush(vmwriter.ARG, thisIndexOf)
		case symboltable.VAR:
			ce.WritePush(vmwriter.LOCAL, thisIndexOf)
		default:
			return nil // TODO:Error,fmt.Errorf("Identifier ...")
		}
		ce.WritePop(vmwriter.POINTER, 0)
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

func (ce *CompilationEngine) CompileKeywordConstTerm(keywordConstTerm *ast.KeywordConstTerm) error {
	switch keywordConstTerm.KeyWord {
	case token.NULL, token.FALSE:
		ce.WritePush(vmwriter.CONST, 0)
		return nil
	case token.TRUE:
		ce.WritePush(vmwriter.CONST, 1)
		ce.WriteArithmetic(vmwriter.NEG)
		return nil
	case token.THIS:
		varKind := ce.KindOf(string(token.THIS))
		indexOf := ce.IndexOf(string(token.THIS))
		switch varKind {
		case symboltable.VAR:
			ce.WritePush(vmwriter.LOCAL, indexOf)
			return nil
		case symboltable.ARGUMENT:
			ce.WritePush(vmwriter.ARG, indexOf)
			return nil
		}
	}
	return nil // TODO: Error
}

func (ce *CompilationEngine) CompileDoStatement(doStatement *ast.DoStatement) error {
	ce.CompileExpressionListStatement(doStatement.ExpressionListStmt)
	ce.WriteCall(fmt.Sprintf("%s.%s", doStatement.ClassName.String(), doStatement.SubroutineName.String()), len(doStatement.ExpressionListStmt.ExpressionList))
	ce.WritePop(vmwriter.TEMP, 0)
	return nil
}

func (ce *CompilationEngine) CompileIfStatement(ifStatement *ast.IfStatement) error {
	ce.incrementLabelFlag()
	ELSE_LABEL, ENDIF_LABEL := fmt.Sprintf("ELSE%d", ce.labelFlag), fmt.Sprintf("ENDIF%d", ce.labelFlag)
	ce.CompileExpression(ifStatement.Condition)
	ce.WriteArithmetic(vmwriter.NOT)

	if ifStatement.Alternative == nil {
		ce.WriteIf(ENDIF_LABEL) // 条件式がfalseであった場合は、ENDIFにjumpする
		for _, stmt := range ifStatement.Consequence.Statements {
			ce.CompileStatement(stmt)
		}
		ce.WriteLabel(ENDIF_LABEL)
		return nil
	}

	ce.WriteIf(ELSE_LABEL) // 条件式がfalseであった場合は、ELSEにjumpする
	for _, stmt := range ifStatement.Consequence.Statements {
		ce.CompileStatement(stmt)
	}
	ce.WriteGoto(ENDIF_LABEL)
	ce.WriteLabel(ELSE_LABEL)
	for _, stmt := range ifStatement.Alternative.Statements {
		ce.CompileStatement(stmt)
	}
	ce.WriteLabel(ENDIF_LABEL)
	return nil
}

func (ce *CompilationEngine) CompileWhileStatement(whileStatement *ast.WhileStatement) error {
	ce.incrementLabelFlag()
	WHILE_LOOP_LABEL, WHILE_END_LABEL := fmt.Sprintf("WHILELOOP%d", ce.labelFlag), fmt.Sprintf("WHILEEND%d", ce.labelFlag)
	ce.WriteLabel(WHILE_LOOP_LABEL)
	ce.CompileExpression(whileStatement.Condition)
	ce.WriteArithmetic(vmwriter.NOT)
	ce.WriteIf(WHILE_END_LABEL) // 条件式がfalseであった場合は、WHILE_END_LABELにjumpする

	for _, stmt := range whileStatement.Statements.Statements {
		ce.CompileStatement(stmt)
	}
	ce.WriteGoto(WHILE_LOOP_LABEL)
	ce.WriteLabel(WHILE_END_LABEL)
	return nil
}

func (ce *CompilationEngine) CompileExpressionListStatement(expressionListStmt *ast.ExpressionListStatement) error {
	for i := range expressionListStmt.ExpressionList {
		ce.CompileExpression(expressionListStmt.ExpressionList[i])
	}
	return nil
}

func (ce *CompilationEngine) incrementLabelFlag() {
	ce.labelFlag++
}
