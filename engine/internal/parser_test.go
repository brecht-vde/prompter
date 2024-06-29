package engine

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	expected := &Template{
		s: []Statement{
			&PlainText{Token: Token{Type: PLAIN, Literal: "This is my first template, "}, Value: "This is my first template, "},
			&Variable{Token: Token{Type: IDENTIFIER, Literal: "Title"}, Identifier: "Title"},
			&PlainText{Token: Token{Type: PLAIN, Literal: ". Welcome "}, Value: ". Welcome "},
			&Variable{Token: Token{Type: IDENTIFIER, Literal: "User"}, Identifier: "User"},
			&PlainText{Token: Token{Type: PLAIN, Literal: "."}, Value: "."}},
	}

	l := NewLexer("This is my first template, {{ Title }}. Welcome {{ User }}.")
	p := NewParser(l)

	template := p.Parse()

	if !reflect.DeepEqual(template, expected) {
		t.Fatalf("expected: \n%q\n, but got \n%q\n", expected, template)
	}
}
