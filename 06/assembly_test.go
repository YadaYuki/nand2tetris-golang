package assembly

import (
	"assembly/value"
	"io/ioutil"
	"strings"
	"testing"
)

func TestAssemble(t *testing.T) {
	testCases := []struct {
		asmFilename  string
		hackFilename string
	}{
		{"add/Add.asm", "add/Add.hack"},
		{"rect/Rect.asm", "rect/Rect.hack"},
		{"max/Max.asm", "max/Max.hack"},
		// {"pong/Pong.asm", "pong/Pong.hack"},
	}
	for _, tt := range testCases {
		asm, _ := ioutil.ReadFile(tt.asmFilename)
		input := string(asm)
		hack, _ := ioutil.ReadFile(tt.hackFilename)
		binaryArrInFile := strings.Split(string(hack), value.LF)
		binaryArr, _ := Assemble(input)
		commandArr := strings.Split(input, value.NEW_LINE)
		for i := range binaryArr {
			if binaryArrInFile[i] != binaryArr[i] {
				t.Fatalf("%s's binary should be %s got,%s", commandArr[i], binaryArrInFile[i], binaryArr[i])
			}
		}
	}
}
