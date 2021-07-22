package codewriter

type CodeWirter struct {
	filename string
	assembly []byte
}

func New(filename string) *CodeWirter {
	return &CodeWirter{filename: filename, assembly: []byte{}}
}
