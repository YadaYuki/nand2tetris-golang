package codewriter

import (
	"VMtranslator/ast"
	"VMtranslator/value"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
)

type CodeWriter struct {
	Filename    string
	Assembly    []byte
	VmClassName string
}

func New(filename string, vmClassName string) *CodeWriter {
	return &CodeWriter{Filename: filename, Assembly: []byte{}, VmClassName: vmClassName}
}

func (codeWriter *CodeWriter) Close() {
	err := ioutil.WriteFile(codeWriter.Filename, codeWriter.Assembly, 0644)
	if err != nil {
		panic(err)
	}
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
	assembly := fmt.Sprintf("@%s", command.LabelName) + value.NEW_LINE
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
	assembly += fmt.Sprintf("@%s", command.LabelName) + value.NEW_LINE + "D;JEQ" + value.NEW_LINE // if D == RAM[SP] == 0 then jump to Label else continue
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
	assembly += "@SP" + value.NEW_LINE + "M=M-1" + value.NEW_LINE // decrement SP
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
