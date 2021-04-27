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
	CLASS       KeyWord = "class"
	METHOD      KeyWord = "method"
	FUNCTION    KeyWord = "function"
	CONSTRUCTOR KeyWord = "constructor"
	FIELD       KeyWord = "field"
	STATIC      KeyWord = "static"
	VAR         KeyWord = "var"
	INT         KeyWord = "int"
	CHAR        KeyWord = "char"
	BOOLEAN     KeyWord = "boolean"
	VOID        KeyWord = "void"
	TRUE        KeyWord = "true"
	FALSE       KeyWord = "false"
	NULL        KeyWord = "null"
	THIS        KeyWord = "this"
	LET         KeyWord = "let"
	DO          KeyWord = "do"
	IF          KeyWord = "if"
	ELSE        KeyWord = "else"
	WHILE       KeyWord = "while"
	RETURN      KeyWord = "return"
	// EOF
)

type Symbol string

var SymbolMap = map[byte]bool{'{': true, '}': true, '(': true, ')': true, '[': true, ']': true, '.': true, ':': true, ',': true, ';': true, '+': true, '-': true, '*': true, '/': true, '&': true, '|': true, '<': true, '>': true, '=': true, '~': true}

const (
	ASSIGN    Symbol = "="
	PLUS      Symbol = "+"
	MINUS     Symbol = "-"
	BANG      Symbol = "!"
	ASTERISK  Symbol = "*"
	SLASH     Symbol = "/"
	LT        Symbol = "<"
	GT        Symbol = ">"
	EQ        Symbol = "=="
	NOT_EQ    Symbol = "!="
	RPAREN    Symbol = ")"
	LPAREN    Symbol = "("
	RBRACE    Symbol = "}"
	LBRACE    Symbol = "{"
	COMMA     Symbol = ","
	SEMICOLON Symbol = ";"
	LBRACKET  Symbol = "["
	RBRACKET  Symbol = "]"
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
