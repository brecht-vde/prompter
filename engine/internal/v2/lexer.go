package v2

import (
	"unicode/utf8"
)

type Lexer struct {
	input   string
	pos     int
	cpos    int
	current rune
	eof     bool
	action  ActionType
}

const (
	M_LBrace = '{'
	M_RBrace = '}'
	M_Colon  = ':'
	M_Space  = ' '
	M_Comma  = ','
	M_Eof    = utf8.RuneError
	M_Quote  = '"'
)

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
)

type ActionType string

const (
	A_None     ActionType = "NONE"
	A_Variable ActionType = "VAR"
	A_Join     ActionType = "JOIN"
)

type Token struct {
	Type    TokenType
	Literal string
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: input,
	}
}

func (l *Lexer) NextToken() Token {
	l.read()

	if l.eof {
		return Token{Type: T_EOF, Literal: ""}
	}

	if l.current == M_LBrace && l.peek() == M_LBrace {
		return l.readType()
	}

	if l.current == M_RBrace && l.peek() == M_RBrace && l.action != A_None {
		l.read()

		switch l.action {
		case A_Variable:
			l.action = A_None
			return Token{Type: T_CloseVar, Literal: "}}"}
		case A_Join:
			l.action = A_None
			return Token{Type: T_CloseJoin, Literal: "}}"}
		}
	}

	if l.action == A_Join && l.current == M_Comma && l.peek() == M_Space {
		return l.readSeparator()
	}

	if l.action != A_None {
		return l.readAction()
	}

	return l.readPlain()
}

//

func (l *Lexer) read() {
	r, s := utf8.DecodeRuneInString(l.input[l.pos:])

	if r == utf8.RuneError {
		l.eof = true
	}

	l.current = r
	l.cpos = l.pos
	l.pos += s
}

func (l *Lexer) peek() rune {
	r, _ := utf8.DecodeRuneInString(l.input[l.pos:])
	return r
}

func (l *Lexer) readType() Token {
	pos := l.cpos

	for {
		l.read()

		if l.current == M_Colon {
			break
		}

		if l.eof {
			// illegal
		}
	}

	t := l.input[pos:l.pos]

	l.read()

	if l.current != M_Space {
		// illegal
	}

	switch t {
	case "{{var:":
		l.action = A_Variable
		return Token{Type: T_OpenVar, Literal: t}
	case "{{join:":
		l.action = A_Join
		return Token{Type: T_OpenJoin, Literal: t}
	}

	return Token{}
}

func (l *Lexer) readAction() Token {
	switch l.action {
	case A_Variable:
		return l.readIdentifier()
	case A_Join:
		return l.readIdentifier()
	default:
		return Token{}
	}
}

func (l *Lexer) readIdentifier() Token {
	pos := l.cpos

	for {
		l.read()

		if (l.action == A_Variable && l.peek() == M_RBrace) ||
			(l.action == A_Join && l.peek() == M_Comma) {
			break
		}

		if l.eof {
			// illegal
		}

		// if not alphanumeric, illegal
	}

	i := l.input[pos:l.pos]

	return Token{Type: T_Identifier, Literal: i}
}

func (l *Lexer) readSeparator() Token {
	l.read()
	l.read()

	if l.current != M_Quote {
		// illegal
		return Token{}
	}

	l.read()

	pos := l.cpos

	for {
		l.read()

		if l.peek() == M_Quote {
			break
		}
	}

	s := l.input[pos:l.pos]

	l.read()

	return Token{Type: T_Separator, Literal: s}
}

func (l *Lexer) readPlain() Token {
	pos := l.cpos

	for {
		l.read()

		if l.current == M_LBrace || l.current == M_Eof {
			break
		}
	}

	p := l.input[pos:l.pos]

	return Token{Type: T_Plain, Literal: p}
}
