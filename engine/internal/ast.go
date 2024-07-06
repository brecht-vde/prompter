package internal

import "strings"

type Template struct {
	s []Statement
}

func (t *Template) String() string {
	b := &strings.Builder{}

	for _, s := range t.s {
		b.WriteString(s.ValueString())
	}

	return b.String()
}

type Node interface {
	TokenLiteral() string
	String() string
	ValueString() string
}

type Statement interface {
	Node
	expressionNode()
}

type PlainText struct {
	Token Token
	Value string
}

func (p *PlainText) TokenLiteral() string {
	return p.Token.Literal
}
func (p *PlainText) expressionNode() {}
func (p *PlainText) String() string {
	return p.Value
}
func (p *PlainText) ValueString() string {
	return p.Value
}

type Variable struct {
	Token      Token
	Identifier string
	Value      string
}

func (v *Variable) TokenLiteral() string {
	return v.Token.Literal
}
func (v *Variable) expressionNode() {}
func (v *Variable) String() string {
	return v.Identifier
}
func (v *Variable) ValueString() string {
	return v.Value
}

type Join struct {
	Token      Token
	Identifier string
	Separator  string
	Value      string
}

func (j *Join) TokenLiteral() string {
	return j.Token.Literal
}
func (j *Join) expressionNode() {}
func (j *Join) String() string {
	return j.Identifier + ": " + j.Separator
}
func (j *Join) ValueString() string {
	return j.Value
}
