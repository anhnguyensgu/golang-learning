package validparentheses

func isValid(s string) bool {
	st := []rune{}

	for _, a := range s {
		if a == '(' || a == '[' || a == '{' {
			st = append(st, a)
		} else if a == ')' || a == ']' || a == '}' {
			if len(st) == 0 {
				return false
			}
			cur := st[len(st)-1]
			if (a == ')' && cur == '(') || (a == ']' && cur == '[') || (a == '}' && cur == '{') {
				st = st[0 : len(st)-1]
			} else {
				return false
			}

		}
	}

	return len(st) == 0
}
