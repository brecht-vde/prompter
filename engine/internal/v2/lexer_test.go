package v2

import "testing"

type TestCase struct {
	input    string
	expected []Token
}

func testLexer(tcs []TestCase, t *testing.T) {
	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			l := NewLexer(tc.input)

			for _, e := range tc.expected {
				token := l.NextToken()

				if token.Type != e.Type {
					t.Fatalf("received: %v, expected: %v", token.Type, e.Type)
				}

				if token.Literal != e.Literal {
					t.Fatalf("received %s, expected: %s", token.Literal, e.Literal)
				}
			}
		})
	}
}

func TestLexerVariables(t *testing.T) {
	tcs := []TestCase{
		{
			input: "{{var: item}}",
			expected: []Token{
				{
					Type:    T_OpenVar,
					Literal: "{{var:",
				},
				{
					Type:    T_Identifier,
					Literal: "item",
				},
				{
					Type:    T_CloseVar,
					Literal: "}}",
				},
				{
					Type:    T_EOF,
					Literal: "",
				},
			},
		},
	}

	testLexer(tcs, t)
}

func TestLexerJoins(t *testing.T) {
	tcs := []TestCase{
		{
			input: `{{join: items, ", "}}`,
			expected: []Token{
				{
					Type:    T_OpenJoin,
					Literal: "{{join:",
				},
				{
					Type:    T_Identifier,
					Literal: "items",
				},
				{
					Type:    T_Separator,
					Literal: ", ",
				},
				{
					Type:    T_CloseJoin,
					Literal: "}}",
				},
				{
					Type:    T_EOF,
					Literal: "",
				},
			},
		},
	}

	testLexer(tcs, t)
}
