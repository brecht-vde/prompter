package engine

import "testing"

func TestVariables(t *testing.T) {
	tests := []struct {
		template  string
		variables map[string]interface{}
		expected  string
	}{
		{
			"Hello {{ var }}",
			map[string]interface{}{"var": "user"},
			"Hello user",
		},
		{
			"Count: {{ var }}",
			map[string]interface{}{"var": 2},
			"Count: 2",
		},
	}

	for _, tt := range tests {
		engine := New()
		result, err := engine.Render(tt.template, tt.variables)

		if err != nil {
			t.Fatalf("could not render template %s. got %s", tt.template, err)
		}

		if result != tt.expected {
			t.Fatalf("template rendered incorrectly. got %s. expected %s", result, tt.expected)
		}
	}
}
