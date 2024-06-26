package engine

import (
	"unicode/utf8"
)

type Lexer struct {
	template string
	length   int
	current  rune
	next     rune
	cpos     int
	npos     int
	tpos     int
	semantic bool
	eof      bool
	illegal  bool
}

func NewLexer(template string) *Lexer {
	l := &Lexer{
		template: template,
	}

	l.length = utf8.RuneCountInString(template)

	r, s := utf8.DecodeRuneInString(l.template[l.cpos:])

	l.current = r
	l.npos = l.cpos + s

	r, _ = utf8.DecodeRuneInString(l.template[l.npos:])

	l.next = r

	return l
}

func (l *Lexer) NextToken() Token {
	if l.eof || l.illegal {
		return Token{Type: EOF, Literal: ""}
	}

	if l.current == '{' && l.next == '{' {
		l.lex()
		l.lex()
		l.semantic = true
		l.tpos = l.cpos
		return Token{Type: DEL_OPEN, Literal: "{{"}
	}

	if l.current == '}' && l.next == '}' {
		l.lex()
		l.lex()
		l.semantic = false
		l.tpos = l.cpos
		return Token{Type: DEL_CLOSE, Literal: "}}"}
	}

	if l.semantic {
		l.skipWhitespace()
		token := l.lexIdentifier()
		l.skipWhitespace()
		return token
	}

	if !l.semantic {
		return l.lexPlain()
	}

	l.illegal = true
	return Token{Type: ILLEGAL, Literal: ""}
}

func (l *Lexer) lex() {
	l.cpos = l.npos

	r, s := utf8.DecodeRuneInString(l.template[l.cpos:])

	l.current = r
	l.npos += s

	r, _ = utf8.DecodeRuneInString(l.template[l.npos:])

	l.next = r

	if l.current == utf8.RuneError {
		l.eof = true
	}
}

func (l *Lexer) skipWhitespace() {
	for l.current == ' ' || l.current == '\t' || l.current == '\n' || l.current == '\r' {
		l.lex()
	}

	l.tpos = l.cpos
}

func (l *Lexer) lexIdentifier() Token {

	for {
		if l.eof {
			break
		}

		if l.current == ' ' || l.current == '}' {
			break
		}

		l.lex()
	}

	// for !l.eof && l.current != ' ' && l.current != '}' {
	// 	l.lex()
	// }

	literal := l.template[l.tpos:l.cpos]
	l.tpos = l.cpos

	return Token{Type: IDENTIFIER, Literal: literal}
}

func (l *Lexer) lexPlain() Token {
	for {
		if l.eof {
			break
		}

		if l.current == '{' && l.next == '{' {
			break
		}

		l.lex()
	}

	if l.current != '{' || l.next != '{' {
		l.lex()
	}

	literal := l.template[l.tpos:l.cpos]
	l.tpos = l.cpos

	return Token{Type: PLAIN, Literal: literal}
}
