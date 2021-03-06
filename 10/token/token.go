package token

// TokenType is type of token
type TokenType string

// Token has memeber tokentype,tokenliteral
type Token struct {
	Type    TokenType
	Literal string
}

const (
	Keyword       TokenType = "KEYWORD"
	SymboL        TokenType = "SYMBOL"
	Identifier    TokenType = "IDENTIFIER"
	IntConst      TokenType = "INT_CONST"
	StartingConst TokenType = "STARTING_CONST"
)

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
)
