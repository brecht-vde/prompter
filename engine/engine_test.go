package engine

import "testing"

func TestVariables(t *testing.T) {
	tests := []struct {
		template  string
		variables map[string]interface{}
		expected  string
	}{
		{
			`Generate a personalized product recommendation message for a customer. The message should include a greeting, an introduction that acknowledges the customer's recent activities or interests, a list of recommended products, and a friendly closing. Use the following details to craft the message:
				Customer Name: {{var: CustomerName}}
				Customer Interests: {{join: CustomerInterests, ", "}}
				Recommended Products: {{join: RecommendedProducts, ", "}}
				Minimum Rating: {{var: MinimumRating}}
				Categories: {{join: Categories, " - "}}`,
			map[string]interface{}{
				"CustomerName":        "My Corporation ltd",
				"CustomerInterests":   []string{"Sports", "Fishing"},
				"RecommendedProducts": []string{"shin guards", "hockey sticks", "dumbells", "bait"},
				"MinimumRating":       4,
				"Categories":          []int{1, 2, 3},
			},
			`Generate a personalized product recommendation message for a customer. The message should include a greeting, an introduction that acknowledges the customer's recent activities or interests, a list of recommended products, and a friendly closing. Use the following details to craft the message:
				Customer Name: My Corporation ltd
				Customer Interests: Sports, Fishing
				Recommended Products: shin guards, hockey sticks, dumbells, bait
				Minimum Rating: 4
				Categories: 1 - 2 - 3`,
		},
	}

	for _, tt := range tests {
		engine := NewEngine()
		result, err := engine.Render(tt.template, tt.variables)

		if err != nil {
			t.Fatalf("could not render template %s. got %s", tt.template, err)
		}

		if result != tt.expected {
			t.Fatalf("template rendered incorrectly. got %s. expected %s", result, tt.expected)
		}
	}
}
