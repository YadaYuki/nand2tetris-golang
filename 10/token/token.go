package token

// TokenType is type of token
type TokenType string

// Token has memeber tokentype,tokenliteral
type Token struct {
	Type    TokenType
	Literal string
}

const (
	SYMBOL        TokenType = "SYMBOL"
	KEYWORD       TokenType = "KEYWORD"
	IDENTIFIER    TokenType = "IDENTIFIER"
	INTCONST      TokenType = "INT_CONST"
	STARTINGCONST TokenType = "STARTING_CONST"
	EOF           TokenType = "EOF"
)

// KeyWord is keyword type
type KeyWord string

const (
	CLASS       KeyWord = "CLASS"
	METHOD      KeyWord = "METHOD"
	FUNCTION    KeyWord = "FUNCTION"
	CONSTRUCTOR KeyWord = "CONSTRUCTOR"
	INT         KeyWord = "INT"
	BOOLEAN     KeyWord = "BOOLEAN"
	CHAR        KeyWord = "CHAR"
	VOID        KeyWord = "VOID"
	VAR         KeyWord = "VAR"
	STATIC      KeyWord = "STATIC"
	FIELD       KeyWord = "FIELD"
	LET         KeyWord = "LET"
	DO          KeyWord = "DO"
	IF          KeyWord = "IF"
	ELSE        KeyWord = "ELSE"
	WHILE       KeyWord = "WHILE"
	RETURN      KeyWord = "RETURN"
	TRUE        KeyWord = "TRUE"
	FALSE       KeyWord = "FALSE"
	NULL        KeyWord = "NULL"
	THIS        KeyWord = "THIS"
	// EOF
)

var SymbolMap = map[byte]bool{'{': true, '}': true, '(': true, ')': true, '[': true, ']': true, '.': true, ':': true, ',': true, ';': true, '+': true, '-': true, '*': true, '/': true, '&': true, '|': true, '<': true, '>': true, '=': true, '~': true}
var KeyWordMap = map[string]KeyWord{
	"class":       token.CLASS,
	"method":      token.METHOD,
	"function":    token.FUNCTION,
	"constructor": token.CONSTRUCTOR,
	"field":       token.FIELD,
	"static":      token.STATIC,
	"var":         token.VAR,
	"int":         token.INT,
	"char":        token.CHAR,
	"boolean":     token.BOOLEAN,
	"void":        token.VOID,
	"true":        token.TRUE,
	"false":       token.FALSE,
	"null":        token.NULL,
	"this":        token.THIS,
	"let":         token.LET,
	"do":          token.DO,
	"if":          token.IF,
	"else":        token.ELSE,
	"while":       token.WHILE,
	"return":      token.RETURN}
