package services

import (
	"testing"

	"api/services"
)

func TestEvaluateExpression(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		expectErr  bool
	}{
		{"2+3", 5, false},
		{"5-3", 2, false},
		{"2*3", 6, false},
		{"6/3", 2, false},
		{"3+4-4", 3, false},
		{"3+4*2", 11, false},
		{"(3+4)*2", 14, false},
		{"3/0", 0, true},
		{"3++4", 0, true},
		{"3*2", 6, false},
		{"", 0, true},
		{"2 + (3 * 4) - 5", 9, false},
	}

	for _, test := range tests {
		result, err := services.EvaluateExpression(test.expression)
		if (err != nil) != test.expectErr {
			t.Errorf("For expression '%s': Expected error: %v, got: %v", test.expression, test.expectErr, err)
		}

		if result != test.expected {
			t.Errorf("For expression '%s': Expected result: %f, got: %f", test.expression, test.expected, result)
		}
	}
}
