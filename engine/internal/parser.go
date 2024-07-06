package internal

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

	for p.c.Type != T_EOF && p.c.Type != T_Illegal {
		p.parseNext()

		if p.c.Type == T_Plain {
			t.s = append(t.s, &PlainText{Token: p.c, Value: p.c.Literal})
		} else if p.c.Type == T_Identifier && p.n.Type == T_Separator {
			t.s = append(t.s, &Join{Token: p.c, Identifier: p.c.Literal, Separator: p.n.Literal})
		} else if p.c.Type == T_Identifier {
			t.s = append(t.s, &Variable{Token: p.c, Identifier: p.c.Literal})
		} else if p.c.Type == T_Illegal {
			p.errors = append(p.errors, p.c.Literal)
		}
	}

	return t
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) initialize() {
	p.parseNext()
}

func (p *Parser) parseNext() {
	p.c = p.n
	p.n = p.l.NextToken()
}
