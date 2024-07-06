package internal

import (
	"reflect"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	expected := &Template{
		s: []Statement{
			&PlainText{Token: Token{Type: T_Plain, Literal: "This is my first template, "}, Value: "This is my first template, "},
			&Variable{Token: Token{Type: T_Identifier, Literal: "Title"}, Identifier: "Title"},
			&PlainText{Token: Token{Type: T_Plain, Literal: ". Welcome "}, Value: ". Welcome "},
			&Variable{Token: Token{Type: T_Identifier, Literal: "User"}, Identifier: "User"},
			&PlainText{Token: Token{Type: T_Plain, Literal: ". Here are some options: "}, Value: ". Here are some options: "},
			&Join{Token: Token{Type: T_Identifier, Literal: "Items"}, Identifier: "Items", Separator: ", "},
			&PlainText{Token: Token{Type: T_Plain, Literal: "."}, Value: "."},
		},
	}

	l := NewLexer(`This is my first template, {{var: Title}}. Welcome {{var: User}}. Here are some options: {{join: Items, ", "}}.`)
	p := NewParser(l)

	template := p.Parse()

	if !reflect.DeepEqual(template, expected) {
		t.Fail()
		t.Logf("expected: \n%q\n, but got \n%q\n", expected, template)
	}

	if p.errors != nil || len(p.errors) > 0 {
		t.Fail()
		t.Logf("expected no errors, but received: \n%s", strings.Join(p.errors, "\n"))
	}
}
