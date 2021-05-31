package vmwriter

import (
	"bytes"
	"io/ioutil"
	"jack_compiler/value"
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
	if !bytes.Equal(vmWriter.VMCode, []byte("push const 0"+value.NEW_LINE)) {
		t.Fatalf("vmCode should be %s. got %s", "push const 0", vmWriter.VMCode)
	}
}

func TestWritePop(t *testing.T) {
	vmWriter := New("test.vm", 0644)
	vmWriter.WritePop(CONST, 0)
	if !bytes.Equal(vmWriter.VMCode, []byte("pop const 0"+value.NEW_LINE)) {
		t.Fatalf("vmCode should be %s. got %s", "pop const 0", vmWriter.VMCode)
	}
}
func TestWriteArithmetic(t *testing.T) {
	vmWriter := New("test.vm", 0644)
	vmWriter.WriteArithmetic(ADD)
	if !bytes.Equal(vmWriter.VMCode, []byte("add"+value.NEW_LINE)) {
		t.Fatalf("vmCode should be %s. got %s", "add", vmWriter.VMCode)
	}
}
