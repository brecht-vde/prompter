package internal

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
		{
			input: "{{var:item}}",
			expected: []Token{
				{
					Type:    T_Illegal,
					Literal: E_UnexpectedCharacterExpectedSpace,
				},
				{
					Type:    T_EOF,
					Literal: "",
				},
			},
		},
		{
			input: "{{var: 123abc}}",
			expected: []Token{
				{
					Type:    T_OpenVar,
					Literal: "{{var:",
				},
				{
					Type:    T_Illegal,
					Literal: E_InvalidIdentifier,
				},
				{
					Type:    T_EOF,
					Literal: "",
				},
			},
		},
		{
			input: "{{var 123}}",
			expected: []Token{
				{
					Type:    T_Illegal,
					Literal: E_UnexpectedCharacterExpectedColon,
				},
				{
					Type:    T_EOF,
					Literal: "",
				},
			},
		},
		{
			input: "{{varr: 123}}",
			expected: []Token{
				{
					Type:    T_Illegal,
					Literal: E_UnknownActionType,
				},
				{
					Type:    T_EOF,
					Literal: "",
				},
			},
		},
		{
			input: "hello {{var: test}}.",
			expected: []Token{
				{
					Type:    T_Plain,
					Literal: "hello ",
				},
				{
					Type:    T_OpenVar,
					Literal: "{{var:",
				},
				{
					Type:    T_Identifier,
					Literal: "test",
				},
				{
					Type:    T_CloseVar,
					Literal: "}}",
				},
				{
					Type:    T_Plain,
					Literal: ".",
				},
				{
					Type:    T_EOF,
					Literal: "",
				},
			},
		},
		{
			input: `hello {{var: User}} {{var: Id}}.`,
			expected: []Token{
				{
					Type:    T_Plain,
					Literal: "hello ",
				},
				{
					Type:    T_OpenVar,
					Literal: "{{var:",
				},
				{
					Type:    T_Identifier,
					Literal: "User",
				},
				{
					Type:    T_CloseVar,
					Literal: "}}",
				},
				{
					Type:    T_Plain,
					Literal: " ",
				},
				{
					Type:    T_OpenVar,
					Literal: "{{var:",
				},
				{
					Type:    T_Identifier,
					Literal: "Id",
				},
				{
					Type:    T_CloseVar,
					Literal: "}}",
				},
				{
					Type:    T_Plain,
					Literal: ".",
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
		{
			input: `hello {{join: items, ", "}}.`,
			expected: []Token{
				{
					Type:    T_Plain,
					Literal: "hello ",
				},
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
					Type:    T_Plain,
					Literal: ".",
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
