package main

import (
	"reflect"
	"testing"
)

func TestGetVmFileList(t *testing.T) {
	testCases := []struct {
		dirPath    string
		vmFileList []string
	}{
		{"FunctionCalls/NestedCall", []string{"FunctionCalls/NestedCall/Sys.vm"}},
		{"FunctionCalls/FibonacciElement", []string{"FunctionCalls/FibonacciElement/Main.vm", "FunctionCalls/FibonacciElement/Sys.vm"}},
	}
	for _, tt := range testCases {
		vmFileList, err := getVmFileListInDir(tt.dirPath)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(tt.vmFileList, vmFileList) {
			t.Fatalf("vmFileList should be %s, but got %s", tt.vmFileList, vmFileList)
		}
	}
}

func TestRemoveExt(t *testing.T) {
	testCases := []struct {
		filename           string
		filenameExtRemoved string
	}{
		{"index.js", "index"},
		{"index.vm", "index"},
	}
	for _, tt := range testCases {
		filenameExtRemoved := removeExt(tt.filename)
		if tt.filenameExtRemoved != filenameExtRemoved {
			t.Fatalf("filenameExtRemoved should be %s, but got %s", tt.filenameExtRemoved, filenameExtRemoved)
		}
	}
}
