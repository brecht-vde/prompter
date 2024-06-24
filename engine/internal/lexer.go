package engine

type Lexer struct {
	template   string
	length     int
	pos        int
	spos       int
	current    string
	isSemantic bool
}

func NewLexer(template string) *Lexer {
	l := &Lexer{
		template: template,
		length:   len(template),
		pos:      1,
		spos:     0,
		current:  string(template[0]),
	}

	return l
}

func (l *Lexer) NextToken() Token {
	if l.pos >= l.length {
		l.spos = l.pos
		return Token{Type: EOF, Literal: ""}
	}

	if l.current == "{" && l.peek() == "{" {
		l.lex()
		l.lex()
		l.isSemantic = true
		l.spos = l.pos
		return Token{Type: DEL_OPEN, Literal: "{{"}
	}

	if l.current == "}" && l.peek() == "}" {
		l.lex()
		l.lex()
		l.isSemantic = false
		l.spos = l.pos
		return Token{Type: DEL_CLOSE, Literal: "}}"}
	}

	if l.isSemantic {
		tok := l.lexSemantic()
		l.spos = l.pos
		return tok
	}

	tok := l.lexPlain()
	l.spos = l.pos
	return tok
}

func (l *Lexer) lex() {
	if l.pos < l.length {
		l.current = string(l.template[l.pos])
		l.pos += 1
	} else {
		l.pos = l.length
		l.current = "EOF"
	}
}

func (l *Lexer) peek() string {
	if l.pos < l.length {
		return string(l.template[l.pos])
	} else {
		return "EOF"
	}
}

func (l *Lexer) delex() {
	if l.pos != l.length {
		l.pos -= 1
		l.current = string(l.template[l.pos])
	}
}

func (l *Lexer) lexPlain() Token {
	for {
		l.lex()

		if (l.current == "{" && l.peek() == "{") || l.current == "EOF" {
			l.delex()
			break
		}
	}

	literal := l.template[l.spos:l.pos]
	return Token{Type: PLAIN, Literal: literal}
}

func (l *Lexer) lexSemantic() Token {
	defer l.skipWhitespace()
	l.skipWhitespace()

	switch {
	case l.isLetter():
		return l.lexIdentifier()
	default:
		return Token{Type: ILLEGAL, Literal: ""}
	}
}

func (l *Lexer) skipWhitespace() {
	for {
		l.lex()

		if !l.isWhitespace() {
			l.delex()
			l.spos = l.pos
			break
		}
	}
}

func (l *Lexer) isWhitespace() bool {
	return l.current == " " || l.current == "\t" || l.current == "\r" || l.current == "\n"
}

func (l *Lexer) isLetter() bool {
	ch := l.current[0]
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func (l *Lexer) lexIdentifier() Token {
	for {
		l.lex()

		if l.isWhitespace() || l.current == "}" {
			l.delex()
			break
		}
	}

	literal := l.template[l.spos:l.pos]
	return Token{Type: IDENTIFIER, Literal: literal}
}
