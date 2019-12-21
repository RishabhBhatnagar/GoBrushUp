package data_structures

func CheckParentheses(expression string) (valid bool) {
	depth := 0
	valid = true
	for _, char := range expression {
		switch char {
		case '(':
			depth++
		case ')':
			depth--
		}
		valid = valid && (depth >= 0)
	}
	return valid && (depth == 0)
}
