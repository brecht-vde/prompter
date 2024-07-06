package internal

import "testing"

func TestEvaluator(t *testing.T) {
	expected := `This is my first template, Awesome Title. Welcome Awesome User. Here are some options: Item 1, Item 2, Item 3.`

	data := map[string]interface{}{
		"Title": "Awesome Title",
		"User":  "Awesome User",
		"Items": []string{"Item 1", "Item 2", "Item 3"},
	}

	l := NewLexer(`This is my first template, {{var: Title}}. Welcome {{var: User}}. Here are some options: {{join: Items, ", "}}.`)
	p := NewParser(l)

	template := p.Parse()

	rendition, err := Eval(*template, data)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if expected != rendition {
		t.Fatalf("expected \n%s\n got \n%s\n", expected, rendition)
	}
}
