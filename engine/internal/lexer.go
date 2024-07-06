package internal

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
	illegal bool
}

const (
	E_UnexpectedEndOfFile              = "unexpected end of file"
	E_UnexpectedCharacterExpectedSpace = "unexpected character, expected ' '"
	E_UnexpectedCharacterExpectedQuote = `unexpected character, expected '"'`
	E_UnexpectedCharacterExpectedColon = "unexpected character, expect ':'"
	E_UnknownActionType                = "unknown action type"
	E_UnknownClosingType               = "unknown closing type"
	E_InvalidIdentifier                = "invalid identifier, identifiers can only contain a-z/A-Z"
)

const (
	M_LBrace = '{'
	M_RBrace = '}'
	M_Colon  = ':'
	M_Space  = ' '
	M_Comma  = ','
	M_Eof    = utf8.RuneError
	M_Quote  = '"'
)

type ActionType string

const (
	A_None     ActionType = "NONE"
	A_Variable ActionType = "VAR"
	A_Join     ActionType = "JOIN"
)

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  input,
		action: A_None,
	}
}

func (l *Lexer) NextToken() (token Token) {
	l.read()

	if l.isEof() {
		token = Token{Type: T_EOF, Literal: ""}
	} else if l.isStartAction() {
		token = l.readType()
	} else if l.isEndAction() {
		token = l.readEnd()
	} else if l.isSeparator() {
		token = l.readSeparator()
	} else if l.isAction() {
		token = l.readAction()
	} else if l.isPlain() {
		token = l.readPlain()
	} else {
		l.illegal = true
		token = Token{Type: T_Illegal, Literal: "unexpected token"}
	}

	return token
}

func (l *Lexer) isEof() bool {
	return l.eof || l.illegal
}

func (l *Lexer) isStartAction() bool {
	return l.current == M_LBrace && l.peek() == M_LBrace
}

func (l *Lexer) isAction() bool {
	return l.action != A_None
}

func (l *Lexer) isPlain() bool {
	return l.action == A_None
}

func (l *Lexer) isEndAction() bool {
	return l.current == M_RBrace && l.peek() == M_RBrace && l.action != A_None
}

func (l *Lexer) isSeparator() bool {
	return l.action == A_Join && l.current == M_Comma && l.peek() == M_Space
}

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
	err := ""
	pos := l.cpos

	for {
		l.read()

		if l.current == M_Colon {
			break
		}

		if l.current == M_Space {
			err = E_UnexpectedCharacterExpectedColon
			break
		}

		if l.eof {
			err = E_UnexpectedEndOfFile
			break
		}
	}

	literal := l.input[pos:l.pos]

	l.read()

	if err == "" && l.current != M_Space {
		err = E_UnexpectedCharacterExpectedSpace
	}

	if err != "" {
		l.illegal = true
		return Token{Type: T_Illegal, Literal: err}
	} else if literal == "{{var:" {
		l.action = A_Variable
		return Token{Type: T_OpenVar, Literal: literal}
	} else if literal == "{{join:" {
		l.action = A_Join
		return Token{Type: T_OpenJoin, Literal: literal}
	} else {
		l.illegal = true
		return Token{Type: T_Illegal, Literal: E_UnknownActionType}
	}
}

func (l *Lexer) readAction() Token {
	switch l.action {
	case A_Variable:
		return l.readIdentifier()
	case A_Join:
		return l.readIdentifier()
	default:
		l.illegal = true
		return Token{Type: T_Illegal, Literal: "unknown action"}
	}
}

func (l *Lexer) readEnd() Token {
	l.read()

	switch l.action {
	case A_Variable:
		l.action = A_None
		return Token{Type: T_CloseVar, Literal: "}}"}
	case A_Join:
		l.action = A_None
		return Token{Type: T_CloseJoin, Literal: "}}"}
	default:
		l.illegal = true
		return Token{Type: T_Illegal, Literal: E_UnknownClosingType}
	}
}

func (l *Lexer) readIdentifier() Token {
	err := ""
	pos := l.cpos

	for {
		l.read()

		if (l.action == A_Variable && l.peek() == M_RBrace) ||
			(l.action == A_Join && l.peek() == M_Comma) {
			break
		}

		if l.eof {
			err = E_UnexpectedEndOfFile
		}

		if (l.current < 'a' || l.current > 'z') && (l.current < 'A' || l.current > 'Z') {
			err = E_InvalidIdentifier
		}
	}

	i := l.input[pos:l.pos]

	if err != "" {
		l.illegal = true
		return Token{Type: T_Illegal, Literal: err}
	}

	return Token{Type: T_Identifier, Literal: i}
}

func (l *Lexer) readSeparator() Token {
	err := ""

	l.read()
	l.read()

	if l.current != M_Quote {
		err = E_UnexpectedCharacterExpectedQuote
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

	if err != "" {
		l.illegal = true
		return Token{Type: T_Illegal, Literal: err}
	}

	return Token{Type: T_Separator, Literal: s}
}

func (l *Lexer) readPlain() Token {
	pos := l.cpos

	for {
		if l.peek() == M_LBrace || l.current == M_Eof {
			break
		}

		l.read()
	}

	p := l.input[pos:l.pos]

	return Token{Type: T_Plain, Literal: p}
}
