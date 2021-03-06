package lexer

// Lexer has member necessary for parsing
type Lexer struct {
	input        string // code input
	position     int
	readPosition int
	ch           byte
}
