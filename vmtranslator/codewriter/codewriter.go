package codewriter

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"vmtranslator/ast"
	"vmtranslator/value"
)

type CodeWriter struct {
	Filename    string
	Assembly    []byte
	VmClassName string
}

func New(filename string) *CodeWriter {
	return &CodeWriter{Filename: filename, Assembly: []byte{}}
}

func (codeWriter *CodeWriter) SetVmClassName(vmClassName string) {
	codeWriter.VmClassName = vmClassName
}

func (codeWriter *CodeWriter) Close() {
	err := ioutil.WriteFile(codeWriter.Filename, codeWriter.Assembly, 0644)
	if err != nil {
		panic(err)
	}
}

func (codeWriter *CodeWriter) WriteInit() error {
	callInitAssembly, err := codeWriter.getInitAssembly()
	if err != nil {
		return err
	}
	codeWriter.writeAssembly(callInitAssembly)
	return nil
}

func (codeWriter *CodeWriter) getInitAssembly() (string, error) {
	assembly := ""
	// SP = 256
	assembly += "@256" + value.NEW_LINE + "D=A" + value.NEW_LINE // set 256 to D
	assembly += "@SP" + value.NEW_LINE + "M=D" + value.NEW_LINE  // set D to M
	// call Sys.init 0
	callInitCommand := &ast.CallCommand{Command: ast.C_CALL, Symbol: ast.CALL, FunctionName: "Sys.init", NumArgs: 0}
	callInitAssembly, err := codeWriter.getCallAssembly(callInitCommand)
	if err != nil {
		return "", err
	}
	assembly += callInitAssembly
	return assembly, nil
}

func (codeWriter *CodeWriter) WritePushPop(command ast.MemoryAccessCommand) error {
	var assembly string
	switch c := command.(type) {
	case *ast.PushCommand:
		pushAssembly, err := codeWriter.getPushAssembly(c)
		if err != nil {
			return err
		}
		assembly = pushAssembly
	case *ast.PopCommand:
		popAssembly, err := codeWriter.getPopAssembly(c)
		if err != nil {
			return err
		}
		assembly = popAssembly
	}
	codeWriter.writeAssembly(assembly)
	return nil
}

func (codeWriter *CodeWriter) WriteArithmetic(command *ast.ArithmeticCommand) error {
	arithmeticAssembly, err := codeWriter.getArithmeticAssembly(command)
	if err != nil {
		return err
	}
	codeWriter.writeAssembly(arithmeticAssembly)
	return nil
}

func (codeWriter *CodeWriter) WriteLabel(command *ast.LabelCommand) error {
	labelAssembly, err := codeWriter.getLabelAssembly(command)
	if err != nil {
		return err
	}
	codeWriter.writeAssembly(labelAssembly)
	return nil
}

func (codeWrite *CodeWriter) getLabelAssembly(command *ast.LabelCommand) (string, error) {
	assembly := fmt.Sprintf("(%s)", command.LabelName) + value.NEW_LINE
	return assembly, nil
}

func (codeWriter *CodeWriter) WriteGoto(command *ast.GotoCommand) error {
	gotoAssembly, err := codeWriter.getGotoAssembly(command)
	if err != nil {
		return err
	}
	codeWriter.writeAssembly(gotoAssembly)
	return nil
}

func (codeWrite *CodeWriter) getGotoAssembly(command *ast.GotoCommand) (string, error) {
	assembly := fmt.Sprintf("@%s", command.LabelName) + value.NEW_LINE + "0;JMP" + value.NEW_LINE // jump to label
	return assembly, nil
}

func (codeWriter *CodeWriter) WriteIf(command *ast.IfCommand) error {
	ifAssembly, err := codeWriter.getIfAssembly(command)
	if err != nil {
		return err
	}
	codeWriter.writeAssembly(ifAssembly)
	return nil
}

func (codeWrite *CodeWriter) getIfAssembly(command *ast.IfCommand) (string, error) {
	assembly := ""
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE                                 // decrement SP
	assembly += "A=M" + value.NEW_LINE + "D=M" + value.NEW_LINE                                   // set RAM[SP] to D
	assembly += fmt.Sprintf("@%s", command.LabelName) + value.NEW_LINE + "D;JNE" + value.NEW_LINE // if D == RAM[SP] != 0 then jump to Label else continue
	return assembly, nil
}

func (codeWriter *CodeWriter) WriteFunction(command *ast.FunctionCommand) error {
	functionAssembly, err := codeWriter.getFunctionAssembly(command)
	if err != nil {
		return err
	}
	codeWriter.writeAssembly(functionAssembly)
	return nil
}

func (codeWriter *CodeWriter) WriteReturn(command *ast.ReturnCommand) error {
	returnAssembly, err := codeWriter.getReturnAssembly(command)
	if err != nil {
		return err
	}
	codeWriter.writeAssembly(returnAssembly)
	return nil
}

func (codeWriter *CodeWriter) getReturnAssembly(command *ast.ReturnCommand) (string, error) {
	assembly := ""
	// FRAME = LCL
	assembly += "@LCL" + value.NEW_LINE + "D=M" + value.NEW_LINE   // set LCL to D
	assembly += "@FRAME" + value.NEW_LINE + "M=D" + value.NEW_LINE // set D to FRAME
	// RET = *(FRAME - 5)
	assembly += "@5" + value.NEW_LINE + "D=A" + value.NEW_LINE + "@FRAME" + value.NEW_LINE + "D=M-D" + value.NEW_LINE // set FRAME - 5 to D
	assembly += "A=D" + value.NEW_LINE + "D=M" + value.NEW_LINE                                                       // set *(FRAME-5) == RAM[FRAME-5] to D
	assembly += "@RETURN" + value.NEW_LINE + "M=D" + value.NEW_LINE                                                   // set D to RETURN
	popArgZeroCommand := &ast.PopCommand{Comamnd: ast.C_POP, Segment: ast.ARGUMENT, Index: 0}
	// *ARG = POP
	assembly += codeWriter.getMemoryAccessPopAssembly(popArgZeroCommand)
	// SP = ARG + 1
	assembly += "@ARG" + value.NEW_LINE + "D=M" + value.NEW_LINE + "D=D+1" + value.NEW_LINE // set ARG+1 to D
	assembly += "@SP" + value.NEW_LINE + "M=D" + value.NEW_LINE                             // set ARG+1 to D
	// THAT = *(FRAME-1)
	assembly += "@FRAME" + value.NEW_LINE + "D=M" + value.NEW_LINE + "D=M-1" + value.NEW_LINE // set FRAME-1 to D
	assembly += "A=D" + value.NEW_LINE + "D=M" + value.NEW_LINE                               // set *(FRAME-1) to D
	assembly += "@THAT" + value.NEW_LINE + "M=D" + value.NEW_LINE                             // set D to THAT
	// THIS = *(FRAME-2)
	assembly += "@FRAME" + value.NEW_LINE + "D=M" + value.NEW_LINE + "@2" + value.NEW_LINE + "D=D-A" + value.NEW_LINE // set FRAME-2 to D
	assembly += "A=D" + value.NEW_LINE + "D=M" + value.NEW_LINE                                                       // set *(FRAME-2) to D
	assembly += "@THIS" + value.NEW_LINE + "M=D" + value.NEW_LINE                                                     // set D to THIS
	// ARG = *(FRAME-3)
	assembly += "@FRAME" + value.NEW_LINE + "D=M" + value.NEW_LINE + "@3" + value.NEW_LINE + "D=D-A" + value.NEW_LINE // set FRAME-3 to D
	assembly += "A=D" + value.NEW_LINE + "D=M" + value.NEW_LINE                                                       // set *(FRAME-3) to D
	assembly += "@ARG" + value.NEW_LINE + "M=D" + value.NEW_LINE                                                      // set D to ARG
	// LCL = *(FRAME-4)
	assembly += "@FRAME" + value.NEW_LINE + "D=M" + value.NEW_LINE + "@4" + value.NEW_LINE + "D=D-A" + value.NEW_LINE // set FRAME-4 to D
	assembly += "A=D" + value.NEW_LINE + "D=M" + value.NEW_LINE                                                       // set *(FRAME-4) to D
	assembly += "@LCL" + value.NEW_LINE + "M=D" + value.NEW_LINE                                                      // set D to LCL
	// goto RETURN
	assembly += "@RETURN" + value.NEW_LINE + "A=M" + value.NEW_LINE + "0;JMP" + value.NEW_LINE
	return assembly, nil
}

func (codeWriter *CodeWriter) getCallAssembly(command *ast.CallCommand) (string, error) {
	assembly := ""
	returnAddressFlag := strconv.Itoa(rand.Intn(1000000))
	returnLabel := "RETURN" + returnAddressFlag
	//push return-address
	assembly += fmt.Sprintf("@%s", returnLabel) + value.NEW_LINE + "D=A" + value.NEW_LINE //set  Return Address to D
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE + "M=D" + value.NEW_LINE  // set D to RAM[SP]
	assembly += "@SP" + value.NEW_LINE + "M=M+1" + value.NEW_LINE                         // increment SP
	//push LCL
	assembly += "@LCL" + value.NEW_LINE + "A=M" + value.NEW_LINE + "D=A" + value.NEW_LINE // set LCL to D
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE + "M=D" + value.NEW_LINE  // set D to RAM[SP]
	assembly += "@SP" + value.NEW_LINE + "M=M+1" + value.NEW_LINE                         // increment SP
	//push ARG
	assembly += "@ARG" + value.NEW_LINE + "A=M" + value.NEW_LINE + "D=A" + value.NEW_LINE // set ARG to D
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE + "M=D" + value.NEW_LINE  // set D to RAM[SP]
	assembly += "@SP" + value.NEW_LINE + "M=M+1" + value.NEW_LINE                         // increment SP
	//push THIS
	assembly += "@THIS" + value.NEW_LINE + "A=M" + value.NEW_LINE + "D=A" + value.NEW_LINE // set THIS to D
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE + "M=D" + value.NEW_LINE   // set D to RAM[SP]
	assembly += "@SP" + value.NEW_LINE + "M=M+1" + value.NEW_LINE                          // increment SP
	//push THAT
	assembly += "@THAT" + value.NEW_LINE + "A=M" + value.NEW_LINE + "D=A" + value.NEW_LINE // set THAT to D
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE + "M=D" + value.NEW_LINE   // set D to RAM[SP]
	assembly += "@SP" + value.NEW_LINE + "M=M+1" + value.NEW_LINE                          // increment SP
	// ARG = SP - n - 5
	assembly += fmt.Sprintf("@%d", command.NumArgs) + value.NEW_LINE + "D=A" + value.NEW_LINE + "@5" + value.NEW_LINE + "D=D+A" + value.NEW_LINE // set (n  + 5)  to D
	assembly += "@SP" + value.NEW_LINE + "D=M-D" + value.NEW_LINE                                                                                // set SP-n-5 (=SP-(n+5)) to D
	assembly += "@ARG" + value.NEW_LINE + "M=D" + value.NEW_LINE                                                                                 // set D to ARG
	// LCL = SP
	assembly += "@SP" + value.NEW_LINE + "D=M" + value.NEW_LINE  // set SP to D
	assembly += "@LCL" + value.NEW_LINE + "M=D" + value.NEW_LINE // set D to LCL
	// goto f
	gotoFCommand := &ast.GotoCommand{Command: ast.C_GOTO, Symbol: ast.GOTO, LabelName: command.FunctionName}
	gotoFuncAssembly, err := codeWriter.getGotoAssembly(gotoFCommand)
	if err != nil {
		return "", err
	}
	assembly += gotoFuncAssembly
	// (return address)
	returnLabelCommand := &ast.LabelCommand{Command: ast.C_LABEL, Symbol: ast.LABEL, LabelName: returnLabel}
	returnLabelAssembly, err := codeWriter.getLabelAssembly(returnLabelCommand)
	if err != nil {
		return "", err
	}
	assembly += returnLabelAssembly
	return assembly, nil
}

func (codeWriter *CodeWriter) WriteCall(command *ast.CallCommand) error {
	callAssembly, err := codeWriter.getCallAssembly(command)
	if err != nil {
		return err
	}
	codeWriter.writeAssembly(callAssembly)
	return nil
}

func (codeWriter *CodeWriter) getFunctionAssembly(command *ast.FunctionCommand) (string, error) {
	assembly := ""
	functionLabelCommand := &ast.LabelCommand{Command: ast.C_LABEL, Symbol: ast.LABEL, LabelName: command.FunctionName}
	labelAssembly, err := codeWriter.getLabelAssembly(functionLabelCommand)
	if err != nil {
		return "", err
	}
	assembly += labelAssembly

	// initialize local variable by 0
	pushZeroConstCommand := &ast.PushCommand{Segment: ast.CONSTANT, Comamnd: ast.C_PUSH, Index: 0}
	for i := 0; i < command.NumLocals; i++ {
		assembly += codeWriter.getPushConstantAssembly(pushZeroConstCommand) // push 0 to stack for initialization.
	}
	return assembly, nil
}

func (codeWriter *CodeWriter) getArithmeticAssembly(arithmeticCommand *ast.ArithmeticCommand) (string, error) {
	switch arithmeticCommand.Symbol {
	case ast.ADD:
		return codeWriter.getAddCommandAssembly(), nil
	case ast.SUB:
		return codeWriter.getSubCommandAssembly(), nil
	case ast.NEG:
		return codeWriter.getNegCommandAssembly(), nil
	case ast.NOT:
		return codeWriter.getNotCommandAssembly(), nil
	case ast.AND:
		return codeWriter.getAndCommandAssembly(), nil
	case ast.OR:
		return codeWriter.getOrCommandAssembly(), nil
	case ast.GT, ast.LT, ast.EQ:
		return codeWriter.getCompareAssembly(arithmeticCommand.Symbol), nil
	}
	return "", fmt.Errorf("%T couldn't convert to arithmeticAssembly", arithmeticCommand)
}

func (codeWriter *CodeWriter) getAddCommandAssembly() string {
	assembly := ""
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE   // read value which SP points to into M
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-1 points to into M
	assembly += "D=M" + value.NEW_LINE                            // set M to D
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-2 points to into M
	assembly += "M=M+D" + value.NEW_LINE                          // set RAM[SP-2] = RAM[SP-2] + RAM[SP-1]
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE // decrement SP
	return assembly
}

func (codeWriter *CodeWriter) getSubCommandAssembly() string {
	assembly := ""
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE   // read value which SP points to into M
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-1 points to into M
	assembly += "D=M" + value.NEW_LINE                            // set M to D
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-2 points to into M
	assembly += "M=M-D" + value.NEW_LINE                          // set RAM[SP-2] = RAM[SP-2] - RAM[SP-1]
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE // decrement SPftemp
	return assembly
}

func (codeWriter *CodeWriter) getAndCommandAssembly() string {
	assembly := ""
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE   // read value which SP points to into M
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-1 points to into M
	assembly += "D=M" + value.NEW_LINE                            // set M to D
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-2 points to into M
	assembly += "M=M&D" + value.NEW_LINE                          // set RAM[SP-2] = RAM[SP-2] and RAM[SP-1]
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE // decrement SP
	return assembly
}

func (codeWriter *CodeWriter) getOrCommandAssembly() string {
	assembly := ""
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE   // read value which SP points to into M
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-1 points to into M
	assembly += "D=M" + value.NEW_LINE                            // set M to D
	assembly += "A=A-1" + value.NEW_LINE                          // set value which SP-2 points to into M
	assembly += "M=M|D" + value.NEW_LINE                          // set RAM[SP-2] = RAM[SP-2] or RAM[SP-1]
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE // decrement SP
	return assembly
}

func (codeWriter *CodeWriter) getNegCommandAssembly() string {
	assembly := ""
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE // read value which SP points to into M
	assembly += "A=A-1" + value.NEW_LINE                        // set value which SP-1 points to into M
	assembly += "M=-M" + value.NEW_LINE                         // set RAM[SP-1] = -RAM[SP-1]
	return assembly
}

func (codeWriter *CodeWriter) getNotCommandAssembly() string {
	assembly := ""
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE // read value which SP points to into M
	assembly += "A=A-1" + value.NEW_LINE                        // set value which SP-1 points to into M
	assembly += "M=!M" + value.NEW_LINE                         // set RAM[SP-1] = !RAM[SP-1]
	return assembly
}

func (codeWriter *CodeWriter) getCompareAssembly(compareCommandSymbol ast.CommandSymbol) string {
	assembly := ""
	// set x(RAM[SP-2]) - y(RAM[SP-1]) to D(==x-y)
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE + "A=M" + value.NEW_LINE + "D=M" + value.NEW_LINE // set RAM[SP-1]=y to D
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE + "A=M" + value.NEW_LINE                          // set RAM[SP-2]=x to M
	assembly += "D=M-D" + value.NEW_LINE                                                                            // set x - y to D
	compareAddressFlag := strconv.Itoa(rand.Intn(1000000))
	// jump based on D
	switch compareCommandSymbol {
	case ast.EQ:
		assembly += "@TRUE" + compareAddressFlag + value.NEW_LINE + "D;JEQ" + value.NEW_LINE
	case ast.GT:
		assembly += "@TRUE" + compareAddressFlag + value.NEW_LINE + "D;JGT" + value.NEW_LINE
	case ast.LT:
		assembly += "@TRUE" + compareAddressFlag + value.NEW_LINE + "D;JLT" + value.NEW_LINE
	}
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE + "M=0" + value.NEW_LINE + "@NEXT" + compareAddressFlag + value.NEW_LINE + "0;JMP" + value.NEW_LINE      // if false set 0 to RAM[SP-2] & jump to NEXT(to prevent TRUE process)
	assembly += "(TRUE" + compareAddressFlag + ")" + value.NEW_LINE + "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE + "M=0" + value.NEW_LINE + "M=-1" + value.NEW_LINE // if true set -1 to RAM[SP-2]
	assembly += "(NEXT" + compareAddressFlag + ")" + value.NEW_LINE                                                                                                      // NEXT Addr
	assembly += "@SP" + value.NEW_LINE + "M=M+1" + value.NEW_LINE                                                                                                        // increment SP
	return assembly
}

func (codeWriter *CodeWriter) getPopAssembly(popCommand *ast.PopCommand) (string, error) {
	switch popCommand.Segment {
	case ast.STATIC:
		return codeWriter.getPopStaticAssembly(popCommand), nil
	case ast.ARGUMENT, ast.LOCAL, ast.THAT, ast.THIS, ast.POINTER, ast.TEMP:
		return codeWriter.getMemoryAccessPopAssembly(popCommand), nil
	}
	return "", nil
}

func (codeWriter *CodeWriter) getPopStaticAssembly(popCommand *ast.PopCommand) string {
	assembly := ""
	// set RAM[SP] to  {VM Classname}.{idx} .
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE + "A=A-1" + value.NEW_LINE + "D=M" + value.NEW_LINE       // set RAM[SP-1] to D
	assembly += fmt.Sprintf("@%s.%d", codeWriter.VmClassName, popCommand.Index) + value.NEW_LINE + "M=D" + value.NEW_LINE // set D(==RAM[SP-1]) to RAM[{Vm Classname}.{idx}]
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE                                                         // decrement SP
	return assembly
}

func (codeWriter *CodeWriter) getMemoryAccessPopAssembly(popCommand *ast.PopCommand) string {
	assembly := ""
	// set RAM[SP] to RAM[{segment} + idx]
	assembly += fmt.Sprintf("@%d", popCommand.Index) + value.NEW_LINE + "D=A" + value.NEW_LINE // set Idx to D
	TEMP_BASE_ADDRESS, POINTER_BASE_ADDRESS := 5, 3
	// set Segment Base Address to A (A == {segment},M=R[{segment}])
	switch popCommand.Segment {
	case ast.LOCAL:
		assembly += "@LCL" + value.NEW_LINE + "A=M" + value.NEW_LINE
	case ast.ARGUMENT:
		assembly += "@ARG" + value.NEW_LINE + "A=M" + value.NEW_LINE
	case ast.THAT:
		assembly += "@THAT" + value.NEW_LINE + "A=M" + value.NEW_LINE
	case ast.THIS:
		assembly += "@THIS" + value.NEW_LINE + "A=M" + value.NEW_LINE
	case ast.TEMP:
		assembly += "@" + strconv.Itoa(TEMP_BASE_ADDRESS) + value.NEW_LINE
	case ast.POINTER:
		assembly += "@" + strconv.Itoa(POINTER_BASE_ADDRESS) + value.NEW_LINE
	}
	assembly += "D=D+A" + value.NEW_LINE                                                                            // set {segment} + Idx to D
	assembly += "@temp" + value.NEW_LINE + "M=D" + value.NEW_LINE                                                   // set {segment} + Idx to RAM[temp]
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE + "A=A-1" + value.NEW_LINE + "D=M" + value.NEW_LINE // set RAM[SP-1] to D
	assembly += "@temp" + value.NEW_LINE + "A=M" + value.NEW_LINE                                                   // set {segment}+Idx to A → A=={segment} + Idx, M == RAM[{segment} + Idx]
	assembly += "M=D" + value.NEW_LINE                                                                              // set D(==RAM[SP-1]) to RAM[{segment} + Idx]
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE                                                   // decrement SP
	return assembly
}

func (codeWriter *CodeWriter) getPushAssembly(pushCommand *ast.PushCommand) (string, error) {
	switch pushCommand.Segment {
	case ast.CONSTANT:
		return codeWriter.getPushConstantAssembly(pushCommand), nil
	case ast.ARGUMENT, ast.LOCAL, ast.THAT, ast.THIS, ast.POINTER, ast.TEMP:
		return codeWriter.getMemoryAccessPushAssembly(pushCommand), nil
	case ast.STATIC:
		return codeWriter.getPushStaticAssembly(pushCommand), nil
	}
	return "", fmt.Errorf("%T couldn't convert to pushAssembly", pushCommand)
}

func (codeWriter *CodeWriter) getPushConstantAssembly(pushCommand *ast.PushCommand) string {
	assembly := ""
	assembly += "@" + strconv.Itoa(pushCommand.Index) + value.NEW_LINE + "D=A" + value.NEW_LINE // set constant value to D
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE                                 // read value which SP points to into M
	assembly += "M=D" + value.NEW_LINE                                                          // set D to M
	assembly += "@SP" + value.NEW_LINE + "M=M+1" + value.NEW_LINE                               // increment SP
	return assembly
}

func (codeWriter *CodeWriter) getMemoryAccessPushAssembly(pushCommand *ast.PushCommand) string {
	assembly := ""
	assembly += "@" + strconv.Itoa(pushCommand.Index) + value.NEW_LINE + "D=A" + value.NEW_LINE // set constant value to D
	TEMP_BASE_ADDRESS, POINTER_BASE_ADDRESS := 5, 3
	// read Segment Base Address to A (A == {segment},M=R[{segment}])
	switch pushCommand.Segment {
	case ast.LOCAL:
		assembly += "@LCL" + value.NEW_LINE + "A=M" + value.NEW_LINE
	case ast.ARGUMENT:
		assembly += "@ARG" + value.NEW_LINE + "A=M" + value.NEW_LINE
	case ast.THAT:
		assembly += "@THAT" + value.NEW_LINE + "A=M" + value.NEW_LINE
	case ast.THIS:
		assembly += "@THIS" + value.NEW_LINE + "A=M" + value.NEW_LINE
	case ast.TEMP:
		assembly += "@" + strconv.Itoa(TEMP_BASE_ADDRESS) + value.NEW_LINE
	case ast.POINTER:
		assembly += "@" + strconv.Itoa(POINTER_BASE_ADDRESS) + value.NEW_LINE
	}
	assembly += "A=A+D" + value.NEW_LINE                                                 // set A =  A + D (A == LCL + idx , M == RAM[LCL + idx])
	assembly += "D=M" + value.NEW_LINE                                                   // set D=M (D == RAM[LCL + idx])
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE + "M=D" + value.NEW_LINE // set D to RAM[sp]
	assembly += "@SP" + value.NEW_LINE + "M=M+1" + value.NEW_LINE                        // increment SP
	return assembly
}

func (codeWriter *CodeWriter) getPushStaticAssembly(pushCommand *ast.PushCommand) string {
	assembly := ""
	assembly += fmt.Sprintf("@%s.%d", codeWriter.VmClassName, pushCommand.Index) + value.NEW_LINE + "D=M" + value.NEW_LINE // set static to D
	assembly += "@SP" + value.NEW_LINE + "A=M" + value.NEW_LINE + "M=D" + value.NEW_LINE                                   // set D to RAM[sp]
	assembly += "@SP" + value.NEW_LINE + "M=M+1" + value.NEW_LINE                                                          // increment SP
	return assembly
}

func (codeWriter *CodeWriter) writeAssembly(assembly string) {
	codeWriter.Assembly = append(codeWriter.Assembly, []byte(assembly)...)
}
