package data_structures

import "testing"

func TestCheckParentheses(t *testing.T) {
	tables := []struct{
		input string
		expected bool
	}{
		{"()()()", true},
		{"((())())()", true},
		{"a + (b * c)", true},
		{" ", true},
		{"(", false},
		{")", false},
		{")(", false},
		{"(()()", false},
		{"(.(.(..).)....", false},
	}
	for _, row := range tables{
		if result := CheckParentheses(row.input); row.expected != result{
			t.Errorf("For Input: %s, Expected: %t, Got: %t", row.input, row.expected, result)
		}
	}
}
