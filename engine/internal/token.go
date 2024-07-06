package internal

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	T_EOF        TokenType = "EOF"
	T_Plain      TokenType = "PLAIN"
	T_OpenVar    TokenType = "OPEN_VAR"
	T_CloseVar   TokenType = "CLOSE_VAR"
	T_Identifier TokenType = "IDENTIFIER"
	T_OpenJoin   TokenType = "OPEN_JOIN"
	T_CloseJoin  TokenType = "CLOSE_JOIN"
	T_Separator  TokenType = "SEPARATOR"
	T_Illegal    TokenType = "ILLEGAL"
)
