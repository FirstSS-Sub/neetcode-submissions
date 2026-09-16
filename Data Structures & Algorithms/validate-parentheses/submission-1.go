func isValid(s string) bool {
    stack := make([]rune, 0, len(s))
	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)
		} else {
			if len(stack) < 1 {
				return false
			}
			if !((r == ')' && stack[len(stack)-1] == '(') ||
				(r == '}' && stack[len(stack)-1] == '{') ||
				(r == ']' && stack[len(stack)-1] == '[')) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
