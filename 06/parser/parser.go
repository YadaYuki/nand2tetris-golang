package parser

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type CommandType int

const (
	A_COMMAND CommandType = iota
	C_COMMAND
	L_COMMAND
)

// Advance アセンブルすべきコマンドがあるかどうかをboolで返す
func Advance(s *bufio.Scanner) (string, error) {
	if HasMoreCommand(s) {
		return s.Text(), nil
	}
	return "", errors.New("error")
}

// HasMoreCommand アセンブルすべきコマンドがあるかどうかをboolで返す
func HasMoreCommand(s *bufio.Scanner) bool {
	return s.Scan()
}

func main() {
	file, err := os.Open("06/rect/Rect.asm")

	if err != nil {
		log.Fatalf("%s", err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for {
		s, err := Advance(fileScanner)
		if err != nil {
			break
		}
		fmt.Println(s)
	}
	fmt.Println(A_COMMAND)

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("%s", err)
	}

}
