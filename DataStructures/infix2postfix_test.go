package data_structures

import "testing"

func TestInfix2postfix(t *testing.T) {
	tables := []struct{
		expression string
		expected string
	}{
		{"(a + b)", "ab+"},
		{"a + b", "ab+"},
		{"(a + b", ""},
	}
	for _, table := range tables{
		if postfix, _ := Infix2postfix(table.expression); postfix != table.expected {
			t.Errorf("postfix incorrect. Expected %s, Got: %s", table.expected, postfix)
		}
	}
}
