package vmwriter

import (
	"bytes"
	"io/ioutil"
	"jackcompiler/value"
	"testing"
)

var filename = "test.vm"

func TestClose(t *testing.T) {
	vmCode := []byte("Hello,World")
	vmWriter := &VMWriter{
		VMCode: vmCode, Filename: filename, perm: 0644,
	}
	vmWriter.Close()
	content, _ := ioutil.ReadFile(filename)
	if !bytes.Equal(content, vmCode) {
		t.Fatalf("vmCode should be %s. got %s", vmCode, content)
	}
}
func TestWriteData(t *testing.T) {
	vmCode := "Hello,World."
	vmWriter := &VMWriter{
		VMCode: []byte(vmCode), Filename: filename, perm: 0644,
	}
	addVmCode := "Good bye, World"
	vmWriter.writeData(addVmCode)
	if !bytes.Equal(vmWriter.VMCode, []byte(vmCode+addVmCode)) {
		t.Fatalf("vmCode should be %s. got %s", vmCode+addVmCode, vmWriter.VMCode)
	}
}

func TestWritePush(t *testing.T) {
	vmWriter := New("test.vm", 0644)
	vmWriter.WritePush(CONST, 0)
	if !bytes.Equal(vmWriter.VMCode, []byte("push constant 0"+value.NEW_LINE)) {
		t.Fatalf("vmCode should be %s. got %s", "push constant 0", vmWriter.VMCode)
	}
}

func TestWritePop(t *testing.T) {
	vmWriter := New("test.vm", 0644)
	vmWriter.WritePop(CONST, 0)
	if !bytes.Equal(vmWriter.VMCode, []byte("pop constant 0"+value.NEW_LINE)) {
		t.Fatalf("vmCode should be %s. got %s", "pop constant 0", vmWriter.VMCode)
	}
}
func TestWriteArithmetic(t *testing.T) {
	vmWriter := New("test.vm", 0644)
	vmWriter.WriteArithmetic(ADD)
	if !bytes.Equal(vmWriter.VMCode, []byte("add"+value.NEW_LINE)) {
		t.Fatalf("vmCode should be %s. got %s", "add", vmWriter.VMCode)
	}
}

func TestWriteLabel(t *testing.T) {
	vmWriter := New("test.vm", 0644)
	vmWriter.WriteLabel("LOOP")
	if !bytes.Equal(vmWriter.VMCode, []byte("label LOOP"+value.NEW_LINE)) {
		t.Fatalf("vmCode should be %s. got %s", "label LOOP", vmWriter.VMCode)
	}
}

func TestWriteGoto(t *testing.T) {
	vmWriter := New("test.vm", 0644)
	vmWriter.WriteGoto("LOOP")
	if !bytes.Equal(vmWriter.VMCode, []byte("goto LOOP"+value.NEW_LINE)) {
		t.Fatalf("vmCode should be %s. got %s", "goto LOOP", vmWriter.VMCode)
	}
}

func TestWriteIf(t *testing.T) {
	vmWriter := New("test.vm", 0644)
	vmWriter.WriteIf("LOOP")
	if !bytes.Equal(vmWriter.VMCode, []byte("if-goto LOOP"+value.NEW_LINE)) {
		t.Fatalf("vmCode should be %s. got %s", "if-goto LOOP", vmWriter.VMCode)
	}
}

func TestWriteCall(t *testing.T) {
	vmWriter := New("test.vm", 0644)
	vmWriter.WriteCall("hogeFunc", 3)
	if !bytes.Equal(vmWriter.VMCode, []byte("call hogeFunc 3"+value.NEW_LINE)) {
		t.Fatalf("vmCode should be %s. got %s", "call hogeFunc 3", vmWriter.VMCode)
	}
}

func TestWriteFunction(t *testing.T) {
	vmWriter := New("test.vm", 0644)
	vmWriter.WriteFunction("hogeFunc", 3)
	if !bytes.Equal(vmWriter.VMCode, []byte("function hogeFunc 3"+value.NEW_LINE)) {
		t.Fatalf("vmCode should be %s. got %s", "function hogeFunc 3", vmWriter.VMCode)
	}
}

func TestWriteReturn(t *testing.T) {
	vmWriter := New("test.vm", 0644)
	vmWriter.WriteReturn()
	if !bytes.Equal(vmWriter.VMCode, []byte("return"+value.NEW_LINE)) {
		t.Fatalf("vmCode should be %s. got %s", "return", vmWriter.VMCode)
	}
}
