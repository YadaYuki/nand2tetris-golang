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

// type Symbol string

var SymbolMap = map[byte]bool{'{': true, '}': true, '(': true, ')': true, '[': true, ']': true, '.': true, ':': true, ',': true, ';': true, '+': true, '-': true, '*': true, '/': true, '&': true, '|': true, '<': true, '>': true, '=': true, '~': true}

const (
	ASSIGN    SymbolType = "="
	PLUS      SymbolType = "+"
	MINUS     SymbolType = "-"
	BANG      SymbolType = "!"
	ASTERISK  SymbolType = "*"
	SLASH     SymbolType = "/"
	LT        SymbolType = "<"
	GT        SymbolType = ">"
	EQ        SymbolType = "=="
	NOT_EQ    SymbolType = "!="
	RPAREN    SymbolType = ")"
	LPAREN    SymbolType = "("
	RBRACE    SymbolType = "}"
	LBRACE    SymbolType = "{"
	COMMA     SymbolType = ","
	SEMICOLON SymbolType = ";"
)

var KeyWordMap = map[string]KeyWord{
	"class":       CLASS,
	"method":      METHOD,
	"function":    FUNCTION,
	"constructor": CONSTRUCTOR,
	"field":       FIELD,
	"static":      STATIC,
	"var":         VAR,
	"int":         INT,
	"char":        CHAR,
	"boolean":     BOOLEAN,
	"void":        VOID,
	"true":        TRUE,
	"false":       FALSE,
	"null":        NULL,
	"this":        THIS,
	"let":         LET,
	"do":          DO,
	"if":          IF,
	"else":        ELSE,
	"while":       WHILE,
	"return":      RETURN,
}
