package vmwriter

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestClose(t *testing.T) {
	filename := "test.vm"
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
	filename := "test.vm"
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
