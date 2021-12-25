package main

import "fmt"

/*
() = 1 0
(()) = 0 1 2 1 0
()() = 1 0 1 0

*/

func loopTopToBotAndBotToTop(s string) int {
	ans := 0
	left := 0
	right := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			if left*2 > ans {
				ans = left * 2
			}
		} else if right > left {
			left, right = 0, 0
		}
	}

	left = 0
	right = 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			if left*2 > ans {
				ans = left * 2
			}
		} else if right < left {
			left, right = 0, 0
		}
	}
	return ans
}

func dynamicProgramming(s string) int {
	dp := make([]int, len(s))

	ans := 0

	for i := 0; i < len(s); i++ {
		dp[i] = 0
	}

	for i := 1; i < len(s); i++ {
		if s[i] == ')' && s[i-1] == '(' {
			if i-2 >= 0 {
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2
			}
		} else if s[i-1] == ')' && s[i] == ')' {
			coverdSecond := i - dp[i-1] - 1
			if coverdSecond >= 0 && s[coverdSecond] == '(' {
				if coverdSecond-1 >= 0 {
					dp[i] = dp[i-1] + 2 + dp[coverdSecond-1]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}

func longestValidParentheses(s string) int {
	stack := []int{-1}
	sum := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				//save new base position
				stack = append(stack, i)
			} else {
				if i-stack[len(stack)-1] > sum {
					sum = i - stack[len(stack)-1]
				}
			}
		}

	}

	return sum
}

func main() {
	fmt.Println(loopTopToBotAndBotToTop("(()"))
	// fmt.Println(longestValidParentheses(")()())"))
	// fmt.Println(longestValidParentheses("()"))
	// fmt.Println(longestValidParentheses("(()"))
	// fmt.Println(longestValidParentheses(""))
	// fmt.Println(longestValidParentheses("(())"))
	// fmt.Println(longestValidParentheses("()()"))
	// fmt.Println(loopTopToBotAndBotToTop("(())()(()))()())))"))
}
