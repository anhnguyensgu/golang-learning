package main

import "fmt"

// https://www.cuemath.com/algebra/sum-of-odd-numbers/

func isPerfectSquareWithFormula(num int) bool {
	i := 1
	for num > 0 {
		num -= i
		i += 2
		if num == 0 {
			return true
		}
	}
	return false
}

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	l := 1
	r := int(num / 2)
	for l <= r {
		mid := l + (r-l)/2
		squa := mid * mid
		if squa == num {
			return true
		}

		if squa > num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquareWithFormula(3))
}
