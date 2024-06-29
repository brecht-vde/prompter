package engine

import "strings"

type Template struct {
	s []Statement
}

func (t *Template) String() string {
	b := &strings.Builder{}

	for _, s := range t.s {
		b.WriteString(s.String())
	}

	return b.String()
}

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type PlainText struct {
	Token Token
	Value string
}

func (p *PlainText) TokenLiteral() string {
	return p.Token.Literal
}
func (p *PlainText) statementNode() {}
func (p *PlainText) String() string {
	return p.Value
}

type Variable struct {
	Token      Token
	Identifier string
}

func (v *Variable) TokenLiteral() string {
	return v.Token.Literal
}
func (v *Variable) statementNode() {}
func (v *Variable) String() string {
	return v.Identifier
}
