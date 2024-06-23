package engine

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	ILLEGAL    = "ILLEGAL"
	EOF        = "EOF"
	IDENTIFIER = "IDENTIFIER"
	DEL_OPEN   = "{{"
	DEL_CLOSE  = "}}"
	PLAIN      = "PLAIN"
)
