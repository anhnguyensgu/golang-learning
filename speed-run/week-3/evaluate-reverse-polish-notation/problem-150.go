package evaluate_reverse_polish_notation

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		a, _ := strconv.Atoi(tokens[0])
		return a
	}

	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	ans := 0
	stack := []int{}
	//closed := false
	idx := 0
	for idx < len(tokens) {
		token := tokens[idx]
		if _, ok := operators[token]; ok {
			cur := len(stack) - 1
			rh := stack[cur]
			lh := stack[cur-1]
			switch token {
			case "+":
				stack[cur-1] = lh + rh
			case "-":
				stack[cur-1] = lh - rh
			case "*":
				stack[cur-1] = lh * rh
			case "/":
				stack[cur-1] = lh / rh
			}
			ans = stack[cur-1]
			stack = stack[0:cur]
		} else {
			s, _ := strconv.Atoi(token)
			stack = append(stack, s)
		}
		idx++
	}
	return ans
}

// 1 2 3 + 1 + +
// 1 2 + 3 *
