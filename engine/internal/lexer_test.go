package engine

import "testing"

func TestLexer(t *testing.T) {
	input := `hello {{ myvar }}!`

	tests := []struct {
		expectedType     Type
		exptectedLiteral string
	}{
		{PLAIN, "hello "},
		{DEL_OPEN, "{{"},
		{IDENTIFIER, "myvar"},
		{DEL_CLOSE, "}}"},
		{PLAIN, "!"},
		{EOF, ""},
	}

	l := NewLexer(input)

	for _, tt := range tests {
		tok := l.NextToken()
		t.Logf("%s: '%s'", tok.Type, tok.Literal)

		if tok.Type != tt.expectedType {
			t.Fatalf("unexpected token type. expected '%s', got '%s'", tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.exptectedLiteral {
			t.Fatalf("unexpected token literal. expected '%s', got '%s'", tt.exptectedLiteral, tok.Literal)
		}
	}
}

func TestLexerV2(t *testing.T) {
	input := `{{ myvar }} - something! {{     yourvar  }}. {{another}}`

	tests := []struct {
		expectedType     Type
		exptectedLiteral string
	}{
		{DEL_OPEN, "{{"},
		{IDENTIFIER, "myvar"},
		{DEL_CLOSE, "}}"},
		{PLAIN, " - something! "},
		{DEL_OPEN, "{{"},
		{IDENTIFIER, "yourvar"},
		{DEL_CLOSE, "}}"},
		{PLAIN, ". "},
		{DEL_OPEN, "{{"},
		{IDENTIFIER, "another"},
		{DEL_CLOSE, "}}"},
	}

	l := NewLexer(input)

	for _, tt := range tests {
		tok := l.NextToken()
		t.Logf("%s: '%s'", tok.Type, tok.Literal)

		if tok.Type != tt.expectedType {
			t.Fatalf("unexpected token type. expected '%s', got '%s'", tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.exptectedLiteral {
			t.Fatalf("unexpected token literal. expected '%s', got '%s'", tt.exptectedLiteral, tok.Literal)
		}
	}
}
