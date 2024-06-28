package engine

import "testing"

func TestLexerValidInputs(t *testing.T) {
	tests := []struct {
		input    string
		expected []struct {
			tokenType Type
			literal   string
		}
	}{
		{
			input: `{{myvar}}`,
			expected: []struct {
				tokenType Type
				literal   string
			}{
				{DEL_OPEN, "{{"},
				{IDENTIFIER, "myvar"},
				{DEL_CLOSE, "}}"},
				{EOF, ""},
			},
		},
		{
			input: `{{ myvar}}`,
			expected: []struct {
				tokenType Type
				literal   string
			}{
				{DEL_OPEN, "{{"},
				{IDENTIFIER, "myvar"},
				{DEL_CLOSE, "}}"},
				{EOF, ""},
			},
		},
		{
			input: `{{myvar }}`,
			expected: []struct {
				tokenType Type
				literal   string
			}{
				{DEL_OPEN, "{{"},
				{IDENTIFIER, "myvar"},
				{DEL_CLOSE, "}}"},
				{EOF, ""},
			},
		},
		{
			input: `{{ myvar }}`,
			expected: []struct {
				tokenType Type
				literal   string
			}{
				{DEL_OPEN, "{{"},
				{IDENTIFIER, "myvar"},
				{DEL_CLOSE, "}}"},
				{EOF, ""},
			},
		},
		{
			input: `text {{myvar}} text.`,
			expected: []struct {
				tokenType Type
				literal   string
			}{
				{PLAIN, "text "},
				{DEL_OPEN, "{{"},
				{IDENTIFIER, "myvar"},
				{DEL_CLOSE, "}}"},
				{PLAIN, " text."},
				{EOF, ""},
			},
		},
		{
			input: `{{myvar}`,
			expected: []struct {
				tokenType Type
				literal   string
			}{
				{DEL_OPEN, "{{"},
				{ILLEGAL, "invalid closing delimiter."},
				{EOF, ""},
			},
		},
		{
			input: `{myvar}}`,
			expected: []struct {
				tokenType Type
				literal   string
			}{
				{ILLEGAL, "unexpected closing delimiter."},
				{EOF, ""},
			},
		},
		{
			input: `myvar}}`,
			expected: []struct {
				tokenType Type
				literal   string
			}{
				{ILLEGAL, "unexpected closing delimiter."},
				{EOF, ""},
			},
		},
		{
			input: `{{myvar`,
			expected: []struct {
				tokenType Type
				literal   string
			}{
				{DEL_OPEN, "{{"},
				{ILLEGAL, "missing closing delimiter."},
				{EOF, ""},
			},
		},
		{
			input: `myvar}}`,
			expected: []struct {
				tokenType Type
				literal   string
			}{
				{ILLEGAL, "unexpected closing delimiter."},
				{EOF, ""},
			},
		},
		{
			input: `text {{myvar text.`,
			expected: []struct {
				tokenType Type
				literal   string
			}{
				{PLAIN, "text "},
				{DEL_OPEN, "{{"},
				{ILLEGAL, "missing closing delimiter."},
				{EOF, ""},
			},
		},
		{
			input: `text {{ myvar} } text.`,
			expected: []struct {
				tokenType Type
				literal   string
			}{
				{PLAIN, "text "},
				{DEL_OPEN, "{{"},
				{ILLEGAL, "invalid closing delimiter."},
				{EOF, ""},
			},
		},
		{
			input: `text { { myvar} } text.`,
			expected: []struct {
				tokenType Type
				literal   string
			}{
				{PLAIN, "text { { myvar} } text."},
				{EOF, ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			l := NewLexer(tt.input)

			for _, e := range tt.expected {
				tok := l.NextToken()

				if tok.Type != e.tokenType {
					t.Fatalf("unexpected token type. expected '%s', got '%s'", e.tokenType, tok.Type)
				}

				if tok.Literal != e.literal {
					t.Fatalf("unexpected token literal. expected '%s', got '%s'", e.literal, tok.Literal)
				}
			}
		})
	}
}
