package engine

type Parser struct {
	l      *Lexer
	c      Token
	n      Token
	errors []string
}

func NewParser(l *Lexer) *Parser {
	return &Parser{
		l: l,
	}
}

func (p *Parser) Parse() *Template {
	t := &Template{}
	t.s = []Statement{}

	p.initialize()

	for p.c.Type != EOF && p.c.Type != ILLEGAL {
		p.parseNext()

		if p.c.Type == PLAIN {
			t.s = append(t.s, &PlainText{Token: p.c, Value: p.c.Literal})
		}

		if p.c.Type == IDENTIFIER {
			t.s = append(t.s, &Variable{Token: p.c, Identifier: p.c.Literal})
		}
	}

	return t
}

func (p *Parser) initialize() {
	p.parseNext()
}

func (p *Parser) parseNext() {
	p.c = p.n
	p.n = p.l.NextToken()
}
