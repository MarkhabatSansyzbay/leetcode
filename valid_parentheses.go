package myFunctions

// 20. Valid Parentheses

func isValid(s string) bool {
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	open := ""
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			open += string(v)
		} else {
			if open == "" {
				return false
			}
			if pairs[v] != rune(open[len(open)-1]) {
				return false
			}
			open = open[:len(open)-1]
		}
	}

	if open != "" {
		return false
	}
	return true
}
