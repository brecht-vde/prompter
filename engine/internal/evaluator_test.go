package engine

import "testing"

func TestEvaluator(t *testing.T) {
	expected := "This is my first template, Awesome Title. Welcome Awesome User."

	data := map[string]interface{}{
		"Title": "Awesome Title",
		"User":  "Awesome User",
	}

	l := NewLexer("This is my first template, {{ Title }}. Welcome {{ User }}.")
	p := NewParser(l)

	template := p.Parse()

	rendition := Eval(*template, data)

	if expected != rendition {
		t.Fatalf("expected \n%s\n got \n%s\n", expected, rendition)
	}
}
