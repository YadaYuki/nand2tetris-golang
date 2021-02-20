package code_writer

import (
	"VMtranslator/parser"
	"errors"
	"strconv"
	"strings"
)

// GetPushPop get
func GetPushPop(commandType parser.CommandType, segment string, index int) (assembly string, err error) {

	if commandType == parser.CPush {
		if segment == "constant" {
			return "@" + strconv.Itoa(index) + "\n" + "D=A\n" + "@SP\n" + "A=M\n" + "M=D", nil
		}
	}
	if commandType == parser.CPop {
		return "", nil
	}
	return "", errors.New("Invalid Command Type")
}

// GetArithmetic returns assembly for arithmetic command
func GetArithmetic(commandStr string) (assembly string, err error) {
	s := strings.TrimSpace(commandStr)
	// TODO: to switch
	if s == "add" {
		return "@SP\n" + "A=M\n" + "D=M\n" + "A=A-1\n" + "D=D+M\n" + "M=D\n" + "@SP\n" + "M=M-1\n", nil
	}
	if s == "sub" {
		return "@SP\n" + "A=M\n" + "D=M\n" + "A=A-1\n" + "M=M-D\n" + "@SP\n" + "M=M-1\n", nil
	}
	if s == "neg" {
		return "@SP\n" + "A=M\n" + "D=M\n" + "M=-D\n", nil
	}
	if s == "and" {
		return "@SP\n" + "A=M\n" + "D=M\n" + "A=A-1\n" + "D=D&M\n" + "M=D\n" + "@SP\n" + "M=M-1\n", nil
	}
	if s == "or" {
		return "@SP\n" + "A=M\n" + "D=M\n" + "A=A-1\n" + "D=D|M\n" + "M=D\n" + "@SP\n" + "M=M-1\n", nil
	}
	if s == "not" {
		return "@SP\n" + "A=M\n" + "D=M\n" + "M=!D\n", nil
	}

	return "", errors.New("invalid arithmetic command")
}
